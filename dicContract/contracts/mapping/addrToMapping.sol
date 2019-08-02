pragma solidity >=0.4.22 <0.7.0;

library addrToMapping {

    //the first floor addrss => mapping
    struct preitmap
    {
        mapping(address => IndexMap) data;
        AddrKeyFlag[] keys;
        uint size;
    }
    struct IndexMap {uint keyIndex; itmap value;}
    struct AddrKeyFlag {address key; bool deleted;}


    //the second floor mapping(string=> address)
    struct itmap
    {
        mapping(string => IndexValue) data;
        KeyFlag[] keys;
        uint size;
    }
    struct IndexValue { uint keyIndex; address value; }
    struct KeyFlag { string key; bool deleted; }
    //cover return true,add return false
    function insert(itmap storage self, string memory key, address value) public returns (bool replaced)
    {
        uint keyIndex = self.data[key].keyIndex;
        self.data[key].value = value;
        if (keyIndex > 0)
            return true;
        else
        {
            keyIndex = self.keys.length++;
            self.data[key].keyIndex = keyIndex + 1;
            self.keys[keyIndex].key = key;
            self.size++;
            return false;
        }
    }

    function insert(preitmap storage self, address key1,string memory key2, address value) public returns(bool replaced)
    {
        uint keyIndex1 = self.data[key1].keyIndex;
        uint keyIndex2 = self.data[key1].value.data[key2].keyIndex;
        self.data[key1].value.data[key2].value = value;
        if (keyIndex1 > 0 )
        {
            if(keyIndex2 > 0)
            {
                return true;
            }
            else
            {
                keyIndex2 = self.data[key1].value.keys.length++;
                self.data[key1].value.data[key2].keyIndex = keyIndex2 + 1;
                self.data[key1].value.keys[keyIndex2].key = key2;
                self.data[key1].value.size++;
                return false;
            }
        }
        else 
        {
            keyIndex2 = self.data[key1].value.keys.length++;
            self.data[key1].value.data[key2].keyIndex = keyIndex2 + 1;
            self.data[key1].value.keys[keyIndex2].key = key2;
            self.data[key1].value.size++;
                
            keyIndex1 = self.keys.length++;
            self.data[key1].keyIndex = keyIndex1 + 1;
            self.keys[keyIndex1].key = key1;
            self.size++;
            return false;
        }
    }

    function remove(itmap storage self, string memory key) public returns (bool success)
    {
        uint keyIndex = self.data[key].keyIndex;
        if (keyIndex == 0)
            return false;
        delete self.data[key];
        self.keys[keyIndex - 1].deleted = true;
        self.size --;
        return true;
    }

    function remove(preitmap storage self, address key) public returns (bool success)
    {
        uint keyIndex = self.data[key].keyIndex;
        if (keyIndex == 0)
            return false;
        //remove submapping
        for (uint256 i = iterate_start(self.data[key].value); iterate_valid(self.data[key].value, i); i = iterate_next(self.data[key].value, i))
        {
            remove(self.data[key].value,self.data[key].value.keys[i].key);
        }
        delete self.data[key];
        self.keys[keyIndex - 1].deleted = true;
        self.size --;
        return true;
    }


    function contains(itmap storage self, string memory key) public view returns (bool)
    {
        return self.data[key].keyIndex > 0;
    }
    function contains(preitmap storage self, address key) public view returns (bool)
    {
        return self.data[key].keyIndex > 0;
    }


    function iterate_start(itmap storage self) public view returns (uint keyIndex)
    {
        return iterate_next(self, uint(-1));
    }
    function iterate_start(preitmap storage self) public view returns (uint keyIndex)
    {
        return iterate_next(self, uint(-1));
    }


    function iterate_valid(itmap storage self, uint keyIndex) public view returns (bool)
    {
        return keyIndex < self.keys.length;
    }
    function iterate_valid(preitmap storage self, uint keyIndex) public view returns (bool)
    {
        return keyIndex < self.keys.length;
    }


    function iterate_next(itmap storage self, uint keyIndex) public view returns (uint r_keyIndex)
    {
        keyIndex++;
        while (keyIndex < self.keys.length && self.keys[keyIndex].deleted)
            keyIndex++;
        return keyIndex;
    }
    function iterate_next(preitmap storage self, uint keyIndex) public view returns (uint r_keyIndex)
    {
        keyIndex++;
        while (keyIndex < self.keys.length && self.keys[keyIndex].deleted)
            keyIndex++;
        return keyIndex;
    }

    function iterate_get(itmap storage self, uint keyIndex) public view returns (string memory key, address value)
    {
        key = self.keys[keyIndex].key;
        value = self.data[key].value;
    }
    function iterate_get(preitmap storage self, uint keyIndex) public view returns (address key, itmap storage value)
    {
        key = self.keys[keyIndex].key;
        value = self.data[key].value;
    }
}

