// Code generated by protoc-gen-gorm. DO NOT EDIT.
// versions:
// 	protoc-gen-gorm v2.0.0
// 	protoc        v3.20.0
// source: example/base/test.proto

package pbconf2

import (
	context "context"
	_ "github.com/mozyy/protoc-gen-gorm/options"
	types "github.com/mozyy/protoc-gen-gorm/types"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	time "time"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConfigGORM struct {
	ID uint32 `gorm:"primaryKey"` // ID
	// name
	Name string // google.protobuf.Timestamp created_at = 2;
	// oneof test_oneof {
	//   string name_first = 8;
	//   string name_last = 9;
	// }
	Config    *Config2GORM
	ConfigID  uint32
	CreatedAt *time.Time       // created_at
	UpdatedAt *time.Time       // updated_at
	DeletedAt *types.DeletedAt `gorm:"index"` // deleted_at
}

// TableName overrides the default tablename generated by GORM
func (ConfigGORM) TableName() string {
	return "oauth_configs"
}

// ToPB  converts the fields of this object to PB object
func (m *ConfigGORM) ToPB(ctx context.Context) *Config {
	to := Config{}
	to.ID = m.ID
	to.Name = m.Name
	to.Config = m.Config.ToPB(ctx)
	to.ConfigID = m.ConfigID
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	to.DeletedAt = m.DeletedAt
	return &to
}

// ToORM  converts the fields of this object to PB object
func (m *Config) ToORM(ctx context.Context) *ConfigGORM {
	to := ConfigGORM{}
	to.ID = m.GetID()
	to.Name = m.GetName()
	to.Config = m.GetConfig().ToORM(ctx)
	to.ConfigID = m.GetConfigID()
	if m.GetCreatedAt() != nil {
		CreatedAt := m.GetCreatedAt().AsTime()
		to.CreatedAt = &CreatedAt
	}
	if m.GetUpdatedAt() != nil {
		UpdatedAt := m.GetUpdatedAt().AsTime()
		to.UpdatedAt = &UpdatedAt
	}
	to.DeletedAt = m.GetDeletedAt()
	return &to
}

type Config2GORM struct {
	ID        uint32 `gorm:"primaryKey"` // ID
	Type      string
	CreatedAt *time.Time       // created_at
	UpdatedAt *time.Time       // updated_at
	DeletedAt *types.DeletedAt `gorm:"index"` // deleted_at
}

// TableName overrides the default tablename generated by GORM
func (Config2GORM) TableName() string {
	return "oauth_config2"
}

// ToPB  converts the fields of this object to PB object
func (m *Config2GORM) ToPB(ctx context.Context) *Config2 {
	to := Config2{}
	to.ID = m.ID
	to.Type = m.Type
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	to.DeletedAt = m.DeletedAt
	return &to
}

// ToORM  converts the fields of this object to PB object
func (m *Config2) ToORM(ctx context.Context) *Config2GORM {
	to := Config2GORM{}
	to.ID = m.GetID()
	to.Type = m.GetType()
	if m.GetCreatedAt() != nil {
		CreatedAt := m.GetCreatedAt().AsTime()
		to.CreatedAt = &CreatedAt
	}
	if m.GetUpdatedAt() != nil {
		UpdatedAt := m.GetUpdatedAt().AsTime()
		to.UpdatedAt = &UpdatedAt
	}
	to.DeletedAt = m.GetDeletedAt()
	return &to
}

type UserGORM struct {
	ID               uint32 `gorm:"primaryKey"`
	CreditCard       []*CreditCardGORM
	CreditCardSignal *CreditCardGORM
	Name             string
	User             []*UserGORM `gorm:"many2many:user_friend"`
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
	DeletedAt        *types.DeletedAt `gorm:"index"`
}

// TableName overrides the default tablename generated by GORM
func (UserGORM) TableName() string {
	return "oauth_users"
}

// ToPB  converts the fields of this object to PB object
func (m *UserGORM) ToPB(ctx context.Context) *User {
	to := User{}
	to.ID = m.ID
	for _, CreditCard := range m.CreditCard {
		to.CreditCard = append(to.CreditCard, CreditCard.ToPB(ctx))
	}
	to.CreditCardSignal = m.CreditCardSignal.ToPB(ctx)
	to.Name = m.Name
	for _, User := range m.User {
		to.User = append(to.User, User.ToPB(ctx))
	}
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	to.DeletedAt = m.DeletedAt
	return &to
}

// ToORM  converts the fields of this object to PB object
func (m *User) ToORM(ctx context.Context) *UserGORM {
	to := UserGORM{}
	to.ID = m.GetID()
	for _, CreditCard := range m.GetCreditCard() {
		to.CreditCard = append(to.CreditCard, CreditCard.ToORM(ctx))
	}
	to.CreditCardSignal = m.GetCreditCardSignal().ToORM(ctx)
	to.Name = m.GetName()
	for _, User := range m.GetUser() {
		to.User = append(to.User, User.ToORM(ctx))
	}
	if m.GetCreatedAt() != nil {
		CreatedAt := m.GetCreatedAt().AsTime()
		to.CreatedAt = &CreatedAt
	}
	if m.GetUpdatedAt() != nil {
		UpdatedAt := m.GetUpdatedAt().AsTime()
		to.UpdatedAt = &UpdatedAt
	}
	to.DeletedAt = m.GetDeletedAt()
	return &to
}

type CreditCardGORM struct {
	ID         uint32 `gorm:"primaryKey"`
	Number     string
	UserGORMID uint32
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	DeletedAt  *types.DeletedAt `gorm:"index"`
}

// TableName overrides the default tablename generated by GORM
func (CreditCardGORM) TableName() string {
	return "oauth_credit_cards"
}

// ToPB  converts the fields of this object to PB object
func (m *CreditCardGORM) ToPB(ctx context.Context) *CreditCard {
	to := CreditCard{}
	to.ID = m.ID
	to.Number = m.Number
	to.UserGORMID = m.UserGORMID
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	to.DeletedAt = m.DeletedAt
	return &to
}

// ToORM  converts the fields of this object to PB object
func (m *CreditCard) ToORM(ctx context.Context) *CreditCardGORM {
	to := CreditCardGORM{}
	to.ID = m.GetID()
	to.Number = m.GetNumber()
	to.UserGORMID = m.GetUserGORMID()
	if m.GetCreatedAt() != nil {
		CreatedAt := m.GetCreatedAt().AsTime()
		to.CreatedAt = &CreatedAt
	}
	if m.GetUpdatedAt() != nil {
		UpdatedAt := m.GetUpdatedAt().AsTime()
		to.UpdatedAt = &UpdatedAt
	}
	to.DeletedAt = m.GetDeletedAt()
	return &to
}
