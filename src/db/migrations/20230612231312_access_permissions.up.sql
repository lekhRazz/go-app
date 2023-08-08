BEGIN;

CREATE TABLE access_permissions (
   id   bigserial primary key not null,
   uuid VARCHAR (36)     not null unique,
   scope VARCHAR(250) unique,
	 api TEXT,
   method VARCHAR(50),
   description TEXT,

	 created_at TIMESTAMP default now(),
	 created_by INT,

	 updated_at TIMESTAMP default now(),
   updated_by INT,

   deleted BOOLEAN default false,
	 deleted_at TIMESTAMP,
	 deleted_by INT
  );


COMMIT;