package repository

import (
	"context"
	"goMeesho/database"
	"goMeesho/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func SaveSMS(ctx context.Context, sms *models.SMS) error {
	collection := database.DB.Collection("sms")

	filter := bson.M{
		"_id": sms.PhoneNumber,
	}

	update := bson.M{
		"$push": bson.M{
			"messages": bson.M{
				"text":      sms.Messages[0].Text,
				"status":    sms.Messages[0].Status,
				"createdAt": sms.Messages[0].CreatedAt,
			},
		},
	}

	opts := options.UpdateOne().SetUpsert(true)

	_, err := collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func GetSMSByPhone(ctx context.Context, phone string) (*models.SMS, error) {
	collection := database.DB.Collection("sms")

	var sms models.SMS
	err := collection.FindOne(
		ctx,
		bson.M{
			"_id": phone,
		},
	).Decode(&sms)

	if err != nil {
		return nil, err
	}

	return &sms, nil
}
