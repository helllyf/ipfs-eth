package ethbridge

//ganache-cli -m "much repair shock carbon improve miss forget sock include bullet interest solution"
import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-ipfs/dicContract/go/router"
	logging "github.com/ipfs/go-log"

	"github.com/BurntSushi/toml"
)


var log = logging.Logger("core/commands")
var confPath = "\\ethConf.toml"
type _bridge struct {
	EthAddr	 		string	 				//账户地址
	PK 				string					//账户私钥
	RouterAddr 		string					//router合约地址
	Url 			string					//区块链节点端口
	UserGroupsAddr 	string
}
var ethConf _bridge

//区块链节点实例化
var client 		*ethclient.Client
var auth 		*bind.TransactOpts
//var callOpt 	*bind.CallOpts
//router合约实例化
var routerSol	*router.Router
//单例模式
var once 		sync.Once

func getConf() {
	if _, err := toml.DecodeFile(confPath, &ethConf); err != nil {
		log.Fatal(err)
	}
}

func InitEth()  {
	privateKey, err := crypto.HexToECDSA(ethConf.PK)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth 			= bind.NewKeyedTransactor(privateKey)
	auth.Nonce 		= big.NewInt(int64(nonce))
	auth.Value 		= big.NewInt(0) // in wei
	auth.GasLimit 	= uint64(3000000) // in units
	auth.GasPrice 	= gasPrice
}

func UpdateAuth() *bind.TransactOpts {
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(ethConf.EthAddr))
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce 		= big.NewInt(int64(nonce))
	return auth
}

//获取ethClient实例
func InitEthClient(rootPath string)*ethclient.Client{
	once.Do(func() {
		if rootPath != "" {
			confPath = rootPath + confPath
			//strings.Replace(confPath,"\\", "/",-1)
		}
		getConf()
		m, err := ethclient.Dial(ethConf.Url)
		if err != nil {
			log.Fatal(err)
		}
		client = m
		InitEth()
	})
	return client
}

func GetRouter() *router.Router{
	if routerSol == nil{
		_address := common.HexToAddress(ethConf.RouterAddr)
		m, err := router.NewRouter(_address, client)
		if err != nil {
			log.Fatal(err)
		}
		routerSol = m
	}
	return routerSol
}



func TestRouter() {
	routerSol = GetRouter()
	tx,err := routerSol.InsertAuto(UpdateAuth())
	if err != nil {
		log.Fatal(err)
	}
	txHash :=  tx.Hash().Hex()
	fmt.Println(txHash)
	result, err := routerSol.Show(nil ,common.HexToAddress(ethConf.EthAddr))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Hex())

	ShowUserGroups(true)
}

func GetBalance() *big.Int{
	account := common.HexToAddress(ethConf.EthAddr)
	client := InitEthClient("")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	return balance
}
//
//func getSolidityJson( path string) string{
//	f,err := os.Open(path)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer func() {
//		if err = f.Close(); err != nil {
//			log.Fatal(err)
//		}
//	}()
//
//	fd, err := ioutil.ReadAll(f)
//	return string(fd)
//}

func IsZeroAddress(addr common.Address)  bool{
	zeroAddr := common.Address{0}
	if addr == zeroAddr {
		return true
	}
	return false
}
