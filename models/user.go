package models

import "time"

type Role string

const (
	AdminRole    Role = "admin"
	OwnerRole    Role = "owner"
	CustomerRole Role = "customer"
)

type User struct {
	ID                uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName         string     `gorm:"size:150:not null" json:"first_name"`
	LastName          string     `gorm:"size:150;not null" json:"last_name"`
	Email             string     `gorm:"size:255;not null" json:"email"`
	Password          string     `gorm:"size:255;not null" json:"-"` // Hide password from JSON
	IsSuperuser       bool       `gorm:"default:false" json:"is_superuser"`
	JoinedAt          time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"joined_at"`
	LastLoginAt       *time.Time `json:"last_login_at"`
	PasswordChangedAt *time.Time `json:"password_changed_at"`
	AvatarKey         *string    `gorm:"size:255" json:"avatar_key"`
	Status            string     `gorm:"size:20;default:active;check:status IN ('active', 'inactive', 'suspended')" json:"status"`
	EmailVerified     bool       `gorm:"default:false" json:"email_verified"`
	Role              Role       `grom:"type:varchar(20);not null;" json:"role"`
}

type OwnerProfile struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint      `json:"user_id"`
	User        User      `gorm:"foreignKey:UserID" json:"user"`
	CompanyName string    `gorm:"size:255;default:null" json:"company_name"`
	PhoneNumber string    `gorm:"size:20;not null" json:"phone_number"`
	Website     string    `gorm:"size:255;default:null" json:"website"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
