Тестовое задание на позицию GO-разработчика
Проект реализует API веб сервера с БД postgressql развернутой у меня локально
Исполняемый файл main.go находится в cmd
все внутренние файлы лежат в папке internal.
для тестирования использовал запросы при помощи curl
Примеры запросов к серверу:
Получение данных из таблицы * вместо * поставть measure или product
curl -v http://localhost:8080/*/
GET запрос по известному id
curl -v http://localhost:8080/product/id
Пример PUT запроса:
curl -v -X PUT http://localhost:8080/measure/id -H "Content-Type: application/json" -d '{"name":"kilogram"}'
POST запрос используется для обновления данных в таблице
curl -v -X POST http://localhost:8080/product/ -H "Content-Type: application/json" -d '{"name":"Orange","quantity":14,"unit_cost":2.01,"measure":2}'
DELETE запрос
curl -v -X DELETE http://localhost:8080/measure/5 
