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
- При необходимости обновить [документацию](#документация)
- Запушить
- Создать Pull request
- Назначить reviewer-а
- ...
- PROFIT

## Полезная информация

### [Protobuf](https://github.com/golang/protobuf)

Для перегенерации моделей выполните следующий скрипт (в *vscode* уже настроен экшн для расширения **saveAndRun**):

```console
$ ./make_protos.sh
```

### Docker

Для запуска проекта целиком в докере выполните следующие команды:

```console
$ docker-compose -f docker-compose/all.debug.yml build
$ docker-compose -f docker-compose/all.debug.yml up
```

Если нужно запустить только сторонние сервисы (mysql, memcached, etc.):

```console
$ docker-compose -f docker-compose/depsonly.yml up
```

### Memcached

Для дебага мемкеша удобно использовать **telnet**:

```console
$ telnet localhost 11211
```

[Список команд](https://github.com/memcached/memcached/wiki/Commands)

### Тестовый пароль/хэш

password -> $2a$08$5.4GFX2fkP7XWYrpDWQFqup6.NC6MejFMEOmgX30gRCu4AsMd/A0G
