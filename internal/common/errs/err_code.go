package response

/*
定义用户管理系统的统一响应格式和错误码

错误码分段规则：
1. 通用错误：0-999（成功、系统错误、参数错误等）
2. 用户模块：1000-1999（注册、登录、注销等）
3. 认证/Token模块：2000-2999（Token生成、校验等）
4. 权限模块：3000-3999（权限校验、管理员限制等）
5. 验证码模块：4000-4999（短信/邮箱验证码相关）
6. 数据校验模块：5000-5999（用户名/密码格式等）
*/
const (
	CodeSuccess           = 0   // 成功
	CodeUnknownError      = 100 // 未知系统错误
	CodeParamError        = 101 // 参数错误（通用，如必填项为空）
	CodeDatabaseError     = 102 // 数据库操作错误（如查询/插入失败）
	CodeRedisError        = 103 // Redis操作错误（如缓存设置/获取失败）
	CodeRequestFreqErr    = 104 // 请求频率过高（限流）
	CodeRequestTimeoutErr = 105 // 请求超时

	CodeUserNotFound     = 1001 // 用户不存在
	CodeUsernameExist    = 1002 // 用户已存在（注册时用户名重复）
	CodeUserPwdError     = 1003 // 登录密码错误
	CodeUserStatusForbid = 1004 // 用户账号被禁用
	CodeUserIDInvalid    = 1005 // 用户ID无效（如非数字/小于1）
	CodeUserPhoneExist   = 1006 // 手机号已绑定其他账号
	CodeUserEmailExist   = 1007 // 邮箱已绑定其他账号
	CodeUserOldPwdError  = 1008 // 修改密码时原密码错误
	CodeUserNotLogin     = 1009 // 用户未登录（需登录才能操作）
	CodeUserLogoutFailed = 1010 // 注销失败（如数据库操作异常）
	CodeUserNull         = 1011 // 用户参数错误

	CodeTokenInvalid     = 2001 // Token无效（篡改/格式错误）
	CodeTokenExpired     = 2002 // Token过期
	CodeTokenEmpty       = 2003 // 请求头中Token为空
	CodeTokenGenerateErr = 2004 // Token生成失败（如JWT签名错误）

	CodePermissionDenied = 3001 // 无操作权限（通用）
	CodeAdminOnly        = 3002 // 仅管理员可执行此操作
	CodeUserSelfOnly     = 3003 // 仅本人可操作（如注销自己账号）

	CodeSmsSendFailed    = 4001 // 短信验证码发送失败
	CodeSmsCodeInvalid   = 4002 // 短信验证码无效/过期
	CodeSmsCodeFreqLimit = 4003 // 验证码发送频率限制（如1分钟1次）
	CodeEmailSendFailed  = 4004 // 邮箱验证码发送失败

	CodeUsernameInvalid = 5001 // 用户名格式错误（3-20位，仅字母/数字）
	CodePwdInvalid      = 5002 // 密码格式错误（≥6位，含字母+数字）
	CodePhoneInvalid    = 5003 // 手机号格式错误（11位数字）
	CodeEmailInvalid    = 5004 // 邮箱格式错误（如无@符号）
)

// CodeMsgMap 错误码对应提示语（一一对应，无重复）
var CodeMsgMap = map[int]string{
	// 通用错误
	CodeSuccess:           "操作成功",
	CodeUnknownError:      "系统未知错误，请稍后重试",
	CodeParamError:        "请求参数错误",
	CodeDatabaseError:     "数据库操作失败",
	CodeRedisError:        "缓存操作失败",
	CodeRequestFreqErr:    "请求频率过高，请1分钟后重试",
	CodeRequestTimeoutErr: "请求超时",

	// 用户模块
	CodeUserNotFound:     "用户不存在",
	CodeUsernameExist:    "用户名已存在",
	CodeUserPwdError:     "密码错误",
	CodeUserStatusForbid: "用户账号已被禁用",
	CodeUserIDInvalid:    "用户ID无效",
	CodeUserPhoneExist:   "手机号已绑定其他账号",
	CodeUserEmailExist:   "邮箱已绑定其他账号",
	CodeUserOldPwdError:  "原密码错误",
	CodeUserNotLogin:     "用户未登录，请先登录",
	CodeUserLogoutFailed: "注销失败，请稍后重试",
	CodeUserNull:         "用户参数错误",

	// 认证/Token模块
	CodeTokenInvalid:     "登录凭证无效，请重新登录",
	CodeTokenExpired:     "登录凭证已过期，请重新登录",
	CodeTokenEmpty:       "请先登录",
	CodeTokenGenerateErr: "登录凭证生成失败",

	// 权限模块
	CodePermissionDenied: "无操作权限",
	CodeAdminOnly:        "仅管理员可执行此操作",
	CodeUserSelfOnly:     "仅可操作本人账号",

	// 验证码模块
	CodeSmsSendFailed:    "验证码发送失败，请稍后重试",
	CodeSmsCodeInvalid:   "验证码无效或已过期",
	CodeSmsCodeFreqLimit: "验证码发送频率过高，请1分钟后重试",
	CodeEmailSendFailed:  "邮箱验证码发送失败",

	// 数据校验模块
	CodeUsernameInvalid: "用户名格式错误（需3-20位，仅字母/数字）",
	CodePwdInvalid:      "密码格式错误（需≥6位，含字母+数字）",
	CodePhoneInvalid:    "手机号格式错误（需11位数字）",
	CodeEmailInvalid:    "邮箱格式错误",
}

// GetMsg 根据错误码获取默认提示语
func GetMsg(code int) string {
	if msg, ok := CodeMsgMap[code]; ok {
		return msg
	}
	return CodeMsgMap[CodeUnknownError]
}
