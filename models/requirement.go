package models

import "gorm.io/gorm"

// Requirement Database model info
// @Description App type information
type Requirement struct {
	*gorm.Model
	ID             uint       `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name           string     `gorm:"not null; " json:"name,omitempty"`
	Description    string     `gorm:"not null; " json:"description,omitempty"`
	Purpose        string     `gorm:"not null; " json:"purpose,omitempty"`
	Status         string     `gorm:"not null; " json:"status,omitempty"`
	BussinessValue uint       `gorm:"not null; " json:"bussiness_value,omitempty"`
	AssignedTo     string     `gorm:"not null; " json:"assigned_to,omitempty"`
	SprintID       uint       `gorm:"foreignkey:SprintID OnDelete:SET NULL" json:"sprint_id,omitempty" swaggertype:"number"`
	Tests          []Test     `gorm:"association_foreignkey:RequirementID constraint:OnUpdate:SET NULL OnDelete:SET NULL" json:"tests,omitempty"`
	Documents      []Document `gorm:"association_foreignkey:RequirementID constraint:OnUpdate:SET NULL OnDelete:SET NULL" json:"documents,omitempty"`
}

// RequirementPost model info
// @Description RequirementPost type information
type RequirementPost struct {
	Name           string `gorm:"not null; " json:"name,omitempty"`
	Description    string `gorm:"not null; " json:"description,omitempty"`
	Purpose        string `gorm:"not null; " json:"purpose,omitempty"`
	Status         string `gorm:"not null; " json:"status,omitempty"`
	BussinessValue uint   `gorm:"not null; " json:"bussiness_value,omitempty"`
	AssignedTo     string `gorm:"not null; " json:"assigned_to,omitempty"`
	SprintID       uint   `gorm:"foreignkey:SprintID OnDelete:SET NULL" json:"sprint_id,omitempty" swaggertype:"number"`
}

// RequirementGet model info
// @Description RequirementGet type information
type RequirementGet struct {
	ID             uint       `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name           string     `gorm:"not null; " json:"name,omitempty"`
	Description    string     `gorm:"not null; " json:"description,omitempty"`
	Purpose        string     `gorm:"not null; " json:"purpose,omitempty"`
	Status         string     `gorm:"not null; " json:"status,omitempty"`
	BussinessValue uint       `gorm:"not null; " json:"bussiness_value,omitempty"`
	AssignedTo     string     `gorm:"not null; " json:"assigned_to,omitempty"`
	Documents      []Document `gorm:"association_foreignkey:RequirementID constraint:OnUpdate:SET NULL OnDelete:SET NULL" json:"documents,omitempty"`
}

// RequirementPut model info
// @Description RequirementPut type information
type RequirementPut struct {
	ID             uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name           string `gorm:"not null; " json:"name,omitempty"`
	Description    string `gorm:"not null; " json:"description,omitempty"`
	Purpose        string `gorm:"not null; " json:"purpose,omitempty"`
	Status         string `gorm:"not null; " json:"status,omitempty"`
	BussinessValue uint   `gorm:"not null; " json:"bussiness_value,omitempty"`
	AssignedTo     string `gorm:"not null; " json:"assigned_to,omitempty"`
}

// RequirementPatch model info
// @Description RequirementPatch type information
type RequirementPatch struct {
	ID             uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name           string `gorm:"not null; " json:"name,omitempty"`
	Description    string `gorm:"not null; " json:"description,omitempty"`
	Purpose        string `gorm:"not null; " json:"purpose,omitempty"`
	Status         string `gorm:"not null; " json:"status,omitempty"`
	BussinessValue uint   `gorm:"not null; " json:"bussiness_value,omitempty"`
	AssignedTo     string `gorm:"not null; " json:"assigned_to,omitempty"`
}
