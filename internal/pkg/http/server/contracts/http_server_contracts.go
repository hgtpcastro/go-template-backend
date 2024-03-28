package contracts

import "context"

type GenericHttpServer[T any] interface {
	Run(config ...func(s *T)) error
	Shutdown(ctx context.Context) error
}
