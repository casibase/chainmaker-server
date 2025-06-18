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

package object

import (
	"encoding/json"
	"fmt"
	"strconv"

	"chainmaker.org/chainmaker/pb-go/v2/common"
	chainmakersdk "chainmaker.org/chainmaker/sdk-go/v2"
)

type ChainConfig struct {
	ChainId     string `json:"chain_id"`
	OrgId       string `json:"org_id"`
	UserKey     string `json:"user_key"`
	UserCert    string `json:"user_cert"`
	SignKey     string `json:"sign_key"`
	SignCert    string `json:"sign_cert"`
	NodeAddr    string `json:"node_addr"`
	TLSHostName string `json:"tls_host_name"`
	AuthType    string `json:"auth_type"`
	EnableTLS   bool   `json:"enable_tls"`
	CACert      string `json:"ca_cert"`
}

type ChainmakerInfo struct {
	ChainConfig    ChainConfig `json:"chain_config"`
	ContractName   string      `json:"contract_name"`
	ContractMethod string      `json:"contract_method"`
	Data           string      `json:"data"`
	TxId           string      `json:"txId"`
}

type ChainmakerTxInfo struct {
	TxId      string `json:"tx_id"`
	Block     string `json:"block"`
	BlockHash string `json:"block_hash"`
	Result    string `json:"result"`
}

func InvokeContract(chainmakerInfo ChainmakerInfo) (ChainmakerTxInfo, error) {
	config := chainmakerInfo.ChainConfig

	client, err := createClient(config)
	if err != nil {
		return ChainmakerTxInfo{}, err
	}

	var dataMap map[string]string
	err = json.Unmarshal([]byte(chainmakerInfo.Data), &dataMap)
	if err != nil {
		return ChainmakerTxInfo{}, fmt.Errorf("InvokeContract() error: %v", err)
	}

	kvPairs := make([]*common.KeyValuePair, 0, len(dataMap))
	for key, value := range dataMap {
		if key == "key" {
			owner, name, err := GetOwnerAndNameFromId(value, "/")
			if err != nil {
				return ChainmakerTxInfo{}, err
			}

			value = GetIdFromOwnerAndName(owner, name, "_")
		}
		kvPairs = append(kvPairs, &common.KeyValuePair{
			Key:   key,
			Value: []byte(value),
		})
	}

	resp, err := client.InvokeContract(chainmakerInfo.ContractName, chainmakerInfo.ContractMethod, "", kvPairs, -1, true)
	if err != nil {
		return ChainmakerTxInfo{}, fmt.Errorf("InvokeContract() error: %v", err)
	}

	if resp.Code != common.TxStatusCode_SUCCESS {
		return ChainmakerTxInfo{}, fmt.Errorf("InvokeContract() error, result = %s, message = %s", resp.ContractResult.Result, resp.ContractResult.Message)
	}

	txId := resp.TxId

	transactionInfo, err := client.GetTxByTxId(txId)
	if err != nil {
		return ChainmakerTxInfo{}, fmt.Errorf("InvokeContract() error: %v", err)
	}

	blockInfo, err := client.GetBlockByHeight(transactionInfo.GetBlockHeight(), true)
	if err != nil {
		return ChainmakerTxInfo{}, err
	}

	block := strconv.FormatUint(transactionInfo.GetBlockHeight(), 10)
	blockHash := blockInfo.GetBlock().GetBlockHashStr()
	return ChainmakerTxInfo{TxId: txId, Block: block, BlockHash: blockHash}, nil
}

func QueryContract(chainmakerInfo ChainmakerInfo) (ChainmakerTxInfo, error) {
	config := chainmakerInfo.ChainConfig

	client, err := createClient(config)
	if err != nil {
		return ChainmakerTxInfo{}, err
	}
	transactionInfo, err := client.GetTxByTxId(chainmakerInfo.TxId)
	if err != nil {
		return ChainmakerTxInfo{}, fmt.Errorf("query contract error: %v", err)
	}

	parameters := transactionInfo.GetTransaction().GetPayload().GetParameters()
	if parameters == nil {
		return ChainmakerTxInfo{}, fmt.Errorf("query contract result is nil")
	}

	resultMap := make(map[string]string)
	for _, parameter := range parameters {
		if parameter.Key == "key" {
			owner, name, err := GetOwnerAndNameFromId(string(parameter.Value), "_")
			if err != nil {
				return ChainmakerTxInfo{}, fmt.Errorf("InvokeContract() error: %v", err)
			}

			resultMap[parameter.Key] = GetIdFromOwnerAndName(owner, name, "/")
		} else {
			resultMap[parameter.Key] = string(parameter.Value)
		}
	}

	resultBytes, err := json.Marshal(resultMap)
	if err != nil {
		return ChainmakerTxInfo{}, fmt.Errorf("marshal result error: %v", err)
	}

	txInfo := ChainmakerTxInfo{
		TxId:   chainmakerInfo.TxId,
		Block:  strconv.FormatUint(transactionInfo.GetBlockHeight(), 10),
		Result: string(resultBytes),
	}
	return txInfo, nil
}

func createClient(config ChainConfig) (*chainmakersdk.ChainClient, error) {
	node := chainmakersdk.NewNodeConfig(
		chainmakersdk.WithNodeAddr(config.NodeAddr),
		chainmakersdk.WithNodeConnCnt(10),
		chainmakersdk.WithNodeUseTLS(config.EnableTLS),
		chainmakersdk.WithNodeCACerts([]string{config.CACert}),
		chainmakersdk.WithNodeTLSHostName(config.TLSHostName),
	)

	client, err := chainmakersdk.NewChainClient(
		chainmakersdk.WithChainClientOrgId(config.OrgId),
		chainmakersdk.WithChainClientChainId(config.ChainId),
		chainmakersdk.WithUserKeyBytes([]byte(config.UserKey)),
		chainmakersdk.WithUserCrtBytes([]byte(config.UserCert)),
		chainmakersdk.WithUserSignKeyBytes([]byte(config.SignKey)),
		chainmakersdk.WithUserSignCrtBytes([]byte(config.SignCert)),
		chainmakersdk.AddChainClientNodeConfig(node),
		chainmakersdk.WithPkcs11Config(&chainmakersdk.Pkcs11Config{
			Enabled: false,
		}),
		chainmakersdk.WithAuthType(config.AuthType),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create chain client: %v", err)
	}

	return client, nil
}
