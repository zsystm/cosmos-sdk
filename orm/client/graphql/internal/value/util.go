package value

import "encoding/base64"

func DecodeBase64Scalar(value string) []byte {
	bytes, err := base64.RawURLEncoding.DecodeString(value)
	if err != nil {
		return nil
	}
	return bytes
}
