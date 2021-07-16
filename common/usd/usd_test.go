package usd

import (
	"fmt"
	"testing"
)

func TestUsd(t *testing.T) {
	// u := it120.
	// fmt.Println(u.GetCurrencyRate("CNY", "USD"))

	c := newSinaCurrency()
	fmt.Println(c.GetCurrencyRate("usd", "cny"))

}
