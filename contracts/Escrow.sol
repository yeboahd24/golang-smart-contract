// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title Simple Escrow Contract
 * @dev Implements an escrow system between a buyer and seller with dispute resolution
 */
contract Escrow {
    // State variables
    address public buyer;
    address public seller;
    address public arbiter;
    uint256 public amount;
    bool public isComplete;
    bool public isDisputed;
    
    // Enum to track escrow state
    enum State { AWAITING_PAYMENT, AWAITING_DELIVERY, COMPLETE, DISPUTED, REFUNDED }
    State public currentState;
    
    // Events
    event FundsDeposited(address indexed from, uint256 amount);
    event FundsReleased(address indexed to, uint256 amount);
    event DisputeRaised(address indexed by, string reason);
    event DisputeResolved(address indexed winner);
    event EscrowComplete();
    
    // Modifiers
    modifier onlyBuyer() {
        require(msg.sender == buyer, "Only buyer can call this function");
        _;
    }
    
    modifier onlySeller() {
        require(msg.sender == seller, "Only seller can call this function");
        _;
    }
    
    modifier onlyArbiter() {
        require(msg.sender == arbiter, "Only arbiter can call this function");
        _;
    }
    
    modifier inState(State _state) {
        require(currentState == _state, "Invalid state");
        _;
    }

    /**
     * @dev Constructor to create escrow agreement
     * @param _seller Address of the seller
     * @param _arbiter Address of the arbiter for dispute resolution
     */
    constructor(address _seller, address _arbiter) {
        require(_seller != address(0) && _arbiter != address(0), "Invalid address");
        buyer = msg.sender;
        seller = _seller;
        arbiter = _arbiter;
        currentState = State.AWAITING_PAYMENT;
    }
    
    /**
     * @dev Deposit funds into escrow
     */
    function depositFunds() external payable onlyBuyer inState(State.AWAITING_PAYMENT) {
        require(msg.value > 0, "Amount must be greater than 0");
        amount = msg.value;
        currentState = State.AWAITING_DELIVERY;
        emit FundsDeposited(msg.sender, msg.value);
    }
    
    /**
     * @dev Release funds to seller
     */
    function releaseFunds() external onlyBuyer inState(State.AWAITING_DELIVERY) {
        currentState = State.COMPLETE;
        isComplete = true;
        payable(seller).transfer(amount);
        emit FundsReleased(seller, amount);
        emit EscrowComplete();
    }
    
    /**
     * @dev Raise a dispute
     * @param reason Reason for dispute
     */
    function raiseDispute(string memory reason) external inState(State.AWAITING_DELIVERY) {
        require(msg.sender == buyer || msg.sender == seller, "Only buyer or seller can raise dispute");
        currentState = State.DISPUTED;
        isDisputed = true;
        emit DisputeRaised(msg.sender, reason);
    }
    
    /**
     * @dev Resolve dispute and release funds
     * @param winner Address of the winning party
     */
    function resolveDispute(address winner) external onlyArbiter inState(State.DISPUTED) {
        require(winner == buyer || winner == seller, "Winner must be buyer or seller");
        if (winner == buyer) {
            currentState = State.REFUNDED;
            payable(buyer).transfer(amount);
            emit FundsReleased(buyer, amount);
        } else {
            currentState = State.COMPLETE;
            payable(seller).transfer(amount);
            emit FundsReleased(seller, amount);
        }
        isComplete = true;
        emit DisputeResolved(winner);
        emit EscrowComplete();
    }
    
    /**
     * @dev Get current balance of contract
     */
    function getBalance() external view returns (uint256) {
        return address(this).balance;
    }
    
    /**
     * @dev Get current state of escrow
     */
    function getState() external view returns (State) {
        return currentState;
    }
}
