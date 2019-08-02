pragma solidity >=0.4.22 <0.7.0;

import "./dictionary.sol";
contract router {
    mapping(address => dictionary) dicRouter;
    constructor() public{

    }
    
    function insert(address dic) public {
        dicRouter[msg.sender] = dictionary(dic);
    }
    
    function insertAuto() public {
        dicRouter[msg.sender] = dicCreator();
    }
    
    function dicCreator() public returns (dictionary addr)
    {
        return new dictionary(msg.sender);
    }
    
    function show(address sender) public view returns (address dic){
        return address(dicRouter[sender]);
    }
}
