package butil

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// TokenRemoveBearer
//
// # 移除Token中的Bearer
//
// 用于移除 Token 中的 "Bearer " 前缀，返回移除后, 前缀的 Token；
// 如果数据不包含 "Bearer " 前缀，则直接返回原数据；
// 该算法支持任何形式的的 Token。
//
// # 参数
//   - getToken 	获取的 Token 整体(string)
//
// # 返回
//   - string 		去除 Bearer 前缀的 Token 或直接返回的 Token
func TokenRemoveBearer(getToken string) string {
	// 检查数据是否包含 "Bearer " 前缀
	if !strings.Contains(getToken, "Bearer ") {
		return strings.Replace(getToken, "Bearer ", "", 1)
	} else {
		return getToken
	}
}

// GenerateRandomString
//
// # 生成随机字符串
//
// 用于生成一个指定长度的随机字符串，返回生成的随机字符串. 生成的字符串为 string 类型；
// 生成的随机字符串为大小写字母和数字的组合。
//
// # 参数:
//   - length 	生成的字符串长度(int)
//
// # 返回:
//   - string 	生成的随机字符串
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		randIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[randIndex.Int64()]
	}
	return string(result)
}
