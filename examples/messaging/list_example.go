package messaging

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/redis/go-redis/v9"
)

/*
* LPush :
* RPush
* LLen
* LRange
* LPop
* LTrim
*
 */

/*
* -------- Blocking list operations -------
* BRPop
* BLPop
*
 */

func MessagingExample(client *redis.Client, ctx *context.Context) {
	client.Del(*ctx, "list_1")
	client.LPush(*ctx, "list_1", "item2", "item1")
	client.RPush(*ctx, "list_1", "item3", "item4")

	l, _ := client.LLen(*ctx, "list_1").Result()
	fmt.Println("Length of list: ", l)

	rn, _ := client.LRange(*ctx, "list_1", 0, math.MaxInt16).Result()
	fmt.Println("Range list : ", rn)

	client.LPop(*ctx, "list_1")
	client.LTrim(*ctx, "list_1", 0, 1)
	lis, _ := client.LRange(*ctx, "list_1", 0, math.MaxInt16).Result()

	fmt.Println("After trim : ", lis)

	// ----- Blocking operations ------

	listName := "list_name"
	client.LPush(*ctx, listName, "item1", "item2", "item3")
	val, _ := client.BRPop(*ctx, time.Second, listName).Result()
	fmt.Println("hh :", val)
}
