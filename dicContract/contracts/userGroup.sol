pragma solidity >=0.4.22 <0.7.0;

import "./mapping/addrToBool.sol";
contract userGroup{
    //add securte as part of value ?
    address public creator;
    bool public deleted = false;
    addrToBool.itmap readmap;
    addrToBool.itmap updatemap;
    addrToBool.itmap changePermissionmap;
    
    
    event _updateUser(address key,bool read,bool update,bool change);
    event _changeStatusOfUserGroup(bool status);
    
    
    modifier onlyChangePermission {
        require(
            addrToBool.contains(changePermissionmap,msg.sender) || msg.sender == creator,
            "Only permisssion user can call this function."
        );
        _;
    }
    
    constructor(address sender) public{
        creator = sender;
        
    }
    
    // function init() public {
    //     addrToBool.insert(changePermissionmap,creator,true);
    // }
    
    function andOrupdateUser(address key,uint8 permisssion) external onlyChangePermission returns (bool res)
    {
        if(permisssion > 7 || permisssion < 0) return false;
        bool read = permisssion % 2 == 1 ? true: false;
        permisssion >> 1;
        bool update = permisssion % 2 == 1 ? true: false;
        permisssion >> 1;
        bool change = permisssion % 2 == 1 ? true: false;
        
        addrToBool.insert(readmap,key,read);
        addrToBool.insert(updatemap,key,update);
        addrToBool.insert(changePermissionmap,key,change);
        emit _updateUser(key,read,update,change);
        return true;
    }
    
    function changeStatusOfUserGroup(bool status) external onlyChangePermission {
        deleted = status;
        emit _changeStatusOfUserGroup(status);
    }
    function getUserPerssion(address key) external view returns (bool read,bool update,bool change)
    {
        if(key == creator) return (true,true,true);
        if(deleted == true) return(true,true,false);
        if(addrToBool.contains(readmap,key) == true) 
            read = readmap.data[key].value;
        else 
            read = false;
            
        if(addrToBool.contains(updatemap,key) == true)
            update = updatemap.data[key].value;
        else
            update = false;
            
        if(addrToBool.contains(changePermissionmap,key) == true)
            change = changePermissionmap.data[key].value;
        else
            change = false;
    }
}


