package entity

type Users struct {
	ID    int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name  string `json:"name" gorm:"size:30;not null"`
	Email string `json:"email" gorm:"size:300;not null"`
}
