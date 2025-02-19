CONFIG=config/database/dbconfig.yml
ENV=development
# format code golang
format:
	go fmt ./...

# tạo swagger
swag:
	swag init

atlas:
	- atlas schema inspect --url "postgres://test:test@localhost:5432/test?sslmode=disable" 
#tạo file sql
a2:
	- atlas schema apply --url "postgres://test:test@localhost:5432/test?sslmode=disable" --to file://schema.hcl --dry-run > output.sql

a3:
	- atlas schema apply --url "postgres://test:test@localhost:5432/test?sslmode=disable" --dev-url "postgres://test:test@localhost:5432/test?sslmode=disable" --to file://schema.hcl

#báo cho Make biết rằng các lệnh dưới không phải là file.
.PHONY: new up down down2 status

# Kiểm tra cách sử dụng
help:
	@echo "Cách sử dụng: make <task> [name/limit]"
	@echo "Các task: new, up, down, down2, status"

# Tạo migration mới (phải truyền tên) make new name=t22
new:
	@if [ "$(name)" = "" ]; then \
		echo "Vui lòng nhập tên migration!"; \
		exit 1; \
	fi
	sql-migrate new -config=$(CONFIG) -env=$(ENV) $(name)

# Chạy migration lên
up:
	sql-migrate up -config=$(CONFIG) -env=$(ENV)

# Rollback toàn bộ migration
down:
	sql-migrate down -config=$(CONFIG) -env=$(ENV) -limit=0

# Rollback một số migration với limit
down2:
	@if [ "$(limit)" = "" ]; then \
		echo "Vui lòng nhập số lượng migrations để rollback (limit)!"; \
		exit 1; \
	fi
	sql-migrate down -config=$(CONFIG) -env=$(ENV) -limit=$(limit)

# Kiểm tra trạng thái migration
status:
	sql-migrate status -config=$(CONFIG) -env=$(ENV)
