\c quizdb

create procedure add_question(topic text, question text, answers text[], correct_index integer)
language plpgsql as $$
declare
    qid integer;
    i integer;
    answer text;
begin
    insert into questions (topic, "text")
    values (topic, question)
    returning id into qid;

    for i in 1..array_length(answers, 1)
    loop
        insert into answers (question_id, "text", is_correct)
        values (qid, answers[i], i-1 = correct_index);
    end loop;
end;
$$;
