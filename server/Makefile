postgresinit:
	docker run --name=chat-db -e POSTGRES_PASSWORD='qwerty' -p 5433:5433 -d postgres     
	
migrateup: 
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/go-chat?sslmode=disable" -verbose up
.PHONY: postgres