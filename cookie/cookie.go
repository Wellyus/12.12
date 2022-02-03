package cookie

import (
	"fmt"
	"log"
	"net/http"
)

func cookie() {
	server := http.Server{
		Addr: "localhost:8888",
	}
	http.HandleFunc("/set_cookie", set_Cookie)
	http.HandleFunc("/get_cookie", getCookie)
	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)
	server.ListenAndServe()
}

func set_Cookie(w http.ResponseWriter, r *http.Request) {
	cookie1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "cookie_from_cookie1",
		HttpOnly: true,
	}
	cookie2 := http.Cookie{Name: "second_cookie",
		Value:    "cookie_from_cookie2",
		HttpOnly: true,
	}
	//w.Header().Set("Set-Cookie", cookie1.String())
	//w.Header().Add("Set-Cookie", cookie2.String())
	http.SetCookie(w, &cookie1)
	http.SetCookie(w, &cookie2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	cookie1, err := r.Cookie("first_cookie")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, cookie1)
	cookies := r.Cookies()
	fmt.Fprintln(w, cookies)
}
