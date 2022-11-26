-- Auth tables
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


-- Education tables
CREATE TABLE IF NOT EXISTS SCHOOLS (
    id serial primary key,
    name text not null,
    location text not null,
    expiration_date date not null
);

CREATE TABLE IF NOT EXISTS TEACHERS(
    id serial primary key,
    user_id serial references USERS (id),
    school_id serial references schools (id)
);

CREATE TABLE IF NOT EXISTS CLASSES(
    id serial primary key,
    teacher_id serial references teachers (id),
    school_id serial references schools (id),
    name text not null,
    code text not null
);

CREATE TABLE IF NOT EXISTS STUDENTS(
    id serial primary key,
    user_id serial references USERS (id),
    class_id serial references classes (id)
);

CREATE TABLE IF NOT EXISTS PARENTS(
    id serial primary key,
    user_id serial references USERS (id),
    student_id serial references students (id)
);

CREATE TABLE IF NOT EXISTS SUBJECTS(
    id serial primary key,
    teacher_id serial references teachers (id),
    name text not null,
    description text not null
);

CREATE TABLE IF NOT EXISTS HOMEWORKS (
    id serial primary key,
    name text not null,
    description text not null,
    deadline timestamp not null,
    subject_id serial references subjects (id)
);

CREATE TABLE IF NOT EXISTS SCHEDULES (
    id serial primary key,
    name text not null,
    subject_id serial references subjects (id),
    class_id serial references classes (id)
);

CREATE TABLE IF NOT EXISTS SCHEDULE_LESSONS(
    id serial primary key,
    schedule_id serial references schedules (id),
    start_time timestamp not null,
    end_time timestamp not null
);

CREATE TABLE IF NOT EXISTS TESTS(
    id serial primary key,
    name text not null,
    description text not null,
    duration int not null,
    created_at timestamp not null,
    teacher_id serial references teachers (id)
);

CREATE TABLE IF NOT EXISTS QUESTIONS(
    id serial primary key,
    test_id serial references tests (id),
    text text not null
);

CREATE TABLE IF NOT EXISTS OPTIONS(
    id serial primary key,
    question_id serial references questions (id),
    text text not null,
    is_correct boolean not null
);

CREATE TABLE IF NOT EXISTS TEST_RESULTS(
    id serial primary key,
    test_id serial references tests (id),
    student_id serial references students (id),
    incorrect_answers int not null,
    correct_answers int not null,
    time_spent int not null,
    pass_date timestamp not null
);

CREATE TABLE IF NOT EXISTS EVENTS (
    id serial primary key,
    title text not null,
    description text not null,
    date timestamp not null
);

CREATE TABLE IF NOT EXISTS PROJECTS(
    id serial primary key,
    name text not null,
    type text not null,
    level text not null,
    description text not null,
    tech_components text not null,
    duration int not null,
    purchase_box_link text not null,
    creator_Id serial references users (id),
    source_code text not null,
    block_code text not null,
    cover_img_url text not null,
    scheme_img_url text not null
);

CREATE TABLE IF NOT EXISTS STUDENT_PROJECTS(
    id serial primary key,
    name text not null,
    project_id serial references projects(id),
    created_at timestamp not null
);

CREATE TABLE IF NOT EXISTS REPORTS(
    id serial primary key,
    information text not null,
    download_link text not null,
    created_at date not null
);

CREATE TABLE IF NOT EXISTS STATISTICS (
    id serial primary key,
    information text not null,
    created_at date not null
);

-- Curriculum tables
CREATE TABLE IF NOT EXISTS CURRICULUMS (
    id serial not null primary key,
    teacher_id serial references teachers (id),
    title text not null,
    description text not null,
    created_at timestamp not null
);

CREATE TABLE IF NOT EXISTS MODULES (
    id serial not null primary key,
    curriculum_id serial references CURRICULUMS (id),
    title text not null,
    description text not null,
    created_at timestamp not null
);

CREATE TABLE IF NOT EXISTS BLOCKS (
    id serial not null primary key,
    module_id serial references MODULES (id),
    title text not null,
    created_at timestamp not null
);

CREATE TABLE IF NOT EXISTS CURRICULUM_LESSONS (
    id serial not null primary key,
    block_id serial references BLOCKS (id),
    title text not null,
    description text not null,
    type text not null,
    content text not null,
    created_at timestamp not null
);

