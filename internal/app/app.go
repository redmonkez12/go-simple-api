package app

import (
	"fe-go/internal/api"
	"fe-go/internal/store"
	"fe-go/migrations"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *pgxpool.Pool
}

func NewApplication() (*Application, error) {
	pgDb, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(migrations.FS, ".")
	if err != nil {
		return nil, err
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	workoutHandler := api.NewWorkoutHandler()

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDb,
	}

	return app, nil
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}
