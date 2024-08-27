package models

import (
	"gorm.io/gorm"
)

// Requirement Database model info
// @Description App type information
type Requirement struct {
	*gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name        string `gorm:"not null; unique;" json:"name,omitempty"`
	Description string `gorm:"not null; unique;" json:"description,omitempty"`
	Tests       []Test `gorm:"association_foreignkey:RequirementID constraint:OnUpdate:SET NULL OnDelete:SET NULL" json:"tests,omitempty"`
}

// RequirementPost model info
// @Description RequirementPost type information
type RequirementPost struct {
	Name        string `gorm:"not null; unique;" json:"name,omitempty"`
	Description string `gorm:"not null; unique;" json:"description,omitempty"`
}

// RequirementGet model info
// @Description RequirementGet type information
type RequirementGet struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name        string `gorm:"not null; unique;" json:"name,omitempty"`
	Description string `gorm:"not null; unique;" json:"description,omitempty"`
	Tests       []Test `gorm:"association_foreignkey:RequirementID constraint:OnUpdate:SET NULL OnDelete:SET NULL" json:"tests,omitempty"`
}

// RequirementPut model info
// @Description RequirementPut type information
type RequirementPut struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name        string `gorm:"not null; unique;" json:"name,omitempty"`
	Description string `gorm:"not null; unique;" json:"description,omitempty"`
}

// RequirementPatch model info
// @Description RequirementPatch type information
type RequirementPatch struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name        string `gorm:"not null; unique;" json:"name,omitempty"`
	Description string `gorm:"not null; unique;" json:"description,omitempty"`
}
