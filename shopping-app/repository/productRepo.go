package repository

import (
	"context"

	"p3-graded-challenge-1-edricemerson/shopping-app/service/product"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	col *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{
		col: db.Collection("products"),
	}
}

func (r *ProductRepository) Create(p product.Product) (product.Product, error) {

	result, err := r.col.InsertOne(context.Background(), p)
	if err != nil {
		return p, err
	}

	p.ID = result.InsertedID.(primitive.ObjectID)

	return p, nil
}

func (r *ProductRepository) GetAll() ([]product.Product, error) {

	cursor, err := r.col.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var products []product.Product
	cursor.All(context.Background(), &products)

	return products, nil
}

func (r *ProductRepository) GetByID(id string) (product.Product, error) {

	objID, _ := primitive.ObjectIDFromHex(id)

	var p product.Product

	err := r.col.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&p)

	return p, err
}

func (r *ProductRepository) Update(id string, p product.Product) (product.Product, error) {

	objID, _ := primitive.ObjectIDFromHex(id)

	_, err := r.col.UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.M{"$set": p},
	)

	return p, err
}

func (r *ProductRepository) Delete(id string) error {

	objID, _ := primitive.ObjectIDFromHex(id)

	_, err := r.col.DeleteOne(context.Background(), bson.M{"_id": objID})

	return err
}
