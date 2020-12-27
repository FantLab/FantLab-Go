# Reamde

На данный момент репозиторий содержит код API для Фантлаба, написанный на Go. В состоянии активной разработки.

## Порядок работы

Все изменения через пулл реквесты. Ветки называем FLGO-XXX, где XXX - номер issue. Для прохождения ревью требуется минимум один аппрув и проверка линтером.

## Документация

[Общая информация](docs/common.md)

[Список методов](docs/generated.md)

## Команды

### Генерация документации

``` console
./generate_docs.sh
```

### Генерация протомоделей

``` console
./make_protos.sh
```

### Docker Compose

Для запуска проекта через docker-compose выполните следующие команды:

``` console
docker-compose -f docker-compose/deps.yml -f docker-compose/app.yml build
docker-compose -f docker-compose/deps.yml -f docker-compose/app.yml up
```

Если нужно запустить только сторонние сервисы (mysql, memcached, redis, etc.):

``` console
docker-compose -f docker-compose/deps.yml up
```

### Запуск из консоли

``` console
export $(xargs < debug.env) && cd sources && go run .
```

## Полезные ссылки

[Proto plugin V2](https://github.com/protocolbuffers/protobuf-go)
[Memcache](https://github.com/memcached/memcached/wiki/Commands)
[WRK](https://github.com/wg/wrk)
[ELK + Docker](https://github.com/deviantony/docker-elk)
[ELK Integration](https://www.elastic.co/blog/how-to-instrument-your-go-app-with-the-elastic-apm-go-agent)

### Тестовый пароль/хэш

password -> $2a$08$5.4GFX2fkP7XWYrpDWQFqup6.NC6MejFMEOmgX30gRCu4AsMd/A0G
