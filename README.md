# Задание

Необходимо написать на языке go сервис http для обработки данных, включающее в
себя api для работы с несколькими сущностями:

1) Список - это последовательный список, в котором хранятся числа (float64) они
могут добавляться, удаляться, скорее всего это будет структура с методами, у
каждого списка есть некий генерируемый uid (можно взять хэш md5 текущего
времени).

2) Списков - может быть много, они могут добавляться и удаляться - то есть
должен быть CRUD без Update.

3) Операция - это структура со свойствами: Имя (name) и указатель на Список
чисел.

Назначение сущности, это привязка функционала который будет происходить с
числами.

Для примера можно реализовать 2 операции: Сумма (сложение списка чисел),
Количество ( количество элементов в выбранной очереди).

У структур должна быть реализация метода GetResult которая возвращает
рассчитанное значение. Каждый вызов GetResult должен приводить к
логированию в таблицу MySql <имя операции><полученное
значение><датавремя> метка времени должна ставиться автоматически.

Для роутинга должны быть реализованы несколько видов вызова:

AddNumberToList(uid, число) - Добавление числа в список.

AddNumberList() - Добавление списка. (возвращает uid)

DeleteNumberList(uid) - Удаление списка чисел.

AddOperationToList(uid, name) - Устанавливает операцию на список.

GetListResult(uid) - Получает значение операции.
