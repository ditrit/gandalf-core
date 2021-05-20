package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/ditrit/gandalf/core/aggregator/database"

	"github.com/ditrit/gandalf/core/aggregator/api/dao"
	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/ditrit/gandalf/core/models"

	"github.com/gorilla/mux"
)

// PivotController :
type PivotController struct {
	databaseConnection *database.DatabaseConnection
}

// NewPivotController :
func NewPivotController(databaseConnection *database.DatabaseConnection) (eventTypeController *PivotController) {
	eventTypeController = new(PivotController)
	eventTypeController.databaseConnection = databaseConnection

	return
}

// List :
func (pc PivotController) List(w http.ResponseWriter, r *http.Request) {
	database := pc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		pivots, err := dao.ListPivot(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, pivots)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Read :
func (pc PivotController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := pc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var pivot models.Pivot
		if pivot, err = dao.ReadPivot(database, id); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, pivot)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// ReadByName :
func (pc PivotController) ReadByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	var pivot models.Pivot
	var err error
	if pivot, err = dao.ReadPivotByName(pc.databaseConnection.GetTenantDatabaseClient(), name); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, pivot)
}
