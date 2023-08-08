BEGIN;

INSERT INTO roles (id, uuid, code, name)
VALUES 
(
  1,
  '2ed17cb0-0885-11ee-ae34-cb474c18a132',
  'SUPER_ADMIN',
  'superadmin'
),
(
  2,
  '471695da-0885-11ee-9d57-e3fccdcdd54a',
  'ADMIN', 
  'admin'
),
(
  3,
  '595c3948-0885-11ee-98dc-3359ac0e883a',
  'USER',
  'end user/customers'
);
COMMIT;