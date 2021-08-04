package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type key int

var userKey key

func main() {
	http.HandleFunc("/", hand)
	http.ListenAndServe(":8080", nil)
}

func hand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("x-waterk-ip", r.RemoteAddr)
	w.WriteHeader(500)

	ctx := context.Background()
	ctxDerived, cancelFun := context.WithTimeout(ctx, time.Second*5)
	ctxVal := context.WithValue(ctxDerived, userKey, "Rafael")

	go writeTerminal(ctxVal, "helloooo")

	time.Sleep(time.Second * 2)
	cancelFun()

	w.Write([]byte("hello my friend"))
}

func writeTerminal(ctx context.Context, phrase string) {
	fmt.Println(ctx)
	value, ok := ctx.Value(userKey).(string)
	if ok != true {
		fmt.Println("deu erro na assertion")
	}
	fmt.Println(value)
	select {
	case <-ctx.Done():
		fmt.Println("Caiu no done")
	}
	fmt.Println("hello")
}
