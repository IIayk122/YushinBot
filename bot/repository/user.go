package repository

import (
	"context"
	"fmt"

	"gopkg.in/tucnak/telebot.v2"
)

func SetUser(user telebot.User) error {
	collection := DB.Collection("Users")
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return fmt.Errorf("Дубль:\n %s", err)
	}
	return nil
}
