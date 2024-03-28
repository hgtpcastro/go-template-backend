package echo

import (
	"context"

	"github.com/hgtpcastro/go-template-backend/internal/pkg/http/server/contracts"
	"github.com/labstack/echo"
)

type httpServerEcho struct {
	echo *echo.Echo
}

func NewEchoHttpServer() contracts.GenericHttpServer[echo.Echo] {
	return &httpServerEcho{
		echo: echo.New(),
	}
}

func (s *httpServerEcho) Start(config ...func(s *echo.Echo)) error {
	return nil
}

func (s *httpServerEcho) Shutdown(ctx context.Context) error {
	return nil
}
