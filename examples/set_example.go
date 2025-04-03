package examples

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

/*
* -------- Set -----------
* SAdd : add items to set
* SMembers : returns members of a set
* SIsMember : return if a set has item
* SCard : get count of items in a set
* SRem : remove item from a set
* SUnion : union of 2 set
 */

/*
* -------- Sorted Set -------
* ZAdd : add item to set
* ZScore : get the score associated with a member of a sorted set
* ZRank :
* ZIncrBy : increase the score of a member a set
* ZRevRangeByScore : get the members within a specific range of the score
* ZRem : remove member from a set
 */

func SetExample(client *redis.Client, ctx *context.Context) {
	// NormalSetExample(client, ctx)
	SortedSetExample(client, ctx)
}

func SortedSetExample(client *redis.Client, ctx *context.Context) {
	// client.ZAdd(*ctx, "ss_1", []*redis.Z{{Member: "ss_1", Score: 1}, {Member: "ss_2", Score: 2}})
	client.ZAdd(*ctx, "ss_1", redis.Z{Member: "ss_1", Score: 1}, redis.Z{Member: "ss_2", Score: 2})

	score, _ := client.ZScore(*ctx, "ss_1", "ss_1").Result()
	fmt.Println("Score of ss_1", score)

	rank, _ := client.ZRank(*ctx, "ss_1", "ss_2").Result()
	fmt.Println("Rank : ", rank)

	client.ZIncrBy(*ctx, "ss_1", 5, "ss_2")

	rangeResult, _ := client.ZRange(*ctx, "ss_1", 0, 7).Result()
	fmt.Println("Range result: ", rangeResult)
}

func NormalSetExample(client *redis.Client, ctx *context.Context) {
	client.SAdd(*ctx, "set_1", []string{"item1", "item2", "item3", "item4"})

	set_1, _ := client.SMembers(*ctx, "set_1").Result()
	fmt.Println("Set example : ", set_1)

	s_isMember, _ := client.SIsMember(*ctx, "set_1", "item3").Result()
	fmt.Println("Is Member : ", s_isMember)

	s_count, _ := client.SCard(*ctx, "set_1").Result()
	fmt.Println("Set count  : ", s_count)

	client.SRem(*ctx, "set_1", []string{"item1", "item3"})
	set_1, _ = client.SMembers(*ctx, "set_1").Result()
	fmt.Println("set after removal: ", set_1)

	client.SAdd(*ctx, "set_2", []string{"i1", "i2", "i3"})
	// union
	unionResult, _ := client.SUnion(*ctx, "set_1", "set_2").Result()
	fmt.Println("union result : ", unionResult)

	// intersect
	// client.SInterStore(ctx context.Context, destination string, keys ...string)
	// client.SInter(ctx context.Context, keys ...string)
}
