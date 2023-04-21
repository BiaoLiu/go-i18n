package i18n

import (
	"github.com/BiaoLiu/go-i18n"
	"golang.org/x/text/language"
)

// InitLang will init i18n support via input language.
func InitLang(lang language.Tag) {
	tag, _, _ := i18n.Supported.Match(lang)
	switch tag {
	case language.AmericanEnglish, language.English:
		initEnUS(lang)
	case language.SimplifiedChinese, language.Chinese:
		initZhCN(lang)
	default:
		initEnUS(lang)
	}
	i18n.InitPrinter(lang)
}
