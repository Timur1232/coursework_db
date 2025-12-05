SELECT id_vgk, id_object, formation_date FROM vgk
WHERE id_vgk = 1;

SELECT id_rescuer, first_name, second_name, id_vgk FROM vgk_rescuers
WHERE id_vgk = 1;

UPDATE vgk SET id_vgk = DEFAULT WHERE id_vgk = 1;

SELECT id_vgk, id_object, formation_date FROM vgk
WHERE id_vgk = (SELECT last_value FROM vgk_id_vgk_seq);

SELECT id_rescuer, first_name, second_name, id_vgk FROM vgk_rescuers
WHERE id_vgk = (SELECT last_value FROM vgk_id_vgk_seq);
