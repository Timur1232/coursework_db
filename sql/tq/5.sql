SELECT r.first_name, r.second_name, v.status FROM vgk_rescuers r
JOIN vgk v ON r.id_vgk = v.id_vgk
WHERE r.experience_years > 15;
