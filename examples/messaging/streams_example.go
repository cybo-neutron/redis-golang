package messaging

import (
	"caching/internal/redis"
	"context"
	"fmt"
	"log"
	"time"

	go_redis "github.com/redis/go-redis/v9"
)

type Stream struct {
	ctx context.Context
}

const (
	STREAM_NAME = "test-stream"
)

func New() *Stream {
	return &Stream{
		context.Background(),
	}
}

func (s *Stream) StreamsExample() {
	client, err := redis.New("localhost:6379")
	if err != nil {
		log.Fatal("Error instantiating redis client ", err)
	}

	s.addToStream(client)
	s.addToStream(client)
	s.addToStream(client)
	s.addToStream(client)
	s.addToStream(client)
	s.addToStream(client)

	go s.handleRead(client)
}

func (s *Stream) addToStream(client *go_redis.Client) {
	id, err := client.XAdd(s.ctx, &go_redis.XAddArgs{Stream: STREAM_NAME, Values: []string{"key-1", "value-1"}}).Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Added to stream : ", id)
}

func (s *Stream) handleRead(client *go_redis.Client) {
	streamItems := client.XRead(s.ctx, &go_redis.XReadArgs{[]string{"test-stream", "$"}, 10, 1 * time.Second, "1"}).Val()

	for _, streamItem := range streamItems {
		for _, item := range streamItem.Messages {
			fmt.Println(item.ID, item.Values)
		}
	}
}
