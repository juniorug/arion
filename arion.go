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
	actorId string
	//assetId        string            `json:"assetId"`
	actorType        string            `json:"actorType"`
	actorName        string            `json:"actorName"`
	aditionalInfoMap map[string]string `json:"aditionalInfoMap"`
}

// Step describes basic details of an step in the SCM
type Step struct {
	stepId string
	//assetId        string            `json:"assetId"`
	stepName  string `json:"stepName"`
	stepOrder string `json:"stepOrder"`
	actorType string `json:"actorType"`
}
