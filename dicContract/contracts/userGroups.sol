pragma solidity >=0.4.22 <0.7.0;

import "./mapping/addrToMapping.sol";
import "./userGroup.sol";

//Everyone can add their own user groups
contract userGroups {
    address public creater;

    addrToMapping.preitmap public ownerToGroups;
    event groups(string name,address value);


    constructor() public{
        creater = msg.sender;
    }

    event _updateUserGroup(address add,string name,address value);
    event _deleteGroupByName(string name);
    event _deleteOwnerRec(address owner,string name);
    function updateUserGroupByAddress(string memory name,address value)
    public
    returns (bool res)
    {
        addrToMapping.insert(ownerToGroups,msg.sender,name,value);
        emit _updateUserGroup(msg.sender,name,value);
        return true;
    }

    function updateUserGroupWithNewGroup(string memory name)
    public
    returns (bool res)
    {
        //if(addrToMapping.contains(ownerToGroups.data[msg.sender].value,name) == true)
        // {
        //     return false;
        // }
        userGroup value = userGroupCreator();
        addrToMapping.insert(ownerToGroups,msg.sender,name,address(value));
        emit _updateUserGroup(msg.sender,name,address(value));
        return true;
    }

    function userGroupCreator() public returns (userGroup addr)
    {
        return new userGroup(msg.sender);
    }

//show usergroups
    function getGroupCount(address owner) public view returns (uint){
        if(addrToMapping.contains(ownerToGroups,owner) == false)
            return 0;
        return ownerToGroups.data[owner].value.size;
    }

    function getGroupByIndex(address owner,uint index) public view returns (string memory,address)
    {
        if(addrToMapping.contains(ownerToGroups,owner) == false)
            return (string(""),address(0));
        return addrToMapping.iterate_get(ownerToGroups.data[owner].value,index);
    }

    // function showGroups(address owner) public returns (bool res)
    // {
    //     if(addrToMapping.contains(ownerToGroups,owner) == false)
    //         return false;
    //     addrToMapping.itmap storage itmap = ownerToGroups.data[owner].value;
    //     for(uint256 i = addrToMapping.iterate_start(itmap);addrToMapping.iterate_valid(itmap,i); i = addrToMapping.iterate_next(itmap,i))
    //     {
    //         (string memory name, address value) = addrToMapping.iterate_get(itmap,i);
    //         emit groups(name, value);
    //     }
    //     return true;
    // }

    function deleteGroupByName(string memory name) public returns (bool res)
    {
        if(addrToMapping.contains(ownerToGroups,msg.sender) == false)
            return false;

        if(addrToMapping.contains(ownerToGroups.data[msg.sender].value,name) == false)
        {
            return false;
        }
        else
        {
            address addr = ownerToGroups.data[msg.sender].value.data[name].value;
            userGroup uG = userGroup(addr);
            uG.changeStatusOfUserGroup(true);
            addrToMapping.remove(ownerToGroups.data[msg.sender].value,name);
            emit _deleteGroupByName(name);
            return true;
        }
    }

    function deleteOwnerRec() public returns(bool res)
    {
        if(addrToMapping.contains(ownerToGroups,msg.sender) == false)
            return false;
        addrToMapping.itmap storage itmap = ownerToGroups.data[msg.sender].value;
        for(uint256 i = addrToMapping.iterate_start(itmap);addrToMapping.iterate_valid(itmap,i); i = addrToMapping.iterate_next(itmap,i))
        {
            (string memory name, address value) = addrToMapping.iterate_get(itmap,i);
            addrToMapping.remove(ownerToGroups.data[msg.sender].value,name);
            emit _deleteOwnerRec(msg.sender,name);
        }
        addrToMapping.remove(ownerToGroups,msg.sender);

    }
}



