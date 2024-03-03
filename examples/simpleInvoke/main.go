package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/NethermindEth/starknet.go/account"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/NethermindEth/starknet.go/utils"
	"github.com/joho/godotenv"
)

// NOTE : Please add in your keys only for testing purposes, in case of a leak you would potentially lose your funds.
var (
	name                  string = "testnet"                                                            //env."name"
	account_addr          string = "0x00655d6457243978cb2f505de1df6e827c3b0781f8503399682fbcd8f369ca15" //Replace it with your account address
	account_cairo_version        = 2                                                                    //Replace  with the cairo version of your account
	privateKey            string = "0x86edbb5d6ddd8f8bf24a8842359d1d0a909fc43f920e9168ecdedc45b9a1d8" //Replace it with your account private key
	public_key            string = "0x5573b134321002f0922e3f215332aa6e10ad14f6b01d1108e4f26adc2858c79"  //Replace it with your account public key
	someContract          string = "0x4c1337d55351eac9a0b74f3b8f0d3928e2bb781e5084686a892e66d49d510d"   //Replace it with the contract that you want to invoke
	contractMethod        string = "increase_value"                                                     //Replace it with the function name that you want to invoke
)

func main() {
	// Loading the env
	godotenv.Load(fmt.Sprintf(".env.%s", name))
	url := os.Getenv("INTEGRATION_BASE") //please modify the .env.testnet and replace the INTEGRATION_BASE with a starknet goerli RPC.
	fmt.Println("Starting simpleInvoke example")

	// Initialising the connection
	clientv02, err := rpc.NewProvider(url)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error dialing the RPC provider: %s", err))
	}

	// Here we are converting the account address to felt
	account_address, err := utils.HexToFelt(account_addr)
	if err != nil {
		panic(err.Error())
	}
	// Initializing the account memkeyStore
	ks := account.NewMemKeystore()
	fakePrivKeyBI, ok := new(big.Int).SetString(privateKey, 0)
	if !ok {
		panic(err.Error())
	}
	ks.Put(public_key, fakePrivKeyBI)

	fmt.Println("Established connection with the client")

	// Here we are setting the maxFee
	maxfee, err := utils.HexToFelt("0x9184e72a000")
	if err != nil {
		panic(err.Error())
	}

	// Initializing the account
	accnt, err := account.NewAccount(clientv02, account_address, public_key, ks, account_cairo_version)
	if err != nil {
		panic(err.Error())
	}

for {
	// Getting the nonce from the account
	nonce, err := accnt.Nonce(context.Background(), rpc.BlockID{Tag: "latest"}, accnt.AccountAddress)
	if err != nil {
                        fmt.Println(err)
                        time.Sleep(40 * time.Second)
                        continue
	}

	// Building the InvokeTx struct
	InvokeTx := rpc.InvokeTxnV1{
		MaxFee:        maxfee,
		Version:       rpc.TransactionV1,
		Nonce:         nonce,
		Type:          rpc.TransactionType_Invoke,
		SenderAddress: accnt.AccountAddress,
	}

	// Converting the contractAddress from hex to felt
	contractAddress, err := utils.HexToFelt(someContract)
	if err != nil {
		panic(err.Error())
	}

	// Building the functionCall struct, where :
	FnCall := rpc.FunctionCall{
		ContractAddress:    contractAddress,                               //contractAddress is the contract that we want to call
		EntryPointSelector: utils.GetSelectorFromNameFelt(contractMethod), //this is the function that we want to call
	}

	// Building the Calldata with the help of FmtCalldata where we pass in the FnCall struct along with the Cairo version
	InvokeTx.Calldata, err = accnt.FmtCalldata([]rpc.FunctionCall{FnCall})
	if err != nil {
		panic(err.Error())
	}

	// Signing of the transaction that is done by the account
	err = accnt.SignInvokeTransaction(context.Background(), &InvokeTx)
	if err != nil {
		panic(err.Error())
	}

	resp, err := accnt.AddInvokeTransaction(context.Background(), InvokeTx)
	if err != nil {
                        fmt.Println(err)
			time.Sleep(40 * time.Second)
			continue
	}
	
	fmt.Println("Transaction hash response : ", resp.TransactionHash)
	time.Sleep(20 * time.Second)

}
}

