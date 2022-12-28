package httpx

import (
	"log"
	"net/http"
)

type MiddlewareFunc func(w http.ResponseWriter, r *http.Request) error

func Uses(ms ...MiddlewareFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, m := range ms {
			if err := m(w, r); err != nil {
				switch e := err.(type) {
				case Error:
					http.Error(w, e.Error(), e.Code())
				default:
					log.Println(err)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
				return
			} else {
				if w.Header().Values("Location") != nil {
					return
				}
			}
		}
	}
}
