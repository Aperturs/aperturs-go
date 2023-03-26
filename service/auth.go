package service

import (
	"apertursGin/database"
	"apertursGin/graph/model"
	"apertursGin/util"
	"context"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.mongodb.org/mongo-driver/mongo"
)
var db = database.Connect()
func UserRegister(ctx context.Context, input model.NewAccount) (*model.LoginResponse, error) {
	// Check Email
	email := *input.Email
	_, err := db.UserGetByEmail(ctx,email)
	if err == nil {
		// if err != record not found
		if err != mongo.ErrNilDocument {
			return nil, err
		}
	}
	createdAccount, err := db.AccountCreate(ctx, input)
	if err != nil {
		return nil, err
	}

	token, err := JwtGenerate(ctx, createdAccount.ID.Hex())
	if err != nil {
		return nil, err
	}
	response :=  model.LoginResponse{
		Token: token,
		ID: createdAccount.ID.Hex(),
	}
	return &response, nil
}

func UserLogin(ctx context.Context, email string, password string) (*model.LoginResponse, error) {
	account, err := db.UserGetByEmail(ctx, email)
	if err != nil {
		// if user not found
		if err == mongo.ErrNilDocument {
			return nil, &gqlerror.Error{
				Message: "Email not found",
			}
		}
		return nil, err
	}

	if err := util.ComparePassword(account.Password, password); err != nil {
		return nil, err
	}

	token, err := JwtGenerate(ctx, account.ID.Hex())
	if err != nil {
		return nil, err
	}

	response :=  model.LoginResponse{
		Token: token,
		ID: account.ID.Hex(),
	}
	return &response, nil
}