package files

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
	"log"
	"net/http"

	auth "github.com/abbot/go-http-auth"
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

func genFileServer(dir string, prefix string) http.HandlerFunc {
	fileServer := http.FileServer(http.Dir(dir))
	fileHandler := http.StripPrefix("/static/", fileServer)

	return func(writer http.ResponseWriter, req *http.Request) {
		fileHandler.ServeHTTP(writer, req)
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
