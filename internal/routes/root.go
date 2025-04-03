package routes

import (
	"caching/internal/redis"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func TestCachingRoutes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method : ", r.Method)
	jsonData, err := json.Marshal(map[string]string{
		"message": "success",
	})
	if err != nil {
		// w.Write([]byte(map[string]string{
		// 	message: "failed",
		// }))
		w.Write([]byte("Failed"))
	}
	w.Write(jsonData)
}

func AddPlayer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		type User struct {
			Name string `json:"name"`
		}

		var user User
		err := decoder.Decode(&user)
		if err != nil {
			fmt.Println("Failed")
			errorJsonData, err := json.Marshal(map[string]string{
				"message": "failed",
			})
			if err != nil {
				w.Write([]byte("Failed"))
			}
			w.Write(errorJsonData)
		}

		userJson, _ := json.Marshal(user)
		fmt.Println("user : ", string(userJson))

		client, err := redis.New("localhost:6379")
		if err != nil {
			fmt.Println("Client creating error", err)
		}

		ctx := context.Background()
		client.SAdd(ctx, "members", user.Name)
		fmt.Println("user added to redis set")

		successJsonData, err := json.Marshal(map[string]string{
			"message": "success",
		})
		if err != nil {
			w.Write([]byte("Failed"))
		}
		w.Write(successJsonData)
	}
}

func EstablishRoutes() {
	http.HandleFunc("/test-cache-route", TestCachingRoutes)
	http.HandleFunc("/add-player", AddPlayer)

	log.Fatal(http.ListenAndServe(":5005", nil))
}
