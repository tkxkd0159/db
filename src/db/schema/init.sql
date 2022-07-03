CREATE DATABASE jsroot WITH OWNER ljs;
GRANT ALL PRIVILEGES ON DATABASE jsroot TO ljs;


CREATE SCHEMA "app1" AUTHORIZATION ljs;

CREATE TABLE app1.contacts(
   id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
   name varchar(255) not null,
   phone varchar(255) not null
);

INSERT INTO app1.contacts(name, phone) VALUES('ljs', '010-2222-3333');