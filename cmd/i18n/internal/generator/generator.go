package generator

import (
	"encoding/json"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

var funcs = template.FuncMap{
	"funcName": func(lang string) string {
		lang = strings.ReplaceAll(lang, "_", "")
		lang = strings.ToUpper(lang[:1]) + lang[1:]
		return lang
	},
}

func Generator(i18nPath string) error {
	translationPath := filepath.Join(i18nPath, "translations")
	fi, err := os.ReadDir(translationPath)
	if err != nil {
		return err
	}
	//if err = utils.CreateFolder(filepath.Join(i18nPath, "i18n")); err != nil {
	//	return err
	//}
	goFile, err := os.Create(filepath.Join(i18nPath, "generated.go"))
	if err != nil {
		return err
	}
	data := make(map[string]*map[string]string)

	for _, v := range fi {
		if !v.IsDir() {
			continue
		}
		dataFiles, err := os.ReadDir(path.Join(translationPath, v.Name()))
		if err != nil {
			return err
		}
		data[v.Name()] = new(map[string]string)

		for _, file := range dataFiles {
			content, err := os.ReadFile(path.Join(translationPath, v.Name(), file.Name()))
			if err != nil {
				return err
			}
			err = json.Unmarshal(content, data[v.Name()])
			if err != nil {
				return err
			}
		}
	}
	err = i18nTmpl.Execute(goFile, struct {
		Data      map[string]*map[string]string
		BackQuote string
	}{
		data,
		"`",
	})
	if err != nil {
		return err
	}
	return nil
}

var i18nTmpl = template.Must(template.New("i18n").Funcs(funcs).Parse(`package i18n

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

{{- range $k, $v := .Data }}
// init{{ funcName $k }} will init {{ $k }} support.
func init{{ funcName $k }}(tag language.Tag) {
	{{- range $k, $v := $v }}
	_ = message.SetString(tag, {{$k}}, {{$v}})
{{- end }}
}
{{- end }}
`))
