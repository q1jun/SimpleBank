package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano()) //确保每次运行生成的值不同

}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	//返回 min + [0,max-min) ,即 [min , max)的随机整数
	return min + rand.Int63n(max - min + 1) //rand.Int63n()返回[0,n-1)之间的随机整数
}

//RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i:= 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

//RandomOwner generates a random owner name
func RandomOwner() string{
	return RandomString(6)
}

//RandomMoney generates a random amount of money
func RandomMoney() int64{
	return RandomInt(0,1000)
}

//RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{
		"EUR", "USD", "CAD",
	}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
