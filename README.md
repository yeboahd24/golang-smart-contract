# Golang Smart Contract Projects

This repository contains a collection of smart contract project ideas that can be implemented using Golang. These projects range from beginner to intermediate level and are designed to help you learn different aspects of blockchain development.

## Project Ideas

### 1. Token Management System
A basic implementation of an ERC-20 like token system.

**Features:**
- Token minting and burning
- Transfer tokens between addresses
- Check token balances
- Set token allowances
- Implementation of standard token interfaces

**Learning Outcomes:**
- Basic smart contract development
- Token standards
- State management
- Event emission

### 2. Decentralized Voting System
A transparent and secure voting mechanism implemented on the blockchain.

**Features:**
- Voter registration and verification
- Secure vote casting
- Real-time vote counting
- Prevention of double voting
- Result verification

**Learning Outcomes:**
- Access control
- Security patterns
- State management
- Event handling

### 3. Simple Escrow Service
A trustless escrow service for handling transactions between two parties.

**Features:**
- Fund deposit mechanism
- Time-locked releases
- Multi-signature approval
- Dispute resolution system
- Automatic fund release on condition completion

**Learning Outcomes:**
- Time-based operations
- Multi-party transactions
- Security patterns
- Fund management

### 4. NFT Marketplace
A basic marketplace for creating and trading NFTs.

**Features:**
- NFT minting
- Marketplace listings
- Buy/Sell mechanisms
- Royalty management
- Token metadata handling

**Learning Outcomes:**
- NFT standards
- Marketplace mechanics
- Asset management
- Metadata handling

### 5. Decentralized Crowdfunding
A platform for creating and managing crowdfunding campaigns.

**Features:**
- Campaign creation
- Pledge management
- Automatic fund release on goal achievement
- Refund mechanism for failed campaigns
- Campaign deadline management

**Learning Outcomes:**
- Time-based conditions
- Fund management
- State transitions
- Event handling

### 6. Supply Chain Tracking
A system for tracking products through various stages of a supply chain.

**Features:**
- Product registration
- Ownership transfer
- Status updates
- History tracking
- Role-based access control

**Learning Outcomes:**
- Complex state management
- Access control
- Event tracking
- Data structures

### 7. Simple Decentralized Exchange (DEX)
A basic implementation of a decentralized token exchange.

**Features:**
- Liquidity pool management
- Token swapping
- Price calculation
- Fee management
- Slippage protection

**Learning Outcomes:**
- Advanced state management
- Mathematical calculations
- Financial concepts
- Complex interactions

## Technical Requirements

### Prerequisites
- Go 1.16 or higher
- Ethereum client (like Geth)
- Solidity compiler
- Smart contract development framework

### Core Dependencies
```go
require (
    github.com/ethereum/go-ethereum v1.10.x
    github.com/spf13/viper
    github.com/stretchr/testify
)
```

### Development Tools
- Truffle/Hardhat for smart contract deployment
- Ganache for local blockchain testing
- Web3.js/Ethers.js for frontend integration
- MetaMask for wallet integration

## Development Setup

### Prerequisites
1. Install Solidity compiler (solc):
```bash
sudo add-apt-repository ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install solc
```

