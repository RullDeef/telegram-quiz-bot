\c quizdb

call add_question(
    'Go',
    'Сколько ключевых слов в языке?',
    array[
        '25',
        '33',
        '44',
        '50'
    ],
    0
);

call add_question(
    'Go',
    'Кто разработал язык программирования Go?',
    array[
        'Microsoft',
        'Google',
        'Apple',
        'Oracle'
    ],
    1
);

call add_question(
    'Go',
    'Как объявить переменную в Go?',
    array[
        'var x int := 10',
        'var x = 10',
        'x := 10',
        'let x = 10'
    ],
    2
);

call add_question(
    'Go',
    'Что означает оператор := в Go?',
    array[
        'Конкатенация строк',
        'Присваивание значения переменной',
        'Инициализация переменной значением по умолчанию',
        'Объявление переменной с присваиванием'
    ],
    3
);

call add_question(
    'Go',
    'Что делает оператор defer в Go?',
    array[
        'Создает новую горутину',
        'Задерживает выполнение функции до тех пор, пока не будет завершено выполнение вызывающей функции',
        'Копирует значение переменной',
        'Ничего не делает, это несуществующая функция'
    ],
    1
);

call add_question(
    'Go',
    'Как объявить функцию без аргументов и без возвращаемого значения в Go?',
    array[
        'void f() {}',
        'function f() {}',
        'func f() {}',
        'func f -> void {}'
    ],
    2
);

call add_question(
    'Go',
    'Как объявить функцию с аргументами и возвращаемым значением в Go?',
    array[
        'func f(x int, y int) int {}',
        'int f(x, y) {}',
        'func f(x int, y int) -> int {}',
        'function f(x, y) -> int {}'
    ],
    0
);

call add_question(
    'Go',
    'Как объявить массив с фиксированным размером в Go?',
    array[
        'var arr [5]int',
        'var arr = [5]int{1, 2, 3, 4, 5}',
        'arr = {1, 2, 3, 4, 5}',
        'var arr []int = {1, 2, 3, 4, 5}'
    ],
    0
);

call add_question(
    'Go',
    'Как объявить срез в Go?',
    array[
        'var s = [5]int{1, 2, 3, 4, 5}',
        'var s []int = {1, 2, 3, 4, 5}',
        'var s []int = make([]int, 5)',
        'var s = []int{1, 2, 3, 4, 5}'
    ],
    3
);

call add_question(
    'Go',
    'Можно ли менять размер среза в Go?',
    array[
        'Да',
        'Нет',
        'Только если срез был объявлен с использованием make',
        'Только если это необходимо для корректной работы программы'
    ],
    0
);

call add_question(
    'Go',
    'Как получить длину среза в Go?',
    array[
        's.length()',
        'length(s)',
        's.len()',
        'len(s)'
    ],
    3
);

call add_question(
    'Go',
    'Как объединить два среза в один в Go?',
    array[
        's1.concat(s2)',
        'concat(s1, s2)',
        'append(s1, s2...)',
        's1 + s2'
    ],
    2
);

call add_question(
    'Go',
    'Как объявить отображение (map) в Go?',
    array[
        'var m map[string]int ',
        'm := map[string]int{}',
        'm := make(map)',
        'var m = [string]int{}'
    ],
    1
);

call add_question(
    'Go',
    'Как добавить пару ключ-значение в map в Go?',
    array[
        'add(m, key, value)',
        'm[key] = value',
        'insert(m, key, value)',
        'append(m, key, value)'
    ],
    1
);

call add_question(
    'Go',
    'Какой символ используется в конце тела функции в Go?',
    array[
        ';',
        '.',
        '}',
        ']'
    ],
    2
);

call add_question(
    'Go',
    'Как конвертировать тип данных в Go?',
    array[
        '(data_type) value',
        'data_type(value)',
        'convert(value, data_type)',
        'value.convert(data_type)'
    ],
    0
);

call add_question(
    'Go',
    'Как объявить цикл for в Go?',
    array[
        'for init; cond; inc {}',
        'for init; cond; inc; {}',
        'for (init; cond; inc) {}',
        'for init, cond, inc {}'
    ],
    0
);

call add_question(
    'Go',
    'Как называется пакет в Go, который содержит точку входа?',
    array[
        'entry',
        'start',
        'main',
        'Пакет можно назвать произвольно'
    ],
    2
);

call add_question(
    'Go',
    'Выберите верное утверждение для программы, запущенной с GOMAXPROCS=5:',
    array[
        'Программа не может создать более пяти потоков ОС',
        'Количество горутин на поток не может превышать пяти',
        'Программа не может одновременно выполняться более чем в пяти потоках ОС',
        'Программа резервирует пять потоков ОС для параллельного выполнения'
    ],
    2
);

call add_question(
    'Go',
    'Какая команда пригодится для поиска узких мест в программе?',
    array[
        'go vet',
        'go tool pprof',
        'go tool objdump',
        'go profile'
    ],
    1
);
