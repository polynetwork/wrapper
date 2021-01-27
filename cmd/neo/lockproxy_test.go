package neo

import (
	"encoding/hex"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/sc"
	"github.com/joeqian10/neo-gogogo/tx"
	"github.com/polynetwork/poly/common"
	"github.com/skyinglyh1/poly_wrapper/config"
	"github.com/skyinglyh1/poly_wrapper/log"
	"math/big"
	"testing"
)

func Test_GetOperator(t *testing.T) {
	config.DefConfig.Init("./config.json")
	invoker, err := NewNeoInvoker(config.DefConfig.NeoUrl, config.DefConfig.NeoWallet, config.DefConfig.NeoWalletPwd)
	if err != nil {
		t.Fatal(err)
	}
	ccmc, _ := ParseNeoAddr("0xe1695b1314a1331e3935481620417ed835669407")
	neoLock, _ := ParseNeoAddr("0xedd2862dceb90b945210372d229f453f2b705f4f")
	//polyNeoWrapper, _ := ParseNeoAddr("0x104057f879009326250ee1f5d60e2efd925024e6")

	// check if initialized
	{
		res, err := invoker.GetStorage(ccmc, hex.EncodeToString([]byte("IsInitGenesisBlock")))
		if err != nil {
			log.Errorf("Initialized height err: %v", err)
			return
		}
		log.Infof("ccmc initialized poly height on neo ccmc is: %s", res)
	}

	{
		res, err := invoker.GetProxyOperator(neoLock)
		if err != nil {
			log.Errorf("GetOperator err: %v", err)
			return
		}
		log.Infof("neoLockProxy operator: %s", res)
	}

	var toChainId uint64 = 79
	{
		res, err := invoker.GetProxyHash(neoLock, toChainId)
		if err != nil {
			log.Errorf("GetProxyHash err: %v", err)
			return
		}
		log.Infof("neoLockProxy: %x, toChainId: %d, toChainProxyHash (little) : %s", neoLock, toChainId, res)
	}

	fromAssetHashs := []string{
		"0x17da3881ab2d050fea414c80b3fa8324d756f60e",
	}
	{
		froms := make([][]byte, 0)
		for _, from := range fromAssetHashs {
			f, _ := ParseNeoAddr(from)
			froms = append(froms, f)
		}
		res, err := invoker.GetAssetHashs(neoLock, toChainId, froms)
		if err != nil {
			log.Errorf("GetProxyHash err: %v", err)
			return
		}
		bals, err := invoker.GetAssetBalances(neoLock, froms)
		if err != nil {
			log.Errorf("GetAssetBalances err: %v", err)
			return
		}

		for i, _ := range fromAssetHashs {
			log.Infof("neoLockProxy GetAssetHashs, toChainId: %d, from: %s, proxyBalance: %s, to(little): %s", toChainId, fromAssetHashs[i], bals[i].String(), res[i])
		}
	}
}

func Test_SendTxFromNeo(t *testing.T) {
	config.DefConfig.Init("./config_test.json")
	invoker, err := NewNeoInvoker(config.DefConfig.NeoUrl, config.DefConfig.NeoWallet, config.DefConfig.NeoWalletPwd)
	if err != nil {
		t.Fatal(err)
	}

	neoLockAddr, err := ParseNeoAddr("")
	if err != nil {
		return
	}
	hrc20AddrOnNeo, err := ParseNeoAddr("config.DefConfig.NeoHrc20")
	if err != nil {
		return
	}

	fromAsset := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: hrc20AddrOnNeo,
	}
	rawFrom, err := helper.AddressToScriptHash(invoker.Acc.Address)
	if err != nil {
		return
	}
	fromAddr := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: rawFrom.Bytes(),
	}
	toChainId := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(79)),
	}
	toAddr := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: common.ADDRESS_EMPTY[:],
	}
	amt := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(1)),
	}
	tb := tx.NewTransactionBuilder(config.DefConfig.NeoUrl)
	sb := sc.NewScriptBuilder()
	sb.MakeInvocationScript(neoLockAddr, "lock", []sc.ContractParameter{fromAsset, fromAddr, toChainId, toAddr, amt})
	script := sb.ToArray()

	itx, err := tb.MakeInvocationTransaction(script, rawFrom, nil, rawFrom, helper.Zero, helper.Zero)
	if err != nil && err.Error() != "" {
		return
	}
	err = tx.AddSignature(itx, invoker.Acc.KeyPair)
	if err != nil {
		return
	}
	response := tb.Client.SendRawTransaction(itx.RawTransactionString())
	if response.HasError() {
		return
	}

	log.Infof("successful to send %d hrc20 from neo to heco", itx.HashString())
}
