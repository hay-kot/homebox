package mid

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/hay-kot/homebox/backend/pkgs/server"
)

// Panic is a middleware that recovers from panics anywhere in the chain and wraps the error.
// and returns it up the middleware chain.
func Panic(develop bool) server.Middleware {
	return func(h server.Handler) server.Handler {
		return server.HandlerFunc(func(w http.ResponseWriter, r *http.Request) (err error) {
			defer func() {
				if rec := recover(); rec != nil {
					trace := debug.Stack()

					if develop {
						err = fmt.Errorf("PANIC [%v]", rec)
						fmt.Printf("%s", string(trace))
					} else {
						err = fmt.Errorf("PANIC [%v] TRACE[%s]", rec, string(trace))
					}

				}
			}()

			return h.ServeHTTP(w, r)
		})
	}
}
