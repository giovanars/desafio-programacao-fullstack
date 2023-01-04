package entity

type TransactionType struct {
	ID   ID
	Name string
}

func NewTransactionType(name string) (*TransactionType, error) {
	t := &TransactionType{
		ID:   NewID(),
		Name: name,
	}

	err := t.Validate()
	if err != nil {
		return nil, err
	}

	return t, nil

}

func (t *TransactionType) Validate() error {
	if t.Name == "" {
		return ErrInvalidEntity
	}

	return nil
}
