CREATE TABLE IF NOT EXISTS ROLES(
    id SERIAL PRIMARY KEY,
    name text not null
);

INSERT INTO ROLES (name) SELECT * FROM (SELECT 'student') AS tmp WHERE NOT EXISTS (SELECT name FROM ROLES WHERE NAME='student') LIMIT 1;
INSERT INTO ROLES (name) SELECT * FROM (SELECT 'teacher') AS tmp WHERE NOT EXISTS (SELECT name FROM ROLES WHERE NAME='teacher') LIMIT 1;
INSERT INTO ROLES (name) SELECT * FROM (SELECT 'parent') AS tmp WHERE NOT EXISTS (SELECT name FROM ROLES WHERE NAME='parent') LIMIT 1;
INSERT INTO ROLES (name) SELECT * FROM (SELECT 'school_admin') AS tmp WHERE NOT EXISTS (SELECT name FROM ROLES WHERE NAME='school_admin') LIMIT 1;
INSERT INTO ROLES (name) SELECT * FROM (SELECT 'admin') AS tmp WHERE NOT EXISTS (SELECT name FROM ROLES WHERE NAME='admin') LIMIT 1;

CREATE TABLE IF NOT EXISTS USERS(
   id SERIAL PRIMARY KEY,
   login text not null,
   role_id SERIAL REFERENCES ROLES (id),
   password text not null,
   name text not null,
   surname text not null,
   email text unique not null
);