package test

import (
	"github.com/XiaoLFeng/go-general-utils/butil"
	"testing"
)

// TestGenerateUUIDFromString
//
// # 通过字符串生成UUID
//
// 用于测试通过字符串生成一个 uuid.UUID；
// 参考方法：butil.GenerateUUIDFromString
func TestGenerateUUIDFromString(t *testing.T) {
	getUUID := butil.GenerateUUIDFromString("hello")
	t.Log(getUUID)
}

// TestConvertStringToUUID
//
// # 字符串转UUID
//
// 用于测试将字符串转换为 uuid.UUID；
func TestConvertStringToUUID(t *testing.T) {
	getUUIDFirst := butil.ConvertStringToUUID("5d41402a-bc4b-2a76-b971-9d911017c592")
	t.Log(getUUIDFirst)
	getUuidSecond := butil.ConvertStringToUUID("5d41402abc4b2a76b9719d911017c592")
	t.Log(getUuidSecond)
	// bad test
	getUuidThird := butil.ConvertStringToUUID("5d41402abc4b2a76b9719d911017c59")
	t.Log(getUuidThird)
}
