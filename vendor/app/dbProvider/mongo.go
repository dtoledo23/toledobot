package dbProvider

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"gopkg.in/mgo.v2"
)

const (
	connectionTimeout = "30"

	// Database schemas
	mongo = "mongodb"
)

var (
	session *mgo.Session
	db      *mgo.Database
)

// GetConnection returns a reference to the database driver to enable executing queries.
func GetConnection() *mgo.Database {
	if db == nil {
		Connect()
	}
	return db
}

// GetSession returns a reference to the session for accessing multiple databases on the given config.
func GetSession() *mgo.Session {
	if session == nil {
		Connect()
	}
	return session
}

// Connect creates a fresh connection to database.
func Connect() {
	session, db = connect()
}

func connect() (*mgo.Session, *mgo.Database) {
	scheme := mongo
	dbName := os.Getenv("MONGO_NAME")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	user := os.Getenv("MONGO_USER")
	pass := os.Getenv("MONGO_PASS")

	connectionURL := url.URL{
		Scheme:   scheme,
		User:     url.UserPassword(user, pass),
		Host:     host + ":" + port,
		Path:     dbName,
		RawQuery: additionalConfig(),
	}
	// sql.Open("postgres", "postgres://toledo:toledo@localhost/toledo?sslmode=disable")
	session, err := mgo.Dial(connectionURL.String())

	if err != nil {
		log.Fatal(fmt.Errorf("Invalid connection arguments. Using connection url '%s' produced error %s", connectionURL.String(), err))
		return nil, nil
	}

	err = session.Ping()

	if err != nil {
		log.Fatal(fmt.Errorf("Could not connect to postgres. Using connection url '%s' produced error %s", connectionURL.String(), err))
		return nil, nil
	}

	log.Println("Connected succesfully to db:", connectionURL.String())

	// Database name provided on connection url.
	db := session.DB("")
	return session, db
}

func additionalConfig() string {
	values := url.Values{}
	return values.Encode()
}
