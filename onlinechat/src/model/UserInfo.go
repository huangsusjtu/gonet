package model

import (
	sql "database/sql"
	. "persistence"
	"strings"
	"util"
)

// represent user model in database
type UserInfo struct {
	Uid         int64
	Name        string
	Password    string
	Email       string
	Phone       string
	Description string
}

/*
* name email phone desc password
 */
var sql_create_user string = "CREATE TABLE IF NOT EXISTS user (id BIGINT UNSIGNED PRIMARY KEY, name VARCHAR(64), email VARCHAR(64), phone VARCHAR(32), description VARCHAR(1024), password BINARY(16));"
var sql_select_user_name string = "select * from user where name = ?;"
var sql_select_user_email string = "select * from user where email = ?;"
var sql_select_user_phone string = "select * from user where phone = ?;"
var sql_insert_user string = "insert into user (id, name, email, phone, description, password) values(?,?,?,?,?,?);"

var sqlStmtSelectName *sql.Stmt = nil
var sqlStmtSelectEmail *sql.Stmt = nil
var sqlStmtSelectPhone *sql.Stmt = nil
var sqlStmtInsertUser *sql.Stmt = nil

// login type
const (
	login_by_name = iota
	login_by_email
	login_by_phone
)

// init the model class
func initUserModel() {
	db := GetSqlConPool().GetDb()
	if db == nil {
		logger.Errorln("UserInfo init db is nil")
		return
	}
	defer GetSqlConPool().RecycleDB(db)

	stmt, err := db.Prepare(sql_create_user)
	if err != nil {
		logger.Errorln("UserInfo Prepare sql_create_user", err.Error())
		return
	}
	stmt.Exec()
	stmt.Close()

	sqlStmtSelectName, err = db.Prepare(sql_select_user_name)
	if sqlStmtSelectName == nil {
		logger.Errorln("UserInfo Prepare sql_select_user_name")
		logger.Errorln("sql_select_user_name", err.Error())
	}
	sqlStmtSelectEmail, err = db.Prepare(sql_select_user_email)
	if sqlStmtSelectEmail == nil {
		logger.Errorln("UserInfo Prepare sql_select_user_email")
		logger.Errorln("sql_select_user_email", err.Error())
	}
	sqlStmtSelectPhone, err = db.Prepare(sql_select_user_phone)
	if sqlStmtSelectPhone == nil {
		logger.Errorln("UserInfo Prepare sql_select_user_phone")
		logger.Errorln("sql_select_user_phone", err.Error())
	}
	sqlStmtInsertUser, err = db.Prepare(sql_insert_user)
	if sqlStmtInsertUser == nil {
		logger.Errorln("UserInfo Prepare sql_insert_user")
		logger.Errorln("sql_insert_user", err.Error())
	}

	logger.Debugln("UserInfo Model init ok")
}

func destroyUserModel() {
	if sqlStmtSelectName != nil {
		sqlStmtSelectName.Close()
		sqlStmtSelectName = nil
	}
	if sqlStmtSelectEmail != nil {
		sqlStmtSelectEmail.Close()
		sqlStmtSelectEmail = nil
	}
	if sqlStmtSelectPhone != nil {
		sqlStmtSelectPhone.Close()
		sqlStmtSelectPhone = nil
	}
	if sqlStmtInsertUser != nil {
		sqlStmtInsertUser.Close()
		sqlStmtInsertUser = nil
	}

	logger.Debugln("UserInfo Model destroy")
}

// check if the user have registed in database
func (u *UserInfo) CheckUserExist() bool {
	if u.Name != "" && u.checkUserNameExist() {
		return true
	}

	if u.Email != "" && u.checkUserEmailExist() {
		return true
	}

	if u.Phone != "" && u.checkUserPhoneExist() {
		return true
	}

	return false
}

