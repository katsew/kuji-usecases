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
	kr := kuji_redis.NewSimpleStrategy(&opt)
	k := kuji.NewKuji()
	k.Use("redis", kuji.KujiStrategyConfig{
		Strategy: kr,
	})

	c := []kuji.KujiCandidate{
		kuji.KujiCandidate{
			Id:     1,
			Weight: 5,
		},
		kuji.KujiCandidate{
			Id:     2,
			Weight: 15,
		},
		kuji.KujiCandidate{
			Id:     3,
			Weight: 80,
		},
	}

	_, err := k.RegisterCandidatesWithKey("redis", "gacha", c)
	if err != nil {
		panic(err)
	}

	PickOne(k)

}

func PickOne(k kuji.Kuji) {
	for i := 0; i < 1000; i++ {
		str, err := k.PickOneByKey("redis", "gacha")
		if err != nil {
			panic(err)
		}
		fmt.Printf("Picked card ID is %s \n", str)
	}
}
