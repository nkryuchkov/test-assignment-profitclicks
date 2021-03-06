[![Go Report Card](http://goreportcard.com/badge/nkryuchkov/test-assignment-profitclicks)](http://goreportcard.com/report/nkryuchkov/test-assignment-profitclicks)
[![GoDoc](https://godoc.org/github.com/nkryuchkov/test-assignment-profitclicks?status.svg)](https://godoc.org/github.com/nkryuchkov/test-assignment-profitclicks)
[![GitHub license](https://img.shields.io/github/license/nkryuchkov/test-assignment-profitclicks.svg)](https://github.com/nkryuchkov/test-assignment-profitclicks/blob/master/LICENSE)

# Как запустить

1) Запустить MySQL.
2) Выбрать пользователя, который будет владеть базой, создав нового при необходимости.
3) Выбрать/создать базу данных у этого пользователя, в которой будут храниться данные.
4) Скопировать `config.example.json` в `config.json`
5) В `config.json` прописать строку соединения к MySQL и изменить другие параметры при необходимости.
6) Скомпилировать и запустить либо запустить через `go run main.go`.

# API

В качестве префикса API необходимо использовать `/api/v1`, т. е. для роута `GET /example` нужно сделать запрос `GET http://host:port/api/v1/example`.

При запросе с методом, отличным от требуемого, возвращается ответ со статусом `405 Method not allowed`.

Реализованы следующие роуты:

- `POST /number?uid=<UID списка>&number=<число>` добавляет число в список UID
- `POST /list` создает новый список и возвращает его UID в формате `{"uid": "<UID списка >"}`
- `DELETE /list?uid=<UID списка>` удаляет список с переданным UID
- `POST /operation?uid=<UID списка>&name=<имя операции>` устанавливает для списка с переданным UID операцию с переданным именем; если такой операции не существует, то возвращается ошибка; допустимые операции — sum и count
- `GET /result?uid=<UID списка>` возврашает результат установленной операции над списком в формате `{"result": <результат>}`

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
