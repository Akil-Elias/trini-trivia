package models

import "github.com/jackc/pgx/pgtype"

type User struct {
	ID        pgtype.UUID
	Email     string `json:"email"`
	Name      string `json:"name"`
	Gender    bool   `json:"gender"`
	Highscore int    `json:"highscore"`
}

type Question struct {
	ID       pgtype.UUID
	Level    int    `json:"level"`
	Category string `json:"category"`
	Number   int    `json:"number"`
	Points   int    `json:"points"`
}
