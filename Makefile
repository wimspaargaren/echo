
gen:
	@oapi-codegen -generate server -package api api/spec.yml  > api/server.go
	@oapi-codegen -generate client,types -package api api/spec.yml  > api/client.go

lint:
	@echo Linting...
	@golangci-lint  -v --concurrency=3 --config=.golangci.yml --issues-exit-code=1 run \
	--out-format=colored-line-number 
