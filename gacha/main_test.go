package main

import (
	"testing"
	"github.com/katsew/kuji-redis"
	"github.com/katsew/kuji"
	"github.com/go-redis/redis"
	"fmt"
)

func BenchmarkPickOne(b *testing.B) {
	opt := redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	}
	kr := kuji_redis.SimpleStrategy(&opt)
	instance := kuji.NewKuji(kr)

	c := []kuji.KujiCandidate{
		kuji.KujiCandidate{
			Id: 1,
			Weight: 5,
		},
		kuji.KujiCandidate{
			Id: 2,
			Weight: 15,
		},
		kuji.KujiCandidate{
			Id: 3,
			Weight: 80,
		},
	}

	_, err := instance.RegisterCandidatesWithKey("simple", c)
	if (err != nil) {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		num, err := instance.PickOneByKey("simple")
		if err != nil {
			panic(err)
		}
		fmt.Println("Pickup string is", num)
	}
}