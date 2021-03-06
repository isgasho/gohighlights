package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/icbd/gohighlights/controllers/api"
	"github.com/icbd/gohighlights/models"
)

// POST /api/v1/sessions
//{
//	email: "",
//	password: ""
//}
func SessionsCreate(c *gin.Context) {
	resp := api.New(c)

	vo := models.SessionVO{}
	if err := c.BindJSON(&vo); err != nil {
		resp.ParametersErr(err)
		return
	}

	u, err := models.UserFindByEmail(vo.Email)
	if err != nil {
		// not found account
		resp.UnauthorizedErr()
		return
	} else {
		u.Password = vo.Password
		if !u.ValidPassword() {
			// password incorrect
			resp.UnauthorizedErr()
			return
		}
	}

	if s, err := u.GenerateSession(); err != nil {
		resp.ParametersErr(err)
	} else {
		resp.Created(s)
	}
}
