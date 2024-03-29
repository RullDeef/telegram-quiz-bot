\c quizdb

alter table users
    add constraint c_users_pk
        primary key (id),
    add constraint c_users_role
        check (("role") in ('USER', 'ADMIN'));

alter table questions
    add constraint c_questions_pk
        primary key (id);

alter table answers
    add constraint c_answers_pk
        primary key (id),
    add constraint c_answers_fk_question
        foreign key (question_id) references questions(id);

alter table "statistics"
    add constraint c_statistics_pk
        primary key (user_id),
    add constraint c_statistics_fk_user
        foreign key (user_id) references users(id);
