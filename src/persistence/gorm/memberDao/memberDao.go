package memberdao

import (
	"time"

	"github.com/hhhhp52/webtest/src/domain"
	"github.com/jinzhu/gorm"
)

const (
	table = "user"
)

type model struct {
	ID        int64      `gorm:"column:id; primary_key"`
	Account   string     `gorm:"column:account; unique_index"`
	Password  string     `gorm:"column:password"`
	Name      string     `gorm:"column:name"`
	Nickname  string     `gorm:"column:nickname"`
	Email     string     `gorm:"column:email"`
	Status    string     `gorm:"column:status; default:'enabled'"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	CreatedBy string     `gorm:"column:created_by"`
	UpdatedBy string     `gorm:"column:updated_by"`
}

// GetAll get a record by account
func GetAll(tx *gorm.DB) []domain.User {
	var rows []model
	err := tx.Table(table).
		Debug().
		Scan(&rows).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}

	if err != nil {
		panic(err)
	}

	var result []domain.User
	for _, row := range rows {
		temp := mapping(tx, &row)
		result = append(result, *temp)
	}
	return result
}

func mapping(tx *gorm.DB, m *model) *domain.User {
	return &domain.User{
		ID:        m.ID,
		Account:   m.Account,
		Password:  m.Password,
		Name:      m.Name,
		Nickname:  m.Nickname,
		Email:     m.Email,
		Status:    m.Status,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		CreatedBy: m.CreatedBy,
		UpdatedBy: m.UpdatedBy,
	}
}
