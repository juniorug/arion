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
	"encoding/json"
	"fmt"
	"strconv"
	"time"

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
	StepOrder        uint              `json:"stepOrder"`
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
type Asset struct {
	AssetId          string            `json:"assetId"`
	AssetItems       []AssetItem       `json:"assetItems"`
	Actors           []Actor           `json:"actors"`
	Steps            []Step            `json:"steps"`
	AditionalInfoMap map[string]string `json:"aditionalInfoMap"`
}
// QueryResult structure used for handling result of query
type QueryResult struct {
	Key    string `json:"Key"`
	Record *AssetItem
}
// InitLedger adds a base set of items to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	
	aditionalInfoMap := make(map[string]string)
	
	actors := []Actor{
		Actor{ActorId: "1", ActorType: "rawMaterial", ActorName: "James Johnson", AditionalInfoMap: aditionalInfoMap},
		Actor{ActorId: "2", ActorType: "manufacturer", ActorName: "Helena Smith", AditionalInfoMap: aditionalInfoMap},
	}

	steps := []Step{
		Step{StepId: "1", StepName: "exploration", StepOrder: 1, ActorType: "rawMaterial", AditionalInfoMap: aditionalInfoMap},
		Step{StepId: "2", StepName: "Production", StepOrder: 2, ActorType: "manufacturer", AditionalInfoMap: aditionalInfoMap},
	}

	assetItems := []AssetItem{
		AssetItem{
			AssetItemId: "1", 
			CurrentOwnerId: "1", 
			ProcessDate: "2020-03-07T15:04:05", 
			DeliveryDate: "", 
			OrderPrice: "", 
			ShippingPrice: "",
			Status: "order initiated",
			AditionalInfoMap: aditionalInfoMap,
		},
		AssetItem{
			AssetItemId: "2", 
			CurrentOwnerId: "1", 
			ProcessDate: "2020-03-07T15:04:05", 
			DeliveryDate: "", 
			OrderPrice: "", 
			ShippingPrice: "",
			Status: "order initiated",
			AditionalInfoMap: aditionalInfoMap,
		},
	}

	assets := []Asset{
		Asset{AssetId: "Gravel", AssetItems: assetItems, Actors: actors, Steps: steps, AditionalInfoMap: aditionalInfoMap},
	}

	for i, asset := range assets {
		assetsAsBytes, _ := json.Marshal(assets)
		err := ctx.GetStub().PutState("ASSET"+strconv.Itoa(i), assetsAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

// CreateActor adds a new Actor to the world state with given details
func (s *SmartContract) CreateActor(ctx contractapi.TransactionContextInterface, actorId string, actorType string, actorName string, aditionalInfoMap map[string]string) error {
	actor := Actor{
		ActorId:   actorId,
		ActorType:  actorType,
		ActorName: actorName,
		AditionalInfoMap:  aditionalInfoMap,
	}
	actorAsBytes, _ := json.Marshal(actor)
	return ctx.GetStub().PutState("ACTOR"+ActorId, actorAsBytes)
}

// CreateStep adds a new Step to the world state with given details
func (s *SmartContract) CreateStep(ctx contractapi.TransactionContextInterface, stepId string, stepName string, stepOrder string, actorType: string, aditionalInfoMap map[string]string) error {
	step := Step{
		StepId:   stepId,
		StepName:  stepName,
		StepOrder: stepOrder,
		ActorType: actorType,
		AditionalInfoMap:  aditionalInfoMap,
	}
	stepAsBytes, _ := json.Marshal(step)
	return ctx.GetStub().PutState("STEP"+stepId, stepAsBytes)
}

// CreateAssetItem adds a new AssetItem to the world state with given details
func (s *SmartContract) CreateAssetItem(ctx contractapi.TransactionContextInterface, assetItemId string, currentOwnerId string, processDate string, deliveryDate: string,  shippingPrice: string,  status: string, aditionalInfoMap map[string]string) error {
	assetItem := AssetItem{
		AssetItemId:   assetItemId,
		CurrentOwnerId:  currentOwnerId,
		ProcessDate: processDate,
		DeliveryDate: deliveryDate,
		OrderPrice:   orderPrice,
		ShippingPrice:  shippingPrice,
		Status: status,
		AditionalInfoMap:  aditionalInfoMap,
	}
	assetItemAsBytes, _ := json.Marshal(assetItem)
	return ctx.GetStub().PutState("ASSET_ITEM"+assetItemId, assetItemAsBytes)
}

// CreateAsset adds a new Asset to the world state with given details
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, assetId string, assetItems []AssetItem, actors []Actor, steps: []Step, aditionalInfoMap map[string]string) error {
	asset := Asset{
		AssetId:   assetId,
		AssetItems:  assetItems,
		Actors: actors,
		Steps: steps,
		AditionalInfoMap:  aditionalInfoMap,
	}
	assetAsBytes, _ := json.Marshal(asset)
	return ctx.GetStub().PutState("ASSET"+assetId, assetAsBytes)
}

// QueryActor returns the actor stored in the world state with given id
func (s *SmartContract) QueryActor(ctx contractapi.TransactionContextInterface, actorId string) (*Actor, error) {
	actorAsBytes, err := ctx.GetStub().GetState("ACTOR"+actorId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if actorAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", actorId)
	}

	actor := new(Actor)
	_ = json.Unmarshal(actorAsBytes, actor)

	return actor, nil
}

// QueryStep returns the step stored in the world state with given id
func (s *SmartContract) QueryStep(ctx contractapi.TransactionContextInterface, stepId string) (*Step, error) {
	stepAsBytes, err := ctx.GetStub().GetState("STEP"+stepId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if stepAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", stepId)
	}

	step := new(Step)
	_ = json.Unmarshal(stepAsBytes, step)

	return step, nil
}

// QueryAssetItem returns the AssetItem stored in the world state with given id
func (s *SmartContract) QueryStep(ctx contractapi.TransactionContextInterface, assetItemId string) (*AssetItem, error) {
	assetItemAsBytes, err := ctx.GetStub().GetState("ASSET_ITEM"+assetItemId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if assetItemAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", assetItemId)
	}

	assetItem := new(AssetItem)
	_ = json.Unmarshal(assetItemAsBytes, assetItem)

	return assetItem, nil
}

// QueryAsset returns the Asset stored in the world state with given id
func (s *SmartContract) QueryStep(ctx contractapi.TransactionContextInterface, assetId string) (*Asset, error) {
	assetAsBytes, err := ctx.GetStub().GetState("ASSET"+assetId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if assetAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", assetId)
	}

	asset := new(Asset)
	_ = json.Unmarshal(assetAsBytes, asset)

	return asset, nil
}
//TODO: create 'allQueries' methods:
//QueryAllActors
//QueryAllSteps
//QueryAllAssetItems
//QueryAllAssets


// ChangeAssetItemOwner updates the owner field of assetItem with given id in world state
func (s *SmartContract) ChangeAssetItemOwner(ctx contractapi.TransactionContextInterface, AssetItemId string, newOwnerId string) error {
	assetItem, err := s.QueryAssetItem(ctx, "ASSET_ITEM"+assetItemId)
	if err != nil {
		return err
	}

	assetItem.CurrentOwnerId = newOwnerId
	assetItem.ProcessDate = time.Now()
	assetItem.OrderPrice = orderPrice
	assetItem.ShippingPrice = shippingPrice
	assetItem.Status = Status
	for key, value := range AditionalInfo {
		assetItem.AditionalInfoMap[key] = value
	}
}