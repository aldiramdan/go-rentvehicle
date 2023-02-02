package auth

import (
	"encoding/json"
	"net/http"

	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/interfaces"
	"github.com/aldiramdan/go-backend/libs"
)

type auth_ctrl struct {
	srvc interfaces.AuthSrvc
}

func NewCtrl(srvc interfaces.AuthSrvc) *auth_ctrl {

	return &auth_ctrl{srvc}

}

func (c *auth_ctrl) Login(w http.ResponseWriter, r *http.Request) {

	var data models.User

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		libs.GetResponse(err.Error(), 401, true)
		return
	}

	c.srvc.Login(&data).Send(w)
}
