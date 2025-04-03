package main

import (
	"caching/examples/messaging"
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
	messaging.PubSubExample()

	// challenges.StringHashChallenge(*client, &ctx)

	routes.EstablishRoutes()
}
