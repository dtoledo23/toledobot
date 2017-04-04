package bot

import (
	"app/dbProvider"
	"log"
	"strings"

	"fmt"

	"gopkg.in/mgo.v2/bson"
)

const (
	userIDKey = "userid"
)

func updateAndGetContext(message string, userID string) *botSession {
	createSessionIfMissing(userID)
	updateContext(message, userID)

	dbSessions := dbProvider.GetConnection().C(botSessionsCollection)

	var session botSession
	err := dbSessions.Find(bson.M{userIDKey: userID}).One(&session)

	if err != nil {
		log.Println(fmt.Errorf("Error retrieving session: %s", err))
	}

	return &session
}

func createSessionIfMissing(userID string) {
	var session botSession
	dbSessions := dbProvider.GetConnection().C(botSessionsCollection)
	err := dbSessions.Find(bson.M{userIDKey: userID}).One(&session)

	if err != nil {
		log.Println("Creating new session for user:", userID)
		session = botSession{UserID: userID}
		dbSessions.Insert(session)
	}
}

func updateContext(message string, userID string) error {
	product := getProductCommand(message)

	if product != "" {
		dbSessions := dbProvider.GetConnection().C(botSessionsCollection)
		_, err := dbSessions.Upsert(bson.M{userIDKey: userID},
			bson.M{"context": map[string]string{"product": product},
				userIDKey: userID,
			})

		if err != nil {
			return fmt.Errorf("Error updating context for userId %s: %s", userID, err)
		}

		log.Printf("Updated context for user %s. Added product '%s' in context", userID, product)
	}

	return nil
}

// Catches if the message is a command to set a product in the context and returns such product
func getProductCommand(message string) string {
	message = strings.Trim(message, "!.,? ")
	product := strings.TrimPrefix(message, productInfoPrefix)

	if product == message {
		return ""
	}

	return strings.Trim(product, "!.,? ")
}
