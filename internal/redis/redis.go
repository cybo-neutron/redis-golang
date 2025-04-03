package redis

import (
	"github.com/redis/go-redis/v9"
)

type RedisOptions struct {
	Addr string
}

func New(url string) (*redis.Client, error) {
	// parsedUrl, err := redis.ParseURL(url)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to parse url %v", url)
	// }
	opts := &redis.Options{
		// Addr: parsedUrl.Addr,
		Addr: url,
	}
	client := redis.NewClient(opts)

	return client, nil
}
