package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"flag"
	"io"
	"os"
)

var key *keystore.KeyStore
var password *string

func main()  {
	password = flag.String("pw", "123456", "")
	flag.Parse()

	fmt.Println("pw:  " + *password)
	key = keystore.NewKeyStore("./keystores", keystore.StandardScryptN, keystore.LightScryptP)
	for i := 0; i < 500; i++  {
		newAccount()
	}
}

func newAccount()  {
	//key := keystore.NewKeyStore("./keystores", keystore.StandardScryptN, keystore.LightScryptP)
	act, err := key.NewAccount(*password)
	if err != nil {
		fmt.Println(err)
	}else {
		fd,_ := os.OpenFile("./keystores/addr.txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
		s := fmt.Sprintf("%s\n", act.Address.Hex())
		io.WriteString(fd,s)

		fmt.Println(act.Address.Hex())
	}
}


