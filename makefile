postgres:
	docker run --name postgres-tasks -p 5432:5432 -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres-container createdb --username=postgres --owner=postgres gorm

.PHONY: postgres createdb