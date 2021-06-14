package model

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// TsVector ...
type TsVector struct {
	Value string
}

// Scan implements the sql.Scanner interface
func (vec *TsVector) Scan(v interface{}) error {
	return nil
}

// GormDataType ...
func (vec TsVector) GormDataType() string {
	return "tsvector"
}

// GormValue ...
func (vec TsVector) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "to_tsvector(?)",
		Vars: []interface{}{vec.Value},
	}
}
