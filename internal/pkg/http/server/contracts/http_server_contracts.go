package contracts

import "context"

type GenericHttpServer[T any] interface {
	Start(config ...func(s *T)) error
	Shutdown(ctx context.Context) error
}
