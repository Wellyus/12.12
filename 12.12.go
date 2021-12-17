package main

import (
	"fmt"

	"github.com/Wellyus/12.12/web"
)

func main() {
	var post = web.Post{
		Content: "hello from ubuntu",
		Author:  "Wellyus",
	}
	post.Create()
	post_, err := web.GetPost(0)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(post_)
}
