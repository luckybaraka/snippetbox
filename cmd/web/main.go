package main

import (
	"database/sql" // New import
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Define an application struct to hold the application-wide dependancies for the
// web application. For now we'll only include feilds for the two custom loggers, but
// we'll add more to it as the build progresses.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	// type config struct {
	// 	addr      string
	// 	staticDir string
	// }

	// var cfg config

	// flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network addrress")
	// flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	addr := flag.String("addr", ":4000", "HTTP Network address")

	//define a new command-line flag for the MySQL DSN string.
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
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

	//To keep the main() function tidy I've put the code for creating a connection
	//pool into the separate openDB() function below. We pass openDB() the DSN
	//from the command-line flag
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	//we also defer a call to db.Close(), so that the connection pool is closed
	//before the main() function exists.
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	//Initliaze a new http.server struct. We set the Addr and Handler fields so that
	//the server uses the same network address and routes as before, and set
	//the ErrorLog field so that the server now uses the custom logger in the event of any issues
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		//call the new app.routes() method to get the servemux containing routes
		Handler: app.routes(),
	}

	//We are learning that GO uses so many ways to show info into the terminal, in all cases it is important to NOT use Fatal() and Panic() anywhere outside our main.go

	//lets comment out how we can actaully store the logs in our project
	// f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	infoLog.Printf("Starting a server on %s", *addr) //informational message
	// Because the err variable is now already declared in the code above, we need
	// to use the assignment operator = here, instead of the := 'declare and assign'
	// operator.
	err = srv.ListenAndServe()
	errorLog.Fatal(err) //error message

	//we are learning that we can actually redirect the logs into something either splunk or on-disk file by using
	//here is wht we did in this application go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log
}

// The openDB() function wraps sql.Open() and returns a sql.DB() connection pool
// for a given DSN
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
