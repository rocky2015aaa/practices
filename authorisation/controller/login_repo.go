package controller

import (
	"errors"
)

type LoginRepo interface {
	save() error
	saveCookie() (int64, error)
	getUser() (*User, error)
	getCookie() string
}

type LoginInfo struct {
	UserName   string `json:"user_name" db:"user_name"`
	Password   string `json:"password" db:"password"`
	CookieData string `json:"cookie_data,omitempty" db:"cookie_data"`
}

func (li *LoginInfo) save() error {
	stmt, err := db.Prepare("INSERT INTO user (user_name,password) VALUES (?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(li.UserName, li.Password)
	if err != nil {
		return err
	}
	return nil
}

func (li *LoginInfo) saveCookie() (int64, error) {
	var affect int64
	stmt, err := db.Prepare("UPDATE user SET cookie_data = ? WHERE user_name = ? and password = ?")
	if err != nil {
		return affect, err
	}

	res, err := stmt.Exec(li.CookieData, li.UserName, li.Password)
	if err != nil {
		return affect, err
	}

	affect, err = res.RowsAffected()
	if err != nil {
		return affect, err
	}

	return affect, nil
}

func (li *LoginInfo) getUser() (*User, error) {
	user := User{LoginInfo: li}
	err := db.QueryRow("SELECT * FROM user WHERE user_name = ? and password = ?", li.UserName, li.Password).Scan(&user.Id, &li.UserName, &li.Password, &li.CookieData, &user.CreatedAt, &user.ModifiedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (li *LoginInfo) getCookie() string {
	return li.CookieData
}

func (li *LoginInfo) validate() error {
	if li.UserName == "" {
		return errors.New("User Name is not valid")
	}
	if li.Password == "" {
		return errors.New("Password is not valid")
	}

	return nil
}
