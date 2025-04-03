package examples

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func StringExamples(client *redis.Client, ctx *context.Context) {
	// Set
	client.Set(*ctx, "string_set", "example-value", 0)

	string_set, _ := client.Get(*ctx, "string_set").Result()
	fmt.Println("string_set : ", string_set)

	// Time to live
	client.Set(*ctx, "string_ttl", "string_ttl", 2*time.Second)
	// time.Sleep(3 * time.Second)
	string_ttl, err := client.Get(*ctx, "string_ttl").Result()
	if err != nil {
		fmt.Println("Error getting value of string_ttl")
	} else {
		fmt.Println("string_ttl : ", string_ttl)
	}

	// Incr and IncrBy
	client.Incr(*ctx, "incr_key")
	client.IncrBy(*ctx, "incr_key", 8)

	incr_key, err := client.Get(*ctx, "incr_key").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("incr_key: ", incr_key)
	}

	// SetArgs
	set_args, err := client.SetArgs(*ctx, "incr_key", 0, redis.SetArgs{Get: true}).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("set_args: ", set_args)
	aftr_setargs, err := client.Get(*ctx, "incr_key").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("aftr_setargs : ", aftr_setargs)

	// GetRange
	get_range, err := client.GetRange(*ctx, "string_set", 0, 4).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("get_range : ", get_range)
}
