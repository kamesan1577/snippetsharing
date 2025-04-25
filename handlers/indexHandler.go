package handlers

import (
	"fmt"
	"net/http"
	"snippetsharing/databases"
)

type indexParam struct {
	Language   string
	Sourcecode string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := GetTemplate("index.html")
	if err != nil {
		return
	}

	// urlからhashを受け取る
	hash := "kakiaf"

	c := databases.Content{}
	content, err := c.Get(hash)
	fmt.Println(content)
	if err != nil {
		fmt.Println(err) // TODO: あとでちゃんとロギングかく
		return
	}
	param := indexParam{
		Language:   content.Language,
		Sourcecode: content.Sourcecode,
	}
	err = tmpl.Execute(w, param)
	if err != nil {
		return
	}
}
