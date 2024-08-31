package butil

import (
	"crypto/md5"
	"github.com/google/uuid"
)

// GenerateUUID
//
// # 生成随机UUID
//
// 用于生成一个随机的 UUID，返回生成的 UUID；
// 生成的 UUID 是一个随机的值，每次生成的 UUID 都是不同的；
//
// # 返回
//   - getUUID 	生成的 UUID(uuid.UUID)
func GenerateUUID() (getUUID uuid.UUID) {
	getUUID, _ = uuid.NewV7()
	return getUUID
}

// GenerateUUIDFromString
//
// # 通过字符串生成UUID
//
// 用于通过字符串生成一个 UUID，生成的 UUID 是一个固定的值，不会随机生成；
// 根据输入参数的内容决定返回的 UUID。
//
//   - 例如：“hello” -> “5d41402a-bc4b-2a76-b971-9d911017c592”。
//
// # 参数:
//   - s 		传入的字符串(string)
//
// # 返回:
//   - getUUID 	生成的 UUID(uuid.UUID)
func GenerateUUIDFromString(s string) (getUUID uuid.UUID) {
	return md5.Sum([]byte(s))
}

// ConvertStringToUUID
//
// # 字符串转UUID
//
// 用于将字符串转换为 uuid.UUID，返回转换后的 uuid.UUID；
// 支持将不含“-”的 UUID（NoDashUUID） 转换为标准 uuid.UUID；
// 若输入的字符串不符合 uuid.UUID 标准或 NoDashUUID 格式，则返回一个空的 uuid.UUID(00000000-0000-0000-0000-000000000000)。
//
//   - 例如：“5d41402a-bc4b-2a76-b971-9d911017c592”(string) -> “5d41402a-bc4b-2a76-b971-9d911017c592”(uuid)；
//   - 例如：“5d41402abc4b2a76b9719d911017c592”(string) -> “5d41402a-bc4b-2a76-b971-9d911017c592”(uuid)；
//   - 例如：“5d41402abc4b2a76b9719d911017c59”(string) -> “00000000-0000-0000-0000-000000000000”(uuid)；
//
// # 参数:
//   - s 		传入的字符串(string)
//
// # 返回:
//   - getUUID 	转换后的 UUID(uuid.UUID)
func ConvertStringToUUID(s string) (getUUID uuid.UUID) {
	// 若为 NoDashUUID 格式，则转换为标准 UUID 格式
	if len(s) == 32 {
		s = s[:8] + "-" + s[8:12] + "-" + s[12:16] + "-" + s[16:20] + "-" + s[20:]
	}
	parse, _ := uuid.Parse(s)
	return parse
}
