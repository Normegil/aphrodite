package model

import (
	"github.com/satori/go.uuid"
	"encoding/json"
)

type ID struct {
	identifier uuid.UUID
}

func NewID() ID {
	return ID{uuid.NewV4()}
}

func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.identifier.String())
}

func (id *ID) UnmarshalJSON(b []byte) error {
	identifier, err := uuid.FromBytes(b)
	if nil != err {
		return err
	}
	id.identifier = identifier
	return nil
}