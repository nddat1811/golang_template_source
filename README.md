# golang_template_source

# Naming convention
* File naming convention
Generally, file names are single lowercase words.
Go follows a convention where source files are all lower case with underscore separating multiple words.
Compound file names are separated with _
File names that begin with “.” or “_” are ignored by the go tool
Test files in Go come with suffix _test.go 

* Functions and variables
Use camel case, exported functions and variables start with uppercase
Constant should use all capital letters and use underscore _ to separate words.
If variable type is bool, its name should start with Has, Is, Can or Allow, etc.
Avoid cryptic abbreviations. Eg: usrAge := 25

## Project Structure (Big Project)
```
my_project/
├── cmd/
│   ├── api/
│   │   └── main.go           # Entry point cho API server
│   ├── worker/
│   │   └── main.go           # Entry point cho worker (nếu có)
│   └── cli/
│       └── main.go           # Entry point cho CLI tool (nếu có)
├── config/
│   ├── database               # Config cho App Engine
│   ├── config.go               
│   └──constant/
├── internal/                 # Không thể truy cập từ bên ngoài
│   ├── handler/              # HTTP handlers (giao tiếp với các adapter bên ngoài)
│   │   ├── user_handler.go
│   │   └── auth_handler.go
│   ├── middleware/           # Middleware (VD: logging, auth, etc.)
│   │   └── auth_middleware.go
│   ├── repository/           # Lớp repository (tương tác với DB hoặc các nguồn dữ liệu)
│   │   ├── user_repository.go
│   │   └── auth_repository.go
│   ├── usecase/              # Business logic (lớp use case)
│   │   ├── user_usecase.go
│   │   └── auth_usecase.go
│   └──  domain/               # Entities và các giá trị liên quan đến domain
│       ├── user.go
│       └── auth.go
├── utils/                      # Thư viện dùng chung giữa các phần (VD: logging, utils)
├── docs/                     # API documentation (VD: Swagger files)
│   └── swagger.json
├── .github/
│   └── workflows/
│       └── ci.yml            # CI/CD workflow file (GitHub Actions hoặc tương tự)
├── go.mod                    # File khai báo dependencies của Go
├── go.sum                    # File lock dependencies
├── Dockerfile                # Dockerfile cho dự án
├── docker-compose.yml        # Docker Compose (nếu cần chạy nhiều services)
├── Makefile                  # Các script tiện ích (build, test, run)
├── .env                      # File môi trường (environment variables)
├── README.md                 # Tài liệu dự án
└── tests/                    # Thư mục chứa test cases
    ├── user_handler_test.go
    └── auth_usecase_test.go
```
- `cmd`: entry points for different app types (e.g., API, worker, CLI).
- `config`: all config files is here
- `internal`: core application logic, inaccessible externally. (cannot import internal/handler: use of internal package not allowed)
- `handler`: handles HTTP requests and routes.
- `middleware`: reusable logic for requests (e.g., auth, logging).
- `repository`: interacts with the database or external data sources.
- `usecase`: business logic implementation.
- `domain`: defines core entities.
- `logger`: custom logging utilities.
- `utils`: reusable shared packages.
- `response`: standardized client responses.
- `docs`: API documentation (e.g., Swagger).
- `.github`: CI/CD workflows for automated processes.
- `tests`: unit and integration test cases.
Key project files:
go.mod: manages Go module dependencies.
Dockerfile: builds the Docker image.
.env: stores environment variables.
Makefile: scripts for build/test/deploy tasks.

sql-migrate up -config=config/database/dbconfig.yml -env=development
sql-migrate down -config=config/database/dbconfig.yml -env=development

sql-migrate new -config=config/database/dbconfig.yml -env=development "migration_name"

go mod init <module_name>
go get -u github.com/gin-gonic/gin
go get -u gorm.io/driver/postgres
go get github.com/joho/godotenv
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go install github.com/swaggo/swag/cmd/swag@latest
--> swag init 
go get github.com/gin-contrib/cors
go install github.com/rubenv/sql-migrate/...@latest
go get -u github.com/golang-jwt/jwt/v5

excel
go get github.com/xuri/excelize/v2
go get github.com/stretchr/testify/mock

hot reload
go install github.com/mitranim/gow@latest
go install github.com/air-verse/air@latest


go get github.com/DATA-DOG/go-sqlmock
go install go.uber.org/mock/mockgen@latest
go install github.com/golang/mock/mockgen@v1.6.0

schedule
go get github.com/go-co-op/gocron/v2

go get
Tải module và chỉnh sửa go.mod, go.sum.
go install
Tạo file thực thi và lưu vào $GOPATH/bin hoặc $GOBIN.

http://localhost:4000/swagger/index.html

go test ./... -cover
go test ./...
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage