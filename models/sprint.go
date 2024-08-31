package models

import "gorm.io/gorm"

// Sprint Database model info
// @Description App type information
type Sprint struct {
	*gorm.Model
	ID           uint          `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name         string        `gorm:"not null;" json:"name,omitempty"`
	Description  string        `gorm:"not null; " json:"description,omitempty"`
	Requirements []Requirement `gorm:"association_foreignkey:RequirementID constraint:OnUpdate:SET NULL OnDelete:SET NULL" json:"requirements,omitempty"`
}

// SprintPost model info
// @Description SprintPost type information
type SprintPost struct {
	Name        string `gorm:"not null;" json:"name,omitempty"`
	Description string `gorm:"not null; " json:"description,omitempty"`
}

// SprintGet model info
// @Description SprintGet type information
type SprintGet struct {
	ID           uint          `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name         string        `gorm:"not null;" json:"name,omitempty"`
	Description  string        `gorm:"not null; " json:"description,omitempty"`
	Requirements []Requirement `gorm:"association_foreignkey:RequirementID constraint:OnUpdate:SET NULL OnDelete:SET NULL" json:"requirements,omitempty"`
}

// SprintPut model info
// @Description SprintPut type information
type SprintPut struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name        string `gorm:"not null;" json:"name,omitempty"`
	Description string `gorm:"not null; " json:"description,omitempty"`
}

// SprintPatch model info
// @Description SprintPatch type information
type SprintPatch struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name        string `gorm:"not null;" json:"name,omitempty"`
	Description string `gorm:"not null; " json:"description,omitempty"`
}
