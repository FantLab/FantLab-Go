### Jenkins

1) Собираем образ из докерфайла

```console
$ docker image build -t jenkins-docker .
```

2) Запускаем контейнер

```console
$ docker container run -d -p 8181:8080 -p 50000:50000 -v /var/run/docker.sock:/var/run/docker.sock -v jenkins_home:/var/jenkins_home --restart=always jenkins-docker
```
