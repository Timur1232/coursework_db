-- =====================[equipment_usage_check]===================== --

-- успешный
SELECT inventory_number, name, status, equipment_type FROM equipment WHERE inventory_number = 1;
INSERT INTO equipment_usage_history VALUES (1, 2, '2024-01-25', '2024-01-26', 'учения');
SELECT inventory_number, id_rescuer, issue_date, purpose FROM equipment_usage_history WHERE inventory_number = 1;

-- провальный
SELECT inventory_number, name, status, equipment_type FROM equipment WHERE inventory_number = 3;
INSERT INTO equipment_usage_history VALUES (3, 2, '2024-01-25', '2024-01-26', 'учения');
