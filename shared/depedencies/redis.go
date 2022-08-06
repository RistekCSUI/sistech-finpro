package depedencies

import (
	"fmt"
	"github.com/RistekCSUI/sistech-finpro/shared/config"
	"github.com/go-redis/redis/v9"
)

func NewRedis(env *config.EnvConfig) *redis.Client {
	redisUrl := fmt.Sprintf(
		"%s:%s",
		env.RedisHost,
		env.RedisPort,
	)

	return redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "",
		DB:       0,
	})
}
