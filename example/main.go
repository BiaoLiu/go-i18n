package main

import (
	"fmt"

	"github.com/BiaoLiu/go-i18n"
	goi18n "github.com/BiaoLiu/go-i18n/example/pkg/i18n"
	"golang.org/x/text/language"
)

func main() {
	goi18n.InitLang(language.English)
	s1 := i18n.Sprintf("Login error. username:%s not found", "foo")
	fmt.Println(s1)

	goi18n.InitLang(language.Chinese)
	s2 := i18n.Sprintf("Incorrect password")
	fmt.Println(s2)
}
