package mysqluserrepo

const getById string = "GetById"
const existsByEmail string = "ExistsByEmail"
const insertUser string = "InsertUser"

var queryCollection map[string]string = map[string]string{
	getById:       "SELECT * FROM users WHERE user_id = ? LIMIT 1",
	existsByEmail: "SELECT EXISTS(SELECT 1 FROM users WHERE user_email = ? LIMIT 1)",
	insertUser:    "INSERT INTO users (user_email) VALUES (?)",
}
