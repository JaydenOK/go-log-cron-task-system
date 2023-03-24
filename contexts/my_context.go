package contexts

import "net/http"

type MyContext struct {
	W http.ResponseWriter
	R *http.Request
}

func New(w http.ResponseWriter, r *http.Request) MyContext {
	myCtx := MyContext{
		W: w,
		R: r,
	}
	return myCtx
}
