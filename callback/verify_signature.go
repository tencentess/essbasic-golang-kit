package callback

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// 验证回调消息中的签名
func main() {
	// 回调消息体
	payload := "**********"
	// secretToken 创建应用号时配置的
	secretToken := "**********"

	// 1. 取出header [Content-Signature]
	signFromHeader := "***********"
	// 2. 验证签名
	hash := "sha256=" + Hmacsha256hex(payload, secretToken)

	//3. 如果验证通过，继续处理。如果不通过，忽略该请求
	fmt.Println(hash == signFromHeader)

}

// Hmacsha256hex hmac sha256
func Hmacsha256hex(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return hex.EncodeToString(hashed.Sum(nil))
}
