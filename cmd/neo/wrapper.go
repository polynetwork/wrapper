package neo

import (
	"fmt"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/sc"
	"github.com/joeqian10/neo-gogogo/tx"
	"github.com/joeqian10/neo-gogogo/wallet"
	"github.com/skyinglyh1/poly_wrapper/log"
	"math/big"
)

func (this *NeoInvoker) Deploy(code []byte) {
	txBuilder := tx.NewTransactionBuilder(this.Cli.Endpoint.String())
	wh := wallet.NewWalletHelper(txBuilder, this.Acc)
	ContractHash, err := wh.DeployContract(code, "0710", "05", true, true, true,
		"neoPolyWrapper", "1.0", "poly", "email", "desc")
	if err != nil {
		log.Errorf("DeployContract err: %v", err)
		return
	}
	log.Infof("Neo Deploy poly wrapper, contract Hash: %s", ContractHash.String())

}

func (this *NeoInvoker) LockProxy(neoLockProxy []byte) (string, error) {
	scriptBuilder := sc.NewScriptBuilder()
	args := []sc.ContractParameter{}
	scriptBuilder.MakeInvocationScript(neoLockProxy, "lockProxy", args)
	script := scriptBuilder.ToArray()

	// create an InvocationTransaction
	response := this.Cli.InvokeScript(helper.BytesToHex(script), "0000000000000000000000000000000000000000")
	if response.HasError() || response.Result.State == "FAULT" {
		log.Errorf("invoke script error: %s", response.Error.Message)
		return "", fmt.Errorf("[LockProxy], InvokeScript err: %v", response.Error)
	}
	for _, stack := range response.Result.Stack {
		stack.Convert()

		if stack.Type == "ByteArray" {
			return stack.Value.(string), nil // little endian
		}
	}
	return "", fmt.Errorf("lockproxy not found")
}

func (this *NeoInvoker) Owner(neoLockProxy []byte) (string, error) {
	scriptBuilder := sc.NewScriptBuilder()
	args := []sc.ContractParameter{}
	scriptBuilder.MakeInvocationScript(neoLockProxy, "owner", args)
	script := scriptBuilder.ToArray()

	// create an InvocationTransaction
	response := this.Cli.InvokeScript(helper.BytesToHex(script), "0000000000000000000000000000000000000000")
	if response.HasError() || response.Result.State == "FAULT" {
		log.Errorf("invoke script error: %s", response.Error.Message)
		return "", fmt.Errorf("[Owner], InvokeScript err: %v", response.Error)
	}
	for _, stack := range response.Result.Stack {
		stack.Convert()

		if stack.Type == "ByteArray" {
			return stack.Value.(string), nil // little endian
		}
	}
	return "", fmt.Errorf("Owner not found")
}

func (this *NeoInvoker) FeeCollector(neoLockProxy []byte) (string, error) {
	scriptBuilder := sc.NewScriptBuilder()
	args := []sc.ContractParameter{}
	scriptBuilder.MakeInvocationScript(neoLockProxy, "feeCollector", args)
	script := scriptBuilder.ToArray()

	// create an InvocationTransaction
	response := this.Cli.InvokeScript(helper.BytesToHex(script), "0000000000000000000000000000000000000000")
	if response.HasError() || response.Result.State == "FAULT" {
		log.Errorf("invoke script error: %s", response.Error.Message)
		return "", fmt.Errorf("[FeeCollector], InvokeScript err: %v", response.Error)
	}
	for _, stack := range response.Result.Stack {
		stack.Convert()

		if stack.Type == "ByteArray" {
			return stack.Value.(string), nil // little endian
		}
	}
	return "", fmt.Errorf("FeeCollector not found")
}

