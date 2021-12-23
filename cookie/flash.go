package cookie

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("hello from setMessage")
	cookie := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &cookie)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message found")
		}
	} else {
		new_cookie := http.Cookie{
			Name:    "flase",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &new_cookie)
		val, _ := base64.URLEncoding.DecodeString(cookie.Value)
		fmt.Fprintln(w, string(val))
	}
}
