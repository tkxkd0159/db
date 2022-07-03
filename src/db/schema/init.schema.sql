DROP SCHEMA IF EXISTS "app1" CASCADE;
CREATE SCHEMA IF NOT EXISTS "app1" AUTHORIZATION ljs;

CREATE TABLE IF NOT EXISTS app1.contacts(
   id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
   name varchar(255) not null,
   phone varchar(255) not null
);

INSERT INTO app1.contacts(name, phone) VALUES('ljs', '010-4444-355333');