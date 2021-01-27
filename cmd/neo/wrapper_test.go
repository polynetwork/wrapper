package neo

import (
	"encoding/hex"
	"fmt"
	"github.com/skyinglyh1/poly_wrapper/config"
	"github.com/skyinglyh1/poly_wrapper/log"
	"io/ioutil"
	"math/big"
	"os"
	"testing"
)

func ReadFile(fileName string) ([]byte, error) {
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

func Test_DeployNeo_Wrapper(t *testing.T) {
	config.DefConfig.Init("./config.json")
	invoker, err := NewNeoInvoker(config.DefConfig.NeoUrl, config.DefConfig.NeoWallet, config.DefConfig.NeoWalletPwd)
	if err != nil {
		t.Fatal(err)
	}
	code, err := ReadFile("TODO")
	if err != nil {
		log.Errorf("ReadFile err: %v", err)
		return
	}
	invoker.Deploy(code)
}

func Test_Get_Poly_Neo_Wrapper(t *testing.T) {
	config.DefConfig.Init("./config.json")
	invoker, err := NewNeoInvoker(config.DefConfig.NeoUrl, config.DefConfig.NeoWallet, config.DefConfig.NeoWalletPwd)
	if err != nil {
		t.Fatal(err)
	}
	//ccmc, _ := ParseNeoAddr("0xe1695b1314a1331e3935481620417ed835669407")
	//neoLock, _ := ParseNeoAddr("0xedd2862dceb90b945210372d229f453f2b705f4f")
	polyNeoWrapper, _ := ParseNeoAddr("0xcd074cd290acc3d73c030784101afbcf40fd86a1")

	// check lockproxy
	{
		res, err := invoker.LockProxy(polyNeoWrapper)
		if err != nil {
			log.Debugf("GetLockProxy err: %v", err)
			return
		}
		log.Debugf("GetLockProxy is: %s", res)
	}
	//
	//{
	//	res, err := invoker.GetProxyOperator(neoLock)
	//	if err != nil {
	//		log.Errorf("GetOperator err: %v", err)
	//		return
	//	}
	//	log.Infof("neoLockProxy operator: %s", res)
	//}
	//
	//var toChainId uint64 = 7
	//{
	//	res, err := invoker.GetProxyHash(neoLock, toChainId)
	//	if err != nil {
	//		log.Errorf("GetProxyHash err: %v", err)
	//		return
	//	}
	//	log.Infof("neoLockProxy: %x, toChainId: %d, toChainProxyHash (little) : %s", neoLock, toChainId, res)
	//}
	//
	//fromAssetHashs := []string{
	//	"0x843e9f7a4ba7e062a53d7bbbe85cb35421704616",
	//	"0xbb01ac51a4c49bd28676274726497ab27ae8f66c",
	//	"0x17da3881ab2d050fea414c80b3fa8324d756f60e",
	//}
	//{
	//	froms := make([][]byte, 0)
	//	for _, from := range fromAssetHashs {
	//		f, _ := ParseNeoAddr(from)
	//		froms = append(froms, f)
	//	}
	//	res, err := invoker.GetAssetHashs(neoLock, toChainId, froms)
	//	if err != nil {
	//		log.Errorf("GetProxyHash err: %v", err)
	//		return
	//	}
	//	bals, err := invoker.GetAssetBalances(neoLock, froms)
	//	if err != nil {
	//		log.Errorf("GetAssetBalances err: %v", err)
	//		return
	//	}
	//
	//	for i, _ := range fromAssetHashs {
	//		log.Infof("neoLockProxy GetAssetHashs, toChainId: %d, from: %s, proxyBalance: %s, to(little): %s", toChainId, fromAssetHashs[i], bals[i].String(), res[i])
	//	}
	//}
}

func Test_Lock_PolyWrapper(t *testing.T) {
	config.DefConfig.Init("./config.json")
	invoker, err := NewNeoInvoker(config.DefConfig.NeoUrl, config.DefConfig.NeoWallet, config.DefConfig.NeoWalletPwd)
	if err != nil {
		t.Fatal(err)
	}
	//ccmc, _ := ParseNeoAddr("0xe1695b1314a1331e3935481620417ed835669407")
	//neoLock, _ := ParseNeoAddr("0xedd2862dceb90b945210372d229f453f2b705f4f")
	polyNeoWrapper, _ := ParseNeoAddr("0xcd074cd290acc3d73c030784101afbcf40fd86a1")

	fromAsset, _ := ParseNeoAddr("0x17da3881ab2d050fea414c80b3fa8324d756f60e") // nneo
	//fromAsset, _ := ParseNeoAddr("0x23535b6fd46b8f867ed010bab4c2bd8ef0d0c64f") // pnweth
	toAddr, _ := hex.DecodeString("352631d51332f8e6657ae94329d268eb68ca26f7")

	txHash, err := invoker.Lock(polyNeoWrapper, fromAsset, 79, toAddr, big.NewInt(2), big.NewInt(1), big.NewInt(0))
	if err != nil {
		log.Errorf("lock err: %v", err)
		return
	}
	log.Infof("txHash: %s", txHash)
}
