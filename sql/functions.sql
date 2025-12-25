-- =====================[Проверка готовности ВГК]===================== --

CREATE OR REPLACE FUNCTION check_vgk_readiness(vgk_id integer)
RETURNS boolean AS $$
DECLARE
    commanders_count integer;
    active_count integer;
BEGIN
    SELECT COUNT(*) INTO commanders_count
    FROM vgk_rescuers
    WHERE id_vgk = vgk_id
        AND position = 'командир ВГК'
        AND status IN ('on_shift', 'on_duty');

    SELECT COUNT(*) INTO active_count
    FROM vgk_rescuers
    WHERE id_vgk = vgk_id
        AND status IN ('on_shift', 'on_duty');

    RETURN commanders_count = 1 AND active_count >= 2;
END;
$$ LANGUAGE plpgsql;

-- =====================[Перенос]===================== --

CREATE OR REPLACE FUNCTION transfer_application_to_rescuer(app_id integer, user_id integer)
RETURNS integer AS $$
DECLARE
    app_record record;
    med_record record;
    new_rescuer_id integer;
BEGIN
    SELECT * INTO app_record FROM applications_for_admission WHERE id_application = app_id;
    IF NOT FOUND THEN
        RAISE EXCEPTION 'Заявление с id % не найдено', app_id;
    END IF;

    SELECT * INTO med_record FROM candidates_medical_parameters WHERE id_application = app_id;
    IF NOT FOUND THEN
        RAISE EXCEPTION 'Медицинские параметры для заявления % не найдены', app_id;
    END IF;

    INSERT INTO vgk_rescuers (
        first_name,
        second_name,
        surname,
        status,
        birth_date,
        home_address,
        id_user
    ) VALUES (
        app_record.first_name,
        app_record.last_name,
        app_record.surname,
        'inactive',
        app_record.birthday_date,
        app_record.home_address,
        user_id
    ) RETURNING id_rescuer INTO new_rescuer_id;

    INSERT INTO vgk_rescuers_documents (
        document_type,
        id_rescuer,
        document_url,
        valid_until
    )
    SELECT
        document_type,
        new_rescuer_id,
        document_url,
        valid_until
    FROM candidates_documents
    WHERE id_application = app_id;

    INSERT INTO vgk_rescuers_medical_parameters (
        date,
        id_rescuer,
        expire_date,
        health_group,
        height,
        weight,
        conclusion,
        note
    ) VALUES (
        med_record.date,
        new_rescuer_id,
        med_record.expire_date,
        med_record.health_group,
        med_record.height,
        med_record.weight,
        'Принят по заявлению',
        med_record.note
    );

    DELETE FROM candidates_documents WHERE id_application = app_id;
    DELETE FROM candidates_medical_parameters WHERE id_application = app_id;
    DELETE FROM applications_for_admission WHERE id_application = app_id;

    RETURN new_rescuer_id;
END;
$$ LANGUAGE plpgsql;

-- =====================[Проверка окончания смены]===================== --

CREATE OR REPLACE FUNCTION is_shift_finished(shift_end_time timestamp)
RETURNS boolean AS $$
BEGIN
    RETURN shift_end_time <= CURRENT_TIMESTAMP;
END;
$$ LANGUAGE plpgsql;

-- =====================[обновление статуса снаряжения]===================== --

CREATE OR REPLACE FUNCTION update_equipment_status_by_check_date()
RETURNS void AS $$
DECLARE
    equip_record RECORD;
    updated_count integer := 0;
BEGIN
    FOR equip_record IN
        SELECT inventory_number, last_check_date, status
        FROM equipment
    LOOP
        IF equip_record.status = 'operational' AND equip_record.last_check_date < CURRENT_DATE - INTERVAL '1 month' THEN
            UPDATE equipment
            SET status = 'needs_repair_service'
            WHERE inventory_number = equip_record.inventory_number;
            updated_count := updated_count + 1;
        ELSIF equip_record.status = 'in_use' THEN
            RAISE NOTICE 'Снаряжение % в использовании, но требует ремонта', equip_record.inventory_number;
        END IF;
    END LOOP;

    RAISE NOTICE 'Обновлено % единиц оборудования', updated_count;
END;
$$ LANGUAGE plpgsql;

-- =====================[обновление статуса транспорта]===================== --

CREATE OR REPLACE FUNCTION update_transport_status_by_check_date()
RETURNS void AS $$
DECLARE
    transp_record RECORD;
    updated_count integer := 0;
BEGIN
    FOR transp_record IN
        SELECT transport_number, last_check_date, status
        FROM transport
    LOOP
        IF transp_record.status = 'operational' AND transp_record.last_check_date < CURRENT_DATE - INTERVAL '1 month' THEN
            UPDATE transport
            SET status = 'needs_repair_service'
            WHERE transport_number = transp_record.transport_number;
            updated_count := updated_count + 1;
        ELSIF transp_record.status = 'in_use' THEN
            RAISE NOTICE 'Транспорт % используется, но требует ремонт', transp_record.transport_number;
        END IF;
    END LOOP;

    RAISE NOTICE 'Обновлено % единиц транспорта', updated_count;
END;
$$ LANGUAGE plpgsql;

-- =====================[обновление статуса членов ВГК]===================== --

CREATE OR REPLACE FUNCTION update_vgk_rescuers_status()
RETURNS void AS $$
DECLARE
    vgk_record RECORD;
BEGIN
    FOR vgk_record IN
        SELECT id_vgk, status FROM vgk
    LOOP
        IF vgk_record.status IN ('inactive', 'dismissed') THEN
            UPDATE vgk_rescuers
            SET status = 'inactive'
            WHERE id_vgk = vgk_record.id_vgk AND status IN ('on_duty', 'on_shift');
        END IF;
    END LOOP;
END;
$$ LANGUAGE plpgsql;
