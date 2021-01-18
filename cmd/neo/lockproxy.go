package neo

import (
	"encoding/hex"
	"fmt"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/sc"
	"github.com/joeqian10/neo-gogogo/tx"
	"github.com/polynetwork/poly/common"
	"github.com/skyinglyh1/poly_wrapper/log"
	"math/big"
)

func (this *NeoInvoker) BindProxyHash(neoLockProxy []byte, toChainId uint64, toProxyHash []byte) error {
	from, err := ParseNeoAddr(this.Acc.Address)
	if err != nil {
		return fmt.Errorf("[BindProxyHash], ParseNeoAddr acct: %s,  err: %v", this.Acc.Address, err)
	}
	fromUint160, err := helper.UInt160FromBytes(from)
	if err != nil {
		return fmt.Errorf("[BindProxyHash], Uint160FromBytes err: %v", err)
	}
	tci := big.NewInt(int64(toChainId))
	toChainIdValue := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *tci,
	}
	toProxyHashValue := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: toProxyHash,
	}
	// build script
	scriptBuilder := sc.NewScriptBuilder()
	args := []sc.ContractParameter{toChainIdValue, toProxyHashValue}
	scriptBuilder.MakeInvocationScript(neoLockProxy, "bindProxyHash", args)
	script := scriptBuilder.ToArray()

	// create an InvocationTransaction
	tb := tx.NewTransactionBuilder(this.Cli.Endpoint.String())

	sysFee := helper.Fixed8FromFloat64(0)
	netFee := helper.Fixed8FromFloat64(0)

	itx, err := tb.MakeInvocationTransaction(script, fromUint160, nil, fromUint160, sysFee, netFee)
	if err != nil {
		return fmt.Errorf("[BindProxyHash] tb.MakeInvocationTransaction error: %s", err)
	}
	// sign transaction
	err = tx.AddSignature(itx, this.Acc.KeyPair)
	if err != nil {
		return fmt.Errorf("[BindProxyHash] tx.AddSignature error: %s", err)
	}

	rawTxString := itx.RawTransactionString()

	// send the raw transaction
	response := this.Cli.SendRawTransaction(rawTxString)
	if response.HasError() {
		return fmt.Errorf("[BindProxyHash] SendRawTransaction error: %s,  RawTransactionString: %s",
			response.ErrorResponse.Error.Message, rawTxString)
	}

	log.Infof("Neo bindProxyHash, txHash: %s", itx.HashString())
	WaitNeoTx(this.Cli, itx.Hash)

	return nil
}

func (this *NeoInvoker) BindAssetHash(neoLockProxy []byte, fromAssetHash []byte, toChainId uint64, toAssetHash []byte) (string, error) {
	from, err := ParseNeoAddr(this.Acc.Address)
	if err != nil {
		return "", fmt.Errorf("[BindAssetHash], ParseNeoAddr acct: %s,  err: %v", this.Acc.Address, err)
	}
	fromUint160, err := helper.UInt160FromBytes(from)
	if err != nil {
		return "", fmt.Errorf("[BindAssetHash], Uint160FromBytes err: %v", err)
	}
	fromAssetHashValue := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: fromAssetHash,
	}
	toChainIdValue := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(toChainId)),
	}
	toAssetHashValue := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: toAssetHash,
	}
	// build script
	scriptBuilder := sc.NewScriptBuilder()
	args := []sc.ContractParameter{fromAssetHashValue, toChainIdValue, toAssetHashValue}
	scriptBuilder.MakeInvocationScript(neoLockProxy, "bindAssetHash", args)
	script := scriptBuilder.ToArray()

	// create an InvocationTransaction
	tb := tx.NewTransactionBuilder(this.Cli.Endpoint.String())

	sysFee := helper.Fixed8FromFloat64(0)
	netFee := helper.Fixed8FromFloat64(0)

	itx, err := tb.MakeInvocationTransaction(script, fromUint160, nil, fromUint160, sysFee, netFee)
	if err != nil {
		return "", fmt.Errorf("[BindAssetHash] tb.MakeInvocationTransaction error: %s", err)
	}
	// sign transaction
	err = tx.AddSignature(itx, this.Acc.KeyPair)
	if err != nil {
		return "", fmt.Errorf("[BindAssetHash] tx.AddSignature error: %s", err)
	}

	rawTxString := itx.RawTransactionString()

	// send the raw transaction
	response := this.Cli.SendRawTransaction(rawTxString)
	if response.HasError() {
		return "", fmt.Errorf("[BindAssetHash] SendRawTransaction error: %s,  RawTransactionString: %s",
			response.ErrorResponse.Error.Message, rawTxString)
	}
	log.Infof("Neo bindAssetHash, txHash: %s", itx.HashString())
	WaitNeoTx(this.Cli, itx.Hash)

	return itx.HashString(), nil
}

