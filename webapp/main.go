package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/aren55555/learn/dec7/webapp/handlers/tweets"
	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println(redisClient)

	connStr := "host=postgres port=5432 dbname=postgres user=postgres password=hockey sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(db)

	// Migrations
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tweets (
    id SERIAL PRIMARY KEY,
    body TEXT NOT NULL
  )`)
	if err != nil {
		panic(err)
	}

	// sample data
	_, err = db.Exec(`INSERT INTO tweets (body)
  VALUES ('hello world');`)
	if err != nil {
		panic(err)
	}

	// HTTP
	tweetsHandler := &tweets.Handler{
		PG: db,
	}
	http.Handle("/tweets", tweetsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
