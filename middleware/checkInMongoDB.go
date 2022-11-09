package middleware

import (
	"net/http"

	"github.com/nehonar/twitteringo/db"
)

/*
CheckInMongoDB check if we have connection with MongoDB before pass to next step
*/
func CheckInMongoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnectionWithPing() == db.StatusConnectionWrong {
			http.Error(w, "Lost connection with MongoDB", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
