package converter

import (
	"time"

	"gorm.io/gorm"
)

func ToGormDeletedAt(t *time.Time) gorm.DeletedAt {
	if t == nil {
		return gorm.DeletedAt{
			Time:  time.Time{},
			Valid: false,
		}
	} else {
		return gorm.DeletedAt{
			Time:  *t,
			Valid: true,
		}
	}
}

func FromGormDeletedAt(t gorm.DeletedAt) *time.Time {
	if t.Valid {
		return &t.Time
	} else {
		return nil
	}
}