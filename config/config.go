/*
* Copyright (C) 2020 The poly network Authors
* This file is part of The poly network library.
*
* The poly network is free software: you can redistribute it and/or modify
* it under the terms of the GNU Lesser General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* The poly network is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
* GNU Lesser General Public License for more details.
* You should have received a copy of the GNU Lesser General Public License
* along with The poly network . If not, see <http://www.gnu.org/licenses/>.
 */
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/polynetwork/poly-io-test/log"
)

const (
	CM_BTCX  = "btcx"
	CM_ETHX  = "ethx"
	CM_BNBX  = "bnbx"
	CM_HTX   = "htx"
	CM_ERC20 = "erc20x"
	CM_BEP20 = "bep20x"
	CM_HRC20 = "hrc20x"
	CM_ONT   = "ontx"
	CM_ONG   = "ongx"
	CM_OEP4  = "oep4x"
)

type BindProxyStruct struct {
	FromChainId uint64 `json:"fromChainId"`
	FromProxy   string `json:"fromProxy"`
	ToChainId   uint64 `json:"toChainId"`
	ToProxy     string `json:"toProxy"`
}

type BindAssetStruct struct {
	FromChainId uint64 `json:"fromChainId"`
	FromAsset   string `json:"fromAsset"`
	ToChainId   uint64 `json:"toChainId"`
	ToAsset     string `json:"toAsset"`
}

//Config object used by ontology-instance
type TestConfig struct {
	NeoChainID uint64 `json:"neoChainId,omitempty"`

	// neo chain conf
	NeoUrl       string `json:"neoUrl,omitempty"`
	NeoWif       string `json:"neoWif,omitempty"`
	NeoWallet    string `json:"neoWallet,omitempty"`
	NeoWalletPwd string `json:"neoWalletPwd,omitempty"`

	ProxyToBind []BindProxyStruct `json:"proxyToBind,omitempty"`
	AssetToBind []BindAssetStruct `json:"assetToBind,omitempty"`
}

//Default config instance
var DefConfig = NewDefaultTestConfig()
var DefaultConfigFile = "./config.json"
var BtcNet *chaincfg.Params

//NewTestConfig retuen a TestConfig instance
func NewTestConfig() *TestConfig {
	return &TestConfig{}
}

func NewDefaultTestConfig() *TestConfig {
	var config = NewTestConfig()
	err := config.Init(DefaultConfigFile)
	if err != nil {
		return &TestConfig{}
	}
	return config
}

//Init TestConfig with a config file
func (conf *TestConfig) Init(fileName string) error {
	err := conf.loadConfig(fileName)
	if err != nil {
		return fmt.Errorf("loadConfig error:%s", err)
	}
	return nil
}

/**
Load JSON Configuration
*/
func (conf *TestConfig) loadConfig(fileName string) error {
	data, err := conf.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, conf)
	if err != nil {
		return fmt.Errorf("json.Unmarshal TestConfig:%s error:%s", data, err)
	}
	return nil
}

/**
Read  File to bytes
*/
func (conf *TestConfig) ReadFile(fileName string) ([]byte, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("OpenFile %s error %s", fileName, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Errorf("File %s close error %s", fileName, err)
		}
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll %s error %s", fileName, err)
	}
	return data, nil
}

/**
Save Test Configuration To json file
*/
func (conf *TestConfig) Save(fileName string) error {
	data, err := json.MarshalIndent(conf, "", "\t")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fileName, data, 0644); err != nil {
		return fmt.Errorf("failed to write conf file: %v", err)
	}
	return nil
}
