package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func authMiddlewareHandler(next http.Handler) http.HandlerFunc {
	missingHeaderErr := fmt.Errorf("missing authorization header")
	missingAuthTokenErr := fmt.Errorf("missing authorization token")
	authTokenExpiredErr := fmt.Errorf("authorization token expired")
	invalidIssuerErr := fmt.Errorf("invalid issuer provided")
	invalidUseErr := fmt.Errorf("invalid token use")

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Authenticating User Request")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Println(missingHeaderErr)
			http.Error(w, missingHeaderErr.Error(), http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			log.Println(missingAuthTokenErr)
			http.Error(w, missingAuthTokenErr.Error(), http.StatusUnauthorized)
			return
		}

		token := []byte(strings.Split(authHeader, "Bearer ")[1])

		log.Println("Fetching Keys")
		set, err := jwk.Fetch(context.Background(), os.Getenv("JWK_URL"))
		if err != nil {
			log.Println("Error fetching keys: ", err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		parsedToken, err := jwt.Parse(token, jwt.WithKeySet(set), jwt.WithValidate(true))
		if err != nil {
			log.Println("error parsing token, could no authenticate: ", err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// Validade claims
		if parsedToken.Expiration().Unix() < time.Now().Unix() {
			log.Println(authTokenExpiredErr)
			http.Error(w, authTokenExpiredErr.Error(), http.StatusUnauthorized)
			return
		}

		if parsedToken.Issuer() != os.Getenv("COGNITO_ISSUER") {
			log.Println(invalidIssuerErr)
			http.Error(w, invalidIssuerErr.Error(), http.StatusUnauthorized)
			return
		}

		if parsedToken.PrivateClaims()["token_use"] != "access" {
			log.Println(parsedToken.PrivateClaims()["token_use"])
			log.Println(invalidUseErr)
			http.Error(w, invalidIssuerErr.Error(), http.StatusUnauthorized)
			return
		}

		log.Println("Access Authorized")
		next.ServeHTTP(w, r)
	}
}
