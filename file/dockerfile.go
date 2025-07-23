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
FROM debian:stable-slim

EXPOSE 80

WORKDIR /app
COPY ./bin/%s /app/%s

ENTRYPOINT ["/app/%s"]
`, projectName, projectName, projectName))
}
