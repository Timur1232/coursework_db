-- должно провалится
SELECT id_rescuer, first_name, second_name, position, experience_years FROM vgk_rescuers WHERE position = 'спасатель' LIMIT 1;

DELETE FROM positions WHERE position_name = 'спасатель';

SELECT id_rescuer, first_name, second_name, position, experience_years FROM vgk_rescuers WHERE position = 'спасатель' LIMIT 1;