func (this *NeoInvoker) GetProxyOperator(neoLockProxy []byte) (string, error) {
	scriptBuilder := sc.NewScriptBuilder()
	args := []sc.ContractParameter{}
	scriptBuilder.MakeInvocationScript(neoLockProxy, "getOperator", args)
	script := scriptBuilder.ToArray()

	// create an InvocationTransaction
	response := this.Cli.InvokeScript(helper.BytesToHex(script), "0000000000000000000000000000000000000000")
	if response.HasError() || response.Result.State == "FAULT" {
		log.Errorf("invoke script error: %s", response.Error.Message)
		return "", fmt.Errorf("[GetProxyOperator], InvokeScript err: %v", response.Error)
	}
	for _, stack := range response.Result.Stack {
		stack.Convert()

		if stack.Type == "ByteArray" {
			x, e := hex.DecodeString(stack.Value.(string))
			if e != nil {
				return "", fmt.Errorf("bytearray: %s, decodestring err: %v", stack.Value.(string), e)
			}
			base58Addr, _ := common.AddressParseFromBytes(x)
			return base58Addr.ToBase58(), nil
		}
	}
	return "", fmt.Errorf("operator not found")
}

func (this *NeoInvoker) GetProxyHash(neoLockProxy []byte, toChainId uint64) (string, error) {
	scriptBuilder := sc.NewScriptBuilder()
	args := []sc.ContractParameter{
		sc.ContractParameter{
			Type:  sc.Integer,
			Value: *big.NewInt(int64(toChainId)),
		},
	}
	scriptBuilder.MakeInvocationScript(neoLockProxy, "getProxyHash", args)
	script := scriptBuilder.ToArray()

	// create an InvocationTransaction
	response := this.Cli.InvokeScript(helper.BytesToHex(script), "0000000000000000000000000000000000000000")
	if response.HasError() || response.Result.State == "FAULT" {
		log.Errorf("invoke script error: %s", response.Error.Message)
		return "", fmt.Errorf("[GetProxyHash], InvokeScript err: %v", response.Error)
	}
	for _, stack := range response.Result.Stack {
		stack.Convert()
		if stack.Type == "ByteArray" {
			return stack.Value.(string), nil
		}
	}
	return "", fmt.Errorf("GetProxyHash not found")
}

func (this *NeoInvoker) GetAssetHashs(neoLockProxy []byte, toChainId uint64, fromAssetHashs [][]byte) ([]string, error) {
	scriptBuilder := sc.NewScriptBuilder()

	for _, from := range fromAssetHashs {
		scriptBuilder.MakeInvocationScript(neoLockProxy, "getAssetHash", []sc.ContractParameter{
			sc.ContractParameter{
				Type:  sc.ByteArray,
				Value: from,
			},
			sc.ContractParameter{
				Type:  sc.Integer,
				Value: *big.NewInt(int64(toChainId)),
			},
		})
	}
	script := scriptBuilder.ToArray()

	// create an InvocationTransaction
	response := this.Cli.InvokeScript(helper.BytesToHex(script), "0000000000000000000000000000000000000000")
	if response.HasError() || response.Result.State == "FAULT" {
		log.Errorf("invoke script error: %s", response.Error.Message)
		return nil, fmt.Errorf("[GetProxyHash], InvokeScript err: %v", response.Error)
	}
	res := make([]string, len(fromAssetHashs))
	for i, stack := range response.Result.Stack {
		stack.Convert()

		if stack.Type == "ByteArray" {
			res[i] = stack.Value.(string)
		}
	}
	return res, nil
}
