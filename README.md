# Reamde

На данный момент репозиторий содержит код API для Фантлаба, написанный на Go. В состоянии активной разработки.

## Документация

[Общая информация](docs/common.md)

[Список методов](docs/generated.md)

Чтобы актуализировать список методов, выполните следующий скрипт:
```console
$ cd sources
$ go run . -gendocs > ../docs/generated.md
```

## Порядок работы

- Завести issue с описанием задачи/бага
- Создать у себя ветку FLGO-{#issue}
- Написать код
- Написать тесты
- При необходимости обновить [документацию](#документация)
- Запушить
- Создать Pull request
- Назначить reviewer-а
- ...
- PROFIT

## Protobuf

### [Плагин V2](https://github.com/protocolbuffers/protobuf-go)

Для перегенерации моделей выполните следующий скрипт (в *vscode* уже настроен экшн для расширения **saveAndRun**):

```console
$ ./make_protos.sh
```

### Docker

Для запуска проекта через docker-compose выполните следующие команды:

```console
$ docker-compose -f docker-compose/deps.yml -f docker-compose/app.yml build
$ docker-compose -f docker-compose/deps.yml -f docker-compose/app.yml up
```

Если нужно запустить только сторонние сервисы (mysql, memcached, redis, etc.):

```console
$ docker-compose -f docker-compose/deps.yml up
```

### Memcached

Для дебага мемкеша удобно использовать **telnet**:

```console
$ telnet localhost 11211
```

[Список команд](https://github.com/memcached/memcached/wiki/Commands)

## WRK

[Нагрузочное тестирование](https://github.com/wg/wrk)

### Тестовый пароль/хэш

password -> $2a$08$5.4GFX2fkP7XWYrpDWQFqup6.NC6MejFMEOmgX30gRCu4AsMd/A0G
