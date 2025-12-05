-- =====================[vgk_status_change_trigger]===================== --

-- успешный
SELECT id_rescuer, first_name, status, id_vgk FROM vgk_rescuers WHERE id_vgk = 1;
UPDATE vgk SET status = 'temporarily_inactive' WHERE id_vgk = 1;
SELECT id_rescuer, first_name, status, id_vgk FROM vgk_rescuers WHERE id_vgk = 1;

-- провальный
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 3;
UPDATE vgk SET status = 'ready' WHERE id_vgk = 3;
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 3;
