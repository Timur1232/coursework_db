-- должно провалится
SELECT id_accident, accident_type, id_object, begin_date_time FROM accidents WHERE id_object = 2;

DELETE FROM objects WHERE id_object = 2;

SELECT id_accident, accident_type, id_object, begin_date_time FROM accidents WHERE id_object = 2;
