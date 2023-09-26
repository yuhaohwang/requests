//go:build ignore
// +build ignore

package examples

import (
	"fmt"

	"github.com/yuhaohwang/requests"
)

func JSONExample() {
	resp, err := requests.Post("http://example.com", requests.JSON(map[string]string{"foo": "bar"}))
	if err != nil {
		panic(err)
	}
	m := make(map[string]interface{})
	if err := resp.JSON(&m); err != nil {
		panic(err)
	}
	fmt.Println(m)
}
