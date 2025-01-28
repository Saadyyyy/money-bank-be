package models

type Card struct {
	CardId    int64 `gorm:"primaryKey;autoIncrement:true"`
	CardName  string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}
