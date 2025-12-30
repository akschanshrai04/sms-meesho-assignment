package services

import (
	"context"
	"goMeesho/models"
	"goMeesho/repository"
)

func GetSMSByPhone(ctx context.Context, phone string) (*models.SMS, error) {
	return repository.GetSMSByPhone(ctx, phone)
}

func SaveSMS(ctx context.Context, sms *models.SMS) error {
	return repository.SaveSMS(ctx, sms)
}
