pragma solidity >=0.4.22 <0.7.0;

import "./userGroup.sol";

contract dictionary {
    address         public  creator;
    //string          public  name;
    //int             public  fileSize = 0;
    bool            public  accessType = false;
    uint            public  createtime;
    uint            public  changetime;  
    _dicMapping     dicmapping;
    userGroup       public  uG;
    
    
    modifier onlyCreator {
        require(
            (msg.sender == creator),
            "Only creator can call this function."
        );
        _;
    }
    
    struct _fileInfo {
        string      name;
        uint        fileType;       //0 dic ,1 file
        uint        createTime;
        uint        changeTime;
        userGroup   accessControl;
        uint        fileSize;
        bool        isRec;
        dictionary  subDic;
        string      fileHash;
    }
    
    struct _dicMapping {
        mapping(string => IndexValue) data;
        KeyFlag[] keys;
        uint size;
    }
    struct IndexValue { uint keyIndex; _fileInfo value; }
    struct KeyFlag { string key; bool deleted; }
    
    
    constructor(address sender) public{
        creator = sender;
    }

    
    function getFileInfo(string memory filename) public view returns (
        bool        isExit,
        string      memory fname,
        uint        fileType,       //0 dic ,1 file
        uint        createTime,
        uint        changeTime,
        address     accessControl,
        uint        fSize,
        bool        isRec,
        address     subDic,
        string      memory fileHash) 
    {
        _fileInfo memory f = dicmapping.data[filename].value;
        
        return (
            dicmapping.data[filename].keyIndex > 0,
            f.name,
            f.fileType,
            f.createTime,
            f.changeTime,
            address(f.accessControl),
            f.fileSize,
            f.isRec,
            address(f.subDic),
            f.fileHash
            );
    }

    function updateAccessControl(address newUG) public onlyCreator {
        accessType = true;
        uG = userGroup(newUG);
        for (uint256 i = iterate_start(dicmapping); iterate_valid(dicmapping, i); i = iterate_next(dicmapping, i))
        {
            string memory _name = dicmapping.keys[i].key;
            _fileInfo storage filePointer = dicmapping.data[_name].value;
            if(filePointer.isRec == true) 
            {
                filePointer.accessControl = uG;
                if( filePointer.fileType == 0)
                {
                    filePointer.subDic.updateAccessControl(newUG);
                }
            }
        }
    }
    
    function getUserPerssion(address key) public view returns (bool read,bool update) {
        if(accessType == false) return (true,true);
        if(key == creator) return (true,true);
        (bool _read,bool _update,bool change) = uG.getUserPerssion(key);
        return (_read,_update);
    }
    
    // function updateDicInfo(string memory newName,address newUG) public returns (bool res)
    // {
    //     (bool read,bool update) = getUserPerssion(msg.sender);
    //     if(update == true)
    //     {
    //         //name = newName;
    //         updateAccessControl(newUG);
    //         return true;
    //     } 
    //     else 
    //         return false;
    // }
    
    function addOrUpdateFile(string memory newName, uint newFileType,address newUG,uint newFileSize,
    bool newIsRec,address newsubDic,string memory newFileBlockHash) public returns (bool res)
    {
        (bool read,bool update) = getUserPerssion(msg.sender);
        if(update == true)
        {
            
            uint keyIndex = dicmapping.data[newName].keyIndex;
            dicmapping.data[newName].value = _fileInfo(newName,newFileType,now,now,
                newIsRec == true ? uG:userGroup(newUG)
                ,newFileSize,newIsRec,dictionary(newsubDic),newFileBlockHash);
                //update
            if (keyIndex > 0)
                return true;
                //add 
            else
            {
                keyIndex = dicmapping.keys.length++;
                dicmapping.data[newName].keyIndex = keyIndex + 1;
                dicmapping.keys[keyIndex].key = newName;
                dicmapping.size++;
                return true;
            }
        } 
        else 
            return false;
    }
    
    function removeFile(string memory filename) public returns (bool res){
        (bool read,bool update) = getUserPerssion(msg.sender);
        if(update == true)
        {
            uint keyIndex = dicmapping.data[filename].keyIndex;
            if (keyIndex == 0)
                return false;
            delete dicmapping.data[filename];
            dicmapping.keys[keyIndex - 1].deleted = true;
            dicmapping.size --;
            return true;
        } 
        else 
            return false;
    }
    
    
    //public use 
    function getNextFile(uint keyIndex) public view returns(
        uint        r_keyIndex,
        string      memory fname,
        uint        fileType,       //0 dic ,1 file
        uint        createTime,
        uint        changeTime,
        address     accessControl,
        uint        fSize,
        bool        isRec,
        address     subDic,
        string      memory fileHash)
    {
        while (keyIndex < dicmapping.keys.length && dicmapping.keys[keyIndex].deleted)
            keyIndex++;
        if(iterate_valid(dicmapping,keyIndex) == true) {
            string memory key = dicmapping.keys[keyIndex].key;
            _fileInfo memory f = dicmapping.data[key].value;
                    return (
                        keyIndex,
                        f.name,
                        f.fileType,
                        f.createTime,
                        f.changeTime,
                        address(f.accessControl),
                        f.fileSize,
                        f.isRec,
                        address(f.subDic),
                        f.fileHash
                        );
        }
    }


    //internal use
    function iterate_start(_dicMapping memory self) internal pure returns (uint keyIndex)
    {
        return iterate_next(self, uint(-1));
    }
    
    function iterate_valid(_dicMapping memory self, uint keyIndex) internal pure returns (bool)
    {
        return keyIndex < self.keys.length;
    }

    function iterate_next(_dicMapping memory self, uint keyIndex) internal pure returns (uint r_keyIndex)
    {
        keyIndex++;
        while (keyIndex < self.keys.length && self.keys[keyIndex].deleted)
            keyIndex++;
        return keyIndex;
    }
}



