package models

type Users struct {
	UserId    int64 `gorm:"primaryKey;autoIncrement:true"`
	Username  string
	Password  string
	Role      int64
	Token     string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}
