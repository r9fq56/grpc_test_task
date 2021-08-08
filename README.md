# grpc_test_task

## Подготовка к запуску
1. С помощью `go mod init` проинициализировать клиентские и серверные каталоги
2. С помощью `go build -i -v -o bin/{server/client}` собрать код
3. Установить postgres, создать базу данных, таблицы развернуть по следующей схеме:

<pre>var schema = `
CREATE TABLE intdumps (
	id char(36) NOT NULL PRIMARY KEY,
	create_date timestamp
);

CREATE TABLE dumpdata (
	id_dump char(36),
	mac char(17),
	name char(50),
	ip char(15),
	dns char(15),
	gateway char(30)
)`</pre>

## Запуск make_interface_dump
1. Перейти в папку make_interface_dump, запустить сервис bin/server
2. Перейти в папку client_make, запустить сервис bin/client
3. В отдельном терминале выполнить запрос с помощью curl, в случае успешного выполнения вернется id записи.

`curl -X POST http://localhost:8081/dump/v1 -d '{"name": "name","mac": "mac", "ip": "ip", "dns": "dns", "gateway": "gateway"}'`

`{"id":"0e672a98-a093-4efe-8723-41541121d246"}`

## Запуск get_interface_dumps
1. Перейти в папку get_interface_dumps, запустить сервис bin/server
2. Перейти в папку client_get, запустить сервис bin/client
3. В отдельном терминале выполнить запрос с помощью curl, где last_count - лимит по последним записям. В слуае успешного выполнения вернется id записи.

`curl http://localhost:8082/dump/v1/{last_count}`

`{"dumps":[{"id":"8bc14b18-9187-40df-9623-cef26866841a","name":"testname222","mac":"test_mac","ip":"127.0.0.2","dns":"test dns","gateway":"test gateway"},{"id":"8f563f58-ca32-4cac-b4d7-f43c6d8bc119","name":"testname222","mac":"test_mac","ip":"127.0.0.2","dns":"test dns","gateway":"test gateway"}]}`
