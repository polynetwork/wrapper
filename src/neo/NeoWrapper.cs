using System;
using System.ComponentModel;
using System.Numerics;
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using Neo.SmartContract.Framework.Services.System;

namespace neo_wrapper
{
    public class NeoWrapper : SmartContract
    {
        // keys to map values
        private static readonly byte[] OwnerKey = new byte[] { 0x01, 0x01 };
        private static readonly byte[] FeeCollectorKey = new byte[] { 0x01, 0x02 };
        private static readonly byte[] NeoLockProxyKey = new byte[] { 0x01, 0x03 };
        private static readonly byte[] PausedKey = new byte[] { 0x01, 0x04 };
        
        // default setting
        static readonly byte[] DefaultOwner = "Aa1WqKgXiwWrCKtUd3ib2N6rAhKMJ8j9jZ".ToScriptHash();
        static readonly BigInteger NeoChainId = 4;
        
        // events
        // OwnershipTransferred(oldOwner, newOwner)
        public static event Action<byte[], byte[]> OwnershipTransferred;
        // Paused(owner)
        public static event Action<byte[]> Paused;
        // Unpaused(owner)
        public static event Action<byte[]> Unpaused;
        // PolyWrapperLock(fromAsset, fromAddress, toChainId, toAddress, netAmount, fee, id)
        public static event Action<byte[], byte[], BigInteger, byte[], BigInteger, BigInteger, BigInteger> PolyWrapperLock;
        // SpeedUp(fromAsset, txHash, fromAddress, fee)
        public static event Action<byte[], byte[], byte[], BigInteger> PolyWrapperSpeedUp;


        public static object Main(string method, object[] args)
        {
            if (Runtime.Trigger == TriggerType.Verification)
            {
                return Runtime.CheckWitness(Owner());
            }
            else if (Runtime.Trigger == TriggerType.Application)
            {
                // authorized function
                if (method == "transferOwnership")
                {
                    byte[] newOwner = (byte[])args[0];
                    return TransferOwnership(newOwner);
                }
                if (method == "owner") return Owner();
                if (method == "setFeeCollector")
                {
                    byte[] feeCollector = (byte[])args[0];
                    return SetFeeCollector(feeCollector);
                }
                if (method == "feeCollector") return FeeCollector();
                if (method == "setLockProxy")
                {
                    byte[] lockProxy = (byte[])args[0];
                    return SetLockProxy(lockProxy);
                }
                if (method == "lockProxy") return LockProxy();
                if (method == "extractFee")
                {
                    byte[] token = (byte[])args[0];
                    return ExtractFee(token);
                }
                if (method == "pause") return Pause();
                if (method == "unpause") return Unpause();
                if (method == "paused") return paused();

                // business logic
                if (method == "lock")
                {
                    byte[] fromAsset = (byte[])args[0];
                    byte[] fromAddress = (byte[])args[1];
                    BigInteger toChainId = (BigInteger)args[2];
                    byte[] toAddress = (byte[])args[3];
                    BigInteger amount = (BigInteger)args[4];
                    BigInteger fee = (BigInteger)args[5];
                    BigInteger id = (BigInteger)args[6];
                    return Lock(fromAsset, fromAddress, toChainId, toAddress, amount, fee, id);
                }
                if (method == "speedUp")
                {
                    byte[] fromAsset = (byte[])args[0];
                    byte[] fromAddress = (byte[])args[1];
                    byte[] txHash = (byte[])args[2];
                    BigInteger fee = (BigInteger)args[3];
                    return SpeedUp(fromAsset, fromAddress, txHash, fee);
                }


            }
            return false;
        }

        [DisplayName("transferOwnership")]
        public static bool TransferOwnership(byte[] newOwner)
        {
            byte[] oldOwner = Owner();
            _assert(Runtime.CheckWitness(oldOwner), "!owner");
            _assert(newOwner.Length == 20, "len != 20");
            Storage.Put(OwnerKey, newOwner);
            OwnershipTransferred(oldOwner, newOwner);
            return true;
        }

        [DisplayName("owner")]
        public static byte[] Owner()
        {
            var owner = Storage.Get(OwnerKey);
            return owner.Length == 20 ? owner : DefaultOwner;
        }

