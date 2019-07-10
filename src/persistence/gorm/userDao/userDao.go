package userdao

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
	Status    string     `gorm:"column:status; default:'enabled'"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	CreatedBy string     `gorm:"column:created_by"`
	UpdatedBy string     `gorm:"column:updated_by"`
}

// New a record
func New(tx *gorm.DB, user *domain.User) {
	err := tx.Table(table).
		Create(&model{
			Account:   user.Account,
			Password:  user.Password,
			Name:      user.Name,
			Nickname:  user.Nickname,
		}).Error

	if err != nil {
		panic(err)
	}
}


// GetByAccount get a record by account
func GetByAccount(tx *gorm.DB, account string) *domain.User {
	result := &model{}
	err := tx.Table(table).
		Where("account = ?", account).
		Scan(result).Error
	
	if gorm.IsRecordNotFoundError(err){
		return nil
	}

	if err != nil {
		panic(err)
	}
	return mapping(tx, result)
}

func mapping(tx *gorm.DB, m *model) *domain.User {
	return &domain.User{
		ID:        m.ID,
		Account:   m.Account,
		Password:  m.Password,
		Name:      m.Name,
		Nickname:  m.Nickname,
		Status:    m.Status,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		CreatedBy: m.CreatedBy,
		UpdatedBy: m.UpdatedBy,
	}
}
