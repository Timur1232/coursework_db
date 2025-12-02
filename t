1. Таблица equipment_types
- type_name (varchar(255)) - Уникальное наименование типа, используемое в качестве ключа.
- equipment_standards_url (varchar(255)) - Ссылка на внешний документ с нормативами.

2. Таблица objects
- id_object (integer) - Уникальный числовой идентификатор объекта.
- object_type (varchar(255)) - Категория объекта (шахта, рудник и т.д.).
- name (varchar(255)) - Название предприятия.
- address (varchar(255)) - Физический адрес расположения.
- phone (varchar(255)) - Контактный телефон.
- email (varchar(255)) - Адрес электронной почты.
- director_full_name (varchar(255)) - ФИО руководителя объекта.

3. Таблица accidents
- id_accident (integer) - Уникальный номер инцидента.
- id_object (integer) - Ссылка на объект, где произошла авария.
- accident_type (varchar(255)) - Классификация происшествия.
- begin_date_time (timestamp) - Дата и время начала ЧП.
- status (varchar(255)) - Текущий статус ликвидации.
- description (text) - Подробное описание ситуации.
- first_estimate (text) - Первоначальная экспертная оценка.
- cause (text) - Установленная причина аварии.

4. Таблица applications_for_admission
- id_application (integer) - Уникальный номер заявки кандидата.
- id_object (integer) - Ссылка на объект, куда подана заявка.
- passport_number (integer) - Серия и номер паспорта.
- first_name (varchar(255)) - Имя кандидата.
- last_name (varchar(255)) - Фамилия кандидата.
- surname (varchar(255)) - Отчество кандидата (может отсутствовать).
- issue_date (date) - Дата выдачи паспорта.
- phone (varchar(255)) - Контактный телефон кандидата.
- email (varchar(255)) - Адрес электронной почты.
- status (varchar(255)) - Статус рассмотрения заявки.
- birthday_date (date) - Дата рождения кандидата.

5. Таблица candidates_documents
- document_type (varchar(255)) - Тип предоставляемого документа (диплом, справка).
- id_application (integer) - Ссылка на заявку кандидата.
- document_url (varchar(255)) - Ссылка на цифровую копию документа.
- valid_until (date) - Дата окончания действия документа.

6. Таблица candidates_medical_parameters
- id_application (integer) - Ссылка на заявку кандидата (один к одному).
- date (date) - Дата прохождения медосмотра.
- expire_date (date) - Дата окончания действия справки.
- health_group (integer) - Присвоенная группа здоровья.
- height (decimal) - Рост кандидата.
- weight (decimal) - Вес кандидата.
- note (text) - Дополнительные медицинские заметки.

7. Таблица vgk
- id_vgk (integer) - Уникальный номер военизированной горноспасательной команды (ВГК).
- id_object (integer) - Ссылка на обслуживаемый объект.
- status (varchar(255)) - Текущий статус команды (сформирована, расформирована).
- formation_date (date) - Дата создания команды.

8. Таблица positions
- position_name (varchar(255)) - Уникальное название должности.
- salary (decimal(10,2)) - Размер оклада с точностью до копеек.
- min_experience_years (integer) - Требуемый минимальный стаж.
- responsibilities (text) - Текстовое описание должностных обязанностей.

9. Таблица vgk_rescuers
- id_rescuer (integer) - Уникальный табельный номер спасателя.
- id_vgk (integer) - Ссылка на ВГК, в которой состоит спасатель.
- position (varchar(255)) - Должность спасателя в команде.
- first_name (varchar(255)) - Имя спасателя.
- second_name (varchar(255)) - Фамилия спасателя.
- surname (varchar(255)) - Отчество спасателя.
- status (varchar(255)) - Статус (работает, уволен, в отпуске).
- birth_date (date) - Дата рождения.
- home_address (varchar(255)) - Домашний адрес.
- experience_years (integer) - Общий стаж работы в горноспасательной службе.

10. Таблица vgk_rescuers_documents
- document_type (varchar(255)) - Тип документа спасателя (удостоверение, допуск).
- id_rescuer (integer) - Ссылка на спасателя-владельца.
- document_url (varchar(255)) - Ссылка на цифровую копию.
- valid_until (date) - Дата окончания действия документа.

11. Таблица vgk_shifts
- shift_start (timestamp) - Дата и время начала дежурства (часть ключа).
- id_vgk (integer) - Ссылка на заступающую команду (часть ключа).
- id_vgk_location (integer) - Ссылка на место несения дежурства (часть ключа).
- shift_end (timestamp) - Дата и время окончания дежурства.

12. Таблица accidents_response_operations
- id_operation (integer) - Уникальный номер операции по ликвидации.
- id_accident (integer) - Ссылка на аварию, для которой создана операция.
- id_leader (integer) - Ссылка на спасателя-руководителя операции.
- start_date_time (timestamp) - Дата и время начала операции.
- end_date_time (timestamp) - Дата и время завершения операции.
- status (varchar(255)) - Текущий статус операции.

