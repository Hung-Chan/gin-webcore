package message

const (
	// Success 回傳成功 .
	Success = 200

	// Error 回傳失敗 .
	Error = 500

	// DataBindError 資料綁定失敗 .
	DataBindError = 999999

	// TokenEmptyString Token 空白 .
	TokenEmptyString = 100001

	// TokenParseError Token 解析錯誤.
	TokenParseError = 100002

	// TokenTimeout Token 時效已過.
	TokenTimeout = 100003

	// LoginAccountNotFound 登入找不到帳號 .
	LoginAccountNotFound = 200001

	// LoginPasswordError 登入密碼錯誤 .
	LoginPasswordError = 200002

	// TokenProduceError 產生Token失敗 .
	TokenProduceError = 200003

	// LoginTokenWriteError 登入Token寫入失敗 .
	LoginTokenWriteError = 200004
)

var errorMessage = map[int]string{
	Success:              "Success",
	Error:                "Error",
	DataBindError:        "Data Bind Error",
	TokenEmptyString:     "Token Empty String",
	TokenParseError:      "Token Parse Error",
	TokenTimeout:         "Token Timeout",
	LoginAccountNotFound: "Login Account Not Found",
	LoginPasswordError:   "Login Password Error",
	TokenProduceError:    "Token Produce Error",
	LoginTokenWriteError: "Login Token Write Error",
}

// ErrorMessage 回傳相關訊息.
func ErrorMessage(code int) string {
	msg, err := errorMessage[code]

	if err {
		return msg
	}

	return errorMessage[code]
}
