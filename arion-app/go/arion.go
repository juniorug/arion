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

// AssetTransferSmartContract provides functions for managing an asset
type AssetTransferSmartContract struct {
	contractapi.Contract
}

// Actor describes basic details of an actor in the SCM
type Actor struct {
	ActorID          string            `json:"actorID"`
	ActorType        string            `json:"actorType"`
	ActorName        string            `json:"actorName"`
	AditionalInfoMap map[string]string `json:"aditionalInfoMap"`
	//assetID        string            `json:"assetID"`
}

// Step describes basic details of an step in the SCM
type Step struct {
	StepID           string            `json:"stepID"`
	StepName         string            `json:"stepName"`
	StepOrder        uint              `json:"stepOrder"`
	ActorType        string            `json:"actorType"`
	AditionalInfoMap map[string]string `json:"aditionalInfoMap"`
	//AssetID      	 string            `json:"assetID"`
}

// AssetItem describes ths single item in the SCM
type AssetItem struct {
	AssetItemID      string            `json:"assetItemID"`
	CurrentOwnerID   string            `json:"currentOwnerID"`
	ProcessDate      string            `json:"processDate"`  //(date which currenct actor acquired the item)
	DeliveryDate     string            `json:"deliveryDate"` //(date which currenct actor received the item)
	OrderPrice       string            `json:"orderPrice"`
	ShippingPrice    string            `json:"shippingPrice"`
	Status           string            `json:"status"`
	AditionalInfoMap map[string]string `json:"aditionalInfoMap"`
	//AssetID      string	`json:"assetID"`
}

