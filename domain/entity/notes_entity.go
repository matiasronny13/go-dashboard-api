package entity

type Notes struct {
	ID    int `gorm:"primaryKey"`
	Title string
}

func (Notes) TableName() string {
	return "rss"
}
