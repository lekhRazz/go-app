BEGIN;

CREATE TABLE admin_login_session_archive (
   id   bigserial primary key not null,
   uuid VARCHAR (36)     not null unique,
   user_id INT,
	 ip VARCHAR(250),
   user_agent VARCHAR(250),
   token text,
   session_secret VARCHAR(36),
	 created_at TIMESTAMP default now(),
	 created_by INT,

	 updated_at TIMESTAMP,
	 updated_by INT,

   deleted BOOLEAN default false,
	 deleted_at TIMESTAMP,
	 deleted_by INT
  );


COMMIT;