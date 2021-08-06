// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.6.0;

import "../eth-no-truffle/libs/token/ERC20/SafeERC20.sol";
import "../eth-no-truffle/libs/token/ERC20/IERC20.sol";
import "../eth-no-truffle/libs/access/Ownable.sol";
import "../eth-no-truffle/libs/utils/ReentrancyGuard.sol";
import "../eth-no-truffle/libs/math/SafeMath.sol";
import "../eth-no-truffle/libs/lifecycle/Pausable.sol";

import "../eth-no-truffle/interfaces/ILockProxy.sol";

contract PolyWrapper is Ownable, Pausable, ReentrancyGuard {
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    uint public chainId;
    address public feeCollector;
    address public FeeToken;

    ILockProxy public lockProxy;

    constructor(address _owner, address _feeToken, uint _chainId) public {
        require(_chainId != 0, "!legal");
        transferOwnership(_owner);
        chainId = _chainId;
        FeeToken = _feeToken;
    }

    function setFeeCollector(address collector) external onlyOwner {
        require(collector != address(0), "emtpy address");
        feeCollector = collector;
    }


    function setLockProxy(address _lockProxy) external onlyOwner {
        require(_lockProxy != address(0));
        lockProxy = ILockProxy(_lockProxy);
        require(lockProxy.managerProxyContract() != address(0), "not lockproxy");
    }

    function pause() external onlyOwner {
        _pause();
    }

    function unpause() external onlyOwner {
        _unpause();
    }


    function extractFee(address token) external {
        require(msg.sender == feeCollector, "!feeCollector");
        if (token == address(0)) {
            payable(msg.sender).transfer(address(this).balance);
        } else {
            IERC20(token).safeTransfer(feeCollector, IERC20(token).balanceOf(address(this)));
        }
    }
    
    function lock(address fromAsset, uint64 toChainId, bytes memory toAddress, uint amount, uint fee, uint id) external payable nonReentrant whenNotPaused {
        
        require(toChainId != chainId && toChainId != 0, "!toChainId");
        require(toAddress.length !=0, "empty toAddress");
        address addr;
        assembly { addr := mload(add(toAddress,0x14)) }
        require(addr != address(0),"zero toAddress");
        
        _pull(fromAsset, amount);

        amount = _checkoutFee(fromAsset, amount, fee);

        _push(fromAsset, toChainId, toAddress, amount);

        emit PolyWrapperLock(fromAsset, msg.sender, toChainId, toAddress, amount, fee, id);
    }

    function speedUp(address fromAsset, bytes memory txHash, uint fee) external payable nonReentrant whenNotPaused {
        _pull(fromAsset, fee);
        emit PolyWrapperSpeedUp(fromAsset, txHash, msg.sender, fee);
    }

    function _pull(address fromAsset, uint amount) internal {
        if (fromAsset == address(0)) {
            require(msg.value == amount, "insufficient msg.value");
        } else {
            IERC20(fromAsset).safeTransferFrom(msg.sender, address(this), amount);
        }
    }

    function _checkoutFee(address fromAsset, uint amount, uint fee) internal returns (uint) {
        if (fromAsset == FeeToken) {
            require(amount > fee, "amount less than fee");
            return amount.sub(fee);
        } else {
            if (FeeToken == address(0)) {
                require(msg.value >= fee, "insufficient msg.value");
            } else {
                IERC20(FeeToken).safeTransferFrom(msg.sender, address(this), amount);
            }
            return amount;
        }
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

    event PolyWrapperLock(address indexed fromAsset, address indexed sender, uint64 toChainId, bytes toAddress, uint net, uint fee, uint id);
    event PolyWrapperSpeedUp(address indexed fromAsset, bytes indexed txHash, address indexed sender, uint efee);

}