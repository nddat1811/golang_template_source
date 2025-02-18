# format code golang
format:
	go fmt ./...

# tạo swagger
swag:
	swag init

# chạy test
test:
	- go test ./...

test_show_covers:
	- go test ./controller/... ./repository/... ./usecase/... -coverprofile=coverage.out
	- go tool cover -html coverage.out

atlas:
	- atlas schema inspect --url "postgres://test:test@localhost:5432/test?sslmode=disable" 
#tạo file sql
a2:
	- atlas schema apply --url "postgres://test:test@localhost:5432/test?sslmode=disable" --to file://schema.hcl --dry-run > output.sql

a3:
	- atlas schema apply --url "postgres://test:test@localhost:5432/test?sslmode=disable" --dev-url "postgres://test:test@localhost:5432/test?sslmode=disable" --to file://schema.hcl


