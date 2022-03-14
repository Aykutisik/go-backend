generate-mocks:
	mockgen -source repository/repository.go -package repository -destination repository/mock_repository.go
	mockgen -source handler/handler.go -package handler -destination handler/mock_handler.go
	mockgen -source service/service.go -package service -destination service/mock_service.go