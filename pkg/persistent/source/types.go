package source

import "github.com/STRRL/growth-of-codes/pkg/persistent/entity"

type Iterator interface {
	HasNext() (bool, error)
	Next() (*entity.FileComplexitySnapshot, error)
}

type Interface interface {
	FetchIterator() (Iterator, error)
	ListAll() ([]entity.FileComplexitySnapshot, error)
	Close() error
}
