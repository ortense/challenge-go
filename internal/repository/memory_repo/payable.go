package memory_repo

import (
	"errors"

	"github.com/ortense/challenge-go/internal/payable"
)

type Payable struct {
	Data map[string]payable.Payable
	Err  error
}

func NewPayablenMemoryRepo() Payable {
	return Payable{
		Data: map[string]payable.Payable{},
		Err:  nil,
	}
}

func (r *Payable) Save(p payable.Payable) error {
	id := p.Id.String()
	_, exist := r.Data[id]

	if r.Err != nil {
		return r.Err
	}

	if exist {
		return errors.New("payable id conflict")
	}

	r.Data[id] = p

	return nil
}

func (r *Payable) List() ([]payable.Payable, error) {
	list := []payable.Payable{}

	if r.Err != nil {
		return list, r.Err
	}

	for _, p := range r.Data {
		list = append(list, p)
	}

	return list, nil
}
