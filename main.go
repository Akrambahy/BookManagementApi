package main

import (
"fmt"
"sync"
"encoding/json"
"strconv"
"string"
"net/http"
)

type Book struct{

	ID int 	`json : "id"`
	Title string `json : "title"`
	Author string `json : "author"`
	PublicationYear int `json : "publication_year"`

}


var (
books= make(map[int]Book)
mutex =& sync.Mutex()
)


func main(){

}