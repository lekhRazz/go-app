BEGIN;

INSERT INTO access_permissions (id, uuid, scope, api, method, description)
VALUES 
(
  1,
  'c8c5d3cc-1ff8-11ee-951f-ff18d97a8da1',
  'read:admins',
  '/api/admins',
  'GET',
  'This permission allows user to list all the admins in the system'
),
(
  2,
  '2de1c4fa-1ff9-11ee-b024-3f48310b60a1',
  'create:admin',
  '/api/admins',
  'POST',
  'This permission allows user to create admin in the system'
),
(
  3,
  '48d49b34-1ff9-11ee-a677-178327abf2d9',
  'update:admin',
  '/api/admins/:uuid',
  'PUT',
  'This permission allows user to update admin details in the system'
),
(
  4,
  '746180aa-1ff9-11ee-9ee2-83124ca69004',
  'read:admin-detail',
  '/api/admins/:uuid',
  'GET',
  'This permission allows user to read admin detail in the system'
),
(
  5,
  '7ab586ae-1ff9-11ee-b713-93f8bbd957b8',
  'delete:admin',
  '/api/admins/:uuid',
  'PATCH',
  'This permission allows user to delete admin detail in the system'
);

COMMIT;