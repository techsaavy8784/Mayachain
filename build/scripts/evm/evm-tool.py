import argparse
import functools
import json
import logging
import os
import socket
import sys
import time
from contextlib import closing
from urllib.parse import urlparse

import requests
from eth_typing import ChecksumAddress
from web3 import HTTPProvider, Web3
from web3.middleware import geth_poa_middleware
from web3.types import TxParams, Wei

########################################################################################
# EVMSetupTool
########################################################################################


class EVMSetupTool:
    """
    EVMSetupTool is a tool to setup a local EVM network for testing purposes. It deploys
    the required router and token contracts and provides a set of convenience actions.
    """

    default_gas = 65000
    gas_per_byte = 68
    zero_address = Web3.toChecksumAddress("0x0000000000000000000000000000000000000000")
    headers = {"content-type": "application/json", "cache-control": "no-cache"}
    admin_key = "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"
    password = ""
    arb_password = "passphrase"
    arb_admin_key = "0xb6b15c8cb491557369f3c7d2c287b053eb229daa9c22138887752191c9520659"
    erc20rune = "0x3155BA85D5F96b2d030a4966AF206230e46849cb"  # mainnet, does not matter

    def __init__(self, chain, base_url):
        # setup web3 client
        self.chain = chain
        self.rpc_url = base_url
        self.web3 = Web3(HTTPProvider(self.rpc_url))
        if chain == "ARB":
            # Setup eth connection for bridging funds
            self.w3eth = Web3(Web3.HTTPProvider("http://ethereum:8545"))
        self.web3.middleware_onion.inject(geth_poa_middleware, layer=0)
        admin_key = self.admin_key

        if self.chain == "ARB":
            admin_key = self.arb_admin_key

        # get admin account address
        self.account = self.web3.eth.account.privateKeyToAccount(admin_key)
        self.addr = self.account.address

        # import admin key (assume loaded if running hardhat)
        if self.web3.net.version != "31337":
            # check if account already exists
            if (
                self.addr not in self.web3.geth.personal.list_accounts()
                and self.chain != "ARB"
            ):
                print("importing admin key...")
                self.web3.geth.personal.import_raw_key(admin_key, "")
                self.web3.geth.personal.import_raw_key(self.arb_admin_key, "")

            # setup admin account
            if (
                self.chain != "AVAX" and self.chain != "ARB"
            ):  # geth creates a random first account so skip it
                coinbase_addr = self.web3.geth.personal.list_accounts()[0]
                if self.web3.eth.getBalance(self.addr) == 0:
                    self.fund_account(coinbase_addr, self.addr, int(1000e18))  # 1k ETH
                # Setup Arb accounts from ETH
                if self.chain == "ETH":
                    self.fund_account(
                        coinbase_addr,
                        "0x3f1Eae7D46d88F08fc2F8ed27FCb2AB183EB2d0E",
                        int(1000000e18),
                    )  # 1m ETH for chain owner in Arb
                    self.fund_account(
                        coinbase_addr,
                        "0x6A568afe0f82d34759347bb36F14A6bB171d2CBe",
                        int(1000e18),
                    )  # 1k ETH for Validator
                    self.fund_account(
                        coinbase_addr,
                        "0xe2148eE53c0755215Df69b2616E552154EdC584f",
                        int(1000e18),
                    )  # 1k ETH for Sequencer
                    self.fund_account(
                        coinbase_addr,
                        "0x5E1497dD1f08C87b2d8FE23e9AAB6c1De833D927",
                        int(1000e18),
                    )  # 1k ETH for l2owner
                    # l1 traffic for arb
                    # self.fund_account(coinbase_addr, "0x058E6C774025ade66153C65672219191c72c7095", int(1000e18))  # 1k ETH for user 1
            else:
                balance = self.web3.eth.getBalance(self.addr)
                print(f"{self.addr} balance: {balance}")

        # unlock default account (assume unlocked if running hardhat)
        if self.web3.net.version != "31337":
            self.web3.eth.defaultAccount = self.addr
            if self.chain != "ARB":
                self.web3.geth.personal.unlock_account(self.addr, self.password)

    def gas_asset(self):
        if self.chain == "AVAX":
            return "AVAX.AVAX"
        elif self.chain == "ETH":
            return "ETH.ETH"
        elif self.chain == "ARB":
            return "ARB.ETH"
        else:
            logging.fatal(f"unknown chain: {self.chain}")

    def fund_account(self, from_address, to_address, amount):
        print(f"funding account: {to_address} {amount}")
        tx: TxParams = {
            "from": Web3.toChecksumAddress(from_address),
            "to": Web3.toChecksumAddress(to_address),
            "value": amount,
            "gas": self.calculate_gas(""),
        }

        # wait for the transaction to be mined
        tx_hash = self.web3.geth.personal.send_transaction(tx, self.password)
        self.web3.eth.waitForTransactionReceipt(tx_hash)

    def calculate_gas(self, msg) -> Wei:
        return Wei(self.default_gas + self.gas_per_byte * len(msg))

    def deploy_init_contracts(self):
        if self.chain == "ARB":
            self.bridge_arb_funds()
            # Check for brdiged funds
            balance = self.web3.eth.getBalance(self.addr)
            while balance == 0:
                balance = self.web3.eth.getBalance(self.addr)
                print(f"Waiting bridge funds to be available in arb: {balance}")
                time.sleep(20)
        if self.chain != "ARB":
            self.deploy_token()
        self.deploy_router()

    def deploy_token(self):
        print("deploying token contract...")
        tx_hash = self.token_contract().constructor().transact()
        receipt = self.web3.eth.waitForTransactionReceipt(tx_hash)
        print(f"Token Contract Address: {receipt.contractAddress}")

    def deploy_router(self):
        print("deploying router contract...")
        router, args = self.router_contract()
        if self.chain != "ARB":
            tx_hash = router.constructor(*args).transact()
        else:
            tx_hash = self.send_router_tx_arb(router)
        receipt = self.web3.eth.waitForTransactionReceipt(tx_hash)
        print(f"Router Contract Address: {receipt.contractAddress}")

    def bridge_arb_funds(self):
        # Set up the transaction
        from_address_with_checksum = Web3.toChecksumAddress(
            "0x3f1Eae7D46d88F08fc2F8ed27FCb2AB183EB2d0E"
        )
        to_address_with_checksum = Web3.toChecksumAddress(
            "0xa5d8d368c4fc06d71724d91561d6f2a880fd4fd9"
        )
        tx = {
            "nonce": self.w3eth.eth.getTransactionCount(from_address_with_checksum),
            "to": to_address_with_checksum,
            "value": self.w3eth.toWei(100000, "ether"),
            "gas": 2000000,
            "gasPrice": self.w3eth.toWei("50", "gwei"),
            "data": "0x0f4d14e9000000000000000000000000000000000000000000000000000082f79cd90000",
            "chainId": 1337,
        }

        tx_2 = {
            "nonce": self.w3eth.eth.getTransactionCount(from_address_with_checksum) + 1,
            "to": to_address_with_checksum,
            "value": self.w3eth.toWei(1000, "ether"),
            "gas": 2000000,
            "gasPrice": self.w3eth.toWei("50", "gwei"),
            "data": "0x0f4d14e9000000000000000000000000000000000000000000000000000082f79cd90000",
            "chainId": 1337,
        }

        # Sign and send the transactions
        signed_tx = self.w3eth.eth.account.signTransaction(tx, self.arb_admin_key)
        tx_hash = self.w3eth.eth.sendRawTransaction(signed_tx.rawTransaction)

        signed_tx_2 = self.w3eth.eth.account.signTransaction(tx_2, self.arb_admin_key)
        tx_hash_2 = self.w3eth.eth.sendRawTransaction(signed_tx_2.rawTransaction)

        # Get the transaction hash
        print("Bridge first hash:", self.w3eth.toHex(tx_hash))
        print("Bridge second hash:", self.w3eth.toHex(tx_hash_2))

        # Wait for balance in arb to change
        balance = self.web3.eth.getBalance(self.addr)
        while balance == 0:
            balance = self.web3.eth.getBalance(self.addr)
            print(f"Waiting bridge funds to be available in arb: {balance}")
            time.sleep(20)

    def send_router_tx_arb(self, router):
        nonce = self.web3.eth.getTransactionCount(self.addr)
        transaction = router.constructor().buildTransaction(
            {
                "chainId": 412346,
                "gas": 2000000,
                "gasPrice": self.web3.toWei("50", "gwei"),
                "nonce": nonce,
            }
        )
        signed_txn = self.web3.eth.account.signTransaction(
            transaction, self.arb_admin_key
        )
        return self.web3.eth.sendRawTransaction(signed_txn.rawTransaction)

    # --------------------------------- helpers ---------------------------------

    def token_contract(self, address=None):
        with open(os.path.join(os.path.dirname(__file__), "token-abi.json")) as f:
            abi = json.load(f)
        with open(os.path.join(os.path.dirname(__file__), "token-bytecode.txt")) as f:
            bytecode = f.read()
        return self.web3.eth.contract(abi=abi, bytecode=bytecode, address=address)

    # NOTE: returns the router contract and the constructor args
    def router_contract(self, address=None):
        abi_file = "router-abi.json"
        bytecode_file = "router-bytecode.txt"
        args = []

        # if on eth the router constructor takes ERC20 RUNE token address
        if self.chain == "ETH":
            abi_file = "eth-" + abi_file
            bytecode_file = "eth-" + bytecode_file
            args = [self.erc20rune]

        # load abi and bytecode
        with open(os.path.join(os.path.dirname(__file__), abi_file), "r") as f:
            abi = json.load(f)
        with open(os.path.join(os.path.dirname(__file__), bytecode_file), "r") as f:
            bytecode = f.read()
        return self.web3.eth.contract(abi=abi, bytecode=bytecode, address=address), args

    # --------------------------------- utility actions ---------------------------------

    @functools.lru_cache
    def get_vault_addr(self) -> ChecksumAddress:
        data = requests.get("http://localhost:1317/thorchain/inbound_addresses").json()
        for vault in data:
            if vault["chain"] == self.chain:
                return Web3.toChecksumAddress(vault["address"])

        raise ValueError(f"could not find {self.chain} vault")

    @functools.lru_cache
    def get_router_addr(self) -> ChecksumAddress:
        data = requests.get("http://localhost:1317/thorchain/inbound_addresses").json()
        for vault in data:
            if vault["chain"] == self.chain:
                return Web3.toChecksumAddress(vault["router"])

        raise ValueError(f"could not find {self.chain} router")

    def token_balance(self, args):
        if args.address is None:
            args.address = self.addr  # default to our address
        if args.token_address is None:
            raise ValueError("token-address is required")

        token = self.token_contract(address=Web3.toChecksumAddress(args.token_address))
        balance = token.functions.balanceOf(Web3.toChecksumAddress(args.address)).call()
        print(f"Token Balance: {balance}")

    def swap_in(self, args):
        if args.agg_address is None:
            raise ValueError("agg-address is required")
        if args.token_address is None:
            raise ValueError("token-address is required")

        # load aggregator contract - swapIn is not consistent across all aggregators
        with open(os.path.join(os.path.dirname(__file__), "aggregator-abi.json")) as f:
            abi = json.load(f)

        # create contract instance
        agg = self.web3.eth.contract(address=args.agg_address, abi=abi)

        # approve spending
        token = self.token_contract(address=Web3.toChecksumAddress(args.token_address))
        approve_tx_hash = token.functions.approve(
            agg.functions.tokenTransferProxy().call(), args.amount
        ).transact()
        approve_receipt = self.web3.eth.waitForTransactionReceipt(approve_tx_hash)
        print(f"Approve Spending Receipt: {approve_receipt}")

        # swap in
        tx_hash = agg.functions.swapIn(
            Web3.toChecksumAddress(self.get_router_addr()),
            Web3.toChecksumAddress(self.get_vault_addr()),
            f"SWAP:THOR.RUNE:{args.thor_address}",
            Web3.toChecksumAddress(args.token_address),
            args.amount,
            0,
            9999999999,
        ).transact()

        receipt = self.web3.eth.waitForTransactionReceipt(tx_hash)
        print(f"Swap-In Receipt: {receipt}")

    def deposit(self, args):
        router, _ = self.router_contract(address=self.get_router_addr())
        memo = args.memo or f"ADD:{self.gas_asset()}:{args.thor_address}"
        tx_hash = router.functions.deposit(
            Web3.toChecksumAddress(self.get_vault_addr()),
            self.zero_address,
            0,
            memo,
        ).transact({"value": Wei(args.amount)})
        receipt = self.web3.eth.waitForTransactionReceipt(tx_hash)
        print(f"Deposit Receipt: {receipt}")

    def deposit_token(self, args):
        if args.token_address is None:
            raise ValueError("token-address is required")
        if args.thor_address is None:
            raise ValueError("thor-address is required")

        router, _ = self.router_contract(address=self.get_router_addr())
        token = self.token_contract(address=Web3.toChecksumAddress(args.token_address))

        tx_hash = token.functions.approve(
            self.get_router_addr(), args.amount
        ).transact()
        receipt = self.web3.eth.waitForTransactionReceipt(tx_hash)
        print(f"Approve Receipt: {receipt}")

        memo = (
            args.memo
            or f"ADD:{args.chain}.TKN-{args.token_address.upper()}:{args.thor_address}"
        )
        tx_hash = router.functions.deposit(
            self.get_vault_addr(),
            Web3.toChecksumAddress(args.token_address),
            args.amount,
            memo,
        ).transact()
        receipt = self.web3.eth.waitForTransactionReceipt(tx_hash)
        print(f"Deposit Receipt: {receipt}")

    def vault_allowance(self, args):
        if args.token_address is None:
            raise ValueError("token-address is required")

        router, _ = self.router_contract(address=self.get_router_addr())
        result = router.functions.vaultAllowance(
            self.get_vault_addr(),
            Web3.toChecksumAddress(args.token_address),
        ).call()
        print(f"Vault Allowance Result: {result}")


