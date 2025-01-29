package models

import (
	"time"

	"gorm.io/gorm"
)

// Sprint Database model info
// @Description App type information
type Sprint struct {
	*gorm.Model
	ID           uint          `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name         string        `gorm:"not null;" json:"name,omitempty"`
	Description  string        `gorm:"not null; " json:"description,omitempty"`
	Status       string        `gorm:"not null; " json:"status,omitempty"`
	StartDate    string        `json:"start_date,omitempty"`
	Duration     uint          `json:"duration,omitempty"`
	Requirements []Requirement `gorm:"association_foreignkey:SprintID constraint:OnUpdate:SET NULL OnDelete:SET NULL" json:"requirements,omitempty"`
	Documents    []Document    `gorm:"association_foreignkey:SprintID constraint:OnUpdate:SET NULL OnDelete:SET NULL" json:"documents,omitempty"`
}

// SprintPost model info
// @Description SprintPost type information
type SprintPost struct {
	Name        string    `gorm:"not null;" json:"name,omitempty"`
	Description string    `gorm:"not null; " json:"description,omitempty"`
	Status      string    `gorm:"not null; " json:"status,omitempty"`
	StartDate   time.Time `json:"start_date,omitempty"`
	Duration    uint      `json:"duration,omitempty"`
}

// SprintGet model info
// @Description SprintGet type information
type SprintGet struct {
	ID           uint          `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name         string        `gorm:"not null;" json:"name,omitempty"`
	Description  string        `gorm:"not null; " json:"description,omitempty"`
	Status       string        `gorm:"not null; " json:"status,omitempty"`
	StartDate    time.Time     `json:"start_date,omitempty"`
	Duration     uint          `json:"duration,omitempty"`
	Requirements []Requirement `gorm:"association_foreignkey:SprintID constraint:OnUpdate:SET NULL OnDelete:SET NULL" json:"requirements,omitempty"`
	Documents    []Document    `gorm:"association_foreignkey:SprintID constraint:OnUpdate:SET NULL OnDelete:SET NULL" json:"documents,omitempty"`
}

// SprintPut model info
// @Description SprintPut type information
type SprintPut struct {
	ID          uint      `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name        string    `gorm:"not null;" json:"name,omitempty"`
	Description string    `gorm:"not null; " json:"description,omitempty"`
	Status      string    `gorm:"not null; " json:"status,omitempty"`
	StartDate   time.Time `json:"start_date,omitempty"`
	Duration    uint      `json:"duration,omitempty"`
}

// SprintPatch model info
// @Description SprintPatch type information
type SprintPatch struct {
	ID          uint      `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name        string    `gorm:"not null;" json:"name,omitempty"`
	Description string    `gorm:"not null; " json:"description,omitempty"`
	Status      string    `gorm:"not null; " json:"status,omitempty"`
	StartDate   time.Time `json:"start_date,omitempty"`
	Duration    uint      `json:"duration,omitempty"`
}
