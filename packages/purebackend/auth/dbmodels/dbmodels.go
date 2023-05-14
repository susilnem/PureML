package dbmodels

import (
	"time"

	commondbmodels "github.com/PureMLHQ/PureML/packages/purebackend/core/common/dbmodels"
	userorgdbmodels "github.com/PureMLHQ/PureML/packages/purebackend/user_org/dbmodels"
	uuid "github.com/satori/go.uuid"
)

type Session struct {
	commondbmodels.BaseModel `gorm:"embedded"`
	UserUUID                 uuid.NullUUID `json:"user_uuid" gorm:"type:uuid"`
	Device                   string        `json:"device"`
	DeviceId                 string        `json:"device_id" gorm:"not null"`
	DeviceLocation           string        `json:"device_location"`
	Approved                 bool          `json:"approved" default:"false"`
	Invalid                  bool          `json:"invalid" default:"false"`

	User userorgdbmodels.User `gorm:"foreignKey:UserUUID"`
}

type Tokens struct {
	commondbmodels.BaseModel `gorm:"embedded"`
	Name                     string    `json:"name" gorm:"unique;not null"`
	UserUUID                 uuid.UUID `json:"user_uuid" gorm:"type:uuid"`
	LastUsedAt               time.Time `json:"last_used_at"`

	User userorgdbmodels.User `gorm:"foreignKey:UserUUID"`
}
