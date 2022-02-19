package env

import (
	"fmt"
	"os"
	"strconv"
)

var (
	// Database
	dbHost = os.Getenv("DATABASE_HOST")
	dbPort = os.Getenv("DATABASE_PORT")
	dbName = os.Getenv("DATABASE_NAME")
	dbUser = os.Getenv("DATABASE_USER")
	dbPwd  = os.Getenv("DATABASE_PASSWORD")
	// Authentication
	jwtSecret    = []byte(os.Getenv("JWT_SECRET"))
	cookieName   = os.Getenv("COOKIE_NAME")
	cookieDomain = os.Getenv("COOKIE_DOMAIN")
	cookiePath   = os.Getenv("COOKIE_PATH")
	secureCookie = setSecureCookie()
	cookieExp    = setCookieExp()
)

func setSecureCookie() bool {
	secureCookie, err := strconv.ParseBool(os.Getenv("SECURE_COOKIE"))
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	return secureCookie
}

func setCookieExp() int {
	cookieExp, err := strconv.Atoi(os.Getenv("COOKIE_EXP"))
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return cookieExp
}

func GetDbVals() (string, string, string, string, string) {
	return dbHost, dbUser, dbPwd, dbName, dbPort
}

func JetSecret() []byte {
	return jwtSecret
}

func GetCookieVals() (string, int, string, string, bool) {
	return cookieName, cookieExp, cookiePath, cookieDomain, secureCookie
}

func CookieName() string {
	return cookieName
}

func CookieExp() int {
	return cookieExp
}
