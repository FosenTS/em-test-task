package entity

type BioInfo struct {
	Entity `json:"-" db:"-" binding:"-"`

	ID         int    `json:"id" db:"id" binding:"required"`
	Name       string `json:"name" db:"name" binding:"required"`
	Surname    string `json:"surname" db:"surname" binding:"required"`
	Patronymic string `json:"patronymic" db:"patronymic" binding:"required"`
	Age        int    `json:"age" db:"age" binding:"required"`
	Gender     string `json:"gender" db:"gender" binding:"required"`
	National   string `json:"national" db:"national" binding:"required"`
}
