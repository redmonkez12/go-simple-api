package migrations

import "embed"

//go:embed 00001_users.sql 00002_workouts.sql 00003_workout_entries.sql
var FS embed.FS
