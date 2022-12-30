package repository

type Company struct {
	ID      string `gorm:"id" json:"id"`
	Name    string `gorm:"name" json:"name"`
	Code    string `gorm:"code" json:"code"`
	Country string `gorm:"country" json:"country"`
	Website string `gorm:"website" json:"website"`
	Phone   string `gorm:"phone" json:"phone"`
}

type ListCompanyParams struct {
	Name    string
	Code    string
	Country string
	Website string
	Phone   string

	Page int
	Size int
}
