// Package sqlscan allows scanning data from *sql.Rows into complex Go types.
/*
sqlscan is a wrapper around github.com/georgysavva/dbscan package.
It contains adapters and proxy functions that are meant to connect database/sql
with dbscan functionality. sqlscan mirrors all capabilities provided by dbscan.
See dbscan docs to get familiar with all details and features.

How to use

The most common way to use sqlscan is by calling QueryAll or QueryOne function,
it's as simple as this:

	type User struct {
		ID    string `db:"user_id"`
		Name  string
		Email string
		Age   int
	}

	db, _ := sql.Open("pgx", "example-connection-url")

	// Use QueryAll to query multiple records.
	var users []*User
	sqlscan.QueryAll(ctx, &users, db, `SELECT user_id, name, email, age FROM users`)
	// users variable now contains data from all rows.

Types that implement sql Scanner

sqlscan plays well with custom types that implement sql.Scanner interface, here is how you can use them:

	type Data struct {
		Title   string
		Text    string
		Counter int
	}

	func (d *Data) Scan(value interface{}) error {
		b, ok := value.([]byte)
		if !ok {
			return errors.New("Data.Scan: value isn't []byte")
		}
		return json.Unmarshal(b, &d)
	}

	type Post struct {
		PostID  string
		OwnerID string
		Data    Data
	}

	type Comment struct {
		CommentID    string
		OwnerID      string
		OptionalData *Data
	}

Note that type implementing sql.Scanner (Data struct in the example above),
can be presented both by value as in Post struct and by a pointer as in Comment struct.
*/
package sqlscan
