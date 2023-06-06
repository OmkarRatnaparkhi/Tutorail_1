package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type SwitchAccountModel struct {
	ParentId   string `json:"ParentId`
	ChildId    string `json:"ChildId"`
	RecordType int    `json:"RecordType"`
}

func AddModelInRedis(NewUser SwitchAccountModel) {

	objectData := SwitchAccountModel{}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	//obj := []SwitchAccountModel{}

	redisgetdata, err := rdb.HGet(ctx, "Switch_Account", NewUser.ParentId+"_SA").Result()
	if err != nil {

		data, err := json.Marshal(NewUser)
		if err != nil {
			return
		}

		err = rdb.HSet(ctx, "Switch_Account", NewUser.ParentId+"_SA", data).Err()
		if err != nil {
			fmt.Println(err)
		}
		return
	} else {
		json.Unmarshal([]byte(redisgetdata), &objectData)
		// if NewUser.ParentId == objectData.ParentId {

		// 	append(redisgetdata,)

		// 	data, err := json.Marshal(NewUser)

		// 	err = rdb.HSet(ctx, "Switch_Account", NewUser.ParentId+"_SA", data).Err()
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}
		// }
	}
	fmt.Print(redisgetdata)
}

func main() {
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr: "127.0.0.1:6379",
	// 	Password: "",
	// 	DB: 0,
	// })

	/*
		// Setting and getting data in redis using Primitive data type (String)
		err := rdb.Set(ctx, "key", "value", 0).Err()
		if err != nil {
			panic(err)
		}

		val, err := rdb.Get(ctx, "key").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("key", val)

		val2, err := rdb.Get(ctx, "key").Result()
		if err == redis.Nil {
			fmt.Println("key2 does not exist")
		} else if err != nil {
			panic(err)
		} else {
			fmt.Println("key2",val2)
		}
	*/

	/*
		err := rdb.Set(ctx, "name", "Omkar Ratnaparkhi",0).Err()
		if err != nil {
			panic(err)
		}

		val, err := rdb.Get(ctx, "name").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("name",val)

		err2 := rdb.Set(ctx, "name1", "Kedar Mali",0).Err()
		if err != nil {
			panic(err2)
		}

		val2, err2 := rdb.Get(ctx, "name1").Result()
		if err2 == redis.Nil{
			fmt.Println("name1 does not exist")
		} else if err2 != nil {
			panic(err2)
		} else {
			fmt.Println("name1", val2)
		}
	*/

	iobj := SwitchAccountModel{"Omkar", "Omkar", 1}
	AddModelInRedis(iobj)
	iobj1 := SwitchAccountModel{"Omkar", "Rahul", 2}
	AddModelInRedis(iobj1)

}
