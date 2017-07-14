package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("appengine", func() {
	Title("The appengine example")
	Description("A simple appengine example")
	Host("localhost:8080")
	Scheme("http")
	BasePath("/")
	Origin("*", func() {
		Methods("GET", "POST", "PUT", "DELETE", "OPTIONS")
		MaxAge(600)
		Credentials()
	})
})
var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("swagger.json", "../swagger/swagger.json")
	Files("swagger/*filepath", "../static/swagger/")
})

var Account = MediaType("application/vnd.account+json", func() {
	Description("Account")
	Attributes(func() {
		Attribute("id", Integer, "id", func() {
			Example(1)
		})
		Attribute("name", String, "name", func() {
			Example("John")
		})
	})
	View("default", func() {
		Attribute("name")
		Required("name")
	})
})

var _ = Resource("Account", func() {
	BasePath("/account")
	DefaultMedia(Account)
	Action("list", func() {
		Description("list")
		Routing(GET(""))
		Response(OK, CollectionOf(Account))
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
	Action("create", func() {
		Description("create")
		Routing(POST(""))
		Payload(func() {
			Param("name")
			Required("name")
		})
		Response(OK)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
	Action("show", func() {
		Description("show")
		Routing(GET("/:id"))
		Params(func() {
			Param("id")
			Param("name")
			Required("id", "name")
		})
		Response(OK)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
	Action("update", func() {
		Description("update")
		Routing(PUT("/:id"))
		Params(func() {
			Param("id", Integer, "id")
			Required("id")
		})
		Payload(func() {
			Param("name")
			Required("name")
		})
		Response(OK)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
	Action("delete", func() {
		Description("delete")
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id", Integer, "id")
			Required("id")
		})
		Payload(func() {
			Param("name")
			Required("name")
		})
		Response(OK)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
})
