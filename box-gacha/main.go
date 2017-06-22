package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/katsew/kuji"
	"github.com/katsew/kuji-redis"
)

func main() {
	opt := redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	}
	kr := kuji_redis.NewShuffleStrategy(&opt)
	config := kuji.KujiStrategyConfig{
		Strategy: kr,
		FailOver: nil,
	}
	k := kuji.NewKuji()
	k.Use("suffle", config)

	c := []kuji.KujiCandidate{
		kuji.KujiCandidate{
			Id:     1,
			Weight: 1,
		},
		kuji.KujiCandidate{
			Id:     2,
			Weight: 1,
		},
		kuji.KujiCandidate{
			Id:     3,
			Weight: 1,
		},
		kuji.KujiCandidate{
			Id:     4,
			Weight: 1,
		},
		kuji.KujiCandidate{
			Id:     5,
			Weight: 1,
		},
		kuji.KujiCandidate{
			Id:     6,
			Weight: 1,
		},
		kuji.KujiCandidate{
			Id:     7,
			Weight: 1,
		},
		kuji.KujiCandidate{
			Id:     8,
			Weight: 1,
		},
		kuji.KujiCandidate{
			Id:     9,
			Weight: 1,
		},
		kuji.KujiCandidate{
			Id:     10,
			Weight: 1,
		},
		kuji.KujiCandidate{
			Id:     11,
			Weight: 1,
		},
		kuji.KujiCandidate{
			Id:     12,
			Weight: 1,
		},
	}

	_, err := k.RegisterCandidatesWithKey("suffle", "box_gacha", c)
	if err != nil {
		panic(err)
	}

	PickAndDeleteOne(k)

}

func PickAndDeleteOne(k kuji.Kuji) {
	for i := 0; i < 1000; i++ {
		str, err := k.PickAndDeleteOneByKey("suffle", "box_gacha")
		if err != nil {
			fmt.Println("No data stored in redis!")
			break
		}
		fmt.Printf("Picked card ID is %s and this card no longer exists in data store.\n", str)
	}
}
