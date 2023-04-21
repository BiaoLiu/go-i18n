package i18n

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)
// initEnUS will init en_US support.
func initEnUS(tag language.Tag) {
	_ = message.SetString(tag, "login error", "login error")
}
// initZhCN will init zh_CN support.
func initZhCN(tag language.Tag) {
	_ = message.SetString(tag, "login error", "登录失败")
}
