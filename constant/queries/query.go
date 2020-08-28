package queries

// User
const (
	FindUserByUsernameOrEmail = `SELECT u.id, u.username, u.email, u.status, r.role, u.password FROM users as u
															 JOIN u_roles r ON u.role_id = r.id
															 WHERE u.username=? OR u.email=?`
	FindAllUser = `SELECT u.id, u.username, u.email, u.status, r.role FROM users as u
								 JOIN u_roles r ON u.role_id = r.id`
	InsertUser = "INSERT INTO `users` (`username`, `email`, `password`, `role_id`) VALUES (?, ?, ?, ?)"
)

// Item Category
const (
	GetItemCategory        = "SELECT id, name FROM item_category WHERE id>=? AND is_deleted != 1 LIMIT ?"
	CountItemCategory      = "SELECT COUNT(id) as total FROM item_category"
	InsertItemCategory     = "INSERT INTO `item_category` (`name`) VALUES (?)"
	UpdateItemCategory     = "UPDATE item_category SET name=?, updated_at=? WHERE id=?"
	SoftDeleteItemCategory = "UPDATE item_category SET is_deleted=1, updated_at=? WHERE id=?"
)

// Item
const (
	GetItem = `SELECT i.id, ic.name as category_name, i.name as item_name, i.name, i.buy_price, i.sell_price,
	   								 i.stock, i.notes, i.created_at, i.updated_at
										 FROM items as i
										 JOIN item_category ic ON i.id = ic.id
										 WHERE i.id >= ? AND i.is_deleted != 1 LIMIT ?`
	CountItem            = "SELECT COUNT(id) as total FROM items"
	FindItemByCategoryID = `SELECT id, name, category_id, buy_price, sell_price,
													stock, created_at, updated_at, notes FROM items
													WHERE category_id=?`
	InsertItem     = "INSERT INTO `items` (`name`, `category_id`, `buy_price`, `sell_price`, `stock`, `notes`) VALUES (?, ?, ?, ?, ?, ?)"
	UpdateItem     = "UPDATE items SET name=?, category_id=?, buy_price=?, sell_price=?, stock=?, notes=?, updated_at=? WHERE id=?"
	SoftDeleteItem = "UPDATE items SET is_deleted=, updated_at=? WHERE id=?"
)
