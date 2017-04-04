package bot

import (
	"app/dbProvider"
	"app/platforms"
	"log"
	"os"

	"strconv"

	"fmt"

	"gopkg.in/mgo.v2/bson"
)

const (
	botAnswersCollection  = "bot"
	botSessionsCollection = "sessions"
	botProductsCollection = "products"

	witAPI              = "api.wit.ai"
	productKeyInContext = "product"
	productInfoPrefix   = "Quiero informacion de "
)

// HandleMessage defines what to do when receiving a message.
func HandleMessage(message, userID string, sender platforms.Sender) {
	witProductsToken := os.Getenv("WIT_PRODUCTS_TOKEN")
	witStoresToken := os.Getenv("WIT_STORES_TOKEN")

	// Update and retrieve user specific session.
	session := updateAndGetContext(message, userID)
	if getProductCommand(message) != "" {
		sender.SendText("Ok! Preguntame lo que quieras.")
		return
	}

	// Botflow controll
	var answer *botAnswer
	var entities map[string]string

	// Answer product based
	if product, exists := session.Context[productKeyInContext]; exists {
		entities = getEntities(message, witProductsToken)
		log.Println("Found product entities:", entities)
		answer = getProductsAnswer(product, entities)
		log.Println("Found product answer:", answer)
	}

	// Product based answer might not be applicable due to:
	// - There is no product set in user context.
	// - There was no answer found.
	// If that happends. We will try to retrieve a store based answer.
	if answer == nil {
		entities = getEntities(message, witStoresToken)
		log.Println("Found stores entities:", entities)
		answer = getStoresAnswer(entities)
	}

	// If not answer found at this point. We got nothing to di.
	if answer == nil {
		sender.SendText("Sorry! I could not understand")
		return
	}

	switch answer.Type {
	case "text":
		sender.SendText(answer.Payload)
	case "image":
		sender.SendImage(answer.Payload)
	default:
		log.Println("Unsupported answer type:", answer.Type)
	}
}

func getStoresAnswer(entities map[string]string) *botAnswer {
	db := dbProvider.GetConnection()

	var answer botAnswer
	err := db.C(botAnswersCollection).Find(bson.M{"entities": entities}).One(&answer)

	if err != nil {
		log.Println("Error retrieving bot answer: ", err)
		return nil
	}

	return &answer
}

func getProductsAnswer(product string, entities map[string]string) *botAnswer {
	var answerPayload string
	attribute := entities["intent"]
	attributeInfo := getProductInfo(product, attribute)
	log.Println("Attribute info:", attributeInfo)
	switch attribute {
	case "cost":
		answerPayload = fmt.Sprintf("El costo es de %s", attributeInfo)

	case "store":
		answerPayload = fmt.Sprintf("Puedes encontrarlo en las sucursales %s", attributeInfo)

	case "description":
		answerPayload = fmt.Sprintf("%s", attributeInfo)

	case "weight":
		answerPayload = fmt.Sprintf("El peso de la pieza es de %s", attributeInfo)

	case "carats":
		answerPayload = fmt.Sprintf("El kilataje es de %s", attributeInfo)

	case "availability":
		answerPayload = fmt.Sprintf("Tenemos %s disponibles", attributeInfo)

	case "wholesale":
		answerPayload = fmt.Sprintf("El precio de mayoreo es de %s en su compra de mas de 10 piezas.", attributeInfo)

	case "special_event":
		answerPayload = fmt.Sprintf("Perfecto para %s", attributeInfo)

	default:
		return nil
	}

	return &botAnswer{
		Entities: entities,
		Payload:  answerPayload,
		Type:     "text",
	}
}

func getProductInfo(product, attribute string) string {
	dbProducts := dbProvider.GetConnection().C(botProductsCollection)

	var foundProduct map[string]interface{}
	dbProducts.Find(bson.M{"name": product}).One(&foundProduct)

	info := foundProduct[attribute]
	log.Println("Info:", info)
	if floatInfo, ok := info.(float64); ok {
		return strconv.FormatFloat(floatInfo, 'f', -1, 64)
	} else if intInfo, ok := info.(int); ok {
		return strconv.FormatInt(int64(intInfo), 10)
	} else if stringInfo, ok := info.(string); ok {
		return stringInfo
	}
	return ""
}
