\c doc

-- user_schema.users definition

CREATE TABLE user_schema.users (
	id bigserial NOT NULL,
	user_id varchar NOT NULL,
	"password" varchar NOT NULL,
	"admin" bool NOT NULL,
	CONSTRAINT users_pk PRIMARY KEY (id)
);
CREATE INDEX users_user_id_idx ON user_schema.users (user_id,"password");
CREATE INDEX users_user_id_idx ON user_schema.users (user_id);