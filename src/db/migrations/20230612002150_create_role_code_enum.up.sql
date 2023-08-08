BEGIN;
create type enum_role_code as ENUM (
'SUPER_ADMIN',
'ADMIN',
'USER'
);
COMMIT;