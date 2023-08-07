package common

import (
	"fmt"
	"math/rand"
	"time"
)

var MyRand *rand.Rand

func init() {
	MyRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GenerateRandomString() string {
	const orderNumberLength = 5
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// 生成随机字符串
	randString := make([]rune, orderNumberLength)
	for i := 0; i < orderNumberLength; i++ {
		randString[i] = letterRunes[MyRand.Intn(len(letterRunes))]
	}

	// 生成时间戳，精确到毫秒
	timestamp := GenSnowflakeInt64()

	// 组合随机字符串和时间戳生成订单号
	orderNumber := fmt.Sprintf("%s%d", string(randString), timestamp)

	return orderNumber
}
