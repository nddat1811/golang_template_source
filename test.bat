go test ./controller/... ./repository/... ./usecase/... -coverprofile=coverage.out
go tool cover -html coverage.out