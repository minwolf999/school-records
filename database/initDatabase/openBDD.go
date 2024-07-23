package initdatabase

import "database/sql"

/*
This function takes no argument

The objective of this function is to open the connection to the bdd.

The function gonna return:
  - a connection to the BDD
  - an error
*/
func OpenBDD() (*sql.DB, error) {
	return sql.Open("sqlite3", "database/database.sqlite3")
}
