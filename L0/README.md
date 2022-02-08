# Level 0

## Описание
![https://img.shields.io/badge/version-v1.0.0-green](https://img.shields.io/badge/version-v1.0.0-green) ![https://img.shields.io/badge/Golang-v1.17.6-blue](https://img.shields.io/badge/Golang-v1.17.6-blue) ![https://img.shields.io/badge/Docker--compose-does%20not%20work-red](https://img.shields.io/badge/Docker--compose-does%20not%20work-red)

ТЗ можно посмотреть в файле [L0.pdf](L0.pdf)

Кратко:
- [x] Сервис которые получает данные по подписки в NATS-streaming и отправляет на сервер по id
  - [x] Данные хранит в Кеш и пишет в postgres
  - [x] При подение берет данные из БД
- [x] Сервер выдает данные по id
  - [x] Интерфейс (html)

## Запуск
Запуск условно можно разделить на 3 этапа:
1. [Postgres](#postgres)
2. [NATS-streaming](#nats-streaming)
3. [Сервис](#сервис)

:exclamation: запуск через docker-copmpose пока не работает
### Postgres
`Postgres` поднимаем в `docker` командой:
```
docker run --name PSQL -p 5432:5432 -e POSTGRES_USER=myuser -e POSTGRES_PASSWORD=qwerty -e POSTGRES_DB=l0 -d postgres:14
```
Можете поменять `--name PSQL` эта опция дает название контейнеру с `Postgres`  
`-e POSTGRES_USER=myuser` можете заменить `myuser` на своего user, так же можно изменить пароль `-e POSTGRES_PASSWORD=qwerty`

:exclamation::exclamation::exclamation: Не меняйте `-p 5432:5432` `-e POSTGRES_DB=l0` эти параметры пока что не возможно настроить в сервисе

После создания заходим в `bash` в контейнере:
```
docker exec -it PSQL bash
```
Если меняли `--name PSQL` введите свое название контейнера или его `ID`

Теперь заходим в `PSQL`
```
psql -U myuser -W -d l0
```
Появиться строка
```
Password: 
```
Введите пароль каторый указывали при запуске контейнера.  
:exclamation: Вы не будете видеть вводимые символы

Теперь создаем таблицу в БД `l0`:
```SQL
CREATE TABEL taskl0(
order_uid text,
data json);
```
На этом все можете выходить из `PSQL` и `bash` сделать это можно командами  
Для PSQL:
```
\q
```
Для Bash
```
exit
```

### NATS-streaming
`NATS-streaming` тоже поднимим в `docker`  
Для этого надо запустить команду:
```
docker run --name NATS -p 4222:4222 -p 8222:8222 -d nats-streaming 
```

:exclamation::exclamation::exclamation: Тут не меняйте порты `-p 4222:4222 -p 8222:8222`
### Сервис

Запустим сервис, находясь в папке `Исходники`, командой:
```
go run L0.go
``` 
В терминали вы должны увидеть: 
