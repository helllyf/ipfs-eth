call truffle migrate --reset
call cd .\contracts\
call abigen --sol userGroups.sol --pkg userGroups --out .\userGroups.go
call cp userGroups.go ..\go\userGroups\

call abigen --sol userGroup.sol --pkg userGroup --out .\userGroup.go
call cp userGroup.go ..\go\userGroup\

call abigen --sol dictionary.sol --pkg dictionary --out .\dictionary.go
call cp dictionary.go ..\go\dictionary\

call cd ..