package model

type File struct {
	Dest string `json:"dest" binding:"required"`
	File string `json:"file" binding:"required"`
}