// Asset is the collection of assetItems in the SCM
type Asset struct {
	AssetID          string            `json:"assetID"`
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
func (s *AssetTransferSmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	aditionalInfoMap := make(map[string]string)

	actors := []Actor{
		Actor{ActorID: "1", ActorType: "rawMaterial", ActorName: "James Johnson", AditionalInfoMap: aditionalInfoMap},
		Actor{ActorID: "2", ActorType: "manufacturer", ActorName: "Helena Smith", AditionalInfoMap: aditionalInfoMap},
	}

	steps := []Step{
		Step{StepID: "1", StepName: "exploration", StepOrder: 1, ActorType: "rawMaterial", AditionalInfoMap: aditionalInfoMap},
		Step{StepID: "2", StepName: "Production", StepOrder: 2, ActorType: "manufacturer", AditionalInfoMap: aditionalInfoMap},
	}

	assetItems := []AssetItem{
		AssetItem{
			AssetItemID:      "1",
			CurrentOwnerID:   "1",
			ProcessDate:      "2020-03-07T15:04:05",
			DeliveryDate:     "",
			OrderPrice:       "",
			ShippingPrice:    "",
			Status:           "order initiated",
			AditionalInfoMap: aditionalInfoMap,
		},
		AssetItem{
			AssetItemID:      "2",
			CurrentOwnerID:   "1",
			ProcessDate:      "2020-03-07T15:04:05",
			DeliveryDate:     "",
			OrderPrice:       "",
			ShippingPrice:    "",
			Status:           "order initiated",
			AditionalInfoMap: aditionalInfoMap,
		},
	}

	assets := []Asset{
		Asset{AssetID: "Gravel", AssetItems: assetItems, Actors: actors, Steps: steps, AditionalInfoMap: aditionalInfoMap},
	}

	for i, asset := range assets {
		assetAsBytes, _ := json.Marshal(asset)
		err := ctx.GetStub().PutState("ASSET_"+strconv.Itoa(i), assetAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

// CreateActor adds a new Actor to the world state with given details
func (s *AssetTransferSmartContract) CreateActor(ctx contractapi.TransactionContextInterface, actorID string, actorType string, actorName string, aditionalInfoMap map[string]string) error {
	actorJSON, err := ctx.GetStub().GetState("ACTOR_" + actorID)

	if err != nil {
		return fmt.Errorf("Failed to read the data from world state: %s", err)
	}

	if actorJSON != nil {
		return fmt.Errorf("The actor %s already exists", actorID)
	}

	actor := Actor{
		ActorID:          actorID,
		ActorType:        actorType,
		ActorName:        actorName,
		AditionalInfoMap: aditionalInfoMap,
	}

	actorAsBytes, err := json.Marshal(actor)

	if err != nil {
		return err
	}

	return ctx.GetStub().PutState("ACTOR_"+actorID, actorAsBytes)
}

// CreateStep adds a new Step to the world state with given details
func (s *AssetTransferSmartContract) CreateStep(ctx contractapi.TransactionContextInterface, stepID string, stepName string, stepOrder uint, actorType string, aditionalInfoMap map[string]string) error {
	stepJSON, err := ctx.GetStub().GetState("STEP_" + stepID)

	if err != nil {
		return fmt.Errorf("Failed to read the data from world state: %s", err)
	}

	if stepJSON != nil {
		return fmt.Errorf("The step %s already exists", stepID)
	}

	step := Step{
		StepID:           stepID,
		StepName:         stepName,
		StepOrder:        stepOrder,
		ActorType:        actorType,
		AditionalInfoMap: aditionalInfoMap,
	}
	stepAsBytes, err := json.Marshal(step)

	if err != nil {
		return err
	}

	return ctx.GetStub().PutState("STEP_"+stepID, stepAsBytes)
}

// CreateAssetItem adds a new AssetItem to the world state with given details
func (s *AssetTransferSmartContract) CreateAssetItem(ctx contractapi.TransactionContextInterface, assetItemID string, currentOwnerID string, processDate string, deliveryDate string, orderPrice string, shippingPrice string, status string, aditionalInfoMap map[string]string) error {
	assetItemJSON, err := ctx.GetStub().GetState("ASSET_ITEM_" + assetItemID)

	if err != nil {
		return fmt.Errorf("Failed to read the data from world state: %s", err)
	}

	if assetItemJSON != nil {
		return fmt.Errorf("The assetItem %s already exists", assetItemID)
	}

	assetItem := AssetItem{
		AssetItemID:      assetItemID,
		CurrentOwnerID:   currentOwnerID,
		ProcessDate:      processDate,
		DeliveryDate:     deliveryDate,
		OrderPrice:       orderPrice,
		ShippingPrice:    shippingPrice,
		Status:           status,
		AditionalInfoMap: aditionalInfoMap,
	}
	assetItemAsBytes, err := json.Marshal(assetItem)

	if err != nil {
		return err
	}

	return ctx.GetStub().PutState("ASSET_ITEM_"+assetItemID, assetItemAsBytes)
}

// CreateAsset adds a new Asset to the world state with given details
func (s *AssetTransferSmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, assetID string, assetItems []AssetItem, actors []Actor, steps []Step, aditionalInfoMap map[string]string) error {
	assetJSON, err := ctx.GetStub().GetState("ASSET_" + assetID)

	if err != nil {
		return fmt.Errorf("Failed to read the data from world state: %s", err)
	}

	if assetJSON != nil {
		return fmt.Errorf("The asset %s already exists", assetID)
	}

	asset := Asset{
		AssetID:          assetID,
		AssetItems:       assetItems,
		Actors:           actors,
		Steps:            steps,
		AditionalInfoMap: aditionalInfoMap,
	}
	assetAsBytes, err := json.Marshal(asset)

	if err != nil {
		return err
	}

	return ctx.GetStub().PutState("ASSET_"+assetID, assetAsBytes)
}

// QueryActor returns the actor stored in the world state with given id
func (s *AssetTransferSmartContract) QueryActor(ctx contractapi.TransactionContextInterface, actorID string) (*Actor, error) {
	actorAsBytes, err := ctx.GetStub().GetState("ACTOR_" + actorID)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if actorAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", actorID)
	}

	actor := new(Actor)
	_ = json.Unmarshal(actorAsBytes, actor)

	return actor, nil
}

// QueryStep returns the step stored in the world state with given id
func (s *AssetTransferSmartContract) QueryStep(ctx contractapi.TransactionContextInterface, stepID string) (*Step, error) {
	stepAsBytes, err := ctx.GetStub().GetState("STEP_" + stepID)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if stepAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", stepID)
	}

	step := new(Step)
	_ = json.Unmarshal(stepAsBytes, step)

	return step, nil
}

// QueryAssetItem returns the AssetItem stored in the world state with given id
func (s *AssetTransferSmartContract) QueryAssetItem(ctx contractapi.TransactionContextInterface, assetItemID string) (*AssetItem, error) {
	assetItemAsBytes, err := ctx.GetStub().GetState("ASSET_ITEM_" + assetItemID)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if assetItemAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", assetItemID)
	}

	assetItem := new(AssetItem)
	_ = json.Unmarshal(assetItemAsBytes, assetItem)

	return assetItem, nil
}

// QueryAsset returns the Asset stored in the world state with given id
func (s *AssetTransferSmartContract) QueryAsset(ctx contractapi.TransactionContextInterface, assetID string) (*Asset, error) {
	assetAsBytes, err := ctx.GetStub().GetState("ASSET_" + assetID)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if assetAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", assetID)
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
func (s *AssetTransferSmartContract) ChangeAssetItemOwner(ctx contractapi.TransactionContextInterface, assetItemID string, newOwnerID string, orderPrice string, shippingPrice string, status string, aditionalInfo map[string]string) error {
	assetItem, err := s.QueryAssetItem(ctx, "ASSET_ITEM_"+assetItemID)
	if err != nil {
		return err
	}

	assetItem.CurrentOwnerID = newOwnerID
	assetItem.ProcessDate = time.Now().Format("2006-01-02 15:04:05")
	assetItem.OrderPrice = orderPrice
	assetItem.ShippingPrice = shippingPrice
	assetItem.Status = status
	for key, value := range aditionalInfo {
		assetItem.AditionalInfoMap[key] = value
	}

	assetItemAsBytes, _ := json.Marshal(assetItem)
	return ctx.GetStub().PutState("ASSET_ITEM_"+assetItemID, assetItemAsBytes)
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(AssetTransferSmartContract))

	if err != nil {
		fmt.Printf("Error create Arion chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting Arion chaincode: %s", err.Error())
	}
}
