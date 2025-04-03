package messaging

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

/*
* BRPop : Waits
 */

func publish(client *redis.Client, ctx *context.Context) {
	client.LPush(*ctx, "list_pub", "val1")
	time.Sleep(10 * time.Second)
	client.LPush(*ctx, "list_pub", "val2")
}

func consume(client *redis.Client, ctx *context.Context) {
	val, err := client.BRPop(*ctx, 1*time.Second, "list_pub").Result()
	if err != nil {
		fmt.Println("Error consuming 1 : ", err)
	} else {
		fmt.Println("consumed value : ", val)
	}

	val, err = client.BRPop(*ctx, 0, "list_pub").Result()
	if err != nil {
		fmt.Println("Error consuming 2 : ", err)
	} else {
		fmt.Println("consumed value : ", val)
	}

	val, err = client.BRPop(*ctx, 10*time.Second, "list_pub").Result()
	if err != nil {
		fmt.Println("Error consuming 3 : ", err)
	} else {
		fmt.Println("consumed value : ", val)
	}
}

func consume2(client *redis.Client, ctx *context.Context) {
	for {
		val, err := client.BRPop(*ctx, 0, "list_pub").Result()
		if err != nil {
			fmt.Println("Error consuming ", err)
		} else {
			fmt.Println("Consumed value : ", val)
		}
	}
}

func ListPubSubExample(client *redis.Client, ctx *context.Context) {
	go publish(client, ctx)
	// go consume(client, ctx)
	go consume2(client, ctx)
}
