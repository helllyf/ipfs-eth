pragma solidity >=0.4.22 <0.7.0;

import "./userGroup.sol";

contract dictionary {
    address     public  creator;
    string      public  name;
    int         public  fileSize = 0;
    bool        public  accessType = false;
    uint        public  createtime;
    uint        public  changetime;  
    _dicList[]   dicList;
    userGroup   public  uG;
    
    
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
        uint        fileBlockHash;
    }
    
    struct _dicList {
        _fileInfo   file;
        bool        deleted;
    }
    
    constructor(address sender) public{
        creator = sender;
    }

    function updateAccessControl(address newUG) public onlyCreator {
        accessType = true;
        uG = userGroup(newUG);
        
        for(uint i = 0;i < dicList.length;i++)
        {
            if(dicList[i].file.isRec == true){
                dicList[i].file.accessControl = uG;
                if( dicList[i].file.fileType == 0)
                {
                     dicList[i].file.subDic.updateAccessControl(newUG);
                }
            }
                
        }
    }
    
    function getUserPerssion(address key) internal view returns (bool read,bool update) {
        if(accessType == false) return (true,true);
        (bool _read,bool _update,bool change) = uG.getUserPerssion(key);
        return (_read,_update);
    }
    
    function updateDicInfo(string memory newName) public returns (bool res)
    {
        (bool read,bool update) = getUserPerssion(msg.sender);
        if(update == true)
        {
            name = newName;
            return true;
        } 
        else 
            return false;
    }
    
    function addFile(string memory newName, uint newFileType,address newUG,uint newFileSize,
    bool newIsRec,address newsubDic,uint newFileBlockHash) public returns (bool res)
    {
        (bool read,bool update) = getUserPerssion(msg.sender);
        if(update == true)
        {
            _dicList memory dic; 
            dic.file = _fileInfo(newName,newFileType,now,now,
                newIsRec == true ? uG:userGroup(newUG)
                ,newFileSize,newIsRec,dictionary(newsubDic),newFileBlockHash);
            dic.deleted = false;
            dicList.push(dic);
            return true;
        } 
        else 
            return false;
    }
    
    function removeFile(uint delIndex) public returns (bool res){
        (bool read,bool update) = getUserPerssion(msg.sender);
        if(update == true)
        {
            delete dicList[delIndex].file;
            dicList[delIndex].deleted = true;
            return true;
        } 
        else 
            return false;
    }
    
    function updateFile(uint delIndex,string memory newName, uint newFileType,address newUG,
        uint newFileSize,bool newIsRec,address newsubDic,uint newFileBlockHash) 
        public
        returns (bool res){
        (bool read,bool update) = getUserPerssion(msg.sender);
        if(dicList[delIndex].deleted == true) return false;
        if(update == true)
        {
            dicList[delIndex].file = _fileInfo(newName,newFileType,now,now,
                newIsRec == true ? uG:userGroup(newUG)
                ,newFileSize,newIsRec,dictionary(newsubDic),newFileBlockHash);
            return true;
        } 
        else 
            return false;    
    }
}



