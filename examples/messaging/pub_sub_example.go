package messaging

import (
	"caching/internal/redis"
	"context"
	"fmt"
)

const PUB_SUB_CHANNEL = "pub_sub_channel"

func PubSubExample() {
	ctx := context.Background()
	client, err := redis.New("localhost:6379")
	if err != nil {
		fmt.Println("Error instantiating client")
		return
	}

	_, err = client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error in connection", err)
	}
}
