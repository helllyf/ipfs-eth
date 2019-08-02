package ethbridge

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-ipfs/dicContract/go/userGroup"
	"github.com/ipfs/go-ipfs/dicContract/go/userGroups"
	"github.com/pkg/errors"
	"math/big"
)

//合约实例
var userGroupsSol  *userGroups.UserGroups
var groupMapping = make(map[string]string)

func GetUserGroups() *userGroups.UserGroups{
	if userGroupsSol == nil{
		_address := common.HexToAddress(ethConf.UserGroupsAddr)
		m, err := userGroups.NewUserGroups(_address, client)
		if err != nil {
			log.Fatal(err)
		}
		userGroupsSol = m
	}
	return userGroupsSol
}

func ShowUserGroups(showInLine bool){
	//用户组

	GetUserGroups()
	_address := common.HexToAddress(ethConf.EthAddr)
	count,err := userGroupsSol.GetGroupCount(nil,_address)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0;i < int(count.Int64());i++{
		_key,_value,err := userGroupsSol.GetGroupByIndex(nil,_address,big.NewInt(int64(i)))
		if err != nil {
			log.Fatal(err)
		}
		if(showInLine == true) {
			fmt.Printf("%s : %s\n", _key, _value.Hex())
		}
		groupMapping[_key] = _value.Hex()
	}
}

func updateUserGroupWithNewGroup(name string){
	tx,err := userGroupsSol.UpdateUserGroupWithNewGroup(UpdateAuth(),name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())
}

func checkGroup(name string) bool {
	ShowUserGroups(false)
	if _,ok :=groupMapping[name]; ok{
		return true
	}
	return false
}

func GetGroupAddress(groupname string) (string) {
	var isNewOne string
	GetUserGroups()
	isExit :=checkGroup(groupname)
	if isExit == false {
		fmt.Println("Group Contract address is not found,. Do you want to new one? (y/n)")
		fmt.Scanln(&isNewOne)
		if(isNewOne != "n") {
			var name string
			fmt.Scanln(&name)
			addGroup_aux(name,false)
		} else {
			// we can not keep going
			log.Fatal(errors.New("Group Contract address is not found,"))
			return ""
		}
	}
	return groupMapping[groupname]
}

func AddGroup(name string)  {
	var isSave string
	GetUserGroups()
	isExit := checkGroup(name)
	if isExit == true {
		fmt.Println("The group already exists. Do you want to overwrite it? (y/n)")
		fmt.Scanln(&isSave)
	}
	if(isSave != "n") {
		addGroup_aux(name,false)
	}
}

func addGroup_aux(name string,isShowAll bool) {
	if name == "" {
		name = "admin"
	}
	_,ok :=userGroupsSol.UpdateUserGroupWithNewGroup(UpdateAuth(),name)
	if ok != nil {
		log.Fatal(ok)
	}
	ShowUserGroups(isShowAll)
	fmt.Printf("new group  name: %s  address: %s\n",name, groupMapping[name])
}

func AddUser( userPk string,name string,permission uint) {
	if(permission > 7 || permission < 0) {
		fmt.Println("The input permission should be between 1 and 7")
		return
	}
	GetUserGroups()
	ShowUserGroups(false)
	_address := common.HexToAddress(groupMapping[name])
	group, err := userGroup.NewUserGroup(_address, client)
	if err != nil {
		log.Fatal(err)
	}
	group.AndOrupdateUser(UpdateAuth(),common.HexToAddress(userPk),uint8(permission))
	fmt.Printf("add user : %s \n group: %s \n permission: %d",userPk,name,permission)
}
