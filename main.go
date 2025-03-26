package main

import (
	"fmt"
	foo "net/http"
)

func main(){
	fmt.Println("Hello World")

	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("HTTP Response status: ", resp.Status)
}