func (this *NeoInvoker) Lock(neoPolyWrapper []byte, fromAssetHash []byte, toChainId uint64, toAddress []byte, amount *big.Int, fee *big.Int, id *big.Int) (string, error) {
	from, err := ParseNeoAddr(this.Acc.Address)
	if err != nil {
		return "", fmt.Errorf("[LockFromWrapper], ParseNeoAddr acct: %s,  err: %v", this.Acc.Address, err)
	}
	fromUint160, err := helper.UInt160FromBytes(from)
	if err != nil {
		return "", fmt.Errorf("[LockFromWrapper], Uint160FromBytes err: %v", err)
	}
	fromAssetHashValue := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: fromAssetHash,
	}
	fromAddressValue := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: from,
	}
	toChainIdValue := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(toChainId)),
	}
	toAddressValue := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: toAddress,
	}
	amountValue := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *amount,
	}
	feeValue := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *fee,
	}
	idValue := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *id,
	}
	// build script
	scriptBuilder := sc.NewScriptBuilder()
	args := []sc.ContractParameter{fromAssetHashValue, fromAddressValue, toChainIdValue, toAddressValue, amountValue, feeValue, idValue}
	scriptBuilder.MakeInvocationScript(neoPolyWrapper, "lock", args)
	script := scriptBuilder.ToArray()

	// create an InvocationTransaction
	tb := tx.NewTransactionBuilder(this.Cli.Endpoint.String())

	sysFee := helper.Fixed8FromFloat64(0)
	netFee := helper.Fixed8FromFloat64(0)

	itx, err := tb.MakeInvocationTransaction(script, fromUint160, nil, fromUint160, sysFee, netFee)
	if err != nil {
		return "", fmt.Errorf("[LockFromWrapper] tb.MakeInvocationTransaction error: %s", err)
	}
	// sign transaction
	err = tx.AddSignature(itx, this.Acc.KeyPair)
	if err != nil {
		return "", fmt.Errorf("[LockFromWrapper] tx.AddSignature error: %s", err)
	}

	rawTxString := itx.RawTransactionString()

	// send the raw transaction
	response := this.Cli.SendRawTransaction(rawTxString)
	if response.HasError() {
		return "", fmt.Errorf("[LockFromWrapper] SendRawTransaction error: %s,  RawTransactionString: %s",
			response.ErrorResponse.Error.Message, rawTxString)
	}
	log.Infof("Neo LockFromWrapper, txHash: %s", itx.HashString())
	WaitNeoTx(this.Cli, itx.Hash)

	return itx.HashString(), nil
}

func (this *NeoInvoker) ExtractFee(neoPolyWrapper []byte, token []byte) (string, error) {
	from, err := ParseNeoAddr(this.Acc.Address)
	if err != nil {
		return "", fmt.Errorf("[ExtractFee], ParseNeoAddr acct: %s,  err: %v", this.Acc.Address, err)
	}
	fromUint160, err := helper.UInt160FromBytes(from)
	if err != nil {
		return "", fmt.Errorf("[ExtractFee], Uint160FromBytes err: %v", err)
	}
	fromAssetHashValue := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: token,
	}

	// build script
	scriptBuilder := sc.NewScriptBuilder()
	args := []sc.ContractParameter{fromAssetHashValue}
	scriptBuilder.MakeInvocationScript(neoPolyWrapper, "extractFee", args)
	script := scriptBuilder.ToArray()

	// create an InvocationTransaction
	tb := tx.NewTransactionBuilder(this.Cli.Endpoint.String())

	sysFee := helper.Fixed8FromFloat64(0)
	netFee := helper.Fixed8FromFloat64(0)

	itx, err := tb.MakeInvocationTransaction(script, fromUint160, nil, fromUint160, sysFee, netFee)
	if err != nil {
		return "", fmt.Errorf("[ExtractFee] tb.MakeInvocationTransaction error: %s", err)
	}
	// sign transaction
	err = tx.AddSignature(itx, this.Acc.KeyPair)
	if err != nil {
		return "", fmt.Errorf("[ExtractFee] tx.AddSignature error: %s", err)
	}

	rawTxString := itx.RawTransactionString()

	// send the raw transaction
	response := this.Cli.SendRawTransaction(rawTxString)
	if response.HasError() {
		return "", fmt.Errorf("[ExtractFee] SendRawTransaction error: %s,  RawTransactionString: %s",
			response.ErrorResponse.Error.Message, rawTxString)
	}
	log.Infof("Neo ExtractFee, txHash: %s", itx.HashString())
	WaitNeoTx(this.Cli, itx.Hash)

	return itx.HashString(), nil
}
