package query

var Login = `SELECT * FROM auth WHERE username = $1;`

var Register = `INSERT INTO auth (username, password, fullname, email) VALUES ($1, $2, $3, $4) RETURNING *;`

var IsExist = `SELECT COUNT(*) FROM auth WHERE username = $1;`

var UpdateUserAccount = `UPDATE auth SET username = $1, password = $2, fullname = $3, email = $4 WHERE id = $5 RETURNING *;`

var IsNotExist = `SELECT COUNT(*) FROM auth WHERE id = $1;`
