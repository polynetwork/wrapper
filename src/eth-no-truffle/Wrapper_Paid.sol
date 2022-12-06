// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.6.12;

import "./libs/token/ERC20/SafeERC20.sol";
import "./libs/token/ERC20/IERC20.sol";
import "./libs/access/Ownable.sol";
import "./libs/utils/ReentrancyGuard.sol";
import "./libs/math/SafeMath.sol";
import "./libs/lifecycle/Pausable.sol";

import "./interfaces/ILockProxy.sol";

contract PolyWrapper is Ownable, Pausable, ReentrancyGuard {
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    uint constant FEE_DECIMALS = 10**8;
    uint public bridgeFeeRate = 30000; // 0.03% by default
    uint public chainId;
    address public relayerFeeCollector;
    address public bridgeFeeCollector;

    ILockProxy public lockProxy;

    constructor(address _owner, address _relayerFeeCollector, address _bridgeFeeCollector, uint _chainId) public {
        require(_chainId != 0, "!legal");
        transferOwnership(_owner);
        relayerFeeCollector = _relayerFeeCollector;
        bridgeFeeCollector = _bridgeFeeCollector;
        chainId = _chainId;
    }

    function setRelayerFeeCollector(address collector) external {
        require(_msgSender() == relayerFeeCollector, "msg sender not current relayerFeeCollector");
        require(collector != address(0), "emtpy address");
        relayerFeeCollector = collector;
    }

    function setBridgeFeeCollector(address collector) external {
        require(_msgSender() == bridgeFeeCollector, "msg sender not current bridgeFeeCollector");
        require(collector != address(0), "emtpy address");
        bridgeFeeCollector = collector;
    }

    function setLockProxy(address _lockProxy) external onlyOwner {
        require(_lockProxy != address(0));
        lockProxy = ILockProxy(_lockProxy);
        require(lockProxy.managerProxyContract() != address(0), "not lockproxy");
    }

    function setBridgeFeeRate(uint _bridgeFeeRate) external onlyOwner {
        bridgeFeeRate = _bridgeFeeRate;
    }

    function pause() external onlyOwner {
        _pause();
    }

    function unpause() external onlyOwner {
        _unpause();
    }

    function rescueFund(address token) external onlyOwner {
        if (token == address(0)) {
            payable(msg.sender).transfer(address(this).balance);
        } else {
            IERC20(token).safeTransfer(_msgSender(), IERC20(token).balanceOf(address(this)));
        }
    }
    
    function lock(address fromAsset, uint64 toChainId, bytes memory toAddress, uint amount, uint relayerFee, uint id) external payable nonReentrant whenNotPaused {
        
        require(toChainId != chainId && toChainId != 0, "!toChainId");
        require(toAddress.length !=0, "empty toAddress");
        address addr;
        assembly { addr := mload(add(toAddress,0x14)) }
        require(addr != address(0),"zero toAddress");
        
        _pull(fromAsset, amount);

        amount = _checkoutFee(fromAsset, amount, relayerFee);

        _push(fromAsset, toChainId, toAddress, amount);

        emit PolyWrapperLock(fromAsset, msg.sender, toChainId, toAddress, amount, relayerFee, id);
    }

    function _pull(address fromAsset, uint amount) internal {
        if (fromAsset == address(0)) {
            require(msg.value == amount, "insufficient ether");
        } else {
            IERC20(fromAsset).safeTransferFrom(msg.sender, address(this), amount);
        }
    }

    // take fee in the form of ether
    function _checkoutFee(address fromAsset, uint amount, uint relayerFee) internal returns (uint) {
        if (fromAsset == address(0)) {
            require(msg.value == amount, "insufficient/too_much ether");
            if (relayerFee != 0) {
                payable(relayerFeeCollector).transfer(relayerFee);
                amount -= relayerFee;
            }
            if (bridgeFeeRate != 0) {
                uint bridgeFee = _getBridgeFee(amount);
                amount -= bridgeFee;
                payable(bridgeFeeCollector).transfer(bridgeFee);
            }
            return amount;
        } else {
            require(msg.value == relayerFee, "insufficient/too_much ether");
            if (relayerFee != 0) {
                payable(relayerFeeCollector).transfer(relayerFee);
            }
            if (bridgeFeeRate != 0) {
                uint bridgeFee = _getBridgeFee(amount);
                amount -= bridgeFee;
                IERC20(fromAsset).safeTransfer(bridgeFeeCollector, bridgeFee);
            }
            return amount;
        }
    }

    function _getBridgeFee(uint amount) internal view returns (uint) {
        return amount * bridgeFeeRate / FEE_DECIMALS;
    }

    function _push(address fromAsset, uint64 toChainId, bytes memory toAddress, uint amount) internal {
        if (fromAsset == address(0)) {
            require(lockProxy.lock{value: amount}(fromAsset, toChainId, toAddress, amount), "lock ether fail");
        } else {
            IERC20(fromAsset).safeApprove(address(lockProxy), 0);
            IERC20(fromAsset).safeApprove(address(lockProxy), amount);
            require(lockProxy.lock(fromAsset, toChainId, toAddress, amount), "lock erc20 fail");
        }
    }

    event PolyWrapperLock(address indexed fromAsset, address indexed sender, uint64 toChainId, bytes toAddress, uint net, uint relayerFee, uint id);
    event PolyWrapperSpeedUp(address indexed fromAsset, bytes indexed txHash, address indexed sender, uint efee);

}