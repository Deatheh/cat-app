# 🚀 Запуск приложения **Cat-App**

## 📌 Предварительные требования
Перед запуском убедитесь, что у вас установлены:
- [Docker](https://www.docker.com/)
- [Go](https://go.dev/)
- [golang-migrate](https://github.com/golang-migrate/migrate)

## 🔧 Команды для запуска

### 1️⃣ Запуск базы данных PostgreSQL
```sh
docker run --name=cat-db -e POSTGRES_PASSWORD=qwerty -p 5436:5432 -d --rm postgres
```

### 2️⃣ Запуск MinIO (локального S3-хранилища)
```sh
docker run -p 9000:9000 -p 9001:9001 --name minio -v D:\minio\data:/data \
-e "MINIO_ROOT_USER=minio" -e "MINIO_ROOT_PASSWORD=minio124" \
minio/minio server /data --console-address ":9001"
```

### 3️⃣ Выполнение миграций
```sh
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
```

### 4️⃣ Запуск приложения
```sh
go run cmd/main.go
```

## 🎯 Дополнительная информация
- **MinIO Web UI** доступен по адресу: [http://localhost:9001](http://localhost:9001)
- **PostgreSQL** работает на порту `5436`
- **Приложение** слушает указанный в `main.go` порт

## 🛠 Полезные команды
### Остановка контейнеров
```sh
docker stop cat-db minio
```

### Очистка всех контейнеров
```sh
docker rm -f $(docker ps -aq)
```

