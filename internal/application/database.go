package application

import "database/sql"

func (app *App) Database(dsn string) *sql.DB {
	if app.db == nil {
		return app.db
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	app.db = db

	return db
}
