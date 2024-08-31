package bcode

// ErrorCode
//
// # 错误码
//
// 错误码，用于返回错误信息；错误码包含输出，状态码，消息；
//
// # 字段
//   - Output: string	输出
//   - Code: int		状态码
//   - Message: string	消息
type ErrorCode struct {
	Output  string `json:"output"`
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

// --------------------------------------------------------------------------------
// 错误码，用于返回错误信息；错误码包含输出，状态码，消息
// --------------------------------------------------------------------------------

var (
	// Success
	//
	// # 成功
	//
	// 用于返回成功响应，返回成功响应码 200，默认消息 "成功"；
	Success = ErrorCode{Output: "Success", Code: 20000, Message: "成功"}

	// BadRequestInvalidInput
	//
	// # 无效的输入
	//
	// 用于返回错误响应，返回错误响应码 40001，默认消息 "无效的输入"；
	BadRequestInvalidInput = ErrorCode{Output: "RequestInvalid", Code: 40001, Message: "无效的输入"}

	// BadRequestMissingParameter
	//
	// # 缺少必需的参数
	//
	// 用于返回错误响应，返回错误响应码 40002，默认消息 "缺少必需的参数"；
	BadRequestMissingParameter = ErrorCode{Output: "RequestInvalid", Code: 40002, Message: "缺少必需的参数"}

	// BadRequestInvalidParameter
	//
	// # 无效的参数值
	//
	// 用于返回错误响应，返回错误响应码 40003，默认消息 "无效的参数值"；
	BadRequestInvalidParameter = ErrorCode{Output: "RequestInvalid", Code: 40003, Message: "无效的参数值"}

	// BadRequestInvalidJson
	//
	// # 无效的 JSON 格式
	//
	// 用于返回错误响应，返回错误响应码 40004，默认消息 "无效的 JSON 格式"；
	BadRequestInvalidJson = ErrorCode{Output: "RequestInvalid", Code: 40004, Message: "无效的 JSON 格式"}

	// Unauthorized
	//
	// # 未授权
	//
	// 用于返回错误响应，返回错误响应码 40101，默认消息 "未授权"；
	Unauthorized = ErrorCode{Output: "Fail", Code: 40101, Message: "未授权"}

	// UnauthorizedTokenExpired
	//
	// # 令牌已过期
	//
	// 用于返回错误响应，返回错误响应码 40102，默认消息 "令牌已过期"；
	UnauthorizedTokenExpired = ErrorCode{Output: "Fail", Code: 40102, Message: "令牌已过期"}

	// UnauthorizedInvalidToken
	//
	// # 无效的令牌
	//
	// 用于返回错误响应，返回错误响应码 40103，默认消息 "无效的令牌"；
	UnauthorizedInvalidToken = ErrorCode{Output: "Fail", Code: 40103, Message: "无效的令牌"}

	// UnauthorizedTokenNotProvided
	//
	// # 未提供令牌
	//
	// 用于返回错误响应，返回错误响应码 40104，默认消息 "未提供令牌"；
	UnauthorizedTokenNotProvided = ErrorCode{Output: "Fail", Code: 40104, Message: "未提供令牌"}

	// Forbidden
	//
	// # 禁止访问
	//
	// 用于返回错误响应，返回错误响应码 40301，默认消息 "禁止访问"；
	Forbidden = ErrorCode{Output: "Deny", Code: 40301, Message: "禁止访问"}

	// ForbiddenAccessDenied
	//
	// # 访问被拒绝
	//
	// 用于返回错误响应，返回错误响应码 40302，默认消息 "访问被拒绝"；
	ForbiddenAccessDenied = ErrorCode{Output: "Deny", Code: 40302, Message: "访问被拒绝"}

	// ForbiddenInsufficientPrivileges
	//
	// # 权限不足
	//
	// 用于返回错误响应，返回错误响应码 40303，默认消息 "权限不足"；
	ForbiddenInsufficientPrivileges = ErrorCode{Output: "Deny", Code: 40303, Message: "权限不足"}

	// NotFound
	//
	// # 资源未找到
	//
	// 用于返回错误响应，返回错误响应码 40401，默认消息 "资源未找到"；
	NotFound = ErrorCode{Output: "Error", Code: 40401, Message: "资源未找到"}

	// NotFoundUser
	//
	// # 用户未找到
	//
	// 用于返回错误响应，返回错误响应码 40402，默认消息 "用户未找到"；
	NotFoundUser = ErrorCode{Output: "Error", Code: 40402, Message: "用户未找到"}

	// NotFoundPage
	//
	// # 页面未找到
	//
	// 用于返回错误响应，返回错误响应码 40403，默认消息 "页面未找到"；
	NotFoundPage = ErrorCode{Output: "Error", Code: 40403, Message: "页面未找到"}

	// MethodNotAllowed
	//
	// # 方法不允许
	//
	// 用于返回错误响应，返回错误响应码 40501，默认消息 "方法不允许"；
	MethodNotAllowed = ErrorCode{Output: "Error", Code: 40501, Message: "方法不允许"}

	// RequestTimeout
	//
	// # 请求超时
	//
	// 用于返回错误响应，返回错误响应码 40801，默认消息 "请求超时"；
	RequestTimeout = ErrorCode{Output: "Timeout", Code: 40801, Message: "请求超时"}

	// Conflict
	//
	// # 冲突
	//
	// 用于返回错误响应，返回错误响应码 40901，默认消息 "冲突"；
	Conflict = ErrorCode{Output: "Conflict", Code: 40901, Message: "冲突"}

	// ConflictResourceAlreadyExists
	//
	// # 资源已存在
	//
	// 用于返回错误响应，返回错误响应码 40902，默认消息 "资源已存在"；
	ConflictResourceAlreadyExists = ErrorCode{Output: "Conflict", Code: 40902, Message: "资源已存在"}

	// Gone
	//
	// # 资源已删除
	//
	// 用于返回错误响应，返回错误响应码 41001，默认消息 "资源已删除"；
	Gone = ErrorCode{Output: "Gone", Code: 41001, Message: "资源已删除"}

	// PayloadTooLarge
	//
	// # 请求实体过大
	//
	// 用于返回错误响应，返回错误响应码 41301，默认消息 "请求实体过大"；
	PayloadTooLarge = ErrorCode{Output: "Error", Code: 41301, Message: "请求实体过大"}

	// UnsupportedMediaType
	//
	// # 不支持的媒体类型
	//
	// 用于返回错误响应，返回错误响应码 41501，默认消息 "不支持的媒体类型"；
	UnsupportedMediaType = ErrorCode{Output: "Error", Code: 41501, Message: "不支持的媒体类型"}

	// UnprocessableEntity
	//
	// # 不可处理的实体
	//
	// 用于返回错误响应，返回错误响应码 42201，默认消息 "不可处理的实体"；
	UnprocessableEntity = ErrorCode{Output: "Invalid", Code: 42201, Message: "不可处理的实体"}

	// UnprocessableEntityValidation
	//
	// # 验证错误
	//
	// 用于返回错误响应，返回错误响应码 42202，默认消息 "验证错误"；
	UnprocessableEntityValidation = ErrorCode{Output: "Invalid", Code: 42202, Message: "验证错误"}

	// TooManyRequests
	//
	// # 请求过多
	//
	// 用于返回错误响应，返回错误响应码 42901，默认消息 "请求过多"；
	TooManyRequests = ErrorCode{Output: "Limit", Code: 42901, Message: "请求过多"}

	// ServerInternalError
	//
	// # 服务器内部错误
	//
	// 用于返回错误响应，返回错误响应码 50001，默认消息 "服务器内部错误"；
	ServerInternalError = ErrorCode{Output: "Error", Code: 50001, Message: "服务器内部错误"}

	// ServerDatabaseError
	//
	// # 数据库错误
	//
	// 用于返回错误响应，返回错误响应码 50002，默认消息 "数据库错误"；
	ServerDatabaseError = ErrorCode{Output: "Error", Code: 50002, Message: "数据库错误"}

	// ServerConfigurationError
	//
	// # 配置错误
	//
	// 用于返回错误响应，返回错误响应码 50003，默认消息 "配置错误"；
	ServerConfigurationError = ErrorCode{Output: "Error", Code: 50003, Message: "配置错误"}

	// ServerDependencyError
	//
	// # 依赖错误
	//
	// 用于返回错误响应，返回错误响应码 50004，默认消息 "依赖错误"；
	ServerDependencyError = ErrorCode{Output: "Error", Code: 50004, Message: "依赖错误"}

	// ServerTimeout
	//
	// # 服务器超时
	//
	// 用于返回错误响应，返回错误响应码 50005，默认消息 "服务器超时"；
	ServerTimeout = ErrorCode{Output: "Timeout", Code: 50005, Message: "服务器超时"}

	// ServerUnknownError
	//
	// # 未知错误
	//
	// 用于返回错误响应，返回错误响应码 50006，默认消息 "未知错误"；
	ServerUnknownError = ErrorCode{Output: "Error", Code: 50006, Message: "未知错误"}

	// NotImplemented
	//
	// # 未实现
	//
	// 用于返回错误响应，返回错误响应码 50101，默认消息 "未实现"；
	NotImplemented = ErrorCode{Output: "Error", Code: 50101, Message: "未实现"}

	// BadGateway
	//
	// # 错误的网关
	//
	// 用于返回错误响应，返回错误响应码 50201，默认消息 "错误的网关"；
	BadGateway = ErrorCode{Output: "Error", Code: 50201, Message: "错误的网关"}

	// ServiceUnavailable
	//
	// # 服务不可用
	//
	// 用于返回错误响应，返回错误响应码 50301，默认消息 "服务不可用"；
	ServiceUnavailable = ErrorCode{Output: "Error", Code: 50301, Message: "服务不可用"}

	// ServiceUnderMaintenance
	//
	// # 服务维护中
	//
	// 用于返回错误响应，返回错误响应码 50302，默认消息 "服务维护中"；
	ServiceUnderMaintenance = ErrorCode{Output: "Error", Code: 50302, Message: "服务维护中"}

	// GatewayTimeout
	//
	// # 网关超时
	//
	// 用于返回错误响应，返回错误响应码 50401，默认消息 "网关超时"；
	GatewayTimeout = ErrorCode{Output: "Timeout", Code: 50401, Message: "网关超时"}
)
