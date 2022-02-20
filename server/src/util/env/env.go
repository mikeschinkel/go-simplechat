package env

import (
	"fmt"
	"os"
	"strconv"
)

var (
	// Database
	dbHost = ""
	dbPort = ""
	dbName = ""
	dbUser = ""
	dbPwd  = ""
	// Json-Web-Token
	jwtSecret = []byte(nil)
	// Cookies
	cookieName   = ""
	cookieDomain = ""
	cookiePath   = ""
	secureCookie = false
	cookieExp    = 0
)

/**
Load env file and call functions that require parsing.
*/
func Init() {
	var err error
	// Database
	dbHost = os.Getenv("DATABASE_HOST")
	dbPort = os.Getenv("DATABASE_PORT")
	dbName = os.Getenv("DATABASE_NAME")
	dbUser = os.Getenv("DATABASE_USER")
	dbPwd = os.Getenv("DATABASE_PASSWORD")
	// Json-Web-Token
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	// Cookie
	cookieName = os.Getenv("COOKIE_NAME")
	cookieDomain = os.Getenv("COOKIE_DOMAIN")
	cookiePath = os.Getenv("COOKIE_PATH")
	secureCookie, err = strconv.ParseBool(os.Getenv("SECURE_COOKIE"))
	if err != nil {
		fmt.Println(err.Error())
	}
	cookieExp, err = strconv.Atoi(os.Getenv("COOKIE_EXP"))
	if err != nil {
		fmt.Println(err.Error())
	}
}

/**
Look in daos/db.go for connection setup.
*/
func GetDbVals() (string, string, string, string, string) {
	return dbHost, dbUser, dbPwd, dbName, dbPort
}

/**
Needed in jwtUtil to sign the token and the session middleware.
*/
func JwtSecret() []byte {
	return jwtSecret
}

/**
Need in jwtUtil to sign the token. Exp jwt and same time as cookie.
*/
func JwtExp() int {
	return cookieExp
}

/**
Use to set the cookie in the auth router
*/
func GetCookieVals() (string, int, string, string, bool) {
	return cookieName, cookieExp, cookiePath, cookieDomain, secureCookie
}

/**
To look up the cookie name in the middleware.
*/
func CookieName() string {
	return cookieName
}
