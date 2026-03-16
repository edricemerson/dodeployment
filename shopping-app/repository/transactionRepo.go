package repository

import (
	"context"

	"p3-graded-challenge-1-edricemerson/shopping-app/service/transaction"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepository struct {
	col *mongo.Collection
}

func NewTransactionRepository(db *mongo.Database) *TransactionRepository {
	return &TransactionRepository{
		col: db.Collection("transactions"),
	}
}

func (r *TransactionRepository) Create(t transaction.Transaction) (transaction.Transaction, error) {

	result, err := r.col.InsertOne(context.Background(), t)
	if err != nil {
		return t, err
	}

	t.ID = result.InsertedID.(primitive.ObjectID)

	return t, nil
}

func (r *TransactionRepository) GetAll() ([]transaction.Transaction, error) {

	cursor, err := r.col.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var transactions []transaction.Transaction
	cursor.All(context.Background(), &transactions)

	return transactions, nil
}

func (r *TransactionRepository) GetByID(id string) (transaction.Transaction, error) {

	objID, _ := primitive.ObjectIDFromHex(id)

	var t transaction.Transaction

	err := r.col.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&t)

	return t, err
}

func (r *TransactionRepository) Update(id string, t transaction.Transaction) (transaction.Transaction, error) {

	objID, _ := primitive.ObjectIDFromHex(id)

	_, err := r.col.UpdateOne(context.Background(),
		bson.M{"_id": objID},
		bson.M{"$set": t},
	)

	return t, err
}

func (r *TransactionRepository) Delete(id string) error {

	objID, _ := primitive.ObjectIDFromHex(id)

	_, err := r.col.DeleteOne(context.Background(), bson.M{"_id": objID})

	return err
}
