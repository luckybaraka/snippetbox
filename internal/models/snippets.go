package models

import (
	"database/sql"
	"time"
)

//define a snippet that hold the data for an individdual snippet. Notice how the fields of the structs
//corresponding to the fields in our MySQL snippets tables?

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Defines a snippetModel type which wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// This insets a new snippet into the DB
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

// This gets a specific snippet with an ID
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// This returns a list of the latest 10 new snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
