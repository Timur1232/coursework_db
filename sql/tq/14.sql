-- =====================[operations_participation_check]===================== --

-- успешный
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 7;
INSERT INTO operations_participations VALUES (7, 7, 'разведка');
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 7;

-- провальный
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 3;
INSERT INTO operations_participations VALUES (3, 3, 'разведка');
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 3;

-- провальный
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 14;
INSERT INTO operations_participations VALUES (14, 5, 'разведка');
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 14;
