-- =====================[vgk_shift_check]===================== --

-- успешный
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 13;
INSERT INTO vgk_shifts VALUES ('2026-02-01 08:00:00', 13, 1, '2026-02-01 20:00:00');
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 13;

-- провальный
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 14;
INSERT INTO vgk_shifts VALUES ('2025-02-01 08:00:00', 14, 3, '2025-02-01 20:00:00');

-- провальный
SELECT id_vgk, status, id_object, formation_date FROM vgk WHERE id_vgk = 12;
INSERT INTO vgk_shifts VALUES ('2026-02-01 08:00:00', 12, 3, '2026-02-01 20:00:00');
