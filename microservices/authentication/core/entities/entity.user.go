package entities

const UserTableName = "_users"

type User struct {
	Id        string
	Name      string
	Avatar    string
	Email     string
	Password  string
	Verified  bool
	Code      string
	Token     string
	Roles     []string
	CreatedAt string
	UpdatedAt string
}
