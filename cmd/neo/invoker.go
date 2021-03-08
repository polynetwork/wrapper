package neo

import (
	"encoding/hex"
	"fmt"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/joeqian10/neo-gogogo/sc"
	"github.com/joeqian10/neo-gogogo/wallet"
	"github.com/ontio/ontology/common"
	"github.com/skyinglyh1/poly_wrapper/log"
	"math/big"
	"strings"
	"time"
)

type NeoInvoker struct {
	Cli *rpc.RpcClient
	Acc *wallet.Account
}

func NewNeoInvoker(url, walletPath, walletPwd string) (invoker *NeoInvoker, err error) {
	invoker = &NeoInvoker{}
	invoker.Cli = rpc.NewClient(url)
	invoker.Acc = GetAccountByPassword(walletPath, walletPwd)
	if invoker.Acc == nil {
		return nil, fmt.Errorf("NewNeoInvoker GetAccountByPassword error")
	}
	return invoker, nil
}

func (this *NeoInvoker) GetAssetBalances(neoLockProxy []byte, fromAssetHashs [][]byte) ([]*big.Int, error) {
	scriptBuilder := sc.NewScriptBuilder()

	for _, from := range fromAssetHashs {
		scriptBuilder.MakeInvocationScript(from, "balanceOf", []sc.ContractParameter{
			sc.ContractParameter{
				Type:  sc.ByteArray,
				Value: neoLockProxy,
			},
		})
	}
	script := scriptBuilder.ToArray()

	// create an InvocationTransaction
	response := this.Cli.InvokeScript(helper.BytesToHex(script), "0000000000000000000000000000000000000000")
	if response.HasError() || response.Result.State == "FAULT" {
		log.Errorf("invoke script error: %s", response.Error.Message)
		return nil, fmt.Errorf("[GetAssetBalances], InvokeScript err: %v", response.Error)
	}
	res := make([]*big.Int, len(fromAssetHashs))
	for i, stack := range response.Result.Stack {
		stack.Convert()

		if stack.Type == "ByteArray" {
			hexNumBs, _ := hex.DecodeString(stack.Value.(string))
			res[i] = common.BigIntFromNeoBytes(hexNumBs)
		}
	}
	return res, nil
}

func (this *NeoInvoker) GetStorage(neoLockProxy []byte, key string) (string, error) {
	addr, _ := common.AddressParseFromBytes(neoLockProxy)

	resp := this.Cli.GetStorage(addr.ToHexString(), key)
	if resp.HasError() {
		return "", fmt.Errorf("resp.Error.Message: %s", resp.Error.Message)
	}
	return resp.Result, nil
}

func WaitNeoTx(cli *rpc.RpcClient, hash helper.UInt256) {
	tick := time.NewTicker(100 * time.Millisecond)
	for range tick.C {
		res := cli.GetTransactionHeight(hash.String())
		if res.HasError() {
			if strings.Contains(res.Error.Message, "Unknown") {
				continue
			}
			log.Errorf("failed to get neo tx: %s", res.Error.Message)
			continue
		}
		if res.Result <= 0 {
			continue
		}
		log.Infof("capture neo tx %s", hash.String())
		break
	}
}
