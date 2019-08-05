// +build testrunmain

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// this abuses go so much that I felt dirty writing this code
// but it is the only way to do it without writing custom compiler that would
// be a clone of go-build with go-test
func TestRunMain(t *testing.T) {
	//args := []string{"get","QmW2WQi7j6c7UgJTarActp7tDNikE4B2qXtFCfLPdsgaTQ/cat.jpg"} //flag.Args()
	//args := []string{"eth","addUserToGroup","0x68dB32D26d9529B2a142927c6f1af248fc6Ba7e9","group4","7"}
	//args := []string{"add","./dist","-r"}
	//args := []string{"files" ,"stat","/he"}
	//args := []string{"files" ,"ls","/"}
	//args := []string{"eth","setPermission","./123","-r","-n","-g","admin"}
	//args := []string{"eth","tell"}
	//args := []string{"eth","showGroups"}
	//args := []string{"eth","ls"}
	args := []string{"object","new","unixfs-dir"}
	os.Args = append([]string{os.Args[0]}, args...)
	ret := mainRet()

	p := os.Getenv("IPFS_COVER_RET_FILE")
	if len(p) != 0 {
		ioutil.WriteFile(p, []byte(fmt.Sprintf("%d\n", ret)), 0777)
	}

	// close outputs so go testing doesn't print anything
	null, _ := os.Open(os.DevNull)
	os.Stderr = null
	os.Stdout = null
}
