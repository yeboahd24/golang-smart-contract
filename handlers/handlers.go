package handlers

import (
	"encoding/json"
	"math/big"
	"net/http"
	"os"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yourusername/smart_contract/contracts"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type DepositRequest struct {
	Amount string `json:"amount"` // Amount in wei
}

type DisputeRequest struct {
	Reason string `json:"reason"`
}

func jsonResponse(w http.ResponseWriter, status int, resp Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func getAuthTransactor() (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	client, err := ethclient.Dial(os.Getenv("INFURA_URL"))
	if err != nil {
		return nil, err
	}

	nonce, err := client.PendingNonceAt(nil, fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(nil)
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	return auth, nil
}

func GetContractBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonResponse(w, http.StatusMethodNotAllowed, Response{Success: false, Error: "Method not allowed"})
		return
	}

	client, err := ethclient.Dial(os.Getenv("INFURA_URL"))
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, Response{Success: false, Error: "Failed to connect to Ethereum client"})
		return
	}

	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	instance, err := contracts.NewContracts(contractAddress, client)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, Response{Success: false, Error: "Failed to instantiate contract"})
		return
	}

	balance, err := instance.GetBalance(&bind.CallOpts{})
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, Response{Success: false, Error: "Failed to get balance"})
		return
	}

	jsonResponse(w, http.StatusOK, Response{
		Success: true,
		Data: map[string]string{
			"balance": balance.String(),
		},
	})
}

func GetContractStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonResponse(w, http.StatusMethodNotAllowed, Response{Success: false, Error: "Method not allowed"})
		return
	}

	client, err := ethclient.Dial(os.Getenv("INFURA_URL"))
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, Response{Success: false, Error: "Failed to connect to Ethereum client"})
		return
	}

	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	instance, err := contracts.NewContracts(contractAddress, client)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, Response{Success: false, Error: "Failed to instantiate contract"})
		return
	}

	seller, err := instance.Seller(nil)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, Response{Success: false, Error: "Failed to get seller"})
		return
	}

	arbiter, err := instance.Arbiter(nil)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, Response{Success: false, Error: "Failed to get arbiter"})
		return
	}

	state, err := instance.GetState(nil)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, Response{Success: false, Error: "Failed to get contract state"})
		return
	}

	// State 2 typically represents an approved state in escrow contracts
	isApproved := state == 2

	jsonResponse(w, http.StatusOK, Response{
		Success: true,
		Data: map[string]interface{}{
			"seller":     seller.Hex(),
			"arbiter":    arbiter.Hex(),
			"isApproved": isApproved,
			"state":      state,
		},
	})
}

func DepositFunds(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonResponse(w, http.StatusMethodNotAllowed, Response{Success: false, Error: "Method not allowed"})
		return
	}

	var req DepositRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonResponse(w, http.StatusBadRequest, Response{Success: false, Error: "Invalid request body"})
		return
	}

	amount, ok := new(big.Int).SetString(req.Amount, 10)
	if !ok {
		jsonResponse(w, http.StatusBadRequest, Response{Success: false, Error: "Invalid amount"})
		return
	}

	client, err := ethclient.Dial(os.Getenv("INFURA_URL"))
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, Response{Success: false, Error: "Failed to connect to Ethereum client"})
		return
	}

	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	instance, err := contracts.NewContracts(contractAddress, client)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, Response{Success: false, Error: "Failed to instantiate contract"})
		return
	}

	auth, err := getAuthTransactor()
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, Response{Success: false, Error: "Failed to create auth transactor"})
		return
	}
	auth.Value = amount

	tx, err := instance.DepositFunds(auth)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, Response{Success: false, Error: "Failed to deposit funds: " + err.Error()})
		return
	}

	jsonResponse(w, http.StatusOK, Response{
		Success: true,
		Data: map[string]string{
			"txHash": tx.Hash().Hex(),
		},
	})
}

func ReleaseFunds(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonResponse(w, http.StatusMethodNotAllowed, Response{Success: false, Error: "Method not allowed"})
		return
	}

	client, err := ethclient.Dial(os.Getenv("INFURA_URL"))
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, Response{Success: false, Error: "Failed to connect to Ethereum client"})
		return
	}

	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	instance, err := contracts.NewContracts(contractAddress, client)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, Response{Success: false, Error: "Failed to instantiate contract"})
		return
	}

	auth, err := getAuthTransactor()
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, Response{Success: false, Error: "Failed to create auth transactor"})
		return
	}

	tx, err := instance.ReleaseFunds(auth)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, Response{Success: false, Error: "Failed to release funds: " + err.Error()})
		return
	}

	jsonResponse(w, http.StatusOK, Response{
		Success: true,
		Data: map[string]string{
			"txHash": tx.Hash().Hex(),
		},
	})
}

func RaiseDispute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonResponse(w, http.StatusMethodNotAllowed, Response{Success: false, Error: "Method not allowed"})
		return
	}

	var req DisputeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonResponse(w, http.StatusBadRequest, Response{Success: false, Error: "Invalid request body"})
		return
	}

	client, err := ethclient.Dial(os.Getenv("INFURA_URL"))
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, Response{Success: false, Error: "Failed to connect to Ethereum client"})
		return
	}

	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	instance, err := contracts.NewContracts(contractAddress, client)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, Response{Success: false, Error: "Failed to instantiate contract"})
		return
	}

	auth, err := getAuthTransactor()
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, Response{Success: false, Error: "Failed to create auth transactor"})
		return
	}

	tx, err := instance.RaiseDispute(auth, req.Reason)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, Response{Success: false, Error: "Failed to raise dispute: " + err.Error()})
		return
	}

	jsonResponse(w, http.StatusOK, Response{
		Success: true,
		Data: map[string]string{
			"txHash": tx.Hash().Hex(),
		},
	})
}
