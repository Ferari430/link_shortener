package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		wrapper := &WrapperWriter{ResponseWriter: w,
			Code: http.StatusOK,
		}

		start := time.Now()

		next.ServeHTTP(wrapper, r)
		log.Println("Запрос выполнился за", time.Since(start), r.Method, r.URL.Path, wrapper.Code)
		log.Println("----------------------------------------------------")

	})
}
