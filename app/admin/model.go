package admin

import (
	"errors"

	"github.com/rohitdas13595/pawzz-hope/db"
	"github.com/rohitdas13595/pawzz-hope/results"
	"github.com/rohitdas13595/pawzz-hope/utils"

	"gorm.io/gorm"
)

// gorm.Model definition
type Admin struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string
}

func (a *Admin) SignUp() (*results.APIResponse[*Admin], error) {
	if a.Email == "" || a.Password == "" {
		return nil, errors.New("missing email or password")
	}
	hashedPassword, err := utils.HashPassword(a.Password)
	if err != nil {
		return nil, err
	}

	a.Password = hashedPassword

	err = db.DB.Create(&a).Error
	if err != nil {
		return nil, err
	}
	result := results.NewAPIResponse[*Admin](200, "OK", a, nil, 0)
	return &result, nil
}

type AdminLoginResult struct {
	gorm.Model
	Email string `json:"email"`
	Token string `json:"token"`
}

func (a *Admin) SignIn() (*AdminLoginResult, error) {

	if a.Email == "" || a.Password == "" {
		return nil, errors.New("missing email or password")
	}

	var admin Admin
	err := db.DB.Where("email = ?", a.Email).First(&admin).Error
	if err != nil {
		return nil, err
	}
	if !utils.CheckPasswordHash(a.Password, admin.Password) {
		return nil, errors.New("wrong password")
	}

	token, err := utils.CreateAdminToken(&utils.AdminCreateToken{
		Email: admin.Email,
		Id:    admin.ID,
	})
	if err != nil {
		return nil, err
	}

	return &AdminLoginResult{
		Email: admin.Email,
		Token: token,
	}, nil

}

func FindAdminByEmail(email string) (*Admin, error) {
	var admin Admin
	err := db.DB.Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}
	admin.Password = ""
	return &admin, nil
}

func FindAdminById(id uint) (*Admin, error) {
	var admin Admin
	err := db.DB.Where("id = ?", id).First(&admin).Error
	if err != nil {
		return nil, err
	}
	admin.Password = ""
	return &admin, nil
}

func GetAllAdmins() ([]Admin, error) {
	var admins []Admin
	err := db.DB.Find(&admins).Error
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func DeleteAdminById(id uint) error {
	err := db.DB.Where("id = ?", id).Delete(&Admin{}).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateAdminPasswordById(id uint, password string, oldPassword string) error {

	admin, err := FindAdminById(id)
	if err != nil {
		return err
	}
	if !utils.CheckPasswordHash(oldPassword, admin.Password) {
		return errors.New("wrong password")
	}
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	admin.Password = hashedPassword
	err = db.DB.Save(&admin).Error
	if err != nil {
		return err
	}
	return nil
}
