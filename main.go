package main

import (
	"flag"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"os"
)

const (
	redisHost  = "REDIS_HOST"
	redisPort  = "REDIS_PORT"
	redisKey   = "REDIS_KEY"
	redisValue = "REDIS_VALUE"
)

func main() {
	var h = flag.String("hostname", os.Getenv(redisHost), "Set Hostname")
	var p = flag.String("port", os.Getenv(redisPort), "Set Port")
	var key = flag.String("key", os.Getenv(redisKey), "Set Key")
	var val = flag.String("value", os.Getenv(redisValue), "Set Value")
	flag.Parse()

	c, err := redis.Dial("tcp", *h + ":" + *p)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer c.Close()

	c.Do("SET", *key, *val)
	s, err := redis.String(c.Do("GET", *key))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("GET %s\n", s)
}
