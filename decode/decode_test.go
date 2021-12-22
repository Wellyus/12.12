package decode

import (
	"testing"
)

func TestDecode(t *testing.T) {
	post, err := decode("test.json")
	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("wrong id, was expecting 1 but got", post.Id)
	}
	if post.Content != "hello ubuntu" {
		t.Error("wrong content, was expecting hello ubutu but got", post.Content)
	}
}

func TestEncode(t *testing.T) {
	t.Skip("skiping encoding now")
}
