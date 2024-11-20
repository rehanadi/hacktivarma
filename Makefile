DATABASE_NAME = phar1
DATABASE_PASSWORD = 
DATABASE_PORT = 3306

up:
	migrate -database "mysql://root:$(DATABASE_PASSWORD)@tcp(localhost:$(DATABASE_PORT))/$(DATABASE_NAME)" -path db/migrations up

down:
	migrate -database "mysql://root:$(DATABASE_PASSWORD)@tcp(localhost:$(DATABASE_PORT))/$(DATABASE_NAME)" -path db/migrations down

all: down up