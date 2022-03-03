package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	config "restaurant/config"
	mCust "restaurant/model/user"
	util "restaurant/util"

	log "github.com/sirupsen/logrus"
)

// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Param Data body mCust.Customer true "customer"
// @Accept json
// @Produce json
// @Success 200 {string} succeed
// @Router /user/customer [post]
func CreateCustomer(c *gin.Context) {
	db, errconn := config.GetConnection()
	if errconn != nil {
		log.Fatal(errconn)
	} else {
		var customer mCust.Customer
		errbind := c.BindJSON(&customer)
		if errbind != nil {
			log.Error(errbind)
			util.HandleError(c, errbind, util.MISSING_REQUEST, "Missing data", http.StatusBadRequest)
		} else {
			if errpersist := customer.Base.PrePersist(); errpersist != nil {
				log.Error(errpersist)
			}
			query := `INSERT INTO customer VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *`
			if errdb := db.QueryRow(query,
				strings.Replace(uuid.Must(uuid.NewRandom()).String(), "-", "", -1),
				customer.CustomerName,
				customer.CustomerEmail,
				customer.CustomerPhone,
				customer.CustomerDOB,
				customer.CustomerAddress,
				customer.Base.CreatedAt,
				customer.Base.UpdatedAt,
				false,
			).Scan(
				&customer.CustomerID,
				&customer.CustomerName,
				&customer.CustomerEmail,
				&customer.CustomerPhone,
				&customer.CustomerDOB,
				&customer.CustomerAddress,
				&customer.Base.CreatedAt,
				&customer.Base.UpdatedAt,
				&customer.IsDeleted,
			); errdb != nil {
				log.Error(errdb)
				util.HandleError(c, errdb, util.INSERT_RES_FAILED, "Please make sure the data is appropriate", http.StatusInternalServerError)
			} else {
				util.HandleSuccess(c, customer, util.INSERT_RES_SUCCEESS, http.StatusCreated)
			}
		}
	}

}
