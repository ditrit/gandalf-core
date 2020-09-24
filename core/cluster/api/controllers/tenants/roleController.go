package tenants

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ditrit/gandalf/core/cluster/api/dao"
	"github.com/ditrit/gandalf/core/cluster/api/utils"
	"github.com/ditrit/gandalf/core/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// RoleController :
type RoleController struct {
	mapDatabase  map[string]*gorm.DB
	databasePath string
}

// NewRoleController :
func NewRoleController(mapDatabase map[string]*gorm.DB, databasePath string) (roleController *RoleController) {
	roleController = new(RoleController)
	roleController.mapDatabase = mapDatabase
	roleController.databasePath = databasePath

	return
}

// List :
func (rc RoleController) List(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(rc.mapDatabase, rc.databasePath, tenant)

	roles, err := dao.ListRole(database)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, roles)
}

// Create :
func (rc RoleController) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(rc.mapDatabase, rc.databasePath, tenant)

	var role models.Role
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&role); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := dao.CreateRole(database, role); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, role)
}

// Read :
func (rc RoleController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(rc.mapDatabase, rc.databasePath, tenant)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var role models.Role
	if role, err = dao.ReadRole(database, id); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, role)
}

// Update :
func (rc RoleController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(rc.mapDatabase, rc.databasePath, tenant)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var role models.Role
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&role); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	role.ID = uint(id)

	if err := dao.UpdateRole(database, role); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, role)
}

// Delete :
func (rc RoleController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(rc.mapDatabase, rc.databasePath, tenant)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if err := dao.DeleteRole(database, id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
