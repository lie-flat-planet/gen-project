package dir

import "fmt"

func init() {
	Dirs = append(Dirs, NewMakefile())
}

type Makefile struct {
	projectName string
}

func NewMakefile() *Makefile {
	return new(Makefile)
}

func (m *Makefile) RootPath() string {
	return ""
}

func (m *Makefile) Path() string {
	return ""
}

func (m *Makefile) File() string {
	return "makefile"
}

func (m *Makefile) FileContent() string {
	return fmt.Sprintf(`PKG = $(shell cat go.mod | grep "^module " | sed -e "s/module //g")
NAME = $(shell basename $(PKG))
VERSION = v$(shell cat .version)
COMMIT_SHA ?= $(shell git rev-parse --short HEAD)

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
CGO_ENABLED ?= 0

GO = CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go

GOBUILD=CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=amd64 go build -a -ldflags "-X ${PKG}/version.Version=${VERSION}+sha.${COMMIT_SHA}"

WORKSPACE ?= $(NAME)
DOCKERBUILD = docker build -t $(NAME) .
GORUN = $(GO) run

build:
	cd ./cmd/$(WORKSPACE) && $(GOBUILD)

docker:
	$(DOCKERBUILD)

save:
	docker save -o $(NAME).tar $(NAME):latest

up:
	cd ./cmd/$(WORKSPACE) && $(GORUN) .

run: download openapi
	cd ./cmd/$(WORKSPACE) && $(GORUN) .

download:
	go mod download -x

# 生成服务端

# 生成swagger
swagger:
	swag init --pd -d ./cmd/$(WORKSPACE) -o ./cmd/$(WORKSPACE)/docs

# 生成openapi
openapi: swagger
	cd ./cmd/$(WORKSPACE) && klctl openapi  -f ./docs/swagger.json

# 生成client
client-swagger:
	swag init --pd -d ./cmd/$(NAME) -o ./cmd/$(NAME)/docs

# 生成client openapi
client-openapi: client-swagger
	cd ./cmd/$(NAME) && klctl openapi  -f ./docs/swagger.json

go-list:
	go list -m -json all

migrate:
	 cd ./cmd/$(WORKSPACE) && $(GORUN) . migrate
data-migrate: # make data-migrate AppID='ae9f7f36-9802-4ffc-a218-3d7bf7ed7631' StartTime='2023-11-7 0:0:0'
	cd ./cmd/$(WORKSPACE) && $(GORUN) . data-migrate $(AppID) $(StartTime)
`)
}

func (m *Makefile) SetProjectName(name string) {
	m.projectName = name
}
