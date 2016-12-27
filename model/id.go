package model

import (
	"encoding/json"

	"github.com/satori/go.uuid"
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
