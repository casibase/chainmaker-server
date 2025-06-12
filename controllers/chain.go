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

	"github.com/casibase/chainserver/object"
)

// InvokeContract
// @Title InvokeContract
// @Description invoke contract
// @Param body object.Chainmaker true "contract invoke parameters"
// @Success 200 {object} object.Response
// @Failure 400 Invalid request
// @router /invoke-contract [post]
func (c *ApiController) InvokeContract() {
	var chainmakerInfo object.ChainmakerInfo
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &chainmakerInfo)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}
	chainmakerTxInfo, err := object.InvokeContract(chainmakerInfo)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.ResponseOk(chainmakerTxInfo)
}

// QueryContract
// @Title QueryContract
// @Description query contract
// @Param body object.Chainmaker true "contract query parameters"
// @Success 200 {object} object.Response
// @Failure 400 Invalid request
// @router /query-contract [post]
func (c *ApiController) QueryContract() {
	var chainmakerInfo object.ChainmakerInfo
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &chainmakerInfo)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}
	chainmakerTxInfo, err := object.QueryContract(chainmakerInfo)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.ResponseOk(chainmakerTxInfo)
}
