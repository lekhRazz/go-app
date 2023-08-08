BEGIN;

CREATE TABLE role_access_permissions_map (
   id   bigserial primary key not null,
   uuid VARCHAR (36)     not null unique,
   role_id INT,
	 access_id INT,

	 created_at TIMESTAMP default now(),
	 created_by INT,

	 updated_at TIMESTAMP default now(),
   updated_by INT,

   deleted BOOLEAN default false,
	 deleted_at TIMESTAMP,
	 deleted_by INT
  );


COMMIT;