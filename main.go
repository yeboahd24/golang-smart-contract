package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"regexp"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/yourusername/smart_contract/contracts"
	"github.com/yourusername/smart_contract/handlers"
)

func deployContract() error {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	// Connect to Ethereum client
	client, err := ethclient.Dial(os.Getenv("INFURA_URL"))
	if err != nil {
		return fmt.Errorf("error connecting to Ethereum client: %v", err)
	}

	// Load private key
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return fmt.Errorf("error loading private key: %v", err)
	}

	// Get public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("error casting public key to ECDSA")
	}

	// Create an authorized transactor
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(1000000) // in units
	auth.GasPrice = gasPrice

	// Check account balance
	balance, err := client.BalanceAt(context.Background(), auth.From, nil)
	if err != nil {
		return err
	}
	
	fmt.Printf("Your wallet address: %s\n", auth.From.Hex())
	balanceInEth := new(big.Float).Quo(new(big.Float).SetInt(balance), new(big.Float).SetInt(big.NewInt(1e18)))
	fmt.Printf("Account balance: %.6f ETH\n", balanceInEth)

	// Deploy contract
	address, tx, _, err := contracts.DeployContracts(
		auth,
		client,
		common.HexToAddress(os.Getenv("SELLER_ADDRESS")),
		common.HexToAddress(os.Getenv("ARBITER_ADDRESS")),
	)
	if err != nil {
		return err
	}

	fmt.Printf("Contract deployed! Address: %s\n", address.Hex())
	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

	// Wait for the transaction to be mined
	fmt.Println("Waiting for transaction to be mined...")
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return err
	}
	
	if receipt.Status == 1 {
		fmt.Println("Contract deployed successfully!")
		fmt.Printf("Contract address: %s\n", address.Hex())
		
		// Update .env file with the new contract address
		envContent, err := os.ReadFile(".env")
		if err != nil {
			return err
		}
		
		re := regexp.MustCompile(`CONTRACT_ADDRESS=.*`)
		updatedContent := re.ReplaceAllString(string(envContent), fmt.Sprintf("CONTRACT_ADDRESS=%s", address.Hex()))
		
		err = os.WriteFile(".env", []byte(updatedContent), 0644)
		if err != nil {
			return err
		}
		
		fmt.Println("Contract address updated in .env file")
	} else {
		return fmt.Errorf("contract deployment failed")
	}

	return nil
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Deploy contract
	if err := deployContract(); err != nil {
		log.Fatal(err)
	}

	// Setup HTTP routes
	http.HandleFunc("/contract/balance", handlers.GetContractBalance)
	http.HandleFunc("/contract/status", handlers.GetContractStatus)
	http.HandleFunc("/contract/deposit", handlers.DepositFunds)
	http.HandleFunc("/contract/release", handlers.ReleaseFunds)
	http.HandleFunc("/contract/dispute", handlers.RaiseDispute)
	
	// Start HTTP server
	port := "8080"
	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
