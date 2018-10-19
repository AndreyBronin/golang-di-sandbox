

package warehouse

type Warehouse interface {

	PutFood() error
	PutProduct() error

	TakeFood() error
	TakeProduct() error

	DumpExpiredFood()
}