package main

import (
	"context"
	"fmt"
	"log"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

func main() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Ошибка подключения:", err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal("Ошибка при отключении:", err)
		}
	}()

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Ошибка проверки подключения:", err)
	}
	fmt.Println("Успешно подключено к MongoDB!")

	collection := client.Database("cybercar_store").Collection("users")

	// Операции CRUD

	newUser := User{Name: "John Doe", Email: "john@example.com"}
	insertResult, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal("Ошибка добавления:", err)
	}
	fmt.Println("Пользователь добавлен с ID:", insertResult.InsertedID)

	// var result User
	// err = collection.FindOne(context.TODO(), bson.M{"name": "John Doe"}).Decode(&result)
	// if err != nil {
	// 	log.Fatal("Ошибка чтения:", err)
	// }
	// fmt.Printf("Найден пользователь: %+v\n", result)

	// filter := bson.M{"name": "John Doe"}
	// update := bson.M{"$set": bson.M{"email": "john.doe@example.com"}}
	// _, err = collection.UpdateOne(context.TODO(), filter, update)
	// if err != nil {
	// 	log.Fatal("Ошибка обновления:", err)
	// }
	// fmt.Println("Пользователь успешно обновлён!")

	// _, err = collection.DeleteOne(context.TODO(), bson.M{"name": "John Doe"})
	// if err != nil {
	// 	log.Fatal("Ошибка удаления:", err)
	// }
	// fmt.Println("Пользователь успешно удалён!")
}


