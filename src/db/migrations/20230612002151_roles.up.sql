BEGIN;

CREATE TABLE roles (
   id   bigserial primary key not null,
   uuid VARCHAR (36)     not null unique,
   code enum_role_code  not null,
	 name VARCHAR(250)  not null unique,
	 created_at TIMESTAMP default now(),
	 created_by INT,

	 updated_at TIMESTAMP default now(),
   updated_by INT,

   deleted BOOLEAN default false,
	 deleted_at TIMESTAMP,
	 deleted_by INT
  );


COMMIT;