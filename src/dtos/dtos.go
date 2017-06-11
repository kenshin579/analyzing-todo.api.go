package dtos

// Todo describes a TODO reminder
type Todo struct {
	TODO   string `json:"TODO" binding:"required"`
	Author string `json:"Author" binding:"required"`
	When   string `json:"When" binding:"required"`
}
