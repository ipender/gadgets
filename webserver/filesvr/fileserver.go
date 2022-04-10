package main

/*
Serve is a very simple static file server in go
Usage:
	-p="8100": port to serve on
	-d=".":    the directory of static files to host
Navigating to http://localhost:8100 will display the index.html or directory
listing file.
*/
import (
	"flag"
	"io"
	"log"
	"net/http"

	auth "github.com/abbot/go-http-auth"
	"golang.org/x/crypto/bcrypt"
)

const (
	link = `<link rel="stylesheet" href="/path/to/style.css">`
)

func main() {
	authenticator := auth.NewBasicAuthenticator("localhost", Secret)
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	fileHandler := genFileServer(*directory, "/static/")
	http.Handle("/static/", auth.JustCheck(authenticator, fileHandler))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func HashPassword(password []byte, cost int) (hash []byte, err error) {
	return bcrypt.GenerateFromPassword(password, cost)
	// if err != nil {
	// 	log.Fatalf("generate password hashing failed, err=%v\n", err)
	// }
}

func genFileServer(dir string, prefix string) http.HandlerFunc {
	fileServer := http.FileServer(http.Dir(dir))
	fileHandler := http.StripPrefix("/static/", fileServer)

	return func(writer http.ResponseWriter, req *http.Request) {
		fileHandler.ServeHTTP(writer, req)
		io.WriteString(writer, link)
	}
}

func Secret(user, realm string) string {

	users := map[string]string{
		"john": "$1$dlPL2MqE$oQmn16q49SqdmhenQuNgs1", // hello
		// "export": "$def$abc$mg.JpERPBSrgcj6HG0/R0.",
	}

	if a, ok := users[user]; ok {
		return a
	}
	return ""
}

type decoratorForHTTPAuth struct {
	realm        string
	userName     string
	passwordHash []byte

	handler http.Handler
}

func (a *decoratorForHTTPAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handler.ServeHTTP(w, r)
}

func (a *decoratorForHTTPAuth) RequireBasicAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Write()
}
