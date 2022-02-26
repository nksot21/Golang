package models

type BaseModel struct {
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	Deleted   bool  `json:"deleted"`
	DeletedAt int64 `json:"deleted_at"`
}
