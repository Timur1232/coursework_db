-- 1
SELECT * FROM objects WHERE object_type = 'surface';

-- 2
SELECT first_name, second_name, status
FROM vgk_rescuers WHERE status = 'on_duty';

-- 3
SELECT accident_type, begin_date_time
FROM accidents WHERE status = 'resolved';

-- 4
SELECT a.accident_type, o.name FROM accidents a
JOIN objects o ON a.id_object = o.id_object
WHERE a.begin_date_time > '2024-03-01';

-- 5
SELECT r.first_name, r.second_name, v.status FROM vgk_rescuers r
JOIN vgk v ON r.id_vgk = v.id_vgk
WHERE r.experience_years > 15;

-- 6
SELECT e.name, e.status, l.address FROM equipment e
LEFT JOIN vgk_locations l ON e.id_vgk_location = l.id_vgk_location
WHERE e.status = 'needs_repair_service';

-- 7 Находит всех спасателей с опытом более 5 лет, которые активны (не уволены)
SELECT id_rescuer, first_name, second_name, experience_years, status
FROM vgk_rescuers
WHERE experience_years > 5 AND status != 'dismissed';

-- 8 Выводит транспортные средства с пробегом более 100 000 км
SELECT transport_number, model, type, mileage, manufacture_date
FROM transport
WHERE mileage > 100000 AND status != 'written_off'
ORDER BY mileage DESC;

-- 9 Выводит оборудование, требующее ремонта или обслуживания
SELECT inventory_number, name, equipment_type, status, last_check_date
FROM equipment
WHERE status IN ('needs_repair_service', 'under_repair')
ORDER BY last_check_date;

-- 10 Находит всех спасателей старше 40 лет
SELECT id_rescuer, first_name, second_name, birth_date,
       EXTRACT(YEAR FROM AGE(CURRENT_DATE, birth_date)) as age
FROM vgk_rescuers
WHERE birth_date <= CURRENT_DATE - INTERVAL '40 years'
ORDER BY birth_date;

-- 11 Показывает аварии с критическим уровнем опасности
SELECT a.id_accident, a.accident_type, a.begin_date_time, at.danger_level
FROM accidents a
JOIN accident_types at ON a.accident_type = at.accident_name
WHERE at.danger_level = 'critical'
ORDER BY a.begin_date_time DESC;

-- 12 Выводит оборудование, которое не проверялось более 1 года
SELECT inventory_number, name, equipment_type, last_check_date,
       CURRENT_DATE - last_check_date as days_since_check
FROM equipment
WHERE last_check_date < CURRENT_DATE - INTERVAL '1 year'
  AND status != 'written_off'
ORDER BY last_check_date;

-- 13 Находит спасателей, которые являются командирами ВГК
SELECT vr.id_rescuer, vr.first_name, vr.second_name, 
       vr.position, vg.formation_date, o.name as object_name
FROM vgk_rescuers vr
JOIN vgk vg ON vr.id_vgk = vg.id_vgk
JOIN objects o ON vg.id_object = o.id_object
WHERE vr.position = 'командир ВГК'
ORDER BY vr.experience_years DESC;

-- 14 Подсчитывает количество аварий по каждому типу с группировкой по уровню опасности
SELECT danger_level, accident_name, COUNT(*) as accidents_count, 
       MIN(begin_date_time) as earliest_accident, MAX(begin_date_time) as latest_accident
FROM accidents a
JOIN accident_types at ON a.accident_type = at.accident_name
GROUP BY danger_level, accident_name
ORDER BY danger_level;

-- 15 Подсчитывает количество спасателей в каждом ВГК, группируя по id_vgk
SELECT id_vgk, COUNT(*) as rescuers_count,
    MIN(birth_date) as oldest_birthday,
    MAX(experience_years) as max_experience
FROM vgk_rescuers
WHERE id_vgk IS NOT NULL
GROUP BY id_vgk;
