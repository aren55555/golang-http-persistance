package models

import (
	"database/sql"
)

type Tweet struct {
	ID   int64  `json:"id"`
	Body string `json:"body"`
}

func AllTweets(db *sql.DB) (tweets []Tweet, err error) {
	tweets = []Tweet{}

	rows, err := db.Query("SELECT id, body FROM tweets")
	if err != nil {
		return
	}

	for rows.Next() {
		tweet := Tweet{}
		if err = rows.Scan(&tweet.ID, &tweet.Body); err != nil {
			return
		}
		tweets = append(tweets, tweet)
	}

	return
}