        [DisplayName("setFeeCollector")]
        public static bool SetFeeCollector(byte[] collector)
        {
            _assert(Runtime.CheckWitness(Owner()), "!owner");
            _assert(collector.Length == 20, "len != 20");
            Storage.Put(FeeCollectorKey, collector);
            return true;
        }

        [DisplayName("feeCollector")]
        public static byte[] FeeCollector()
        {
            return Storage.Get(FeeCollectorKey);
        }

        [DisplayName("setLockProxy")]
        public static bool SetLockProxy(byte[] lockProxy)
        {
            _assert(Runtime.CheckWitness(Owner()), "!owner");
            _assert(lockProxy.Length == 20, "len != 20");
            Storage.Put(NeoLockProxyKey, lockProxy);
            return true;
        }
        [DisplayName("lockProxy")]
        public static byte[] LockProxy()
        {
            return Storage.Get(NeoLockProxyKey);
        }

        [DisplayName("extractFee")]
        public static bool ExtractFee(byte[] token)
        {
            _assert(Runtime.CheckWitness(FeeCollector()), "!feeCollector");
            byte[] self = ExecutionEngine.ExecutingScriptHash;
            BigInteger fee = ((Func<string, object[], BigInteger>)token.ToDelegate())("balanceOf", new object[] { self });
            if (fee > 0)
            {
                byte[] feeCollector = FeeCollector();
                bool result = ((Func<string, object[], bool>)token.ToDelegate())("transfer", new object[] { self, feeCollector, fee });
                _assert(result, "trasnfer fail");
            }
            return true;
        }

        [DisplayName("lock")]
        public static bool Lock(byte[] fromAsset, byte[] fromAddress, BigInteger toChainId, byte[] toAddress, BigInteger amount, BigInteger fee, BigInteger id)
        {
            _assert(Runtime.CheckWitness(fromAddress), "!fromAddress");
            _assert(toChainId != NeoChainId && toChainId != 0, "!toChainId");
            _assert(amount > fee, "amount less than fee");
            _assert(!paused(), "paused");
            _pull(fromAsset, fromAddress, amount);
            _push(fromAsset, ExecutionEngine.ExecutingScriptHash, toChainId, toAddress, amount - fee);
            PolyWrapperLock(fromAsset, fromAddress, toChainId, toAddress, amount - fee, fee, id);
            return true;
        }

        [DisplayName("speedUp")]
        public static bool SpeedUp(byte[] fromAsset, byte[] fromAddress, byte[] txHash, BigInteger fee)
        {
            _assert(Runtime.CheckWitness(fromAddress), "!fromAddress");
            _assert(!paused(), "paused");
            _pull(fromAsset, fromAddress, fee);
            PolyWrapperSpeedUp(fromAsset, txHash, fromAddress, fee);
            return true;
        }

        private static bool _pull(byte[] fromAsset, byte[] fromAddress, BigInteger amount)
        {
            byte[] self = ExecutionEngine.ExecutingScriptHash;
            bool result = ((Func<string, object[], bool>)fromAsset.ToDelegate())("transfer", new object[] { fromAddress, self, amount });
            _assert(result, "_pull fail");
            return true;
        }
        private static bool _push(byte[] fromAsset, byte[] fromAddress, BigInteger toChainId, byte[] toAddress, BigInteger amount)
        {
            byte[] lockProxy = LockProxy();
            bool result = ((Func<string, object[], bool>)lockProxy.ToDelegate())("lock", new object[] { fromAsset, fromAddress, toChainId, toAddress, amount });
            _assert(result, "lock fail");
            return true;
        }

        [DisplayName("pause")]
        public static bool Pause()
        {
            _assert(Runtime.CheckWitness(Owner()), "!owner");
            Storage.Put(PausedKey, new byte[] { 0x01 });
            Paused(Owner());
            return true;
        }
        [DisplayName("unpause")]
        public static bool Unpause()
        {
            _assert(Runtime.CheckWitness(Owner()), "!owner");
            Storage.Delete(PausedKey);
            Unpaused(Owner());
            return true;
        }
        [DisplayName("paused")]
        public static bool paused()
        {
            byte[] paused = Storage.Get(PausedKey);
            return paused.Length == 1;
        }
        private static void _assert(bool condition, string message)
        {
            if (!condition)
            {
                Runtime.Notify("Fault:" + message);
                throw new Exception(message);
            }
        }


    }
}
