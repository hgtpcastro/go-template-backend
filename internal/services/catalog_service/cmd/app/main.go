package main

import (
	"github.com/hgtpcastro/go-template-backend/internal/pkg/http/server/echo"
)

func main() {
	s := echo.NewEchoHttpServer()
	s.Start()
}
