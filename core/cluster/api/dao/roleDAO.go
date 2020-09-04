package dao

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

type RoleDAO struct {
	GandalfDatabase *gorm.DB
}

func NewRoleDAO(gandalfDatabase *gorm.DB) (roleDAO *RoleDAO) {
	roleDAO = new(RoleDAO)
	roleDAO.GandalfDatabase = gandalfDatabase

	return
}

func (rd RoleDAO) List() (roles []models.Role, err error) {
	err = rd.GandalfDatabase.Find(&roles).Error

	return
}

func (rd RoleDAO) Create(role models.Role) (err error) {
	err = rd.GandalfDatabase.Create(&role).Error

	return
}

func (rd RoleDAO) Read(id int) (role models.Role, err error) {
	err = rd.GandalfDatabase.First(&role, id).Error

	return
}

func (rd RoleDAO) Update(role models.Role) (err error) {
	err = rd.GandalfDatabase.Save(&role).Error

	return
}

func (rd RoleDAO) Delete(id int) (err error) {
	var role models.Role
	err = rd.GandalfDatabase.Delete(&role, id).Error

	return
}