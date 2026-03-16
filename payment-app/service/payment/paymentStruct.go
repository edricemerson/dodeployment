package payment

import "go.mongodb.org/mongo-driver/bson/primitive"

type Payment struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TransactionID string             `json:"transaction_id"`
	Amount        int                `json:"amount"`
	Status        string             `json:"status"`
}
