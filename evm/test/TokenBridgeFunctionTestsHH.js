const { expect } = require("chai");
const { ethers } = require("hardhat");
const h = require("./helpers/evmHelpers");
var assert = require('assert');
const abiCoder = new ethers.utils.AbiCoder();


describe("TokenBridge - Function Tests", async function () {

    let blobstream, accounts, guardian, tbridge, token, blocky0,
        valTs, valParams, valSet;
    const UNBONDING_PERIOD = 86400 * 7 * 3; // 3 weeks
    const WITHDRAW1_QUERY_DATA_ARGS = abiCoder.encode(["bool", "uint256"], [false, 1])
    const WITHDRAW1_QUERY_DATA = abiCoder.encode(["string", "bytes"], ["TRBBridge", WITHDRAW1_QUERY_DATA_ARGS])
    const WITHDRAW1_QUERY_ID = h.hash(WITHDRAW1_QUERY_DATA)
    const EVM_RECIPIENT = "0x88dF592F8eb5D7Bd38bFeF7dEb0fBc02cf3778a0"
    const LAYER_RECIPIENT = "tellor1zy50vdk8fdae0var2ryjhj2ysxtcm8dp2qtckd"

    beforeEach(async function () {
        // init accounts
        accounts = await ethers.getSigners();
        guardian = accounts[10]
        val1 = ethers.Wallet.createRandom()
        val2 = ethers.Wallet.createRandom()
        initialValAddrs = [val1.address,val2.address]
        initialPowers = [1, 2]
        threshold = 2
        blocky = await h.getBlock()
        valTimestamp = blocky.timestamp - 2
        newValHash = await h.calculateValHash(initialValAddrs, initialPowers)
        valCheckpoint = h.calculateValCheckpoint(newValHash, threshold, valTimestamp)
        // deploy contracts
        blobstream = await ethers.deployContract(
            "BlobstreamO", [
            threshold,
            valTimestamp,
            UNBONDING_PERIOD,
            valCheckpoint,
            guardian.address
        ]
        )
        token = await ethers.deployContract("TellorPlayground")
        oldOracle = await ethers.deployContract("TellorPlayground")
        tbridge = await ethers.deployContract("TestTokenBridge", [token.address,blobstream.address, oldOracle.address])
        blocky0 = await h.getBlock()
        // fund accounts
        await token.faucet(accounts[0].address)
    });

    it("constructor", async function () {
        assert.equal(await tbridge.token(), await token.address)
        assert.equal(await tbridge.bridge(), await blobstream.address)
        expect(Number(await tbridge.depositLimitUpdateTime())).to.be.closeTo(Number(blocky0.timestamp), 1)
        expectedDepositLimit = BigInt(100e18) * BigInt(2) / BigInt(10)
        assert.equal(BigInt(await tbridge.depositLimitRecord()), expectedDepositLimit);
    })
    it("withdrawFromLayer", async function () {
        depositAmount = h.toWei("20")
        await h.expectThrow(tbridge.depositToLayer(depositAmount, LAYER_RECIPIENT)) // not approved
        await token.approve(await tbridge.address, h.toWei("100"))
        await tbridge.depositToLayer(depositAmount, LAYER_RECIPIENT)
        value = h.getWithdrawValue(EVM_RECIPIENT,LAYER_RECIPIENT,20)
        blocky = await h.getBlock()
        timestamp = blocky.timestamp - 2
        aggregatePower = 3
        attestTimestamp = timestamp + 1
        previousTimestamp = 0
        nextTimestamp = 0
        newValHash = await h.calculateValHash(initialValAddrs, initialPowers)
        valCheckpoint = await h.calculateValCheckpoint(newValHash, threshold, valTimestamp)
        dataDigest = await h.getDataDigest(
            WITHDRAW1_QUERY_ID,
            value,
            timestamp,
            aggregatePower,
            previousTimestamp,
            nextTimestamp,
            valCheckpoint,
            attestTimestamp
        )
        currentValSetArray = await h.getValSetStructArray(initialValAddrs, initialPowers)
        sig1 = await h.layerSign(dataDigest, val1.privateKey)
        sig2 = await h.layerSign(dataDigest, val2.privateKey)
        sigStructArray = await h.getSigStructArray([sig1, sig2])
        oracleDataStruct = await h.getOracleDataStruct(
            WITHDRAW1_QUERY_ID,
            value,
            timestamp,
            aggregatePower,
            previousTimestamp,
            nextTimestamp,
            attestTimestamp
        )
        await h.advanceTime(43200)
        await tbridge.withdrawFromLayer(
            oracleDataStruct,
            currentValSetArray,
            sigStructArray,
            1,
        )
        recipientBal = await token.balanceOf(EVM_RECIPIENT)
        expectedBal = 20e12 // 20 loya
        assert.equal(recipientBal.toString(), expectedBal)
    })
    it("depositToLayer", async function () {
        depositAmount = h.toWei("1")
        assert.equal(await token.balanceOf(await accounts[0].address), h.toWei("1000"))
        await h.expectThrow(tbridge.depositToLayer(depositAmount, LAYER_RECIPIENT)) // not approved
        await token.approve(await tbridge.address, h.toWei("900"))
        await h.expectThrow(tbridge.depositToLayer(0, LAYER_RECIPIENT)) // zero amount
        await h.expectThrow(tbridge.depositToLayer(h.toWei("21"), LAYER_RECIPIENT)) // over limit
        await tbridge.depositToLayer(depositAmount, LAYER_RECIPIENT)
        blocky1 = await h.getBlock()
        tbridgeBal = await token.balanceOf(await tbridge.address)
        assert.equal(tbridgeBal.toString(), h.toWei("1"))
        userBal = await token.balanceOf(await accounts[0].address)
        assert.equal(userBal.toString(), h.toWei("999"))
        expectedDepositLimit = BigInt(100e18) * BigInt(2) / BigInt(10) - BigInt(depositAmount)
        assert.equal(BigInt(await tbridge.depositLimitRecord()), expectedDepositLimit);
        await tbridge.refreshDepositLimit()
        assert.equal(BigInt(await tbridge.depositLimitRecord()), expectedDepositLimit);
        assert.equal(await tbridge.depositId(), 1)
        depositDetails = await tbridge.deposits(1)
        assert.equal(depositDetails.amount.toString(), depositAmount)
        assert.equal(depositDetails.recipient, LAYER_RECIPIENT)
        assert.equal(depositDetails.sender, await accounts[0].address)
        assert.equal(depositDetails.blockHeight, blocky1.number)
        assert.equal(await tbridge.depositId(), 1)
        await h.advanceTime(43200)
        expectedDepositLimit2 = (BigInt(100e18) + BigInt(depositAmount)) * BigInt(2) / BigInt(10)
        await tbridge.refreshDepositLimit()
        assert.equal(BigInt(await tbridge.depositLimitRecord()), expectedDepositLimit2);
    })
    it("depositLimit", async function () {
        expectedDepositLimit = BigInt(100e18) * BigInt(2) / BigInt(10)
        await tbridge.refreshDepositLimit()
        assert.equal(BigInt(await tbridge.depositLimitRecord()), expectedDepositLimit);
        await token.approve(await tbridge.address, h.toWei("900"))
        depositAmount = h.toWei("2")
        await tbridge.depositToLayer(depositAmount, LAYER_RECIPIENT)
        expectedDepositLimit = BigInt(100e18) * BigInt(2) / BigInt(10) - BigInt(depositAmount)
        await tbridge.refreshDepositLimit()
        assert.equal(BigInt(await tbridge.depositLimitRecord()), expectedDepositLimit);
        await h.advanceTime(43200)
        expectedDepositLimit2 = (BigInt(100e18) + BigInt(depositAmount)) / BigInt(5)
        await tbridge.refreshDepositLimit()
        assert.equal(BigInt(await tbridge.depositLimitRecord()), expectedDepositLimit2);
    })
    it("claim extraWithdraw", async function () {
        await tbridge.refreshDepositLimit()
        expectedDepositLimit = BigInt(100e18) * BigInt(2) / BigInt(10)
        depositAmount = expectedDepositLimit
        await h.expectThrow(tbridge.depositToLayer(depositAmount, LAYER_RECIPIENT)) // not approved
        await token.approve(await tbridge.address, h.toWei("100"))
        await tbridge.depositToLayer(depositAmount, LAYER_RECIPIENT)
        await h.advanceTime(43200) 
        await token.approve(await tbridge.address, h.toWei("100"))
        await tbridge.depositToLayer(depositAmount, LAYER_RECIPIENT)
        let _addy = await accounts[2].address
        value = h.getWithdrawValue(_addy,LAYER_RECIPIENT,40000000)
        blocky = await h.getBlock()
        timestamp = blocky.timestamp - 2
        aggregatePower = 3
        attestTimestamp = timestamp + 1
        previousTimestamp = 0
        nextTimestamp = 0
        newValHash = await h.calculateValHash(initialValAddrs, initialPowers)
        valCheckpoint = await h.calculateValCheckpoint(newValHash, threshold, valTimestamp)
        dataDigest = await h.getDataDigest(
            WITHDRAW1_QUERY_ID,
            value,
            timestamp,
            aggregatePower,
            previousTimestamp,
            nextTimestamp,
            valCheckpoint,
            attestTimestamp
        )
        currentValSetArray = await h.getValSetStructArray(initialValAddrs, initialPowers)
        sig1 = await h.layerSign(dataDigest, val1.privateKey)
        sig2 = await h.layerSign(dataDigest, val2.privateKey)
        sigStructArray = await h.getSigStructArray([sig1, sig2])
        oracleDataStruct = await h.getOracleDataStruct(
            WITHDRAW1_QUERY_ID,
            value,
            timestamp,
            aggregatePower,
            previousTimestamp,
            nextTimestamp,
            attestTimestamp
        )
        await h.advanceTime(43200)
        await tbridge.refreshDepositLimit()
        let _limit = await tbridge.depositLimit.call()
        assert(await token.balanceOf(_addy) == 0)
        await tbridge.withdrawFromLayer(
            oracleDataStruct,
            currentValSetArray,
            sigStructArray,
            1,
        )
        recipientBal = await token.balanceOf(_addy)
        assert(recipientBal - _limit == 0, "token balance should be correct")
        tokensToClaim = await tbridge.tokensToClaim(accounts[2].address)
        assert(tokensToClaim == BigInt(40e18) - BigInt(recipientBal), "tokensToClaim should be correct")
        await h.expectThrow(tbridge.claimExtraWithdraw(await accounts[2].address))
        await h.advanceTime(43200)
        await tbridge.claimExtraWithdraw(await accounts[2].address);
        await h.expectThrow(tbridge.claimExtraWithdraw(await accounts[2].address))
        recipientBal = await token.balanceOf(await accounts[2].address)
        assert(recipientBal == BigInt(40e18), "token balance should be correct")
        await h.advanceTime(43200)
        await tbridge.refreshDepositLimit()
        _limit = await tbridge.depositLimit()
        assert(BigInt(await tbridge.depositLimitRecord()) - expectedDepositLimit == BigInt(0));
        assert(_limit == expectedDepositLimit, "deposit Limit should be correct")
        tokensToClaim = await tbridge.tokensToClaim(accounts[2].address)
        assert(tokensToClaim == BigInt(0), "tokensToClaim should be correct")
    })
    
})
