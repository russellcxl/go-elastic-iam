package middlewares

import "github.com/gin-gonic/gin"

/*
	basic auth that sends user:password in base64 encoding like so,

	func authorizationHeader(user, password string) string {
		base := user + ":" + password
		return "Basic " + base64.StdEncoding.EncodeToString(bytesconv.StringToBytes(base))
	}
	
*/
func Auth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"user1": "123",
	})
}