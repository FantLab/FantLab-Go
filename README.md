# Reamde

На данный момент репозиторий содержит код API для Фантлаба, написанный на Go. В состоянии активной разработки.

## Порядок работы

- Завести issue с описанием задачи/бага
- Создать у себя ветку FLGO-{#issue}
- Написать код
- Запушить
- Создать Pull request
- Назначить reviewer-а
- ...
- PROFIT

## Protobuf

```console
$ protoc --go_out=. proto/*.proto
```

## Docker

```console
$ docker build -t flgo .
$ docker-compose -f docker-compose/fullstack.debug.yml up
```
