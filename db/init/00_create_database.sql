create database quizdb;

-- connect to created database
\c quizdb

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
