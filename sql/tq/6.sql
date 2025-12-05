SELECT e.name, e.status, l.address FROM equipment e
LEFT JOIN vgk_locations l ON e.id_vgk_location = l.id_vgk_location
WHERE e.status = 'needs_repair_service';
