package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func InitRedis(addr string, password string) *redis.Pool {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				fmt.Println("Redis dial tcp failed.")
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					fmt.Println("Redis AUTH failed.")
					c.Close()
					return nil, err
				}
			}
			return c, nil
		},
	}

	conn := pool.Get()
	if conn == nil {
		fmt.Println("Redis get conn from pool failed.")
		return nil
	}

	if _, err := conn.Do("PING"); err != nil {
		fmt.Println("redis PING failed.")
		return nil
	}
	return pool
}

func KEYS(conn redis.Conn, key string) []string {

	reply, err := conn.Do("KEYS", key)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}

	itemsT := reply.([]interface{})
	var items []string
	for _, v := range itemsT {
		items = append(items, string(v.([]byte)))
	}
	return items
}

func SINTER(conn redis.Conn, key string) []string {

	reply, err := conn.Do("SINTER", key)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}

	itemsT := reply.([]interface{})
	var items []string
	for _, v := range itemsT {
		items = append(items, string(v.([]byte)))
	}
	return items
}

func SREM(conn redis.Conn, key string, value string) bool {

	reply, err := conn.Do("SREM", key, value)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if reply.(int64) == 1 {
		return true
	} else {
		return false
	}
}

func HGET(conn redis.Conn, key string, field string) map[string]string {

	reply, err := conn.Do("HGET", key, field)
	if err != nil {
		fmt.Println(err)
		return map[string]string{}
	}

	itemsT := reply.([]interface{})
	if len(itemsT)%2 != 0 {
		return map[string]string{}
	}

	m := make(map[string]string)
	for i, _ := range itemsT {
		if i%2 == 0 {
			m[string(itemsT[i].([]byte))] = string(itemsT[i+1].([]byte))
		}
	}
	return m
}

func HDEL(conn redis.Conn, key string, v1, v2, v3, v4 string) bool {

	reply, err := conn.Do("HDEL", key, v1, v2, v3, v4)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if 4 == reply.(int64) {
		return true
	} else {
		return false
	}
}

func SET(conn redis.Conn, key string, value int) bool {

	reply, err := conn.Do("SET", key, value)
	if err != nil {
		fmt.Println(err)
		return false
	}

	ok := reply.(string)
	if ok == "OK" {
		return true
	} else {
		return false
	}
}

func GET(conn redis.Conn, key string) []uint8 {

	reply, err := conn.Do("GET", key)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if reply != nil {
		ok := reply.([]uint8)
		return ok
	}
	return nil
}

func INCR(conn redis.Conn, key string) error {

	_, err := conn.Do("INCR", key)
	if err != nil {
		return err
	}
	return nil
}

func DECR(conn redis.Conn, key string) error {

	_, err := conn.Do("DECR", key)
	if err != nil {
		return err
	}
	return nil
}

func HGETALL(conn redis.Conn, key string) map[string]string {

	reply, err := conn.Do("HGETALL", key)
	if err != nil {
		fmt.Println(err)
		return map[string]string{}
	}

	itemsT := reply.([]interface{})
	if len(itemsT)%2 != 0 {
		return map[string]string{}
	}

	m := make(map[string]string)
	for i, _ := range itemsT {
		if i%2 == 0 {
			m[string(itemsT[i].([]byte))] = string(itemsT[i+1].([]byte))
		}
	}
	return m
}
