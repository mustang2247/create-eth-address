package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var key *keystore.KeyStore

func main()  {
	//password = flag.String("pw", "", "")
	//flag.Parse()
	//
	//fmt.Println("pw:  " + *password)
	//key = keystore.NewKeyStore("./keystores", keystore.StandardScryptN, keystore.LightScryptP)
	key = keystore.NewKeyStore("./keystores", keystore.StandardScryptN, keystore.StandardScryptP)
	for i := 0; i < 500; i++  {
		//newAccount()
		createKs("mus-1104")
	}
}

func createKs(password string) {
	//ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := key.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	jsonBytes, err := ioutil.ReadFile(account.URL.Path)
	if err != nil {
		log.Fatal(err)
	}

	pKey, err := keystore.DecryptKey(jsonBytes, password)
	if err != nil {
		log.Fatal(err)
	}

	privateKey := hex.EncodeToString(crypto.FromECDSA(pKey.PrivateKey))
	fmt.Println(privateKey)

	fd,_ := os.OpenFile("./keystores/addr.txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	s := fmt.Sprintf("%s\n", account.Address.Hex() + "	" + privateKey )
	io.WriteString(fd,s)


	fmt.Println(account.Address.Hex() + "	" + privateKey)


	//secp256k1.Sign()



}


