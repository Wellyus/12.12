package main

import (
	"fmt"
	"github/Wellyus/12.12/web"
)

func main() {
	var post = web.Post{
		Content: "hello from ubuntu",
		Author:  "Wellyus",
	}
	post.Create()
	post_ := web.Get(0)
	fmt.Println(post_)
}
