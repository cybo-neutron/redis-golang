package messaging

import (
	"caching/internal/redis"
	"context"
	"fmt"
	"time"

	go_redis "github.com/redis/go-redis/v9"
)

const PUB_SUB_CHANNEL = "pub_sub_channel"

var ctx = context.Background()

func InitSubcriber(client *go_redis.Client) {
	fmt.Println("Initializing subscriber")
	subscriber := client.Subscribe(ctx, PUB_SUB_CHANNEL)
	defer subscriber.Close()

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			fmt.Println("Error while receiving message : ", err)
			panic(err)
		}
		fmt.Println("Received message : ", msg.Channel, msg.Payload)
	}
}

func PubSubExample() {
	client, err := redis.New("localhost:6379")
	if err != nil {
		fmt.Println("Error instantiating client")
		return
	}

	_, err = client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error in connection", err)
		return
	}

	go InitSubcriber(client)
	time.Sleep(10 * time.Second)
	client.Publish(ctx, PUB_SUB_CHANNEL, "test-message")
}
