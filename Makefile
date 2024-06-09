# note: call scripts from /scripts
migration_up:
	migrate -path db/migrations/ -database "postgresql://postgres:password@localhost:5432/finder_job?sslmode=disable" -verbose up

migration_down:
	migrate -path db/migrations/ -database "postgresql://postgres:password@localhost:5432/finder_job?sslmode=disable" -verbose down

migration_fix:
	migrate -path db/migrations/ -database "postgresql://postgres:password@localhost:5432/finder_job?sslmode=disable" force VERSION

build:
	go build -o bin/main cmd/app/main.go

run:
	go run cmd/app/main.go


