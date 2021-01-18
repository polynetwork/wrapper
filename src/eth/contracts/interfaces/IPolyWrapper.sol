pragma solidity >=0.6.0;

interface IPolyWrapper {
    function feeCollector() external view returns (address);
    function lockProxy() external view returns (address);
    function paused() external view returns (bool);

    function lock(
        address fromAsset,
        uint64 toChainId, 
        bytes calldata toAddress,
        uint amount,
        uint fee
    ) external payable;

    function speedUp(
        address fromAsset, 
        bytes calldata txHash,
        uint fee
    ) external payable;
    event PolyWrapperLock(address indexed fromAsset, address indexed sender, uint64 toChainId, bytes toAddress, uint net, uint fee);
    event PolyWrapperSpeedUp(address indexed fromAsset, bytes indexed txHash, address indexed sender, uint efee);
}