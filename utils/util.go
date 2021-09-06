package utils

import (
	"math/rand"
	"time"
)


// 定义一个根据数值返回一个随机长度相等于数值的字符串
func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwsyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	res := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range res {
		res[i] = letters[rand.Intn(len(letters))]
	}
	return string(res)
}