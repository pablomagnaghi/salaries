package domain

type Salary struct {
	ID            int64   `gorm:"column:primaryKey" json:"id"`
	Name          string  `gorm:"column:name" json:"name" binding:"required"`
	Salary        float64 `gorm:"column:salary" json:"salary,string" binding:"required"`
	Currency      string  `gorm:"column:currency" json:"currency" binding:"required"`
	OnContract    bool    `gorm:"column:on_contract" json:"on_contract,string"`
	Department    string  `gorm:"column:department" json:"department" binding:"required"`
	SubDepartment string  `gorm:"column:sub_department" json:"sub_department" binding:"required"`
}
