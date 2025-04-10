package core_features

import (
	"caching/internal/redis"
	"context"
	"fmt"
	"strconv"
	"time"
)

var ctx = context.Background()

func no_pipeline() {
	client, err := redis.New("localhost:6379")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	startTime := time.Now()

	for i := 0; i < 100; i++ {
		client.Set(ctx, "key-"+strconv.Itoa(i), "value-"+strconv.Itoa(i), 2*time.Minute)
	}

	fmt.Println("Time taken for no_pipeline : ", time.Since(startTime))
}

func pipeline() {
	client, err := redis.New("localhost:6379")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	startTime := time.Now()

	pipeline := client.Pipeline()

	for i := 0; i < 100; i++ {
		pipeline.Set(ctx, "pkey-"+strconv.Itoa(i), "pvalue-"+strconv.Itoa(i), 2*time.Minute)
	}

	_, err = pipeline.Exec(ctx)
	if err != nil {
		fmt.Println("Error executing pipeline")
		// return
	}

	fmt.Println("Time taken for pipeline : ", time.Since(startTime))
}

func PipelineExample() {
	no_pipeline()
	pipeline()
}
