package database

import (
	"apertursGin/graph/model"
	"apertursGin/util"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/joho/godotenv/autoload"
)

type DB struct {
	client *mongo.Client
	database *mongo.Database
	accountCollection  *mongo.Collection
} 
func Connect() *DB{
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoURL := os.Getenv("MONGO_DB_URL")
	log.Println(mongoURL)
	if mongoURL==""{
		log.Fatal("Error laoding Mongo url from env ")
	}
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err!=nil{
		log.Fatalln(err.Error() + "This is the error for mongo db connect")
	}
	database := client.Database("aperturs")
	
	return &DB{
		client: client,
		database: database,
		accountCollection: database.Collection("accounts"),

	}
}


func (db *DB) AccountCreate(ctx context.Context , input model.NewAccount) (*model.Account,error){
	if input.Email == nil && input.Password==nil { 
		return nil,fmt.Errorf("Password and / or email is null")
	}
	password := *input.Password
	newPassword := util.HashPassword(password)
	input.Password = &newPassword

	account := model.Account{
		ID: primitive.NewObjectID(),
		Email: *input.Email,
		Role: model.RoleTypeUser,
		Password: *input.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_,err := db.accountCollection.InsertOne(ctx,account)
	if err!=nil{
		return nil,err
	}
	return &account,nil
	
}
func (db *DB) UserGetByID(ctx context.Context, id string) (*model.Account, error) {
	objectId,_ := primitive.ObjectIDFromHex(id)
	var account model.Account
	err := db.accountCollection.FindOne(ctx,bson.M{"_id":objectId}).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil,err;
		}
		panic(err)
	}
	return &account,nil
}
func (db *DB) UserGetByEmail(ctx context.Context, email string) (*model.Account, error) {
	var account model.Account
	err := db.accountCollection.FindOne(ctx,bson.M{"email":email}).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil,err;
		}
		panic(err)
	}
	return &account,nil
}