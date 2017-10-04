package models

import (
	"crypto/sha256"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"position_api/db"
)

type User struct {
	gorm.Model
	NickName  string `gorm:"not null"`
	UserName  string `gorm:"not null"`
	Age       int
	Pwd       string
	Email     string
	Gender    int
	Summary   string
	Phone     string `gorm:"unique"`
	IsDelete  bool
	OpenId    string
	AvatarUrl string
}

func init() {
	// db.RegisterMigration(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	//db.Conn.AutoMigrate(&User{})
}

func WXLogin(OpenId string) (user *User, err error) {
	user = &User{}
	//e := db.Conn.Where("open_id = ?", OpenId).First(&user)
	rows := db.Conn.Exec("SELECT id FROM user")

	fmt.Println(rows)

	//err = db.Error
	if err != nil {
		return nil, err
	}
	return
}

func UserLogin(email string, password string) (user *User, err error) {
	user = &User{}
	encryptedPassword := userEncrypedPassword(password)
	db := db.Conn.Where("email = ? AND encrypted_password = ?", email, encryptedPassword).First(user).Scan(user)
	err = db.Error
	if err != nil {
		return nil, err
	}
	return
}

func userEncrypedPassword(password string) (encryptedPassword string) {
	saltedPassword := fmt.Sprintf("%s_%s", password, "this is slat")
	encryptedPassword = fmt.Sprintf("%x", sha256.Sum256([]byte(saltedPassword)))
	return
}

func FindUser(id uint) (*User, error) {
	user := &User{}
	user.ID = id
	res := db.Conn.Debug().First(&user, "id=?", id)
	err := res.Error
	if err != nil {
		beego.BeeLogger.Error("Error finding user with id %s: %v", id, res.Error.Error())
	}
	return user, err
}
func CreateWXUser(NickName string, Age int, Email string, Gender int, Phone string, OpenId string, AvatarUrl string) (*User, error) {
	user := &User{NickName: NickName, Age: Age, Email: Email, Gender: Gender, Phone: Phone, OpenId: OpenId, AvatarUrl: AvatarUrl}
	res := db.Conn.Debug().Create(user)
	err := res.Error
	return user, err
}

func CreateUser(UserName string, Age int, Pwd string, Email string, Gender int, Summary string, Phone string, AvatarUrl string) (*User, error) {
	user := &User{UserName: UserName, Age: Age, Pwd: Pwd, Email: Email, Gender: Gender, Summary: Summary, Phone: Phone, AvatarUrl: AvatarUrl}
	res := db.Conn.Debug().Create(user)
	err := res.Error
	return user, err
}

func UpdateUser(id string, UserName *string, Age *int, Pwd *string, Email *string, Gender *int, Summary *string, Phone *string, AvatarUrl *string) (*User, error) {
	user := &User{}
	res := db.Conn.Debug().First(user, "id=?", id)
	if res.Error != nil {
		beego.BeeLogger.Error("Error finding user with id %s: %v", id, res.Error.Error())
	}
	if UserName != nil {
		user.UserName = *UserName
	}

	if Email != nil {
		user.Email = *Email
	}
	if Age != nil {
		user.Age = *Age
	}
	if Gender != nil {
		user.Gender = *Gender
	}
	if Phone != nil {
		user.Phone = *Phone
	}
	if AvatarUrl != nil {
		user.AvatarUrl = *AvatarUrl
	}
	if Summary != nil {
		user.Summary = *Summary
	}
	if Pwd != nil {
		user.Pwd = userEncrypedPassword(*Pwd)
	}
	res = db.Conn.Debug().Save(user)
	err := res.Error
	return user, err
}

func DeleteUser(id uint) (User, error) {
	user := &User{}
	// res := db.Conn.Debug().Delete(&user)
	res := db.Conn.Where("id = ?", id).First(user).Scan(user)
	err := res.Error
	user.IsDelete = true
	db := db.Conn.Debug().Save(user)
	err = db.Error
	return *user, err
}
