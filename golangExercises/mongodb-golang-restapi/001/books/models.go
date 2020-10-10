package books

//Book ...
type Book struct {
	Isbn   int    `json:"isbn" bson:"isbn"`
	Title  string `json:"title" bson:"title"`
	Author string `json:"author" bason:"author"`
}
