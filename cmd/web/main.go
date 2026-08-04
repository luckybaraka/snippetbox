package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	// type config struct {
	// 	addr      string
	// 	staticDir string
	// }

	// var cfg config

	// flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network addrress")
	// flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	addr := flag.String("addr", ":4000", "HTTP Network address")
	// importantly we use flag.Parse() function to parse the command line flag
	//This reads in the comamnd-line flag value and assign it to the addr
	//Otherwise it will always default value ":4000", If any errors are
	//encountered during parsing the application will be terminated
	flag.Parse()

	//use the log.New() to create a loggter for writting information messages. This taskes
	//three params: the destination to write the logs to (os.Stdout), a string
	// prefix for message (INFO followed by a tab), and flags to indicate what
	//additonal information to include (local date and time). Note that flags
	//are joined using the bitwise OR operator |.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile|log.LUTC)

	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Register the other application routes as normal.
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	infoLog.Printf("Starting a server on %s", *addr) //informational message
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err) //error message

	//we are learning that we can actually redirect the logs into something either splunk or on-disk file by using
	//here is wht we did in this application go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log
}
