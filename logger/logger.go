package logger

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
)

func Println(ctx context.Context, msg ...string) {
	v, ok := ctx.Value("requestID").(int64)
	if !ok {
		fmt.Println("requsetID not found")
		return
	}
	fmt.Printf("[%v] - > %s \n", v, msg)
}

func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := rand.Int63()
		ctx := r.Context()
		ctx = context.WithValue(ctx, "requestID", reqID)
		f(w, r.WithContext(ctx))
	}
}
