// +build !test

package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/golang/glog"
)

func josePart(v interface{}) string {
	switch t := v.(type) {
	case []byte:
		return joseBase64UrlEncode(t)
	case Header:
		return joseBase64UrlEncode(mustMarshal(t))
	case ClaimSet:
		return joseBase64UrlEncode(mustMarshal(t))
	default:
		panic(fmt.Errorf("Could not convert to jose part %f", t))
	}
}

func joseBase64UrlEncode(b []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(b), "=")
}

func mustMarshal(v interface{}) []byte {
	bs, err := json.Marshal(v)
	if err != nil {
		glog.Errorf("unable to marshal %t: %s", v, err)
		panic(err)
	}
	return bs
}
