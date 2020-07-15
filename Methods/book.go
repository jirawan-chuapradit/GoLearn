package main

import (
	"fmt"
	"strconv"
	"time"
)

type book struct {
	product
	published interface{}
}

func (b *book) print() {
	b.product.print()
	p := format(b.published)
	fmt.Printf("\t - (%v)\n",p)
	//fmt.Printf("%-15s: %s  - (%v)\n", b.title, b.price.string(), p)
}

func format(v interface{}) string {
	var t int
	switch  v:=v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	default:
		return "unknow"
	}

	//Mon Jan 2 15:04:05 -0700 MST 2006
	const layout = "2006 - Jan"
	u := time.Unix((int64(t)),0)
	return u.Format(layout)

}
