ifneq (,$(wildcard .env))
	include .env
	export
endif

migrate-up:
	migrate -path ./migrations -database "${DB_URL}" up

migrate-down:
	migrate -path ./migrations -database "${DB_URL}" down

migrate-force:
	migrate -path ./migrations -database "${DB_URL}" -verbose force 1
