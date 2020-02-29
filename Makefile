#---* Makefile *---#



## docker-build		: 	build docker container
docker-build:
	@docker build -t qorpress/qorpress-auth-example --no-cache .
.PHONY: docker-build

## docker-run		: 	run web interface.
docker-run:
	@docker run -ti -p 4000:4000 -v $(PWD)/.config/gopress.yml:/opt/qor/.config/gopress.yml qorpress/qorpress-auth-example
.PHONY: docker-run

## deps			: 	install dependencies
deps:
	@GO111MODULE=off go get -u -f github.com/qor/bindatafs/...
	@go mod vendor
.PHONY: deps

run_bindatafs:
	@go run config/compile/compile.go
	@go run -tags bindatafs main.go
.PHONY: run_bindatafs

build: deps
	@go build main.go
.PHONY: build

build_bindatafs: dep
	@go run config/compile/compile.go
	@go build -tags bindatafs main.go
.PHONY: build_bindatafs

## help			: 	Print commands help.
.PHONY: help
help : Makefile
	@sed -n 's/^##//p' $<

# https://stackoverflow.com/a/6273809/1826109
%:
	@: