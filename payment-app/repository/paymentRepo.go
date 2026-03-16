package repository

import (
	"context"
	"strings"

	"p3-graded-challenge-1-edricemerson/payment-app/service/payment"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	col *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		col: db.Collection("payments"),
	}
}

func (r *MongoRepository) Create(p payment.Payment) (payment.Payment, error) {

	result, err := r.col.InsertOne(context.Background(), p)
	if err != nil {
		return p, err
	}

	// assign generated MongoDB ID
	p.ID = result.InsertedID.(primitive.ObjectID)

	return p, nil
}

func (r *MongoRepository) ReadByTransactionID(txID string) (p payment.Payment, err error) {

	err = r.col.FindOne(context.Background(), bson.M{
		"transaction_id": txID,
	}).Decode(&p)

	if err != nil {
		if strings.Contains(err.Error(), "no documents") {
			err = nil
			return
		}
	}

	return
}
