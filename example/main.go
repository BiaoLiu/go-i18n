package main

import (
	"fmt"

	goi18n "github.com/BiaoLiu/go-i18n"
	"github.com/BiaoLiu/go-i18n/example/pkg/i18n"
	"golang.org/x/text/language"
)

func main() {
	i18n.InitLang(language.English)
	s1 := goi18n.Sprintf("login error")
	fmt.Println(s1)

	i18n.InitLang(language.Chinese)
	s2 := goi18n.Sprintf("login error")
	fmt.Println(s2)
}
