package dto

import "em-test-task/internal/domain/entity"

type BioInfoCreate struct {
	Create `json:"-" binding:"-"`

	Name       string `json:"name" db:"name" binding:"required"`
	Surname    string `json:"surname" db:"surname" binding:"required"`
	Patronymic string `json:"patronymic" db:"patronymic" binding:"required"`
	Age        int    `json:"age" db:"age" binding:"required"`
	Gender     string `json:"gender" db:"gender" binding:"required"`
	National   string `json:"national" db:"national" binding:"required"`
}

func (bIC *BioInfoCreate) ToEntity(id int) *entity.BioInfo {
	return &entity.BioInfo{
		ID:         id,
		Name:       bIC.Name,
		Surname:    bIC.Surname,
		Patronymic: bIC.Patronymic,
		Age:        bIC.Age,
		Gender:     bIC.Gender,
		National:   bIC.National,
	}
}
