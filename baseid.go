package baseid

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/google/uuid"
)

type BaseID struct {
	uuid.UUID
}

func New() *BaseID {
	return &BaseID{
		UUID: uuid.New(),
	}
}

func (id *BaseID) Bytes() []byte {
	return []byte(id.UUID[:])
}

func (id *BaseID) Hex() string {
	return hex.EncodeToString(id.Bytes())
}

func (id *BaseID) Base64() string {
	return base64.StdEncoding.EncodeToString(id.Bytes())
}
