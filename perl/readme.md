## Как использовать

1) Пропишите в deps.txt ваши модули

2) Соберите образ

```console
$ docker build -t flperl .
```

3) Запустите контейнер

```console
$ docker run --rm -i --network=host -v $(pwd):/app flperl perl script.pl
```
