package examples

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type User struct {
	Name       string `redis:"name"`
	Profession string `redis:"profession"`
}

func HashExample(client *redis.Client, ctx *context.Context) {
	client.HSet(*ctx, "hash:1", "name", "John wick", "profession", "assacination")
	client.HSet(*ctx, "hash:2", []string{"name", "Naruto", "profession", "ninja"})
	client.HSet(*ctx, "hash:3", map[string]string{"name": "Tanjero", "profession": "Demon Slayer"})

	// returns only a string
	name_1, _ := client.HGet(*ctx, "hash:1", "name").Result()
	// returns array of string
	hash_2, _ := client.HMGet(*ctx, "hash:1", "name", "profession").Result()
	// return map[string]string
	hash_3, _ := client.HGetAll(*ctx, "hash:3").Result()

	fmt.Println(name_1)
	fmt.Println(hash_2)
	fmt.Println(hash_3)

	// Scan
	var user User
	client.HGetAll(*ctx, "hash:3").Scan(&user)
	fmt.Println(user.Name)
	fmt.Println(user.Profession)

	// get all keys
	hkeys, _ := client.HKeys(*ctx, "hash:1").Result()
	fmt.Println("hkeys : ", hkeys)
}
