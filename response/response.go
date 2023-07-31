package response

type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func Ok (data interface{}) {
	// c.JSON(http.StatusOK, gin.H{"message":"success", data})
}