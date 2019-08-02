const libaddrToBool = artifacts.require("addrToBool");
const libaddrToMapping = artifacts.require("addrToMapping");
const userGroups = artifacts.require("userGroups");

module.exports = function(deployer) {
	deployer.deploy(libaddrToBool);
	deployer.deploy(libaddrToMapping);
	deployer.link(libaddrToBool,userGroups);
	deployer.link(libaddrToMapping,userGroups);
	deployer.deploy(userGroups);
};
