const PolyWrapper = artifacts.require("PolyWrapper");
const utils = require('./utils')

module.exports = async function (deployer, network, accounts) {
    console.log("2_deploy_polywrapper.js, network: ", network)

    const deployedConfig = utils.getConfigContractAddresses();
    if (deployedConfig[network].lockproxy == "") {
        console.log("pls config network: ", network, "in deployed.json");
        return
    }
    
    const admin = accounts[0];
    // ethereum
    await deployer.deploy(PolyWrapper, admin, '2');
    // // bsc
    // await deployer.deploy(PolyWrapper, admin, '6');
    // // heco
    // await deployer.deploy(PolyWrapper, admin, '7');

    let config = utils.getContractAddresses();;
    
    config[network] = {
        owner: admin,
        polywrapper: PolyWrapper.address,
        lockproxy: deployedConfig[network].lockproxy
    };
    
    utils.writeContractAddresses(config);

    let pw = await PolyWrapper.at(config[network].polywrapper);
    await pw.setLockProxy(config[network].lockproxy, {from: admin});
    await pw.setFeeCollector(admin, {from: admin});

}