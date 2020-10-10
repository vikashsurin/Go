package books

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"util/config"

	"gopkg.in/mgo.v2/bson"
)

//AllBooks ...
func AllBooks() ([]Book, error) {
	books := []Book{}
	err := config.BooksColl.Find(bson.M{}).All(&books)
	if err != nil {
		panic(err)
	}
	return books, nil
}

//OneBook ...
func OneBook(r *http.Request) (Book, error) {
	book := Book{}
	bk := Book{}
	bs, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bs, &bk)
	if err != nil {
		panic(err)
	}

	err = config.BooksColl.Find(bson.M{"isbn": &bk.Isbn}).One(&book)
	if err != nil {
		return book, err
	}
	return book, nil

}

//PutBook ...
func PutBook(r *http.Request) (Book, error) {
	book := Book{}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return book, errors.New("500. internal server error")
	}
	err = json.Unmarshal(b, &book)
	if err != nil {
		panic(err)
	}
	err = config.BooksColl.Insert(book)
	if err != nil {
		return book, errors.New("500. internal server error." + err.Error())

	}
	return book, nil

}

//UpdateBook ...
func UpdateBook(r *http.Request) (Book, error) {
	book := Book{}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return book, errors.New("500. internal server error")
	}
	err = json.Unmarshal(b, &book)
	if err != nil {
		panic(err)
	}

	err = config.BooksColl.Update(bson.M{"isbn": book.Isbn}, &book)
	if err != nil {
		return book, err
	}
	return book, nil

}

//DeleteBook ...
func DeleteBook(r *http.Request) error {
	book := Book{}
	bs, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bs, &book)
	if err != nil {
		panic(err)
	}

	err = config.BooksColl.Remove(bson.M{"isbn": book.Isbn})
	if err != nil {
		return errors.New("500. internal server error")
	}
	return nil
}
