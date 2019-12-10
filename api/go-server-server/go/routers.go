/*
 * 基于区块链的供应链金融平台
 *
 * Blockchain
 *
 * API version: 1.0.0
 * Contact: chris-ju@qq.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"AddAccount",
		strings.ToUpper("Post"),
		"/api/account/",
		AddAccount,
	},

	Route{
		"DeleteAccount",
		strings.ToUpper("Delete"),
		"/api/account/{accountID}",
		DeleteAccount,
	},

	Route{
		"GetAccountByID",
		strings.ToUpper("Get"),
		"/api/account/{accountID}",
		GetAccountByID,
	},

	Route{
		"GetAllArticle",
		strings.ToUpper("Get"),
		"/api/account/",
		GetAllArticle,
	},

	Route{
		"TransferArticle",
		strings.ToUpper("Put"),
		"/api/account/{accountID}",
		TransferArticle,
	},

	Route{
		"AdminLogin",
		strings.ToUpper("Post"),
		"/admin/login",
		AdminLogin,
	},

	Route{
		"DeleteAccountByAdmin",
		strings.ToUpper("Delete"),
		"/admin/api/account/",
		DeleteAccountByAdmin,
	},

	Route{
		"GetAccountsByAdmin",
		strings.ToUpper("Get"),
		"/admin/api/account/",
		GetAccountsByAdmin,
	},

	Route{
		"GetCompanysByAdmin",
		strings.ToUpper("Get"),
		"/admin/api/company/",
		GetCompanysByAdmin,
	},

	Route{
		"SetCompany",
		strings.ToUpper("Post"),
		"/admin/api/company/",
		SetCompany,
	},

	Route{
		"CreateUser",
		strings.ToUpper("Post"),
		"/api/user/",
		CreateUser,
	},

	Route{
		"QueryUser",
		strings.ToUpper("Get"),
		"/api/user/",
		QueryUser,
	},

	Route{
		"UpdateUser",
		strings.ToUpper("Put"),
		"/api/user/",
		UpdateUser,
	},
}
