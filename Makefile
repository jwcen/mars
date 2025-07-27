.PHONY: mock
mock:
	@mockgen -source=./internal/apiserver/service/user.go -package=svcmocks -destination=./internal/apiserver/service/mocks/user.mock.go
	@mockgen -source=./internal/apiserver/repository/user.go -package=svcmocks -destination=./internal/apiserver/repository/mocks/user.mock.go
	@mockgen -source=./internal/apiserver/repository/dao/user.go -package=svcmocks -destination=./internal/apiserver/repository/dao/mocks/user.mock.go
	@go mod tidy