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
	GetItemCategory        = "SELECT id, category_name FROM item_category WHERE id>=? AND is_deleted != 1 LIMIT ?"
	CountItemCategory      = "SELECT COUNT(id) as total FROM item_category WHERE is_deleted != 1"
	InsertItemCategory     = "INSERT INTO `item_category` (`category_name`) VALUES (?)"
	UpdateItemCategory     = "UPDATE item_category SET category_name=? WHERE id=?"
	SoftDeleteItemCategory = "UPDATE item_category SET is_deleted=1 WHERE id=?"
)

// Item
const (
	GetItem = `SELECT i.id, ( SELECT category_name FROM item_category
													  WHERE id = i.category_id
													 ) as category_name,
						i.item_name, i.buy_price, i.sell_price, i.stock,
	          i.notes, i.created_at, i.updated_at
						FROM items as i
						WHERE i.id > ? AND i.is_deleted != 1 LIMIT ?`
	CountItem            = "SELECT COUNT(id) as total FROM items WHERE is_deleted != 1"
	FindItemByCategoryID = `SELECT id, item_name, category_id, buy_price, sell_price,
													stock, created_at, updated_at, notes FROM items
													WHERE category_id=?`
	InsertItem     = "INSERT INTO `items` (`item_name`, `category_id`, `buy_price`, `sell_price`, `stock`, `notes`) VALUES (?, ?, ?, ?, ?, ?)"
	UpdateItem     = "UPDATE items SET item_name=?, category_id=?, buy_price=?, sell_price=?, stock=?, notes=? WHERE id=?"
	SoftDeleteItem = "UPDATE items SET is_deleted=1 WHERE id=?"
	FindItemByID   = `SELECT id, item_name, category_id, buy_price, sell_price,
										stock, created_at, updated_at, notes FROM items
										WHERE id=?`
)

const (
	InsertItemHistory = "INSERT INTO `item_history` (`item_id`, `buy_price`, `sell_price`, `stock`, `notes`) VALUES (?, ?, ?, ?, ?)"
)
