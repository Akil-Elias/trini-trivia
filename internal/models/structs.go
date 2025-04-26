package models

import (
	"log"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx"
)

type User struct {
	ID     uuid.UUID
	Email  string `json:"email"`
	Name   string `json:"name"`
	Gender bool   `json:"gender"`
	Rank   int    `json:"rank"`
	Points int    `json:"points"`
}

type Question struct {
	ID           uuid.UUID
	Category     string   `json:"category"`
	Number       int      `json:"number"`
	Body         string   `json:"body"`
	Answer       string   `json:"answer"`
	WrongAnswers []string `json:"wrong_answers"`
}

type Server struct {
	Addr   string
	Logger *log.Logger
	DB     *pgx.Conn
	Router *http.ServeMux
}
