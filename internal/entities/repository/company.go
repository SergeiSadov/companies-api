package repository

type Company struct {
	ID      int    `gorm:"id"`
	Name    string `gorm:"name"`
	Code    string `gorm:"code"`
	Country string `gorm:"country"`
	Website string `gorm:"website"`
	Phone   string `gorm:"phone"`
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
