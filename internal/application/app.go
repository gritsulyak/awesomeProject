package application

import "database/sql"

type App struct {
	db *sql.DB
}

func NewApp() *App {
	return &App{}
}

func (app *App) Start() {

}

func (app *App) Stop() {

}
