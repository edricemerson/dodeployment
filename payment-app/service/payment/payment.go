package payment

type Repository interface {
	Create(p Payment) (Payment, error)
	ReadByTransactionID(txID string) (p Payment, err error)
}
