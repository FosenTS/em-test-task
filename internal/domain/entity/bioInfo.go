package entity

type BioInfo struct {
	Entity `json:"-" db:"-" binding:"-"`

	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Surname    string `json:"surname" db:"surname"`
	Patronymic string `json:"patronymic" db:"patronymic"`
	Age        int    `json:"age" db:"age"`
	Gender     string `json:"gender" db:"gender"`
	National   string `json:"national" db:"national"`
}
