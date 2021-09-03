# ova_route_api

protoc --proto_path=. -I vendor.protogen \
	--go_out=pkg/api --go_opt=paths=import \
	--go-grpc_out=pkg/api --go-grpc_opt=paths=import \
	--grpc-gateway_out=pkg/api \
	--grpc-gateway_opt=logtostderr=true \
	--grpc-gateway_opt=paths=import \
	--swagger_out=allow_merge=true,merge_file_name=api:swagger \
	api/api.proto



    protoc --proto_path=. \
	--go_out=pkg/api --go_opt=paths=import \
	--go-grpc_out=pkg/api --go-grpc_opt=paths=import \
	api/api.proto



	1) inotify через Viper
	2) graceful shutdown
	3) улучшить работу с логами
	4) работать с контекстом
	5) трасировка
	6) счетчики\промет
	7) хелсчек\готовность\redy (кубер)
	8) пагинация в запросах
	9) интеграционные тесты с поднятием сервиса в контейнере