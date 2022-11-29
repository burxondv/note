package repo

type User struct {
	ID          int64
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	ImageURL    string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}

type GetAllUsersParams struct {
	Limit      int32
	Page       int32
	Search     string
}

type GetAllUsersResult struct {
	Users []*User
	Count int32
}

type UserStorageI interface {
	Create(u *User) (*User, error)
	Get(id int64) (*User, error)
	GetAll(params *GetAllUsersParams) (*GetAllUsersResult, error)
	Update(u *User) (*User, error)
	Delete(id int64) error
}