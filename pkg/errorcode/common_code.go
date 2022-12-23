package errorcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "伺服器錯誤")
	InvalidParams             = NewError(10000001, "參數錯誤")
	NotFound                  = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "授權失敗，找不到 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "授權失敗，Token 錯誤")
	UnauthorizedTokenTimeout  = NewError(10000005, "授權失敗，Token 逾時")
	UnauthorizedTokenGenerate = NewError(10000006, "授權失敗，Token 生成失敗")
	TooManyRequests           = NewError(10000007, "請求過多")
)