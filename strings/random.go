package strings

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// Random 获取随机字符串
func Random(l int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	if l > 17 {
		result = append(result, byte(time.Now().Unix()))
		l -= 10
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// MD5 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
