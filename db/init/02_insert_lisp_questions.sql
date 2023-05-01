\c quizdb

call add_question(
    'Lisp',
    'Является ли Lisp функциональным языком программирования?',
    array[
        'Да',
        'Нет',
        'Зависит от реализации',
        'Lisp не является языком программирования'
    ],
    0
);

call add_question(
    'Lisp',
    'Какой тип данных используется для представления списков в Lisp?',
    array[
        'nil',
        'cons',
        'item',
        'tuple'
    ],
    1
);

call add_question(
    'Lisp',
    'Как создаётся пустой список в Lisp?',
    array[
        'none',
        'null',
        'nil',
        'void'
    ],
    2
);

call add_question(
    'Lisp',
    'Какой оператор используется для добавления элемента в начало списка в Lisp?',
    array[
        'insert',
        'push',
        'append',
        'cons'
    ],
    3
);

call add_question(
    'Lisp',
    'Какой оператор используется для извлечения первого элемента списка в Lisp?',
    array[
        'car',
        'cdr',
        'pop',
        'front'
    ],
    0
);

call add_question(
    'Lisp',
    'Какой оператор используется для извлечения последнего элемента списка Lisp?',
    array[
        'cdr',
        'last',
        'pop',
        'end'
    ],
    1
);

call add_question(
    'Lisp',
    'Что означает аббревиатура "LISP"?',
    array[
        'Logic Information System Programming',
        'Lisp Programming Language',
        'List Processing',
        'List Programming'
    ],
    2
);

call add_question(
    'Lisp',
    'Какой оператор используется для сравнения двух чисел в Lisp?',
    array[
        'cmp',
        'compare',
        'eq',
        '='
    ],
    3
);

call add_question(
    'Lisp',
    'Какой оператор используется для объявления функции в Lisp?',
    array[
        'defun',
        'function',
        'lambda',
        'declare'
    ],
    0
);

call add_question(
    'Lisp',
    'Какой оператор используется для вызова функции со списком параметров в Lisp?',
    array[
        '(call fn_name fn_args)',
        '(funcall fn_name fn_args)',
        '(apply fn_name fn_args)',
        '(invoke fn_name fn_args)'
    ],
    1
);

call add_question(
    'Lisp',
    'Какой оператор используется для создания анонимной функции в Lisp?',
    array[
        'defun',
        'function',
        'lambda',
        'declare'
    ],
    2
);

call add_question(
    'Lisp',
    'Какой тип данных используется для представления идентификаторов в Lisp?',
    array[
        'atom',
        'keyword',
        'string',
        'symbol'
    ],
    3
);

call add_question(
    'Lisp',
    'Каким образом выполняется код в Lisp?',
    array[
        'Интерпретируется',
        'Компилируется в машинный код',
        'Транслируется в язык C',
        'Транслируется в язык Java'
    ],
    0
);

call add_question(
    'Lisp',
    'Какой оператор используется для проверки наличия элемента в списке в Lisp?',
    array[
        'in',
        'member',
        'contains',
        'elementof'
    ],
    1
);

call add_question(
    'Lisp',
    'Какой оператор используется для удаления дубликатов из списка в Lisp?',
    array[
        'uniq',
        'unique',
        'delete-duplicates',
        'remove-duplicates'
    ],
    2
);

call add_question(
    'Lisp',
    'Какой оператор используется для объединения двух списков в Lisp?',
    array[
        'union',
        'merge',
        'join',
        'concatenate'
    ],
    3
);

call add_question(
    'Lisp',
    'Какой оператор используется для нахождения длины списка в Lisp?',
    array[
        'length',
        'size',
        'len',
        'count'
    ],
    0
);

call add_question(
    'Lisp',
    'Какой оператор используется для выполнения условных выражений в Lisp?',
    array[
        'loop',
        'cond',
        'match',
        'switch'
    ],
    1
);

call add_question(
    'Lisp',
    'Какой оператор используется для выполнения циклов в Lisp?',
    array[
        'repeat',
        'while',
        'loop',
        'for'
    ],
    2
);

call add_question(
    'Lisp',
    'Приведите равнозначную альтернативу форме (if a b c):',
    array[
        '(and (or a b) c)',
        '(and a (or b c))',
        '(or a (and b c))',
        '(or (and a b) c)'
    ],
    3
);
