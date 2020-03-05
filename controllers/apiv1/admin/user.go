package admin

import (
	"encoding/json"
	"goldfish/app"
	"goldfish/controllers"
	"goldfish/domain/models"
	"goldfish/domain/services"
	"goldfish/types"
)

type UserController struct {
	*controllers.Controller
}

func (u UserController) Post(ctx *app.Context) error {
	model := types.UserAdd{}
	json.Unmarshal(ctx.Model, &model)
	if err := model.Validate(); err != nil {
		return app.NewErrValidation(err.Error())
	}

	user := &models.User{
		MobileNumber: model.MobileNumber,
		FirstName:    model.FirstName,
		LastName:     model.LastName,
		GroupID:      model.GroupID,
		Password:     []byte(model.Password),
	}

	return services.NewUserService(ctx.DBSession).Save(user)
}