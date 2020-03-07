/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an asset
type SmartContract struct {
	contractapi.Contract
}

// Actor describes basic details of an actor in the SCM
type Actor struct {
	ActorId          string            `json:"actorId"`
	ActorType        string            `json:"actorType"`
	ActorName        string            `json:"actorName"`
	AditionalInfoMap map[string]string `json:"aditionalInfoMap"`
	//assetId        string            `json:"assetId"`
}

// Step describes basic details of an step in the SCM
type Step struct {
	StepId           string            `json:"stepId"`
	StepName         string            `json:"stepName"`
	StepOrder        string            `json:"stepOrder"`
	ActorType        string            `json:"actorType"`
	AditionalInfoMap map[string]string `json:"aditionalInfoMap"`
	//AssetId      	 string            `json:"assetId"`
}

type AssetItem struct {
	AssetItemId      string            `json:"assetItemId"`
	CurrentOwnerId   string            `json:"currentOwnerId"`
	ProcessDate      string            `json:"processDate"`  //(date which currenct actor acquired the item)
	DeliveryDate     string            `json:"deliveryDate"` //(date which currenct actor received the item)
	OrderPrice       string            `json:"orderPrice"`
	ShippingPrice    string            `json:"shippingPrice"`
	Status           string            `json:"status"`
	AditionalInfoMap map[string]string `json:"aditionalInfoMap"`
	//AssetId      string	`json:"assetId"`
}

type asset struct {
	AssetId          string            `json:"assetId"`
	AssetItems       []AssetItem       `json:"assetItems"`
	Actors           []Actor           `json:"actors"`
	Steps            []Step            `json:"steps"`
	AditionalInfoMap map[string]string `json:"aditionalInfoMap"`
}
