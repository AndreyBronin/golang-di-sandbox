package core

import (
	"time"
)

type Object interface {
	Name() string
	Weight() uint32
}

type Food interface {
	Object

	Expiration() time.Time
}

type Product interface {
	Object
}