func (u *UserInfo) Register() bool {
	result, err := sqlStmtInsertUser.Exec(util.NewUuid(), u.Name, u.Email, u.Phone, u.Description, []byte(u.Password))
	if err != nil {
		logger.Errorln(err)
		return false
	}

	if num, _ := result.RowsAffected(); num > 0 {
		return true
	}

	return false
}

// check whether username  have exist in database
func (u *UserInfo) checkUserNameExist() bool {
	// db check
	rows, err := sqlStmtSelectName.Query(u.Name)
	if err != nil {
		logger.Errorln("UserInfo sqlStmtSelectName query err", err.Error())
		return true
	}
	defer rows.Close()

	if rows.Next() {
		return true
	}
	return false
}

// check whether useremail  have exist in database
func (u *UserInfo) checkUserEmailExist() bool {
	// db check
	rows, err := sqlStmtSelectEmail.Query(u.Email)
	if err != nil {
		logger.Errorln("sqlStmtSelectEmail query err", err.Error())
		return true
	}
	defer rows.Close()

	if rows.Next() {
		return true
	}

	return false
}

// check whether useremail  have exist in database
func (u *UserInfo) checkUserPhoneExist() bool {
	// db check
	rows, err := sqlStmtSelectPhone.Query(u.Phone)
	if err != nil {
		logger.Errorln("sqlStmtSelectPhone query err")
		return true
	}
	defer rows.Close()

	if rows.Next() {
		return true
	}
	return false
}

// Auth by name and password
func (u *UserInfo) AuthByName() (result bool, errmsg string) {
	if strings.EqualFold(u.Name, "") {
		return false, "user name can not be empty"
	}

	// db read and auth
	rows, err := sqlStmtSelectName.Query(u.Name)
	if err != nil {
		logger.Errorln("UserInfo sqlStmtSelectName query err", err.Error())
		return false, "server internel error"
	}
	defer rows.Close()

	if rows.Next() {
		var password []byte = make([]byte, 16)
		rows.Scan(&u.Uid, &u.Name, &u.Email, &u.Phone, &u.Description, &password)
		if strings.EqualFold(u.Password, string(password)) {
			u.Password = ""
			return true, "Login by name success"
		}
		return false, "Your password is not match your name"
	}
	return false, "Your name does not exist"
}

// Auth by name and password
func (u *UserInfo) AuthByEmail() (bool, string) {
	if strings.EqualFold(u.Email, "") {
		return false, "user email can not be empty"
	}

	// db read and auth
	rows, err := sqlStmtSelectEmail.Query(u.Email)
	if err != nil {
		logger.Errorln("UserInfo sqlStmtSelectEmail query err", err.Error())
		return false, "server internel error"
	}
	defer rows.Close()

	if rows.Next() {
		var password []byte = make([]byte, 16)
		rows.Scan(&u.Uid, &u.Name, &u.Email, &u.Phone, &u.Description, &password)
		if strings.EqualFold(u.Password, string(password)) {
			u.Password = ""
			return true, "Login by email success"
		}
		return false, "Your password is not match your Email"
	}
	return false, "Your email does not exist"
}

// Auth by name and password
func (u *UserInfo) AuthByPhone() (result bool, errmsg string) {
	if strings.EqualFold(u.Phone, "") {
		return false, "user phone can not be empty"
	}

	// db read and auth
	rows, err := sqlStmtSelectPhone.Query(u.Phone)
	if err != nil {
		logger.Errorln("UserInfo sqlStmtSelectPhone query err", err.Error())
		return false, "server internel error"
	}
	defer rows.Close()

	if rows.Next() {
		var password []byte = make([]byte, 16)
		rows.Scan(&u.Uid, &u.Name, &u.Email, &u.Phone, &u.Description, &password)
		if strings.EqualFold(u.Password, string(password)) {
			u.Password = ""
			return true, "Login by phone success"
		}
		return false, "Your password is not match your Phone"
	}
	return false, "Your phone does not exist"
}
