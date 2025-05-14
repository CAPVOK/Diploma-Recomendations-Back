# Thesis Backend

### Инструкция по запуску

1. В `/config` создать файл конфигурации `dev.yaml`
2. Скопировать содержимое `/config/example.yaml` в файл конфигурации `dev.yaml`
3. Запустить `docker-compose` командой `docker-compose up --build -d`
    1. На порту 5432 развернется PostgreSQL
4. Запустить проект командой `go run cmd/main.go`

### Полезная информация при локальной разработке
1. Swagger-документация доступна по адресу http://localhost:8080/swagger/index.html#/
3. Креды для подключения к базе данных указаны в конфигурационных файлах (по умолчанию имя пользователя: `root`, пароль: `dbpassword123`)
4. Чтобы сгенерировать сваггер, находясь в корне, необходимо написать следующую команду `swag init -g cmd/application/application.go --parseDependency`
