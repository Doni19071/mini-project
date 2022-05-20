package appModel

type AdminModel interface {
	GetUsernameAndPassword(username string, password string) (Admin, error)
	AddAdmin(Admin) (Admin, error)
	EditAdmin(int, Admin) (Admin, error)
}
