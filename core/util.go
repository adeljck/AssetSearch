package core

import "encoding/base64"

func ToBase64(query string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(query))
}
