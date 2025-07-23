package file

import (
	"fmt"
	"github.com/lie-flat-planet/gen-project/global"
	"strings"
)

type DockerfileDev struct{}

func (d *DockerfileDev) Name() string {
	return "Dockerfile_dev"
}

func (d *DockerfileDev) Content() string {
	projectName := global.GetProjectName()

	return strings.TrimSpace(fmt.Sprintf(`
FROM golang:1.23 AS builder

WORKDIR /go/src
COPY ./ ./

# build
RUN make build WORKSPACE=%s

# runtime
FROM alpine:latest

ARG PROJECT_NAME=%s

COPY --from=builder /go/src/cmd/${PROJECT_NAME}/${PROJECT_NAME} /go/bin/${PROJECT_NAME}

EXPOSE 80

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/%s"]
`, projectName, projectName, projectName))
}
