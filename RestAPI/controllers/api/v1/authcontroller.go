package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/Despenrado/ElCharge/RestAPI/services/api"
	"github.com/Despenrado/ElCharge/RestAPI/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"gopkg.in/gorilla/mux.v1"
)

// AuthController ...
type AuthController struct {
	service api.Service
	jwtKey  string
	rClient *redis.Client
}

type Claims struct {
	UID string `json:"_id"`
	jwt.StandardClaims
}

// NewAuthController constructor
func NewAuthController(s api.Service, rc *redis.Client) *AuthController {
	return &AuthController{
		service: s,
		rClient: rc,
	}
}

// SetJWTKey ...
func (c *AuthController) SetJWTKey(key string) {
	c.jwtKey = key
}

func (c *AuthController) CreateUser() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := &models.User{}
		err := json.NewDecoder(r.Body).Decode(u)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		u, err = c.service.User().CreateUser(u)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		utils.Respond(w, r, http.StatusCreated, u)
	})
}

func (c *AuthController) Login() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := &models.User{}
		err := json.NewDecoder(r.Body).Decode(u)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		u, err = c.service.User().Login(u)
		if err != nil {
			utils.Error(w, r, http.StatusUnauthorized, utils.ErrIncorrectEmailOrPassword)
			return
		}
		token, err := c.createTokenString(u.ID)
		if err != nil {
			utils.Error(w, r, http.StatusInternalServerError, err)
			return
		}
		// http.SetCookie(w, cookie)
		w.Header().Set("Authorization", "Bearer "+token)
		utils.Respond(w, r, http.StatusOK, u)
	})
}

func (c *AuthController) Logout() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			utils.Error(w, r, http.StatusUnauthorized, errors.New("Missing Authorization Header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(c.jwtKey), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				utils.Error(w, r, http.StatusUnauthorized, err)
				return
			}
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		if !tkn.Valid {
			utils.Error(w, r, http.StatusUnauthorized, err)
			return
		}
		_, err = c.rClient.Get(tokenString[37:]).Result()
		if err != redis.Nil {
			utils.Error(w, r, http.StatusUnauthorized, errors.New("Invalid token"))
			return
		}
		params := mux.Vars(r)
		uid, ok := params["id"]
		if !ok {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrWrongRequest)
			return
		}
		if uid != claims.UID {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrWrongRequest)
			return
		}
		c.rClient.Set(tokenString[37:], tokenString, 5*time.Minute)
		utils.Respond(w, r, http.StatusOK, nil)
	})
}

func (c *AuthController) createTokenString(uid string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		UID: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(c.jwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
	// return &http.Cookie{
	// 		Name:    "token",
	// 		Value:   tokenString,
	// 		Expires: expirationTime,
	// 	},
	// 	nil
}

func (c *AuthController) CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			utils.Error(w, r, http.StatusUnauthorized, errors.New("Missing Authorization Header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(c.jwtKey), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				utils.Error(w, r, http.StatusUnauthorized, err)
				return
			}
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		if !tkn.Valid {
			utils.Error(w, r, http.StatusUnauthorized, err)
			return
		}
		_, err = c.rClient.Get(tokenString[37:]).Result()
		if err != redis.Nil {
			utils.Error(w, r, http.StatusUnauthorized, errors.New("Invalid token"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

// func (c *AuthController) checkToken(w http.ResponseWriter, r *http.Request) error {
// 	// cookieToken, err := r.Cookie("token")
// 	// if err != nil {
// 	// 	if err == http.ErrNoCookie {
// 	// 		utils.Error(w, r, http.StatusUnauthorized, err)
// 	// 		return
// 	// 	}
// 	// }
// 	// tokenString := cookieToken.Value
// 	tokenString := r.Header.Get("Authorization")
// 	if len(tokenString) == 0 {
// 		utils.Error(w, r, http.StatusUnauthorized, errors.New("Missing Authorization Header"))
// 		return errors.New("Missing Authorization Header")
// 	}
// 	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
// 	claims := &Claims{}
// 	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(c.jwtKey), nil
// 	})
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			utils.Error(w, r, http.StatusUnauthorized, err)
// 			return err
// 		}
// 		utils.Error(w, r, http.StatusBadRequest, err)
// 		return err
// 	}
// 	if !tkn.Valid {
// 		utils.Error(w, r, http.StatusUnauthorized, err)
// 		return err
// 	}
// 	_, err = c.rClient.Get(tokenString[44:]).Result()
// 	if err != redis.Nil {
// 		utils.Error(w, r, http.StatusUnauthorized, err)
// 		return err
// 	}
// 	// // We ensure that a new token is not issued until enough time has elapsed
// 	// // In this case, a new token will only be issued if the old token is within
// 	// // 30 seconds of expiry. Otherwise, return a bad request status
// 	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 30*time.Second {
// 	// 	tokenString, err = c.renewToken(claims)
// 	// 	if err != nil {
// 	// 		utils.Error(w, r, http.StatusBadRequest, err)
// 	// 		return
// 	// 	}
// 	// 	w.Header().Set("Authorization", "Bearer "+tokenString)
// 	// }
// 	return nil
// }

func (c *AuthController) RenewToken(claims *Claims) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(c.jwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
