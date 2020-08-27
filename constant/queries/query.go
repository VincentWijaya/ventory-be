package queries

const (
	FindUserByUsernameOrEmail = "SELECT u.id, u.username, u.email, u.status, r.role, u.password FROM users as u JOIN u_roles r ON u.role_id = r.id WHERE u.username=? OR u.email=?"
	FindAllUser               = "SELECT u.id, u.username, u.email, u.status, r.role FROM users as u JOIN u_roles r ON u.role_id = r.id"
	InsertUser                = "INSERT INTO `users` (`username`, `email`, `password`, `role_id`) VALUES (?, ?, ?, ?)"
)

const (
	GetItemCategory = "SELECT id, name FROM item_category WHERE id=? LIMIT ?"
)
