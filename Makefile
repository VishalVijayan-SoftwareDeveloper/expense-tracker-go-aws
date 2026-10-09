DB_URL=postgres://postgres:postgres@host.docker.internal:5432/expense_db?sslmode=disable

migrate-up:
	docker run --rm \
	-v ${PWD}/migrations:/migrations \
	migrate/migrate \
	-path=/migrations \
	-database "$(DB_URL)" \
	up

migrate-down:
	docker run --rm \
	-v ${PWD}/migrations:/migrations \
	migrate/migrate \
	-path=/migrations \
	-database "$(DB_URL)" \
	down 1


migrate-force:
	migrate -path migrations -database "$(DB_URL)" force 1

up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f

build:
	docker compose build

restart:
	docker compose down
	docker compose up -d

