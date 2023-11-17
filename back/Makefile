db_container:
	docker run --name iasked_db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecret -d postgres:16.1-alpine3.18

create_db:
	docker exec -it iasked_db createdb --username=root --owner=root iasked

drop_db:
	docker exec -it iasked_db dropdb iasked

start_db_container:
	docker start iasked_db

stop_db_container:
	docker stop iasked_db

connect_db_container:
	docker exec -it iasked_db psql -d iasked

migrate_up:
	migrate -path ./db/migration -database "postgresql://root:mysecret@localhost:5432/iasked?sslmode=disable" -verbose up

migrate_down:
	migrate -path ./db/migration -database "postgresql://root:mysecret@localhost:5432/iasked?sslmode=disable" -verbose down

.PHONY: db_container create_db drop_db start_db_container stop_db_container connect_db_container