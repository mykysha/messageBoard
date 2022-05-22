package db

import (
	"database/sql"
	"errors"
	"fmt"

	// postgres driver.
	_ "github.com/lib/pq"
	"github.com/nndergunov/messageBoard/server/pkg/db/models"
	"github.com/nndergunov/messageBoard/server/pkg/domain"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type DB struct {
	db *sql.DB
}

func NewDB(url string) (*DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("connecting to database: %w", err)
	}

	return &DB{db: db}, nil
}

func (d DB) getUserID(nickname string) (int, error) {
	user, err := models.Users(qm.Where("nickname=?", nickname)).One(d.db)
	if err != nil {
		return 0, fmt.Errorf("reading user table: %w", err)
	}

	return user.ID, nil
}

func (d DB) getUserByID(userID int) (*models.User, error) {
	user, err := models.Users(qm.Where("id=?", userID)).One(d.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from table: %w", err)
	}

	return user, nil
}

func (d DB) InsertUser(nickname string) (int, error) {
	id, err := d.getUserID(nickname)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return id, nil
		}
	}

	var user models.User

	user.Nickname = nickname

	err = user.Insert(d.db, boil.Infer())
	if err != nil {
		return 0, fmt.Errorf("inserting user: %w", err)
	}

	id, err = d.getUserID(nickname)
	if err != nil {
		return 0, fmt.Errorf("reading element id: %w", err)
	}

	return id, nil
}

func (d DB) InsertMessage(msg domain.Message) error {
	var message models.Message

	message.Text = msg.Text
	message.Time = msg.Time

	id, err := d.getUserID(msg.Author)
	if err != nil {
		return fmt.Errorf("getting user id failed: %w", err)
	}

	message.AuthorID = id

	err = message.Insert(d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("inserting message: %w", err)
	}

	return nil
}

func (d DB) ReadAllMessages() ([]domain.Message, error) {
	dbMessages, err := models.Messages().All(d.db)
	if err != nil {
		return nil, fmt.Errorf("reading messages from db failed: %w", err)
	}

	messages := make([]domain.Message, len(dbMessages))

	for id, dbMsg := range dbMessages {
		messages[id].Text = dbMsg.Text
		messages[id].Time = dbMsg.Time

		author, err := d.getUserByID(dbMsg.AuthorID)
		if err != nil {
			return nil, fmt.Errorf("reading author failed: %w", err)
		}

		messages[id].Author = author.Nickname
	}

	return messages, nil
}
