package hashid

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speps/go-hashids"
)

const (
	_salt      = "qwevf34tfgrbT$%YHb4%$Y&$G@T$%Y$Tg"
	_minLength = 16
	_alphabt   = "qwertyuiopasdfghjklzxcvbnm1234567890"
)

var seed *hashids.HashID

func init() {
	idData := hashids.NewData()
	idData.Salt = _salt
	idData.MinLength = _minLength
	idData.Alphabet = _alphabt
	var err error
	seed, err = hashids.NewWithData(idData)
	if err != nil {
		panic("initialize hash id failed:" + err.Error())
	}

}

type ID int64

func (i ID) MarshalText() (text []byte, err error) {
	if i < 0 {
		return nil, errors.New("err: ID(%d) must greater than 0")
	}
	if i == 0 {
		return
	}
	str, err := seed.EncodeInt64([]int64{int64(i)})
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func (i *ID) UnmarshalText(text []byte) error {
	is, err := seed.DecodeInt64WithError(string(text))
	if err != nil {
		return fmt.Errorf("unmarshal id failed, %w", err)
	}
	if len(is) != 1 {
		return fmt.Errorf("bad unmarshal id length, %d", len(is))
	}
	*i = ID(is[0])
	return nil
}

func (i ID) MarshalJSON() ([]byte, error) {
	text, err := i.MarshalText()
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(text))
}

func (i *ID) UnmarshalJSON(data []byte) error {
	var text string
	err := json.Unmarshal(data, &text)
	if err != nil {
		return err
	}
	return i.UnmarshalText([]byte(text))
}
