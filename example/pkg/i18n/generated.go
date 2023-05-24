package i18n

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)
// initEnUS will init en_US support.
func initEnUS(tag language.Tag) {
	_ = message.SetString(tag, "Incorrect password", "Incorrect password")
	_ = message.SetString(tag, "Login error. username:%s not found", "Login error. username:%s not found")
}
// initZhCN will init zh_CN support.
func initZhCN(tag language.Tag) {
	_ = message.SetString(tag, "login error. username:%s not found", "登录失败. 用户名:%s 不存在")
	_ = message.SetString(tag, "password error", "password error")
}
