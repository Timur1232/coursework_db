SELECT a.accident_type, o.name FROM accidents a
JOIN objects o ON a.id_object = o.id_object
WHERE a.begin_date_time > '2024-03-01';
