package query

var Login = `SELECT * FROM auth WHERE username = $1;`

var Register = `INSERT INTO auth (username, password, fullname, email) VALUES ($1, $2, $3, $4) RETURNING *;`

var IsExist = `SELECT COUNT(*) from auth WHERE username = $1;`
