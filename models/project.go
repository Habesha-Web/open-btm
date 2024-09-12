package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"open-btm.com/utils"
)

// Project Database model info
// @Description App type information
type Project struct {
	*gorm.Model
	ID           uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name         string `gorm:"not null; unique; " json:"name,omitempty"`
	DatabaseName string `gorm:"not null; unique; " json:"database_name,omitempty"`
	Description  string `gorm:"not null; " json:"description,omitempty"`
	UUID         string `gorm:"constraint:not null; unique; type:string;" json:"uuid"`
}

type ProjectUsers struct {
	ID        uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	UserUUID  string `gorm:"not null; " json:"user_uuid,omitempty"`
	ProjectID uint   `gorm:"foreignkey:ProjectID OnDelete:SET NULL" json:"project_id,omitempty" swaggertype:"number"`
}

func (project *Project) BeforeCreate(tx *gorm.DB) (err error) {
	gen, _ := uuid.NewV7()
	id := gen.String()
	project.UUID = id
	dabname, _ := utils.GenerateRandomString(7)
	project.DatabaseName = dabname
	return
}
