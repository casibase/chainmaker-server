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

package main

import (
	"fmt"

	"github.com/beego/beego"
	"github.com/beego/beego/plugins/cors"
	_ "github.com/casibase/chainserver/routers"
	"github.com/casibase/chainserver/util"
)

func main() {
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.RunMode = "dev"

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	port := beego.AppConfig.DefaultInt("httpport", 13900)

	err := util.StopOldInstance(port)
	if err != nil {
		panic(err)
	}

	beego.Run(fmt.Sprintf(":%v", port))
}
