package xhttp

import (
	"fmt"
	"testing"
)

func TestHttp(t *testing.T) {
	bs, err := Get("https://api.it120.cc/gooking/forex/rate?fromCode=CNY&toCode=USD", map[string]string{
		"Content-Type": "application/json",
	})
	if err == nil {
		fmt.Println(string(bs))
	}
}
