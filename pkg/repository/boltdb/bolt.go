package boltdb

import (
	"errors"
	"strconv"

	"github.com/bekzod003/pocket-telegram-bot/pkg/repository"
	"github.com/boltdb/bolt"
)

type TokenRepository struct {
	db *bolt.DB
}

// Constructor
func NewTokenRepository(db *bolt.DB) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}

func (r *TokenRepository) Save(chatId int64, token string, bucket repository.Bucket) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put([]byte(intToByte(chatId)), []byte(token))
	})
}

func (r *TokenRepository) Get(chatId int64, bucket repository.Bucket) (string, error) {
	var token string
	err := r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		data := b.Get([]byte(intToByte(chatId)))
		token = string(data)
		return nil
	})
	if err != nil {
		return "", err
	}

	if token == "" {
		return "", errors.New("token is not found")
	}

	return token, nil
}

func intToByte(n int64) []byte {
	return []byte(strconv.FormatInt(n, 10))
}
