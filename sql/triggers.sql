-- =====================[Валидация статуса ВГК]===================== --

CREATE OR REPLACE FUNCTION validate_vgk_status_change()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.status = 'dismissed' THEN
        UPDATE vgk_rescuers
        SET id_vgk = NULL,
            status = 'inactive'
        WHERE id_vgk = NEW.id_vgk;
        RETURN NEW;
    END IF;

    IF NEW.status = 'inactive' THEN
        UPDATE vgk_rescuers
        SET status = 'inactive'
        WHERE id_vgk = NEW.id_vgk;
        RETURN NEW;
    END IF;

    IF NOT check_vgk_readiness(OLD.id_vgk) THEN
        RAISE EXCEPTION 'ВГК (%) не укомплектована. Требуется командир и минимум 2 активных спасателя.', OLD.id_vgk;
    END IF;

    UPDATE vgk_rescuers
    SET status = NEW.status
    WHERE id_vgk = NEW.id_vgk;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER vgk_status_change_trigger
AFTER UPDATE OF status ON vgk
FOR EACH ROW
WHEN (OLD.status IS DISTINCT FROM NEW.status)
EXECUTE FUNCTION validate_vgk_status_change();

-- =====================[Обновление статуса при добавлении]===================== --

CREATE OR REPLACE FUNCTION update_vgk_status()
RETURNS TRIGGER AS $$
DECLARE
    vgk_id integer;
    is_ready boolean;
    vgk_status vgk_status_enum;
BEGIN
    IF NEW.id_vgk IS NOT NULL THEN
        vgk_id := NEW.id_vgk;

        SELECT status INTO vgk_status
        FROM vgk
        WHERE id_vgk = vgk_id;

        IF vgk_status = 'dismissed' THEN
            NEW.status := 'inactive';
            RETURN NEW;
        END IF;

        is_ready := check_vgk_readiness(vgk_id);

        IF is_ready AND vgk_status = 'inactive' THEN
            UPDATE vgk
            SET status = 'on_duty'
            WHERE id_vgk = vgk_id;
        ELSIF NOT is_ready AND vgk_status IN ('on_duty', 'on_shift') THEN
            UPDATE vgk
            SET status = 'inactive'
            WHERE id_vgk = vgk_id;
            NEW.status := 'inactive';
        END IF;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER trg_update_vgk_after_rescuer_insert
BEFORE INSERT ON vgk_rescuers
FOR EACH ROW
EXECUTE FUNCTION update_vgk_status();

CREATE OR REPLACE TRIGGER trg_update_vgk_after_rescuer_delete
AFTER DELETE ON vgk_rescuers
FOR EACH ROW
EXECUTE FUNCTION update_vgk_status();

CREATE OR REPLACE TRIGGER trg_update_vgk_after_rescuer_update
AFTER UPDATE OF id_vgk, position, status ON vgk_rescuers
FOR EACH ROW
WHEN (NEW.id_vgk IS NOT NULL
    AND (OLD.id_vgk IS DISTINCT FROM NEW.id_vgk
    OR OLD.position IS DISTINCT FROM NEW.position
    OR OLD.status IS DISTINCT FROM NEW.status))
EXECUTE FUNCTION update_vgk_status();

-- =====================[Валидация перед отправкой на операцию]===================== --

CREATE OR REPLACE FUNCTION check_and_set_departure()
RETURNS TRIGGER AS $$
DECLARE
    vgk_status vgk_status_enum;
    vgk_object integer;
    op_status operation_status_enum;
    object_id integer;
BEGIN
    SELECT id_object INTO object_id FROM objects
    WHERE id_object =
        (SELECT id_object FROM accidents
        WHERE id_accident =
            (SELECT id_accident FROM accidents_response_operations
            WHERE id_operation = NEW.id_operation));

    SELECT status, id_object INTO vgk_status, vgk_object FROM vgk WHERE id_vgk = NEW.id_vgk;

    IF object_id != vgk_object THEN
        RAISE EXCEPTION 'ВГК (id = %) должна отправляться на свои объекты', NEW.id_vgk;
    END IF;

    IF vgk_status != 'on_duty' AND vgk_status != 'on_shift' THEN
        RAISE EXCEPTION 'ВГК (id = %) должна быть готова для отправки на операцию (id = %)', NEW.id_vgk, NEW.id_operation;
    END IF;

    IF NOT check_vgk_readiness(NEW.id_vgk) THEN
        RAISE EXCEPTION 'ВГК (id = %) не укомплектована для отправки на операцию (id = %)', NEW.id_vgk, NEW.id_operation;
    END IF;

    SELECT status INTO op_status
    FROM accidents_response_operations
    WHERE id_operation = NEW.id_operation;

    IF op_status IN ('completed', 'failed') THEN
        RAISE EXCEPTION 'Нельзя назначить ВГК (id = %) на завершенную операцию (id = %)', NEW.id_vgk, NEW.id_operation;
    END IF;

    UPDATE vgk SET status = 'on_departure' WHERE id_vgk = NEW.id_vgk;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER operations_participation_check
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

CREATE OR REPLACE TRIGGER equipment_usage_check
BEFORE INSERT ON equipment_usage_history
FOR EACH ROW
EXECUTE FUNCTION check_equipment_status();

-- =====================[Валидация транспорта перед выдачей]===================== --

CREATE OR REPLACE FUNCTION check_transport_status()
RETURNS TRIGGER AS $$
DECLARE
    transp_status equipment_status_enum;
BEGIN
    SELECT status INTO transp_status
    FROM transport
    WHERE transport_number = NEW.transport_number;

    IF transp_status != 'operational' THEN
        RAISE EXCEPTION 'Транспорт % не подходит для выдачи. Текущий статус: %', NEW.transport_number, transp_status;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER transport_usage_check
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

    IF vgk_status != 'on_duty' THEN
        RAISE EXCEPTION 'ВГК (id = %) должна быть готова для отправки на смену', NEW.id_vgk;
    END IF;

    IF NOT check_vgk_readiness(NEW.id_vgk) THEN
        RAISE EXCEPTION 'ВГК (id = %) не укомплектована для отправки на смену', NEW.id_vgk;
    END IF;

    UPDATE vgk SET status = 'on_shift' WHERE id_vgk = NEW.id_vgk;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER vgk_shift_check
BEFORE INSERT ON vgk_shifts
FOR EACH ROW
EXECUTE FUNCTION check_and_set_shift();