########################################################################################
# Helpers
########################################################################################


def check_socket(host, port):
    with closing(socket.socket(socket.AF_INET, socket.SOCK_STREAM)) as sock:
        if sock.connect_ex((host, port)) == 0:
            return True
        else:
            return False


########################################################################################
# Main
########################################################################################


def main():
    # config
    default_rpc = {
        "AVAX": "http://avalanche:9650/ext/bc/C/rpc",
        "ETH": "http://ethereum:8545",
        "ARB": "http://arbitrum:8547",
    }

    # parse args
    parser = argparse.ArgumentParser()
    parser.add_argument("--chain", help="chain name", choices=default_rpc.keys())
    parser.add_argument(
        "--action",
        help="action to perform",
        choices=[
            "deploy",
            "deposit",
            "token-balance",
            "deposit-token",
            "vault-allowance",
            "swap-in",
        ],
    )

    # only used for extended commands
    parser.add_argument("--address", help="the address")
    parser.add_argument("--token-address", help="the token address")
    parser.add_argument("--vault-address", help="the vault address")
    parser.add_argument("--agg-address", help="the aggregator address")
    parser.add_argument(
        "--thor-address",
        help="the memo",
        default="tthor1uuds8pd92qnnq0udw0rpg0szpgcslc9p8lluej",  # cat
    )
    parser.add_argument("--memo", help="the memo for the deposit call (default is add)")

    # defaults are scoped to other flags
    args, _ = parser.parse_known_args()
    parser.add_argument(
        "--amount",
        help="the amount",
        type=int,
        # 100k USD or 1 ETH
        default=int(1000e6 if args.action == "swap-in" else 1e18),
    )
    parser.add_argument("--rpc", help="rpc address", default=default_rpc[args.chain])
    args = parser.parse_args()

    # check that the port is open
    t = urlparse(args.rpc)
    for i in range(1, 30):
        if check_socket(t.hostname, t.port):
            time.sleep(1)
            break
        if i == 30:
            logging.error(f"{args.chain}: {t.hostname}:{t.port} not open")
            sys.exit(1)
        time.sleep(1)

    # run the action
    setup_tool = EVMSetupTool(args.chain, args.rpc)
    mux = {
        "deploy": setup_tool.deploy_init_contracts,
        "deposit": lambda: setup_tool.deposit(args),
        "token-balance": lambda: setup_tool.token_balance(args),
        "deposit-token": lambda: setup_tool.deposit_token(args),
        "vault-allowance": lambda: setup_tool.vault_allowance(args),
        "swap-in": lambda: setup_tool.swap_in(args),
    }
    if args.action:
        mux[args.action]()


if __name__ == "__main__":
    main()
