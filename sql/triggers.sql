-- =====================[Валидация статуса ВГК]===================== --

CREATE OR REPLACE FUNCTION validate_vgk_status_change()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.status = 'disbanded' THEN
        UPDATE vgk_rescuers
        SET id_vgk = NULL,
            status = 'inactive'
        WHERE id_vgk = OLD.id_vgk;
        RETURN NEW;
    END IF;

    IF NEW.status = 'temporarily_inactive' THEN
        UPDATE vgk_rescuers
        SET status = 'inactive'
        WHERE id_vgk = OLD.id_vgk;
        RETURN NEW;
    END IF;

    IF NOT check_vgk_readiness(OLD.id_vgk) THEN
        RAISE EXCEPTION 'ВГК не укомплектована. Требуется командир и минимум 2 активных спасателя.';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER vgk_status_change_trigger
BEFORE UPDATE ON vgk
FOR EACH ROW
WHEN (OLD.status IS DISTINCT FROM NEW.status)
EXECUTE FUNCTION validate_vgk_status_change();

-- =====================[Валидация перед отправкой на операцию]===================== --

CREATE OR REPLACE FUNCTION check_and_set_departure()
RETURNS TRIGGER AS $$
DECLARE
    vgk_status vgk_status_enum;
    op_status operation_status_enum;
BEGIN
    SELECT status INTO vgk_status FROM vgk WHERE id_vgk = NEW.id_vgk;

    IF vgk_status != 'ready' THEN
        RAISE EXCEPTION 'ВГК должна быть в статусе ready для отправки на операцию';
    END IF;

    IF NOT check_vgk_manning(NEW.id_vgk) THEN
        RAISE EXCEPTION 'ВГК не укомплектована для отправки на операцию';
    END IF;

    SELECT status INTO op_status
    FROM accidents_response_operations
    WHERE id_operation = NEW.id_operation;

    IF op_status IN ('completed', 'failed') THEN
        RAISE EXCEPTION 'Нельзя назначить ВГК на завершенную операцию';
    END IF;

    UPDATE vgk SET status = 'on_departure' WHERE id_vgk = NEW.id_vgk;

    UPDATE vgk_rescuers
    SET status = 'on_departure'
    WHERE id_vgk = NEW.id_vgk;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER operations_participation_check
BEFORE INSERT ON operations_participations
FOR EACH ROW
EXECUTE FUNCTION check_and_set_departure();

-- =====================[Валидация снаряжения перед выдачей]===================== --

CREATE OR REPLACE FUNCTION check_equipment_status()
RETURNS TRIGGER AS $$
DECLARE
    equip_status equipment_status_enum;
BEGIN
    SELECT status INTO equip_status
    FROM equipment
    WHERE inventory_number = NEW.inventory_number;

    IF equip_status != 'operational' THEN
        RAISE EXCEPTION 'Снаряжение % не в статусе operational. Текущий статус: %', NEW.inventory_number, equip_status;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER equipment_usage_check
BEFORE INSERT ON equipment_usage_history
FOR EACH ROW
EXECUTE FUNCTION check_equipment_status();

-- =====================[Валидация транспорта перед выдачей]===================== --

CREATE OR REPLACE FUNCTION check_transport_status()
RETURNS TRIGGER AS $$
DECLARE
    transp_status transport_status_enum;
BEGIN
    SELECT status INTO transp_status
    FROM transport
    WHERE transport_number = NEW.transport_number;

    IF transp_status != 'operational' THEN
        RAISE EXCEPTION 'Транспорт % не в статусе operational. Текущий статус: %', NEW.transport_number, transp_status;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER transport_usage_check
BEFORE INSERT ON transport_usage_history
FOR EACH ROW
EXECUTE FUNCTION check_transport_status();

-- =====================[Проверка ВГК перед выходом на смену]===================== --

CREATE OR REPLACE FUNCTION check_and_set_shift()
RETURNS TRIGGER AS $$
DECLARE
    vgk_status vgk_status_enum;
BEGIN
    IF is_shift_finished(NEW.shift_end) THEN
        RAISE EXCEPTION 'Нельзя назначить ВГК на завершенную смену';
    END IF;

    SELECT status INTO vgk_status FROM vgk WHERE id_vgk = NEW.id_vgk;

    IF vgk_status != 'ready' THEN
        RAISE EXCEPTION 'ВГК должна быть готова для отправки на смену';
    END IF;

    IF NOT check_vgk_manning(NEW.id_vgk) THEN
        RAISE EXCEPTION 'ВГК не укомплектована для отправки на смену';
    END IF;

    UPDATE vgk SET status = 'on_shift' WHERE id_vgk = NEW.id_vgk;

    UPDATE vgk_rescuers
    SET status = 'on_shift'
    WHERE id_vgk = NEW.id_vgk;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER vgk_shift_check
BEFORE INSERT ON vgk_shifts
FOR EACH ROW
EXECUTE FUNCTION check_and_set_shift();
