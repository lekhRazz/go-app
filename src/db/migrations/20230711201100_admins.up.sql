BEGIN;

CREATE TABLE admins (
   id   bigserial primary key not null,
   uuid VARCHAR (36)     not null unique,
   first_name VARCHAR(250),
	 middle_name VARCHAR(250),
	 last_name VARCHAR(250),
	 gender enum_gender not null,
	 email VARCHAR(250) not null,
	 phone VARCHAR(250),
	 country_code VARCHAR(250),
   password TEXT not null,
 	 suspended BOOLEAN,
	 suspend_reason text,
	 verified BOOLEAN,
	 role_id INT,
	 
	 created_at TIMESTAMP default now(),
	 created_by INT,

	 updated_at TIMESTAMP,
	 updated_by INT,

   deleted BOOLEAN default false,
	 deleted_at TIMESTAMP,
	 deleted_by INT
  );


COMMIT;