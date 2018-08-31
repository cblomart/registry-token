package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/golang/glog"
)

// JosePart return the object as a JWT JOSE part.
func JosePart(v interface{}) string {
	switch t := v.(type) {
	case []byte:
		return JoseBase64UrlEncode(t)
	case *Header:
		return JoseBase64UrlEncode(MustMarshal(t))
	case *ClaimSet:
		return JoseBase64UrlEncode(MustMarshal(t))
	default:
		panic(fmt.Errorf("Could not convert to jose part %v", t))
	}
}

// JoseBase64UrlEncode encodes a byte arrays as Base64 for JWT
func JoseBase64UrlEncode(b []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(b), "=")
}

// MustMarshal marshall an objects to json not tollerating issues
func MustMarshal(v interface{}) []byte {
	bs, err := json.Marshal(v)
	if err != nil {
		glog.Errorf("unable to marshal %t: %s", v, err)
		panic(err)
	}
	return bs
}
