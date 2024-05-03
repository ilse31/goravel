package pkg

import (
	"net/http"

	rip "github.com/vikram1565/request-ip"
)

// func GetIPAdress to get the ip address
func GetIPAdress(r *http.Request) string {
	return rip.GetClientIP(r)
}
