\c postgres
drop database t;
create database t;
\c t

\i sql/init.sql
\i sql/functions.sql
\i sql/triggers.sql
\i sql/test_data.sql
