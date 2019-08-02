package ethbridge

import (
	"errors"
	"fmt"

	"math/big"
	"os"
	"strings"

	"github.com/disiqueira/gotree"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-ipfs/dicContract/go/dictionary"
)

type fileType struct {
	name 	string		//文件名
	h 		string		//hash
	size 	int			//大小
	ug 		string		//权限组
	t 		int 		//类型
}

type dicMap map[string]*DicType
type fileMap map[string]*fileType

type DicType struct {
	fileType
	subfiles 	fileMap
	subdics 	dicMap
}

var dicRoot DicType
func PrintEthDicRec(root gotree.Tree,dicSolc *dictionary.Dictionary) {
	var keyIndex int = 0
	for {
		res,err := dicSolc.GetNextFile(nil,big.NewInt(int64(keyIndex)))
		if err != nil {
			fmt.Println(err)
		}
		//遍历结束
		if res.Fname == "" {
			break
		}
		if res.FileType.Cmp(big.NewInt(0)) == 0 {
			sub := root.Add(res.Fname)
			subDicSolc,err := dictionary.NewDictionary(res.SubDic,client)
			if err != nil {
				log.Fatal(err)
			}
			PrintEthDicRec(sub,subDicSolc)
		}else{
			root.Add(res.Fname)
		}
		keyIndex++
	}

}
func PrintEthDic()  {
	root := gotree.New("")
	dicSolc := getDicSolc()
	PrintEthDicRec(root,dicSolc)
	fmt.Println(root.Print())
}


func RecPrint(tree gotree.Tree,pointerDic *DicType) {
	for _, v := range pointerDic.subfiles {
		tree.Add(v.name)
	}
	for _, v := range pointerDic.subdics {
		sub := tree.Add(v.name)
		RecPrint(sub,v)
	}
}

func PrintDic() {
	root := gotree.New("")
	for _, v := range dicRoot.subfiles {
		root.Add(v.name)
	}
	for _, v := range dicRoot.subdics {
		sub := root.Add(v.name)
		RecPrint(sub,v)
	}

	fmt.Println(root.Print())
}
func SetPermission(path string ,groupName string,h string) {
	f, _ := os.Stat(path)
	fmt.Println(f.Size())
	s := f.Size()
	isDir := f.IsDir()
	pName :=  strings.Split(path,"/")

	var pointDic *DicType
	pointDic = &dicRoot
	if pointDic == nil {
		pointDic = new(DicType) //头结点，不存数据
	}
	for k,v :=range pName{
		if pointDic.subfiles == nil {
			pointDic.subfiles = make(fileMap)
		}
		if pointDic.subdics == nil {
			pointDic.subdics = make(dicMap)
		}

		if k == len(pName) - 1 {
			if !isDir {

				if pointDic.subfiles[v] == nil {
					pointDic.subfiles[v] = new(fileType)
				}
				pointDic.subfiles[v].name 	= v
				pointDic.subfiles[v].size 	= int(s)
				pointDic.subfiles[v].ug	  	= groupName
				pointDic.subfiles[v].t		= 1
				pointDic.subfiles[v].h		= h
			} else {
				if pointDic.subdics[v] == nil {
					pointDic.subdics[v] = new(DicType)
				}
				pointDic.subdics[v].name 	= v
				pointDic.subdics[v].t 		= 0
				pointDic.subdics[v].ug		= groupName
				pointDic.subdics[v].size	= int(s)
				pointDic.subdics[v].h		=h

			}
			return
		}

		if pointDic.subdics[v] == nil {
			pointDic.subdics[v] = new(DicType)
		}
		pointDic.subdics[v].name = v
		pointDic.subdics[v].t 	= 0
		pointDic.subdics[v].ug	= groupName
		pointDic.subdics[v].size = 0
		pointDic = pointDic.subdics[v]
	}
}

