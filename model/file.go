package model

type FileDetails struct {
	Name string `json:"name" binding:"required"`
	Dest string `json:"dest" binding:"required"`
	File string `json:"file" binding:"required"`
}
