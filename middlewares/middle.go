package middlewares

import "net/http"

type Middle func(http.HandlerFunc) http.HandlerFunc

func Handle(hd http.HandlerFunc, mid ...Middle) http.HandlerFunc {
	for i := len(mid); i > 0; i-- {
		hd = mid[i-1](hd)
	}

	return hd
}