13. Таблица operations_participations
- id_vgk (integer) - Ссылка на команду-участницу (часть ключа).
- id_operation (integer) - Ссылка на операцию (часть ключа).
- assigned_task (varchar(255)) - Задача, поставленная перед командой.

14. Таблица operations_reports
- id_report (integer) - Уникальный номер отчёта.
- id_operation (integer) - Ссылка на операцию, по которой составлен отчёт.
- report_date_time (timestamp) - Дата и время создания отчёта.
- description (text) - Содержание отчёта.

15. Таблица trainings
- date (date) - Дата проведения тренировки (часть ключа).
- id_object_location (integer) - Ссылка на объект, где проходит тренировка (часть ключа).
- id_instructor (integer) - Ссылка на спасателя-инструктора.
- topic (varchar(255)) - Тематика занятия.
- description (text) - Детальное описание плана тренировки.

16. Таблица trainings_participations
- date (date) - Ссылка на дату тренировки (часть ключа).
- id_rescuer (integer) - Ссылка на спасателя-участника (часть ключа).
- notes (text) - Заметки инструктора о результатах спасателя.

17. Таблица certifications_passings
- date (date) - Дата проведения аттестации (часть ключа).
- id_rescuer (integer) - Ссылка на аттестуемого спасателя (часть ключа).
- result (boolean) - Флаг успешности прохождения (сдал/не сдал).
- topic (varchar(255)) - Тема или вид аттестации.

18. Таблица vgk_rescuers_medical_parameters
- date (date) - Дата прохождения медосмотра (часть ключа).
- id_rescuer (integer) - Ссылка на спасателя (часть ключа).
- expire_date (date) - Дата, до которой действует освидетельствование.
- health_group (integer) - Группа здоровья по результатам.
- height (decimal) - Рост спасателя.
- weight (decimal) - Вес спасателя.
- conclusion (varchar(255)) - Медицинское заключение (годен/не годен).
- note (text) - Примечания врача.

19. Таблица vgk_service_room
- id_service_room (integer) - Уникальный номер сервисного помещения.
- id_responsible (integer) - Ссылка на ответственного за помещение спасателя.
- purpose (varchar(255)) - Назначение помещения (склад, мастерская).
- address (varchar(255)) - Адрес расположения помещения.

20. Таблица vgk_locations
- id_vgk_location (integer) - Уникальный номер места дислокации.
- id_responsible (integer) - Ссылка на спасателя, ответственного за локацию.
- address (varchar(255)) - Физический адрес точки базирования.
- status (varchar(255)) - Статус точки (активна, на ремонте).

21. Таблица equipment
- inventory_number (integer) - Уникальный инвентарный номер.
- id_vgk_location (integer) - Ссылка на текущее место хранения.
- equipment_type (varchar(255)) - Тип оборудования (ссылка на справочник).
- name (varchar(255)) - Наименование конкретного экземпляра.
- status (varchar(255)) - Состояние (исправно, в ремонте, списано).
- last_check_date (date) - Дата последней поверки/проверки.

22. Таблица transport
- transport_number (integer) - Уникальный номер транспортного средства.
- id_vgk_location (integer) - Ссылка на текущее место стоянки.
- model (varchar(255)) - Модель ТС.
- type (varchar(255)) - Тип ТС (автомобиль, квадроцикл).
- status (varchar(255)) - Техническое состояние.
- manufacture_date (date) - Дата выпуска.
- mileage (decimal) - Общий пробег.
- last_check_date (date) - Дата последнего техосмотра.

23. Таблица equipment_usage_history
- inventory_number (integer) - Ссылка на использованное оборудование (часть ключа).
- id_rescuer (integer) - Ссылка на спасателя, взявшего оборудование (часть ключа).
- issue_date (date) - Дата выдачи.
- return_date (date) - Дата возврата.
- purpose (varchar(255)) - Цель использования (дежурство, тренировка).

24. Таблица transport_usage_history
- transport_number (integer) - Ссылка на использованный транспорт (часть ключа).
- id_rescuer (integer) - Ссылка на спасателя-водителя (часть ключа).
- departure_date (date) - Дата выезда.
- return_date (date) - Дата возврата.
- purpose (varchar(255)) - Цель поездки.

25. Таблица equipment_service_history
- inventory_number (integer) - Ссылка на обслуживаемое оборудование (часть ключа).
- id_service_room (integer) - Ссылка на сервисное помещение (часть ключа).
- reason (varchar(255)) - Причина обслуживания (плановое, ремонт).
- serve_date (date) - Дата проведения работ.
- status (varchar(255)) - Результат обслуживания.

26. Таблица transport_service_history
- transport_number (integer) - Ссылка на обслуживаемый транспорт (часть ключа).
- id_service_room (integer) - Ссылка на сервисное помещение (часть ключа).
- reason (varchar(255)) - Причина обслуживания.
- serve_date (date) - Дата проведения работ.
- status (varchar(255)) - Итоговый статус после обслуживания.
