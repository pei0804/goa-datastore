package controller

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/tikasan/goa-datastore/app"
	"github.com/tikasan/goa-datastore/model"
)

// AccountController implements the Account resource.
type AccountController struct {
	*goa.Controller
}

// NewAccountController creates a Account controller.
func NewAccountController(service *goa.Service) *AccountController {
	return &AccountController{Controller: service.NewController("AccountController")}
}

// Create runs the create action.
func (c *AccountController) Create(ctx *app.CreateAccountContext) error {
	// AccountController_Create: start_implement

	// Put your logic here
	a := model.Account{
		Name: ctx.Payload.Name,
	}
	adb := model.AccountDB{}
	err := adb.Add(ctx, &a)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	// AccountController_Create: end_implement
	res := &app.Account{}
	return ctx.OK(res)
}

// Delete runs the delete action.
func (c *AccountController) Delete(ctx *app.DeleteAccountContext) error {
	// AccountController_Delete: start_implement

	// Put your logic here

	// AccountController_Delete: end_implement
	res := &app.Account{}
	return ctx.OK(res)
}

// List runs the list action.
func (c *AccountController) List(ctx *app.ListAccountContext) error {
	// AccountController_List: start_implement

	// Put your logic here

	// AccountController_List: end_implement
	res := app.AccountCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *AccountController) Show(ctx *app.ShowAccountContext) error {
	// AccountController_Show: start_implement

	// Put your logic here
	adb := model.AccountDB{}
	a, err := adb.Get(ctx, ctx.ID)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	// AccountController_Show: end_implement
	res := &app.Account{}
	fmt.Println(a)
	return ctx.OK(res)
}

// Update runs the update action.
func (c *AccountController) Update(ctx *app.UpdateAccountContext) error {
	// AccountController_Update: start_implement

	// Put your logic here

	// AccountController_Update: end_implement
	res := &app.Account{}
	return ctx.OK(res)
}
