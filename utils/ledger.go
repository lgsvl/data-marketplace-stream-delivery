//
// Copyright (c) 2019 LG Electronics Inc.
// SPDX-License-Identifier: Apache-2.0
//

package utils

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lgsvl/data-marketplace-stream-delivery/resources"
)

//BlockchainClient is an interface that allows interacting with the blockchain
//go:generate counterfeiter -o ../fakes/fake_blockchain_client.go . BlockchainClient
type BlockchainClient interface {
	CheckContractID(string, map[string]string) (bool, error)
}

type hyperledger struct {
	ctx        context.Context
	logger     *log.Logger
	httpClient *http.Client
	endpoint   string
}

func NewHyperledgerClient(c context.Context, l *log.Logger) BlockchainClient {

	return &hyperledger{
		httpClient: &http.Client{},
		ctx:        c,
		logger:     l,
		endpoint:   os.Getenv("CHAIN_SERVICE_HOST"),
	}
}

func NewHyperledgerClientWithEndpoint(c context.Context, l *log.Logger, e string) BlockchainClient {

	return &hyperledger{
		httpClient: &http.Client{},
		ctx:        c,
		logger:     l,
		endpoint:   e,
	}
}

func (h *hyperledger) CheckContractID(contractID string, auth map[string]string) (bool, error) {
	checkContractURL := FormatURL(h.endpoint, resources.CheckContractIDPath, contractID)

	response, err := HttpExecuteWithHeader(h.logger, h.httpClient, "GET", checkContractURL, nil, auth)
	if err != nil {
		h.logger.Printf("utils.HttpExecute failed")
		return false, fmt.Errorf("utils.HttpExecute failed")
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		h.logger.Printf("error-finding-contract")
		return false, fmt.Errorf("error-finding-contract-response-code-%#v", response.StatusCode)
	}
	return true, nil
}
