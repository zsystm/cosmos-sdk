package math

import "cosmossdk.io/collections"

var IntValue collections.ValueCodec[Int] = intValueCodec{}

type intValueCodec struct{}

func (intValueCodec) Encode(value Int) ([]byte, error) {
	return value.Marshal()
}

func (intValueCodec) Decode(b []byte) (Int, error) {
	i := new(Int)
	err := i.Unmarshal(b)
	if err != nil {
		return Int{}, err
	}
	return *i, nil
}

func (intValueCodec) Stringify(value Int) string {
	return value.String()
}

func (i intValueCodec) ValueType() string {
	return "cosmossdk.io/math.Int"
}
