package handlers

import (
	"html/template"
	"os"
)

// 指定されたファイル名に一致するテンプレートをtemplatesディレクトリから取得する
func GetTemplate(filename string) (*template.Template, error) {
	exe, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	templDir := "/templates/"
	templPath := exe + templDir + filename
	tmpl := template.Must(template.ParseFiles(templPath))
	return tmpl, nil
}
