package gorm

import (
	"database/sql/driver"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

func (d *DeletedAt) toGORMDeletedAt() gorm.DeletedAt {
	return gorm.DeletedAt{Time: d.GetTime().AsTime(), Valid: d.GetTime().IsValid()}
}
func (d *DeletedAt) GormDataType() string {
	return "time"
}

func (d *DeletedAt) Scan(src interface{}) error {
	gd := &gorm.DeletedAt{}
	err := gd.Scan(src)
	if err != nil {
		return err
	}
	d.Time = timestamppb.New(gd.Time)
	return nil
}
func (d *DeletedAt) Value() (driver.Value, error) {
	return d.toGORMDeletedAt().Value()
}

// func (d *DeletedAt) CreateClauses(f *schema.Field) []clause.Interface {
// 	return gd.CreateClauses(f)
// }
func (d *DeletedAt) QueryClauses(f *schema.Field) []clause.Interface {
	return d.toGORMDeletedAt().QueryClauses(f)
}
func (d *DeletedAt) UpdateClauses(f *schema.Field) []clause.Interface {
	return d.toGORMDeletedAt().UpdateClauses(f)
}
func (d *DeletedAt) DeleteClauses(f *schema.Field) []clause.Interface {
	return d.toGORMDeletedAt().DeleteClauses(f)
}

// d *DeletedAt schema.DeleteClausesInterface
