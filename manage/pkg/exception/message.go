package exception

var Msg = map[int]string{
	SUCCESS:                        "操作成功",
	ERROR:                          "操作失败",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "解析TOKEN失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "TOKEN过期",
	ERROR_AUTH_CREAT_TOKEN:         "TOKEN创建失败",
	ERROR_AUTH:                     "TOKEN错误",
	ERROR_WRONG_ELEMENT:            "用户名或密码错误",
	ERROR_SAME_NAME:                "用户名已被注册",
	ERROR_ENCODE_FAIL:              "加密错误",
}

func GetMsg(code int) string {
	msg, ok := Msg[code]
	if ok {
		return msg
	}
	return Msg[ERROR]
}
