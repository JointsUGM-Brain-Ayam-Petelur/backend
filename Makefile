.PHONY: migrate-create

migrate-create:
	@test -n "$(name)" || (echo "Usage: make migrate-create name=create_users_table"; exit 1)
	migrate create -ext sql -dir migrations -seq $(name)