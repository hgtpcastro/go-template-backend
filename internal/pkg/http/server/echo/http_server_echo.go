package echo

import (
	"context"

	server "github.com/hgtpcastro/go-template-backend/internal/pkg/http/server/contracts"
	"github.com/labstack/echo"
)

type httpServerEcho struct {
	echo *echo.Echo
}

func NewEchoHttpServer() server.GenericHttpServer[echo.Echo] {
	return &httpServerEcho{
		echo: echo.New(),
	}
}

func (s *httpServerEcho) Run(config ...func(s *echo.Echo)) error {
	return nil
}

func (s *httpServerEcho) Shutdown(ctx context.Context) error {
	return nil
}
