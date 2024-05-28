// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "../bridge/BlobstreamO.sol";
import { LayerTransition } from "./LayerTransition.sol";

/// @title TokenBridge
/// @dev This is the tellor token bridge to move tokens from
/// Ethereum to layer.  No one needs to do this.  The only reason you 
/// move your tokens over is to become a reporter/validator/tipper.  It works by
/// using layer itself as the bridge and then reads the lightclient contract for 
/// bridging back.  There is a long delay in bridging back (enforced by layer) of 21 days
contract TokenBridge is LayerTransition{
    /*Storage*/
    BlobstreamO public bridge;
    uint256 public depositId;//counterOfHowManydeposits have been made
    uint256 public depositLimitUpdateTime;//last time the limit was updated
    uint256 public depositLimitRecord;//amount you can bridge per limit period
    uint256 public constant DEPOSIT_LIMIT_UPDATE_INTERVAL = 12 hours;
    uint256 public constant INITIAL_LAYER_TOKEN_SUPPLY = 100 ether; // update this as needed
    uint256 public constant MAX_ATTESTATION_AGE = 12 hours;//max of how old an attestation can be for withdrawal from layer
    uint256 public immutable DEPOSIT_LIMIT_DENOMINATOR = 100e18 / 20e18; // 100/depositLimitPercentage

    mapping(uint256 => bool) public withdrawalClaimed;
    mapping(uint256 => DepositDetails) public deposits;

    struct DepositDetails {
        address sender;
        string recipient;
        uint256 amount;
        uint256 blockHeight;
    }

    /*Events*/
    event Deposit(uint256 _depositId, address _sender, string _recipient, uint256 _amount);
    event Withdrawal(uint256 _depositId, string _sender, address _recipient, uint256 _amount);

    /*Functions*/
    /// @notice constructor
    /// @param _token address of tellor token for bridging
    /// @param _blobstream address of BlobstreamO for data bridge
    /// @param _tellorFlex address of oracle(tellorFlex) on chain
    constructor(address _token, address _blobstream, address _tellorFlex) LayerTransition(_tellorFlex, _token){
        bridge = BlobstreamO(_blobstream);
         _refreshDepositLimit();
    }

    /// @notice deposits tokens from Ethereum to layer
    /// @param _amount amount of tokens to bridge over
    /// @param _layerRecipient your cosmos address on layer (don't get it wrong!!)
    function depositToLayer(uint256 _amount, string memory _layerRecipient) external {
        require(_amount > 0, "TokenBridge: amount must be greater than 0");
        require(token.transferFrom(msg.sender, address(this), _amount), "TokenBridge: transferFrom failed");
        require(_amount <= _refreshDepositLimit(), "TokenBridge: amount exceeds deposit limit");
        depositId++;
        depositLimitRecord -= _amount;
        deposits[depositId] = DepositDetails(msg.sender, _layerRecipient, _amount, block.number);
        emit Deposit(depositId, msg.sender, _layerRecipient, _amount);
    }

    /// @notice This withdraws tokens from layer to mainnet Ethereum
    /// @param _attest The data being verified
    /// @param _valset array of current validator set
    /// @param _sigs Signatures
    /// @param _depositId depositId from the layer side
    function withdrawFromLayer(
        OracleAttestationData calldata _attest,
        Validator[] calldata _valset,
        Signature[] calldata _sigs,
        uint256 _depositId
    ) external {
        require(_attest.queryId == keccak256(abi.encode("TRBBridge", abi.encode(false, _depositId))), "TokenBridge: invalid queryId");
        require(!withdrawalClaimed[_depositId], "TokenBridge: withdrawal already claimed");
        require(block.timestamp - _attest.report.timestamp > 12 hours, "TokenBridge: premature attestation");
        //isAnyConsesnusValue here
        bridge.verifyOracleData(_attest, _valset, _sigs);
        require(block.timestamp - _attest.attestationTimestamp <= MAX_ATTESTATION_AGE , "Attestation is too old");
        require(_attest.report.aggregatePower >= bridge.powerThreshold(), "Report aggregate power must be greater than or equal to _minimumPower");
        withdrawalClaimed[_depositId] = true;    
        (address _recipient, string memory _layerSender,uint256 _amountLoya) = abi.decode(_attest.report.value, (address, string, uint256));
        uint256 _amountConverted = _amountLoya * 1e12; 
        require(token.transfer(_recipient, _amountConverted), "TokenBridge: transfer failed");
        emit Withdrawal(_depositId, _layerSender, _recipient, _amountConverted);
    }

    /// @notice refreshes the deposit limit every 12 hours so no one can spam layer with new tokens
    function _refreshDepositLimit() internal returns (uint256) {
        if (block.timestamp - depositLimitUpdateTime > DEPOSIT_LIMIT_UPDATE_INTERVAL) {
            uint256 _layerTokenSupply = token.balanceOf(address(this)) + INITIAL_LAYER_TOKEN_SUPPLY;
            depositLimitRecord = _layerTokenSupply / DEPOSIT_LIMIT_DENOMINATOR;
            depositLimitUpdateTime = block.timestamp;
        }
        return depositLimitRecord;
    }
}
