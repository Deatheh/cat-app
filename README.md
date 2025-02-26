Для запуска приложения нужно находится в директории проекта и использвать следующие команды:
<br />
docker run --name=cat-db -e POSTGRES_PASSWORD=qwerty -p 5436:5432 -d --rm postgres <br />
docker run -p 9000:9000 -p 9001:9001 --name minio -v D:\minio\data:/data -e "MINIO_ROOT_USER=minio" -e "MINIO_ROOT_PASSWORD=minio124"  minio/minio server /data --console-address ":9001"  <br />
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up <br />
go run cmd/main.go <br />
