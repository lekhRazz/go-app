BEGIN;
create type enum_gender as ENUM (
'MALE',
'FEMALE',
'OTHERS'
);
COMMIT;