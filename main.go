package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	apiUrl := "https://jsonplaceholder.typicode.com/posts"

	res,err := http.Get(apiUrl)

	if err != nil {
		fmt.Println("Error Making Http Request")
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error Reading Res Body")
		return
	}
	var posts []Post
	err = json.Unmarshal(body,&posts)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Posts from JSONPlaceholder:")

	for _ ,post := range posts[:5] {
		fmt.Printf("\nPost ID: %d\nUser ID: %d\nTitle: %s\nBody: %s\n",
			post.ID, post.UserID, post.Title, post.Body)
	}


}