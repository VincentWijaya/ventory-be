package queries

const (
	FindUserByUsernameOrEmail = `SELECT u.id, u.username, u.email, u.status, r.role, u.password FROM users as u
															 JOIN u_roles r ON u.role_id = r.id
															 WHERE u.username=? OR u.email=?`
	FindAllUser = `SELECT u.id, u.username, u.email, u.status, r.role FROM users as u
								 JOIN u_roles r ON u.role_id = r.id`
	InsertUser = "INSERT INTO `users` (`username`, `email`, `password`, `role_id`) VALUES (?, ?, ?, ?)"
)

const (
	GetItemCategory = "SELECT id, name FROM item_category WHERE id>=? AND is_deleted != 1 LIMIT ?"
	GetItem         = `SELECT i.id, ic.name as category_name, i.name as item_name, i.name, i.buy_price, i.sell_price,
	   								 i.stock, i.notes, i.created_at, i.updated_at
										 FROM items as i
										 JOIN item_category ic ON i.id = ic.id
										 WHERE i.id >= ? AND i.is_deleted != 1 LIMIT ?`
	CountItem         = "SELECT COUNT(id) as total FROM items"
	CountItemCategory = "SELECT COUNT(id) as total FROM item_category"
)
