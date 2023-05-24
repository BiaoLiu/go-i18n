# go-i18n

支持i8n的golang库


## 安装

```
go get -u github.com/BiaoLiu/go-i18n/cmd/i18n
```

CLI工具：

```
go install github.com/BiaoLiu/go-i18n/cmd/i18n@latest
```

## 使用

```
s := i18n.Sprintf("login error. username:%s not found", "foo")
i18n.Printf("password error")
```

### 提取待翻译字符串

使用CLI工具提取待翻译字符串，并使用language库生成不同语言环境的字符串。待翻译的字符将被提取到`translations/en_US` 目录下，以 JSON 文件的形式存储

```shell
i18n gen --output ./pkg/i18n
```

输出

```shell
.
└── i18n
    ├── generated.go
    ├── i18n.go
    └── translations
        ├── en_US
           └── data.json
       
```

en_US/data.json

```
{
  "\"Incorrect password\"": "\"Incorrect password\"",
  "\"Login error. username:%s not found\"": "\"Login error. username:%s not found\""
}
```

### 翻译字符串

在`translations`目录下添加不同语言的翻译，如：添加中文翻译，translations目录下添加zh_CN目录与data.json文件

```
.
├── en_US
│   └── data.json
└── zh_CN
    └── data.json
```

zh_CN/data.json

```
{
  "\"Incorrect password\"": "\"密码错误\"",
  "\"Login error. username:%s not found\"": "\"登录失败. 用户名:%s 不存在\"",
}
```

### 应用已翻译字符串

将翻译完毕之后的字符串应用到程序中，使用language库生成不同语言环境的字符串

```
i18n gen --output ./pkg/i18n
```

输出

generated.go

```
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
```

I18n.go

```
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
```

此时，通过调用InitLang方法传入语言参数，可以按照语言来进行初始化了