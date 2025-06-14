// Copyright 2025 The Casibase Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package routers
// @APIVersion 1.70.0
// @Title chainserver RESTful API
// @Description Swagger Docs of chainserver API
// @Contact admin@casibase.org
// @SecurityDefinition AccessToken apiKey Authorization header
// @Schemes https,http
// @ExternalDocs Find out more about chainserver
// @ExternalDocsUrl https://casibase.org/
package routers

import (
	"github.com/beego/beego"
	"github.com/casibase/chainserver/controllers"
)

func init() {
	initAPI()
}

func initAPI() {
	ns := beego.NewNamespace("/api",
		beego.NSInclude(
			&controllers.ApiController{},
		),
	)
	beego.AddNamespace(ns)

	beego.Router("/api/invoke-contract", &controllers.ApiController{}, "POST:InvokeContract")
	beego.Router("/api/query-contract", &controllers.ApiController{}, "POST:QueryContract")

	beego.Router("/api/get-form-data", &controllers.ApiController{}, "POST:GetFormData")
}
