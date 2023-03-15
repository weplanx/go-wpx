package sessions

import (
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"github.com/weplanx/utils/values"
)

var Provides = wire.NewSet(
	wire.Struct(new(Controller), "*"),
)

func New(options ...Option) *Service {
	x := new(Service)
	for _, v := range options {
		v(x)
	}
	return x
}

type Option func(x *Service)

func SetNamespace(v string) Option {
	return func(x *Service) {
		x.Namespace = v
	}
}

func SetRedis(v *redis.Client) Option {
	return func(x *Service) {
		x.Redis = v
	}
}

func SetDynamicValues(v *values.DynamicValues) Option {
	return func(x *Service) {
		x.DynamicValues = v
	}
}
