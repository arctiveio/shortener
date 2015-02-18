package db

import (
	"errors"

	tiedot "github.com/HouzuoGuo/tiedot/db"
)

var cachedDB tiedot.DB

const collectionName = "short_url"

var indexes = []string{"short_url", "url"}

type UrlModel struct {
	Url      string `json:"url" required:"true"`
	ShortUrl string `json:"short_url"`
}

func (self UrlModel) Insert() error {
	urlDB := getDB()
	coll := urlDB.Use(collectionName)

	x := map[string]interface{}{}
	x["url"] = self.Url
	x["short_url"] = self.ShortUrl

	_, err := coll.Insert(x)
	return err
}

func getDB() *tiedot.DB {
	return &cachedDB
}

func GetOne(key, value string, url_object *UrlModel) error {
	query := map[string]interface{}{
		"eq": value,
		"in": []interface{}{key},
	}

	urlDB := getDB()
	coll := urlDB.Use(collectionName)

	queryResult := make(map[int]struct{})
	if err := tiedot.EvalQuery(query, coll, &queryResult); err != nil {
		panic(err)
	} else {
		for id := range queryResult {
			doc, err := coll.Read(id)
			if err != nil {
				panic(err)
			}
			url_object.Url = doc["url"].(string)
			url_object.ShortUrl = doc["short_url"].(string)
			return nil
		}
	}

	return errors.New("Object not found")
}

func Count(key, value string) int {
	results := 0

	query := map[string]interface{}{
		"eq": value,
		"in": []interface{}{key},
	}

	urlDB := getDB()
	coll := urlDB.Use(collectionName)

	queryResult := make(map[int]struct{})
	if err := tiedot.EvalQuery(query, coll, &queryResult); err != nil {
		panic(err)
	} else {
		results = len(queryResult)
	}

	return results
}

func InitDB(path string) {

	urlDB, err := tiedot.OpenDB(path)
	if err != nil {
		panic(err)
	}

	var hasCol bool
	for _, name := range urlDB.AllCols() {
		if name == collectionName {
			hasCol = true
			break
		}
	}

	if !hasCol {
		if err := urlDB.Create(collectionName); err != nil {
			panic(err)
		}
	}

	// ****** Index *********

	coll := urlDB.Use(collectionName)

	for _, indexPath := range indexes {
		var indexFound bool
		for _, path := range coll.AllIndexes() {
			if path[0] == indexPath {
				indexFound = true
				break
			}
		}

		if !indexFound {
			if err := coll.Index([]string{indexPath}); err != nil {
				panic(err)
			}
		}
	}

	cachedDB = *urlDB
}
