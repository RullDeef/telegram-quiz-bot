create database quizdb_test;

-- connect to created database
\c quizdb_test

create table users (
    id serial,
    nickname text,
    telegram_id text,
    role text
);

create table quizzes (
    id serial,
    topic text,
    creator_id integer,
    created_at timestamp
);

create table questions (
    id serial,
    "text" text
);

create table answers (
    id serial,
    question_id integer,
    "text" text,
    is_correct boolean
);

create table statistics (
    user_id integer,
    quizzes_completed integer,
    mean_quiz_complete_time real,
    mean_question_reply_time real,
    correct_replies integer,
    correct_replies_percent integer
);

insert into users values (-1, "Jacob", "some_id1", "USER")