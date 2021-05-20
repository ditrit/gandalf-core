package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ditrit/gandalf/core/aggregator/database"

	"github.com/ditrit/gandalf/core/aggregator/api/dao"
	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/ditrit/gandalf/core/models"

	"github.com/gorilla/mux"
)

// ConnectorProductController :
type ConnectorProductController struct {
	databaseConncction *database.DatabaseConnection
}

// NewConnectorProductController :
func NewConnectorProductController(databaseConncction *database.DatabaseConnection) (connectorProductController *ConnectorProductController) {
	connectorProductController = new(ConnectorProductController)
	connectorProductController.databaseConncction = databaseConncction

	return
}

// List :
func (cc ConnectorProductController) List(w http.ResponseWriter, r *http.Request) {
	database := cc.databaseConncction.GetTenantDatabaseClient()
	if database != nil {
		connectorProducts, err := dao.ListConnectorProduct(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, connectorProducts)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Create :
func (cc ConnectorProductController) Create(w http.ResponseWriter, r *http.Request) {
	database := cc.databaseConncction.GetTenantDatabaseClient()
	if database != nil {
		var connectorProduct models.ConnectorProduct
		dccoder := json.NewDecoder(r.Body)
		if err := dccoder.Decode(&connectorProduct); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		if err := dao.CreateConnectorProduct(database, connectorProduct); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, connectorProduct)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Read :
func (cc ConnectorProductController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := cc.databaseConncction.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var connectorProduct models.ConnectorProduct
		if connectorProduct, err = dao.ReadConnectorProduct(database, id); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, connectorProduct)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// ReadByName :
func (cc ConnectorProductController) ReadByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	var connectorProduct models.ConnectorProduct
	var err error
	if connectorProduct, err = dao.ReadConnectorProductByName(cc.databaseConncction.GetTenantDatabaseClient(), name); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, connectorProduct)
}

// Update :
func (cc ConnectorProductController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := cc.databaseConncction.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var connectorProduct models.ConnectorProduct
		dccoder := json.NewDecoder(r.Body)
		if err := dccoder.Decode(&connectorProduct); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		connectorProduct.ID = uint(id)

		if err := dao.UpdateConnectorProduct(database, connectorProduct); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, connectorProduct)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Delete :
func (cc ConnectorProductController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := cc.databaseConncction.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}

		if err := dao.DeleteConnectorProduct(database, id); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}
