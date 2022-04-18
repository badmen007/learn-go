package main

import (
	"fmt"
	"io"
	"net/http"
)

// 只能读取一次
func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		// 记住要返回，不然还会执行后面的代码
		return
	}
	fmt.Fprintf(w, "read the data: %s \n", string(body))

	// 尝试再次读取， 啥也读不到， 但是不会报错
	body, err = io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		return
	}
	fmt.Fprintf(w, "read the data: %s \n", string(body))
}

func main() {

}
