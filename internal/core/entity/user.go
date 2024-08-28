package entity

type User struct {
	ID        uint64 `gorm:"column:id"`
	Name      string `gorm:"column:name"`
	Cellphone string `gorm:"column:cellphone"`
}

func TableName() string {
	return "users"
}
