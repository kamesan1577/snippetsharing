package handlers

import (
	"fmt"
	"net/http"
	"snippetsharing/databases"
	"snippetsharing/library"
	"snippetsharing/models"
)

type shareParam struct {
	Url string
}

func ShareHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := GetTemplate("index.html")
	if err != nil {
		return
	}

	hash := library.GetHash()
	sourcecode := "hoge" // リクエストから取得する
	language := "fugascript"
	cM := models.Content{
		Hash:       hash,
		Sourcecode: sourcecode,
		Language:   language,
	}
	cDb := databases.Content{}
	content, err := cDb.Create(&cM)
	fmt.Println(content)
	if err != nil {
		fmt.Println(err) // TODO: あとでちゃんとロギングかく
		return
	}
	param := shareParam{
		Url: "www.example.com/" + content.Hash,
	}
	err = tmpl.Execute(w, param)
	if err != nil {
		return
	}
}
