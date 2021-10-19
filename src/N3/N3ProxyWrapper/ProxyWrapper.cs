using System;
using System.Numerics;
using Neo;
using Neo.SmartContract;
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Native;
using Neo.SmartContract.Framework.Services;

namespace N3ProxyWrapper
{
    [ContractPermission("*")]
    public class ProxyWrapper : SmartContract
    {
        //TODO:check default lock proxy
        [InitialValue("0x4487494dac2f7eb68bdae009cacd6de88243e542", ContractParameterType.Hash160)]
        private static readonly UInt160 lockProxy = default;
        //TODO:check default superowner
        [InitialValue("NVGUQ1qyL4SdSm7sVmGVkXetjEsvw2L3NT", ContractParameterType.Hash160)]
        private static readonly UInt160 superOwner = default;
        //TODO:check default feeCollector
        [InitialValue("NVGUQ1qyL4SdSm7sVmGVkXetjEsvw2L3NT", ContractParameterType.Hash160)]
        private static readonly UInt160 feeCollector = default;

        private static readonly StorageMap OwnerMap = new StorageMap(Storage.CurrentContext, "owner");
        private static readonly StorageMap ProxyMap = new StorageMap(Storage.CurrentContext, "proxy");
        private static readonly StorageMap PauseMap = new StorageMap(Storage.CurrentContext, "pause");

        public static event Action<object> notify;
        public static event Action<UInt160, UInt160, BigInteger, byte[], BigInteger, BigInteger, object> polyWrapperLock;
        public static void _deploy(object data, bool update)
        {
            OwnerMap.Put("superOwner", superOwner);
            OwnerMap.Put("feeCollector", feeCollector);
            ProxyMap.Put("proxy", lockProxy);
        }

        public static void OnNEP17Payment(UInt160 from, BigInteger amount, object data)
        {
            return;
        }

        public static bool CheckSuperOwner() => Runtime.CheckWitness((UInt160)OwnerMap.Get("superOwner"));

        public static bool TransferOwnerShip(UInt160 newOwner)
        {
            Assert(CheckSuperOwner(), "Forbidden");
            OwnerMap.Put("superOwner", newOwner);
            return true;
        }

        public static bool Update(ByteString nefFile, string manifest)
        {
            Assert(CheckSuperOwner(), "Forbidden");
            ContractManagement.Update(nefFile, manifest);
            return true;
        }

        public static bool SetFeeCollector(UInt160 address)
        {
            Assert(CheckSuperOwner(), "Forbidden");
            OwnerMap.Put("feeCollector", address);
            return true;
        }

        [Safe]
        public static UInt160 FeeCollector()
        {
            return (UInt160)OwnerMap.Get("feeCollector");
        }

        public static bool Pause()
        {
            Assert(CheckSuperOwner(), "Forbidden");
            PauseMap.Put("global", 1);
            return true;
        }

        public static bool Unpause()
        {
            Assert(CheckSuperOwner(), "Forbidden");
            PauseMap.Put("global", 0);
            return true;
        }

        [Safe]
        public static bool IsPause()
        {
            ByteString rawState = PauseMap.Get("global");
            return rawState is null ? false : (BigInteger)rawState == 1;
        }

        public static bool SetProxy(UInt160 address)
        {
            Assert(CheckSuperOwner(), "Forbidden");
            ProxyMap.Put("proxy", address);
            return true;
        }

        [Safe]
        public static UInt160 Proxy()
        {
            return (UInt160)ProxyMap.Get("proxy");
        }

        public static bool ExtractFee(UInt160 token)
        {
            Assert(Runtime.CheckWitness(FeeCollector()), "Forbidden");
            BigInteger balance = (BigInteger)Contract.Call(token, "balanceOf", CallFlags.All, Runtime.ExecutingScriptHash);
            Contract.Call(token, "transfer", CallFlags.All, Runtime.ExecutingScriptHash, FeeCollector(), balance, null);
            return true;
        }

        public static bool Lock(UInt160 fromAsset, UInt160 fromAddress, BigInteger toChainId, byte[] toAddress, BigInteger amount, BigInteger fee, BigInteger id)
        {
            Assert(!IsPause(), "Paused");
            Assert(Runtime.CheckWitness(fromAddress), "Forbidden");
            if (fee != 0)
            {
                Contract.Call(fromAsset, "transfer", CallFlags.All, fromAddress, Runtime.ExecutingScriptHash, fee, null);
            }
            Contract.Call(Proxy(), "lock", CallFlags.All, fromAsset, fromAddress, toChainId, toAddress, amount - fee);
            polyWrapperLock(fromAsset, fromAddress, toChainId, toAddress, amount - fee, fee, id);
            return true;
        }

        public static void Assert(bool condition, string message)
        {
            if (!condition)
            {
                notify(message);
                throw new Exception();
            }
        }
    }
}
