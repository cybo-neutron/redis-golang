package main

import (
	"caching/examples/core_features"
	"caching/internal/routes"
)

func main() {
	// ctx := context.Background()
	//
	// client, err := redis.New("localhost:6379")
	// if err != nil {
	// 	fmt.Println("Error establishing connection", err)
	// }

	// examples.StringExamples(client, &ctx)
	// examples.HashExample(client, &ctx)
	// examples.SetExample(client, &ctx)
	// messaging.MessagingExample(client, &ctx)
	// messaging.ListPubSubExample(client, &ctx)
	// messaging.PubSubExample()
	// messaging.New().StreamsExample()
	core_features.PipelineExample()

	// challenges.StringHashChallenge(*client, &ctx)

	routes.EstablishRoutes()
}
