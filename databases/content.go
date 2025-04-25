package databases

import (
	"snippetsharing/models"
)

type Content struct {
}

// hashからコンテンツを取得する
func (content *Content) Get(hash string) (*models.Content, error) {

	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	c := models.Content{}
	row := db.QueryRow(`SELECT * FROM content WHERE hash = ?`, hash)
	if row.Err() != nil {
		return &models.Content{}, err
	}
	err = row.Scan(&c.ID, &c.Hash, &c.Sourcecode, &c.Language)
	if err != nil {
		return &models.Content{}, err
	}
	return &c, err
}

// コンテンツを登録する
func (content *Content) Create(c *models.Content) (*models.Content, error) {

	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// INSERT文を実行
	query := `INSERT INTO content (hash, sourcecode, language) VALUES (?, ?, ?)`
	result, err := db.Exec(query, c.Hash, c.Sourcecode, c.Language)
	if err != nil {
		return nil, err
	}

	// 挿入されたIDを取得
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	c.ID = int(id)

	return c, nil
}
