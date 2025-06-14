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

package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/casibase/chainserver/object"
)

// GetFormData
// @Title GetFormData
// @Tag Form API
// @Description get form data
// @Param owner query string true "The owner of form"
// @Success 200 {array} object.Form The Response object
// @router /get-form-data [post]
func (c *ApiController) GetFormData() {
	var formObj object.Form
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &formObj)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	limitStr := c.Input().Get("pageSize")
	pageStr := c.Input().Get("p")

	fmt.Printf("GetFormData(), form: %s, limitStr: %s, pageStr: %s\n", formObj.Name, limitStr, pageStr)

	var pagedData []map[string]string
	var totalCount int

	if formObj.Name == "patients" {
		pagedData, totalCount, err = object.GetPatients(formObj, limitStr, pageStr)
	} else if formObj.Name == "hospitals" {
		pagedData, totalCount, err = object.GetHospitals(formObj, limitStr, pageStr)
	} else {
		c.ResponseError(fmt.Sprintf("unsupported form: %s", formObj.Name))
		return
	}

	if err != nil {
		c.ResponseError(err.Error())
		return
	}
	c.ResponseOk(pagedData, totalCount)
}
