package entity

type AddStaffInput struct {
	Name     string
	Email    string
	Position string
}

type Staff struct {
	StaffID  int64
	Name     string
	Email    string
	Position string
}
