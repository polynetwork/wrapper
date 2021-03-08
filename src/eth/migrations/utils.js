const fs = require('fs')

function getConfigContractAddresses() {
    return JSON.parse(
        fs.readFileSync(`${process.cwd()}/deployments/deployed.json`).toString()
    )
}

function getContractAddresses() {
    return JSON.parse(
        fs.readFileSync(`${process.cwd()}/deployments/development.json`).toString()
    )
}

function writeContractAddresses(contractAddresses) {
    fs.writeFileSync(
        `${process.cwd()}/deployments/development.json`,
        JSON.stringify(contractAddresses, null, 4) // Indent 4 spaces
    )
}

module.exports = {
    getConfigContractAddresses,
    getContractAddresses,
    writeContractAddresses
}
