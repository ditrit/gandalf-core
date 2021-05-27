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

// ProductConnectorController :
type ProductConnectorController struct {
	databaseConncction *database.DatabaseConnection
}

// NewProductConnectorController :
func NewProductConnectorController(databaseConncction *database.DatabaseConnection) (productConnectorController *ProductConnectorController) {
	productConnectorController = new(ProductConnectorController)
	productConnectorController.databaseConncction = databaseConncction

	return
}

// List :
func (cc ProductConnectorController) List(w http.ResponseWriter, r *http.Request) {
	database := cc.databaseConncction.GetTenantDatabaseClient()
	if database != nil {
		productConnectors, err := dao.ListProductConnector(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, productConnectors)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Read :
func (cc ProductConnectorController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := cc.databaseConncction.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var productConnector models.ProductConnector
		if productConnector, err = dao.ReadProductConnector(database, id); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, productConnector)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// ReadByName :
func (cc ProductConnectorController) ReadByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	var productConnector models.ProductConnector
	var err error
	if productConnector, err = dao.ReadProductConnectorByName(cc.databaseConncction.GetTenantDatabaseClient(), name); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, productConnector)
}
