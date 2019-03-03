package warehouse

type Warehouse interface {

	PutFood(string) error
	PutProduct() error

	TakeFood() error
	TakeProduct() error

	DumpExpiredFood()
}