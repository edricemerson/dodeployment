package transaction

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ProductID primitive.ObjectID `json:"product_id"`
	Quantity  int                `json:"quantity"`
	Total     int                `json:"total"`
	Status    string             `json:"status"`
}
