pragma solidity >=0.4.22 <0.7.0;

library strToAddr {
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

    function contains(itmap storage self, string memory key) public view returns (bool)
    {
        return self.data[key].keyIndex > 0;
    }
    function iterate_start(itmap storage self) public view returns (uint keyIndex)
    {
        return iterate_next(self, uint(-1));
    }
    function iterate_valid(itmap storage self, uint keyIndex) public view returns (bool)
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
    function iterate_get(itmap storage self, uint keyIndex) public view returns (string memory key, address value)
    {
        key = self.keys[keyIndex].key;
        value = self.data[key].value;
    }
}
