OntCversion = '2.0.0'
"""
Smart contract for wrap locking cross chain asset between Ontology and other chains provided by poly
"""

from ontology.interop.System.Action import RegisterAction
from ontology.interop.System.Storage import Put, GetContext, Get, Delete
from ontology.interop.System.ExecutionEngine import GetExecutingScriptHash
from ontology.interop.System.Runtime import CheckWitness
from ontology.libont import bytearray_reverse
from ontology.interop.System.App import DynamicAppCall

# Keys
OWNER_KEY = "owner"
FEE_COLLECTOR_KEY = "feeCollector"
LOCK_PROXY_KEY = "lockProxy"
PAUSE_KEY = "pause"
# Constant
OntChainIdOnPoly = 3

# Event
TransferOwnership = RegisterAction("TransferOwnership", "oldOwner", "newOwner")
PolyWrapperLock = RegisterAction("PolyWrapperLock", "fromAsset", "msgSender", "toChainId", "toAddress", "amount", "fee",
                                 "id")


def Main(operation, args):
    if operation == "init":
        assert (len(args) == 3)
        owner = args[0]
        feeCollector = args[1]
        lockProxy = args[2]
        return init(owner, feeCollector, lockProxy)
    if operation == "setFeeCollector":
        assert (len(args) == 1)
        return setFeeCollector(args[0])
    if operation == "getFeeCollector":
        return getFeeCollector()
    if operation == "setLockProxy":
        assert (len(args) == 1)
        lockProxy = args[0]
        return setLockProxy(lockProxy)
    if operation == "getLockProxy":
        return getLockProxy()
    if operation == "pause":
        return pause()
    if operation == "unpause":
        return unpause()
    if operation == "ifPause":
        return ifPause()
    if operation == "lock":
        assert (len(args) == 7)
        fromAddress = args[0]
        fromAsset = args[1]
        toChainId = args[2]
        toAddress = args[3]
        amount = args[4]
        fee = args[5]
        id = args[6]
        return lock(fromAddress, fromAsset, toChainId, toAddress, amount, fee, id)

    if operation == "transferOwnership":
        assert (len(args) == 1)
        newOwner = args[0]
        return transferOwnership(newOwner)
    if operation == "setLockProxy":
        assert (len(args) == 1)
        return setLockProxy(args[0])

    return True


def init(owner, feeCollector, lockProxy):
    """
    owner: address type
    feeCollector: address receiving fee
    lockProxy: lockProxy hash
    """
    assert (CheckWitness(owner))
    Put(GetContext(), OWNER_KEY, owner)
    Put(GetContext(), FEE_COLLECTOR_KEY, feeCollector)
    Put(GetContext(), LOCK_PROXY_KEY, bytearray_reverse(lockProxy))
    TransferOwnership("", owner)
    return True


def setFeeCollector(feeCollector):
    """
    :param feeCollector: address
    :return:
    """
    assert (CheckWitness(getOwner()))
    Put(GetContext(), FEE_COLLECTOR_KEY, feeCollector)
    return True


def getFeeCollector():
    return Get(GetContext(), FEE_COLLECTOR_KEY)


def setLockProxy(lockProxy):
    """
    :param lockProxy: ont lock proxy
    :return:
    """
    assert (CheckWitness(getOwner()))
    Put(GetContext(), LOCK_PROXY_KEY, bytearray_reverse(lockProxy))
    return True


def getLockProxy():
    return Get(GetContext(), LOCK_PROXY_KEY)


def pause():
    assert (CheckWitness(getOwner()))
    Put(GetContext(), PAUSE_KEY, True)
    return True


def unpause():
    assert (CheckWitness(getOwner()))
    Delete(GetContext(), PAUSE_KEY)
    return True


def ifPause():
    return Get(GetContext(), PAUSE_KEY)


def lock(fromAddress, fromAsset, toChainId, toAddress, amount, fee, id):
    """
    :param fromAddress: from Account
    :param fromAsset: asset address, not hash, should be reversed hash
    :param toChainId: !=3
    :param toAddress: bytearray
    :param amount: > fee
    :param fee: >= 0
    :param id: like uin in eth
    :return:
    """
    assert (CheckWitness(fromAddress))
    assert (not ifPause())
    assert (toChainId != 0 and toChainId != OntChainIdOnPoly)
    assert (len(toAddress) > 0)
    assert (amount > fee)
    assert (fee > 0)

    lockProxy = getLockProxy()
    toAssethash = DynamicAppCall(lockProxy, 'getAssetHash', [fromAsset, toChainId])
    assert (len(toAssethash) > 0)

    # call lock contract lock
    res = DynamicAppCall(fromAsset, "approve", [fromAddress, lockProxy, amount - fee])
    assert (res == True)
    res = DynamicAppCall(lockProxy, 'lock', [fromAsset, fromAddress, toChainId, toAddress, amount - fee])
    assert (res == True)

    # transfer fee to fee collector
    feeCollector = getFeeCollector()
    assert (len(feeCollector) == 20)
    res = DynamicAppCall(fromAsset, 'transfer', [fromAddress, feeCollector, fee])
    assert (res == True)
    PolyWrapperLock(fromAsset, fromAddress, toChainId, toAddress, amount - fee, fee, id)
    return True


def transferOwnership(newOwner):
    assert (CheckWitness(getOwner()))
    Put(GetContext(), OWNER_KEY, newOwner)
    TransferOwnership(getOwner(), newOwner)
    return True


def getOwner():
    return Get(GetContext(), OWNER_KEY)
