package queries

const (
	FindUserByUsernameAndPassword = "SELECT u.id, u.username, u.email, u.status, r.role FROM users as u JOIN u_roles r ON u.role_id = r.id WHERE u.username=? AND u.password=? AND u.status=1"
	FindAllUser                   = "SELECT u.id, u.username, u.email, u.status, r.role FROM users as u JOIN u_roles r ON u.role_id = r.id"
	InsertUser                    = "INSERT INTO `users` (`username`, `email`, `password`, `role_id`) VALUES (?, ?, ?, ?)"
)
