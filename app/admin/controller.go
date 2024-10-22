package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohitdas13595/pawzz-hope/results"
)

func AdminController(router *gin.RouterGroup) {
	adminRouter := router.Group("/admin")
	{
		adminRouter.POST("/signup", signUp)
	}

}

func signUp(c *gin.Context) {
	var newAdmin Admin

	err := c.BindJSON(&newAdmin)
	if err != nil {
		r := results.NewAPIResponse[any](400, "Bad Request", nil, err, 0)
		c.JSON(http.StatusBadRequest, r)
	}

	// zlog.Logger().Info("signing up ....................................................................", zap.String("email", newAdmin.Email))

	admin, err := newAdmin.SignUp()
	if err != nil {
		r := results.NewAPIResponse[any](500, err.Error(), nil, err, 0)
		c.JSON(http.StatusBadRequest, r)
		return
	}
	c.JSON(http.StatusOK, admin)

}
