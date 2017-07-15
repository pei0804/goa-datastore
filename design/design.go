package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

const DefineTrait = "DefineTrait"

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
	Trait(DefineTrait, func() {
		Response(Unauthorized, ErrorMedia)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
})
var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("swagger.json", "../swagger/swagger.json")
	Files("swagger/*filepath", "../static/swagger/")
})

var User = MediaType("application/vnd.user+json", func() {
	Description("user")
	Attributes(func() {
		Attribute("id", Any, "id(int64)", func() {
			Example(4909628655665152)
		})
		Attribute("id_str", String, "id(string)", func() {
			Example("4909628655665152")
		})
		Attribute("name", String, "name", func() {
			Example("John")
		})
		Required("id", "id_str", "name")
	})
	View("default", func() {
		Attribute("id")
		Attribute("id_str")
		Attribute("name")
		Required("id", "name")
	})
})

var _ = Resource("User", func() {
	BasePath("/users")
	DefaultMedia(User)
	Action("list", func() {
		Description("list")
		Routing(GET(""))
		Params(func() {
			Param("name")
			Required("name")
		})
		Response(OK, CollectionOf(User))
		UseTrait(DefineTrait)
	})
	Action("create", func() {
		Description("create")
		Routing(POST(""))
		Payload(func() {
			Param("name")
			Required("name")
		})
		Response(Created, User)
		UseTrait(DefineTrait)
	})
	Action("show", func() {
		Description("show")
		Routing(GET("/:id"))
		Params(func() {
			Param("id")
			Required("id")
		})
		Response(OK, User)
		UseTrait(DefineTrait)
	})
	Action("update", func() {
		Description("update")
		Routing(PUT("/:id"))
		Params(func() {
			Param("id")
			Required("id")
		})
		Payload(func() {
			Param("name")
			Required("name")
		})
		Response(OK, User)
		UseTrait(DefineTrait)
	})
	Action("delete", func() {
		Description("delete")
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id")
			Required("id")
		})
		Response(NoContent, User)
		UseTrait(DefineTrait)
	})
})