2. Install Go Ethereum tools:
```bash
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

### Generate Go Bindings

1. Create build directory:
```bash
mkdir -p build
```

2. Compile Solidity contract:
```bash
solc --abi --bin contracts/Escrow.sol -o build/
```

3. Generate Go bindings:
```bash
abigen --bin=build/Escrow.bin --abi=build/Escrow.abi --pkg=contracts --out=contracts/escrow.go
```

### Project Structure
```
smart_contract/
├── contracts/
│   ├── Escrow.sol        # Solidity smart contract
│   └── escrow.go         # Generated Go bindings
├── build/
│   ├── Escrow.abi        # Contract ABI
│   └── Escrow.bin        # Contract bytecode
├── main.go               # Main application
├── go.mod               # Go module file
└── .env                 # Environment configuration
```

### Environment Setup
Create a `.env` file with the following:
```env
INFURA_URL=your_infura_url_here
PRIVATE_KEY=your_private_key_here
CONTRACT_ADDRESS=your_deployed_contract_address
```

### Running the Application

1. Install dependencies:
```bash
go mod tidy
```

2. Run the application:
```bash
go run main.go
```

### Testing in Remix IDE

1. Deploy the contract with three different addresses:
   - Buyer (automatically set to the deploying address)
   - Seller (provide a different address)
   - Arbiter (provide another different address)

2. Test functions:
   - `depositFunds()`: Send ETH from buyer
   - `releaseFunds()`: Release funds to seller
   - `raiseDispute()`: Raise a dispute
   - `resolveDispute()`: Resolve disputes (arbiter only)
   - `getBalance()`: Check contract balance
   - `getState()`: Check current state

### Contract States
- `AWAITING_PAYMENT` (0): Initial state
- `AWAITING_DELIVERY` (1): After buyer deposits funds
- `COMPLETE` (2): After successful transaction
- `DISPUTED` (3): When dispute is raised
- `REFUNDED` (4): When funds are returned to buyer

## API Endpoints

### 1. Get Contract Balance
Retrieves the current balance of the smart contract.

**Endpoint:** `GET /contract/balance`

**Response:**
```json
{
    "success": true,
    "data": {
        "balance": "string" // Balance in wei
    }
}
```

### 2. Get Contract Status
Retrieves the current status of the smart contract.

**Endpoint:** `GET /contract/status`

**Response:**
```json
{
    "success": true,
    "data": {
        "state": "number",
        "isApproved": "boolean",
        "seller": "string",
        "arbiter": "string"
    }
}
```

### 3. Deposit Funds
Deposits funds into the smart contract.

**Endpoint:** `POST /contract/deposit`

**Request Body:**
```json
{
    "amount": "string" // Amount in wei
}
```

**Response:**
```json
{
    "success": true,
    "data": {
        "txHash": "string" // Transaction hash
    }
}
```

### 4. Release Funds
Releases the funds to the seller.

**Endpoint:** `POST /contract/release`

**Response:**
```json
{
    "success": true,
    "data": {
        "txHash": "string" // Transaction hash
    }
}
```

### 5. Raise Dispute
Raises a dispute on the contract.

**Endpoint:** `POST /contract/dispute`

**Request Body:**
```json
{
    "reason": "string"
}
```

**Response:**
```json
{
    "success": true,
    "data": {
        "txHash": "string" // Transaction hash
    }
}
```

**Error Response (for all endpoints):**
```json
{
    "success": false,
    "error": "Error message"
}
```

Common error messages:
- "Method not allowed"
- "Invalid request body"
- "Failed to connect to Ethereum client"
- "Failed to instantiate contract"
- "Failed to create auth transactor"
- "Failed to get balance"
- "Failed to get contract state"
- "Failed to get seller/arbiter"
- "Failed to deposit/release funds"
- "Failed to raise dispute"

## Getting Started

1. Clone this repository
```bash
git clone <repository-url>
```

2. Install dependencies
```bash
go mod tidy
```

3. Set up your local blockchain network or connect to a testnet

4. Configure your environment variables
```bash
cp .env.example .env
# Edit .env with your configuration
```

5. Run tests
```bash
go test ./...
```

## Getting Test ETH

To deploy and test smart contracts on the Sepolia testnet, you'll need some test ETH. Here are some options to get test ETH:

### Option 1: Faucets requiring Alchemy account
1. [Alchemy Builder Faucet](https://sepoliafaucet.com/)
   - Requires creating a free Alchemy account
   - Get 0.5 Sepolia ETH per day
   - No Mainnet ETH required

### Option 2: Mining Faucet
1. [Sepolia PoW Faucet](https://sepolia-faucet.pk910.de/)
   - Mine in your browser to get test ETH
   - No account or Mainnet ETH required
   - Get ETH based on mining time

### Option 3: Social Authentication
1. [Paradigm Faucet](https://faucet.paradigm.xyz/)
   - Requires Twitter authentication
   - Get 0.5 Sepolia ETH per request
   - No Mainnet ETH required

Note: Some faucets (like Infura, QuickNode, and MetaMask) require you to have 0.001 ETH on Ethereum Mainnet before they'll give you test ETH. The options above don't have this requirement.

Steps to get test ETH:
1. Choose one of the faucets above
2. Follow their specific requirements (create account, mine, or authenticate)
3. Enter your wallet address: The address from your private key in `.env`
4. Request or mine for test ETH
5. Wait for the transaction to be processed

## Project Structure
```
├── contracts/       # Smart contract source files
├── migrations/      # Deployment scripts
├── pkg/            # Shared Go packages
├── cmd/            # Command-line applications
├── internal/       # Internal packages
├── scripts/        # Utility scripts
└── test/           # Test files
```

## Best Practices

1. **Security**
   - Always use the latest stable versions of dependencies
   - Implement proper access control
   - Use safe mathematical operations
   - Include comprehensive tests
   - Conduct security audits

2. **Code Quality**
   - Follow Go best practices and conventions
   - Document your code thoroughly
   - Use meaningful variable and function names
   - Implement proper error handling
   - Write unit and integration tests

3. **Gas Optimization**
   - Optimize data structures
   - Batch operations when possible
   - Use appropriate data types
   - Minimize storage operations

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Disclaimer

These smart contracts are provided as educational examples and should not be used in production without proper security audits and testing.



