package file

import (
	"fmt"
	"github.com/lie-flat-planet/gen-project/global"
	"strings"
)

type Dockerfile struct{}

func (d *Dockerfile) Name() string {
	return "Dockerfile"
}

func (d *Dockerfile) Content() string {
	projectName := global.GetProjectName()

	return strings.TrimSpace(fmt.Sprintf(`
FROM golang:1.23 AS builder

WORKDIR /go/src
COPY ./ ./

# build
RUN make build WORKSPACE=%s

# runtime
FROM alpine:latest

ARG PROJECT_NAME=test-gen-project

COPY --from=builder /go/src/cmd/${PROJECT_NAME}/${PROJECT_NAME} /go/bin/${PROJECT_NAME}

EXPOSE 80

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/%s"]
`, projectName, projectName))
}
