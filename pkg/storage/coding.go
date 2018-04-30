package storage

import (
	"bytes"
	"encoding/gob"

	"code.ysitd.cloud/auth/account/third_party/forked/github.com/RangelReale/osin"
)

func init() {
	gob.Register(map[string]interface{}{})
	gob.Register(osin.AuthorizeData{})
	gob.Register(osin.AccessData{})
	gob.Register(&osin.DefaultClient{})
}

func encode(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decode(data []byte, v interface{}) error {
	err := gob.NewDecoder(bytes.NewBuffer(data)).Decode(v)
	return err
}
