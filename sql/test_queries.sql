-- ===================== ПРОСТЫЕ ЗАПРОСЫ С WHERE =====================
SELECT * FROM objects WHERE object_type = 'surface';

SELECT first_name, second_name, status FROM vgk_rescuers WHERE status = 'on_duty';

SELECT accident_type, begin_date_time FROM accidents WHERE status = 'resolved';

-- ===================== ЗАПРОСЫ С JOIN И WHERE =====================
SELECT a.accident_type, o.name FROM accidents a 
JOIN objects o ON a.id_object = o.id_object 
WHERE a.begin_date_time > '2024-03-01';

SELECT r.first_name, r.second_name, v.status FROM vgk_rescuers r 
JOIN vgk v ON r.id_vgk = v.id_vgk 
WHERE r.experience_years > 15;

SELECT e.name, e.status, l.address FROM equipment e 
LEFT JOIN vgk_locations l ON e.id_vgk_location = l.id_vgk_location 
WHERE e.status = 'needs_repair_service';

-- ===================== UPDATE С КАСКАДНЫМ ОБНОВЛЕНИЕМ =====================

SELECT id_object, name FROM objects;
SELECT id_accident, accident_type, id_object, begin_date_time FROM accidents;
UPDATE objects SET id_object = DEFAULT WHERE id_object = 1;
SELECT id_object, name FROM objects;
SELECT id_accident, accident_type, id_object, begin_date_time FROM accidents;

SELECT id_vgk, id_object, formation_date FROM vgk
WHERE id_vgk = 1;

SELECT id_rescuer, first_name, second_name, id_vgk FROM vgk_rescuers
WHERE id_vgk = 1;

UPDATE vgk SET id_vgk = DEFAULT WHERE id_vgk = 1;

SELECT id_vgk, id_object, formation_date FROM vgk
WHERE id_vgk = (SELECT last_value FROM vgk_id_vgk_seq);

SELECT id_rescuer, first_name, second_name, id_vgk FROM vgk_rescuers
WHERE id_vgk = (SELECT last_value FROM vgk_id_vgk_seq);

-- ===================== DELETE КАСКАДНОЕ =====================
-- Удаление 1

SELECT document_type, id_application, document_url, valid_until FROM candidates_documents WHERE id_application = 1;

DELETE FROM applications_for_admission WHERE id_application = 1;

SELECT document_type, id_application, document_url, valid_until FROM candidates_documents WHERE id_application = 1;

-- Удаление 2
SELECT document_type, id_rescuer, document_url, valid_until FROM vgk_rescuers_documents WHERE id_rescuer = 1;

DELETE FROM vgk_rescuers WHERE id_rescuer = 1;

SELECT document_type, id_rescuer, document_url, valid_until FROM vgk_rescuers_documents WHERE id_rescuer = 1;

-- ===================== DELETE С ЗАПРЕТОМ (RESTRICT) =====================
-- Удаление 1 (должно провалиться)
SELECT id_accident, accident_type, id_object, begin_date_time FROM accidents WHERE id_object = 2;

DELETE FROM objects WHERE id_object = 2;

SELECT id_accident, accident_type, id_object, begin_date_time FROM accidents WHERE id_object = 2;

-- Удаление 2 (должно провалиться)
SELECT id_rescuer, first_name, second_name, position, experience_years FROM vgk_rescuers WHERE position = 'спасатель' LIMIT 1;

DELETE FROM positions WHERE position_name = 'спасатель';

SELECT id_rescuer, first_name, second_name, position, experience_years FROM vgk_rescuers WHERE position = 'спасатель' LIMIT 1;

-- ===================== ТРИГГЕР vgk_status_change_trigger =====================
-- Успешный
SELECT id_rescuer, first_name, status, id_vgk FROM vgk_rescuers WHERE id_vgk = 1;
UPDATE vgk SET status = 'temporarily_inactive' WHERE id_vgk = 1;
SELECT id_rescuer, first_name, status, id_vgk FROM vgk_rescuers WHERE id_vgk = 1;

-- Провальный
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 3;
UPDATE vgk SET status = 'ready' WHERE id_vgk = 3;
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 3;

-- ===================== ТРИГГЕР operations_participation_check =====================
-- Успешный
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 7;
INSERT INTO operations_participations VALUES (7, 7, 'разведка');
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 7;

-- Провальный
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 3;
INSERT INTO operations_participations VALUES (3, 3, 'разведка');
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 3;

--
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 14;
INSERT INTO operations_participations VALUES (14, 5, 'разведка');
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 14;

-- ===================== ТРИГГЕР equipment_usage_check =====================
-- Успешный

SELECT inventory_number, name, status, equipment_type FROM equipment WHERE inventory_number = 1;
INSERT INTO equipment_usage_history VALUES (1, 2, '2024-01-25', '2024-01-26', 'учения');
SELECT inventory_number, id_rescuer, issue_date, purpose FROM equipment_usage_history WHERE inventory_number = 1;

-- Провальный
SELECT inventory_number, name, status, equipment_type FROM equipment WHERE inventory_number = 3;
INSERT INTO equipment_usage_history VALUES (3, 2, '2024-01-25', '2024-01-26', 'учения');

-- ===================== ТРИГГЕР transport_usage_check =====================
-- Успешный
SELECT transport_number, model, type, status FROM transport WHERE transport_number = 1;
INSERT INTO transport_usage_history VALUES (1, 2, '2024-01-25', '2024-01-26', 'доставка');
SELECT transport_number, id_rescuer, departure_date, purpose FROM transport_usage_history WHERE transport_number = 1;

-- Провальный
SELECT transport_number, model, type, status FROM transport WHERE transport_number = 9;
INSERT INTO transport_usage_history VALUES (9, 2, '2024-01-25', '2024-01-26', 'доставка');

-- ===================== ТРИГГЕР vgk_shift_check =====================
-- Успешный
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 13;
INSERT INTO vgk_shifts VALUES ('2026-02-01 08:00:00', 13, 1, '2026-02-01 20:00:00');
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 13;

-- Провальный
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 14;
INSERT INTO vgk_shifts VALUES ('2025-02-01 08:00:00', 14, 3, '2025-02-01 20:00:00');

SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 12;
INSERT INTO vgk_shifts VALUES ('2026-02-01 08:00:00', 12, 3, '2026-02-01 20:00:00');
