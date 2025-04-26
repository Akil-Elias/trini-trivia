package db

import (
	"context"

	"github.com/Akil-Elias/trini-trivia/internal/models"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx"
)

func InsertQuestion(conn *pgx.Conn, q models.Question) error {
	_, err := conn.Exec(` INSERT INTO questions (id, category, number, body, answer, wrong_answers)
        VALUES ($1, $2, $3, $4, $5, $6)`, context.Background(),
		q.ID, q.Category, q.Number, q.Body, q.Answer, q.WrongAnswers)

	return err
}

func GetQuestionByID(conn *pgx.Conn, id uuid.UUID) (models.Question, error) {
	var q models.Question

	row := conn.QueryRow(`
        SELECT id, category, number, body, answer, wrong_answers
        FROM questions
        WHERE id = $1
    `, context.Background(), id)

	err := row.Scan(&q.ID, &q.Category, &q.Number, &q.Body, &q.Answer, &q.WrongAnswers)
	return q, err
}

func GetAllQuestions(conn *pgx.Conn) ([]models.Question, error) {
	var questions []models.Question

	rows, err := conn.Query(`
        SELECT id, category, number, body, answer, wrong_answers
        FROM questions
    `, context.Background())

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var q models.Question
		err := rows.Scan(&q.ID, &q.Category, &q.Number, &q.Body, &q.Answer, &q.WrongAnswers)
		if err != nil {
			return nil, err
		}
		questions = append(questions, q)
	}

	return questions, rows.Err()
}
