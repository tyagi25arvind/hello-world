package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("public" + r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		defer f.Close()
		io.Copy(w, f)
	})
	http.ListenAndServe(":8000", nil)
}

/*
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
}
*/
/*
import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Ji"))
	})
	http.ListenAndServe(":8000", nil)
}
*/

/*
func NotFoundHandler() Handler
func RedirectHandler(url string, code int) Handler
func StripPrefix(prefix string, h handler) Handler
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler
func FileServer (root FileSystem) Handler
type FileSystem interface {
	Open(name string) (File, error)
}
type Dir string
func (d Dir) Open(name string) (File, error)
*/

/* func main() {
	http.Handle("/", &myHandler{greeting: "Hello"})
	http.ListenAndServe(":8000", nil)
}

type myHandler struct {
	greeting string
}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v world", mh.greeting)))
}
*/
