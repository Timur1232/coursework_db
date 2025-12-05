SELECT id_object, name FROM objects;
SELECT id_accident, accident_type, id_object, begin_date_time FROM accidents;
UPDATE objects SET id_object = DEFAULT WHERE id_object = 1;
SELECT id_object, name FROM objects;
SELECT id_accident, accident_type, id_object, begin_date_time FROM accidents;