/*
判断是否有根目录
若没有，没有自动创建
*/
func getDicAddress() (string,error){
	var isAddOne string
	routerSol = GetRouter()
	result, err := routerSol.Show(nil ,common.HexToAddress(ethConf.EthAddr))
	if err != nil {
		log.Fatal(err)
	}


	if IsZeroAddress(result) {
		fmt.Println("The dic contract is not exists. Do you want to new one? (y/n)")
		fmt.Scanln(&isAddOne)
	} else {
		return result.Hex(),nil
	}

	if(isAddOne != "n") {
		_,err := routerSol.InsertAuto(UpdateAuth())
		if err != nil {
			log.Fatal(err)
		}
		result, err := routerSol.Show(nil ,common.HexToAddress(ethConf.EthAddr))
		if err != nil {
			log.Fatal(err)
		}
		return result.Hex(),nil
	}else{
		return result.Hex(),errors.New("Dic Contract address is not found")
	}
}

func getDicSolc() *dictionary.Dictionary{
	addr,err := getDicAddress()
	if err != nil {
		log.Fatal(err)
	}
	dicSolc,err := dictionary.NewDictionary(common.HexToAddress(addr),client)
	if err != nil {
		log.Fatal(err)
	}
	return dicSolc
}

func CommitToEth() {
	dicSolc := getDicSolc()
	//头节点不存数据
	for _, v := range dicRoot.subfiles {
		commitToEthFile(v,dicSolc)
	}
	for _, v := range dicRoot.subdics {
		_dicAddr := commitToEthDicRec(v,dicSolc)
		addr := GetGroupAddress(v.ug)
		tx,err := dicSolc.AddOrUpdateFile(UpdateAuth(),v.name,big.NewInt(0),common.HexToAddress(addr),
			big.NewInt(int64(v.size)),true,_dicAddr,v.h)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("dic info:",v.name,";tx:",tx.Hash().Hex(),"\n")
	}
	PrintDic()
}

func commitToEthDicRec(pointerDic *DicType,parDicSolc *dictionary.Dictionary) common.Address{
	//get dic address and contract
	dicAddr,dicSolc,err := newDicContract(pointerDic,parDicSolc)
	if err != nil {
		log.Fatal(err)
	}

	//update her files
	for _, v := range pointerDic.subfiles {
		commitToEthFile(v,dicSolc)
	}
	//Recursively add subdirectories ,update her subdirectories' address
	for _, v := range  pointerDic.subdics {
		_dicAddr := commitToEthDicRec(v,dicSolc)
		addr := GetGroupAddress(v.ug)
		tx,err := dicSolc.AddOrUpdateFile(UpdateAuth(),v.name,big.NewInt(0),common.HexToAddress(addr),
			big.NewInt(int64(v.size)),true,_dicAddr,v.h)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("dic info:",v.name,";tx:",tx.Hash().Hex(),"\n")
	}
	return dicAddr
}

func commitToEthFile(pointerFile *fileType,dicSolc *dictionary.Dictionary) {
	addr := GetGroupAddress(pointerFile.ug)
	tx,err := dicSolc.AddOrUpdateFile(UpdateAuth(),pointerFile.name,big.NewInt(1),common.HexToAddress(addr),
		big.NewInt(int64(pointerFile.size)),false,common.Address{0},pointerFile.h)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file info:",pointerFile.name,";tx:",tx.Hash().Hex(),"\n")
}

func newDicContract(pointerDic *DicType,parDicSolc *dictionary.Dictionary) (common.Address,*dictionary.Dictionary,error){
	res,err := parDicSolc.GetFileInfo(nil,pointerDic.name)
	if err != nil {
		fmt.Println(err)
	}
	if res.IsExit != false && !IsZeroAddress(res.SubDic) {
		_dic,err := dictionary.NewDictionary(res.SubDic,client)
		return res.SubDic,_dic,err
	}

	dicAddr,_,dic,err := dictionary.DeployDictionary(UpdateAuth(), client, common.HexToAddress(ethConf.EthAddr))
	if err != nil {
		log.Fatal(err)
	}
	_,err = dic.UpdateAccessControl(UpdateAuth(),common.HexToAddress(pointerDic.ug))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("add new dic contract:",pointerDic.name,"address:", dicAddr.Hex())
	return dicAddr,dic,nil
}