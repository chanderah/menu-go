package middleware

type Error struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

// func ErrorHandler() gin.HandlerFunc {
// 	return func (c *gin.Context)  {
// 		c.Next()
// 		for _, err:= range c.Errors {
// 			switch e:= err.Err.Error(type) {
// 				case Error:
// 					c.AbortWithStatusJSON(e., e)
// 				default:
// 					c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "Something went wrong."})
// 			}
// 		}
// 	}
// }

