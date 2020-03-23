### Jenkins

1) Собираем образ из докерфайла

```console
$ docker image build -t jenkins-docker .
```

2) Запускаем контейнер

```console
docker container run -d -p 8181:8080 -v /var/run/docker.sock:/var/run/docker.sock -v jenkins_home:/var/jenkins_home --restart=always --name jenkins jenkins-docker
```

3) Чтобы включить SSL

* Подключиться к терминалу контейнера
* Перейти в **/var/jenkins_home**
* Сгенерировать сертификат:

```console
keytool -genkey -keyalg RSA -alias selfsigned -keystore jenkins_keystore.jks -storepass <store_password> -keysize 2048
```

И перезапустить контейнер

```console
sudo docker container run -d -p 8181:8443 -v /var/run/docker.sock:/var/run/docker.sock -v jenkins_home:/var/jenkins_home --restart=always --name jenkins jenkins-docker --httpsPort=8443 --httpsKeyStore=/var/jenkins_home/jenkins_keystore.jks --httpsKeyStorePassword=<store_password>
```
