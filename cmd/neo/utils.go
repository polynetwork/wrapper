package neo

import (
	"encoding/hex"
	"fmt"
	"github.com/joeqian10/neo-gogogo/wallet"
	"github.com/ontio/ontology/common"
	"github.com/skyinglyh1/poly_wrapper/log"
	"strings"
)

func ParseNeoAddr(s string) ([]byte, error) {
	if strings.Contains(s, "0x") {
		rb, err := hex.DecodeString(strings.TrimPrefix(s, "0x"))
		if err != nil {
			return nil, fmt.Errorf("ParseNeoAddr err: %v", err)
		}
		return common.ToArrayReverse(rb), nil
	} else if strings.Contains(s, "A") {
		addr, err := common.AddressFromBase58(s)
		if err != nil {
			return nil, fmt.Errorf("AddressFromBase58 err: %v", err)
		}
		return addr[:], nil
	}
	return hex.DecodeString(s)
}
func GetAccountByPassword(walletPath, pwd string) *wallet.Account {
	// open the NEO wallet
	w, err := wallet.NewWalletFromFile(walletPath) //
	if err != nil {
		log.Errorf("[GetAccountByPassword] Failed to open NEO wallet")
		return nil
	}

	if pwd == "" {
		log.Errorf("pls provide neo wallet pwd")
		return nil
	}
	err = w.DecryptAll(pwd)
	if err != nil {
		log.Errorf("[GetAccountByPassword] Failed to decrypt NEO account")
		return nil
	}
	if len(w.Accounts) == 0 {
		log.Errorf("[GetAccountByPassword] empty account")
		return nil
	}
	return w.Accounts[0]
}
