-- =====================[transport_usage_check]===================== --

-- успешный
SELECT transport_number, model, type, status FROM transport WHERE transport_number = 1;
INSERT INTO transport_usage_history VALUES (1, 2, '2024-01-25', '2024-01-26', 'доставка');
SELECT transport_number, id_rescuer, departure_date, purpose FROM transport_usage_history WHERE transport_number = 1;

-- провальный
SELECT transport_number, model, type, status FROM transport WHERE transport_number = 9;
INSERT INTO transport_usage_history VALUES (9, 2, '2024-01-25', '2024-01-26', 'доставка');
