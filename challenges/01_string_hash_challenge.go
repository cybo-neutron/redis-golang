package challenges

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func StringHashChallenge(client redis.Client, ctx *context.Context) {
	// challenge : create a redis hash with fields 'name','email','city'. Then delete 'city' from the hash
	client.HSet(*ctx, "sh:1", map[string]string{"name": "John Wick", "email": "baba_yaga@gmail.com", "city": "New York"})

	client.HDel(*ctx, "sh:1", "city")

	sh_1, _ := client.HGetAll(*ctx, "sh:1").Result()
	fmt.Println("sh_1 : ", sh_1)
}
