package entity

import "time"

type Transaction struct {
	ID          ID
	TypeId      ID
	Date        time.Time
	ProductName string
	Value       float64
	SellerName  string
}

func NewTransaction(typeId ID, date time.Time, productName string, value float64, sellerName string) (*Transaction, error) {
	t := &Transaction{
		ID:          NewID(),
		TypeId:      typeId,
		Date:        date,
		ProductName: productName,
		Value:       value,
		SellerName:  sellerName,
	}

	err := t.Validate()
	if err != nil {
		return nil, err
	}

	return t, nil

}

func (t *Transaction) Validate() error {
	if t.ProductName == "" || t.Value <= 0 || t.SellerName == "" {
		return ErrInvalidEntity
	}

	return nil
}
