[![Go Report Card](https://goreportcard.com/badge/trezor/blockbook)](https://goreportcard.com/report/trezor/blockbook)
[![FLO Release](https://github.com/ranchimall/blockbook/actions/workflows/build-release-flo-deb.yml/badge.svg?branch=flo)](https://github.com/ranchimall/blockbook/actions/workflows/build-release-flo-deb.yml)

# IMPORTANT Changes needed to migrate to blockbook from flosight
1. blockbook URL link
2. API URIs (path)
3. response object path for tx vin floID (for processes that need to find sender id from tx details)
4. The link for fetching details of transactions for every ADDRESS has changed slightly

- For point 3

  We need to change

  vin[x].addr 
  to
  vin[x].addresses[0]

===

For UI 

- OLD BLOCK LINK
  https://flosight.duckdns.org/block/48fb49a89317c0852eb894b0251ae3c830621d416b2481ecd1d541d0b01e5ab8

  https://flosight.duckdns.org/api/block/48fb49a89317c0852eb894b0251ae3c830621d416b2481ecd1d541d0b01e5ab8

- NEW BLOCK LINK
  https://blockbook.ranchimall.net/block/48fb49a89317c0852eb894b0251ae3c830621d416b2481ecd1d541d0b01e5ab8

  https://blockbook.ranchimall.net/api/block/48fb49a89317c0852eb894b0251ae3c830621d416b2481ecd1d541d0b01e5ab8

- OLD tx LINK
  https://flosight.duckdns.org/tx/ae7886df36832824e0e0f37cd003150cc7222e35bc7320d226f1561a598ce39a

  https://flosight.duckdns.org/api/tx/ae7886df36832824e0e0f37cd003150cc7222e35bc7320d226f1561a598ce39a

- NEW tx LINK
  https://blockbook.ranchimall.net/tx/ae7886df36832824e0e0f37cd003150cc7222e35bc7320d226f1561a598ce39a

  https://blockbook.ranchimall.net/api/tx/ae7886df36832824e0e0f37cd003150cc7222e35bc7320d226f1561a598ce39a

- Make this change for sender address in tx API here

  vin[x].addr 
  to
  vin[x].addresses[0]

  OLD Address LINK
  https://flosight.duckdns.org/address/F9RgXE1WYxHwvejktpi6EnjnscEkpiWxvS

  https://flosight.duckdns.org/api/address/F9RgXE1WYxHwvejktpi6EnjnscEkpiWxvS

NEW ADDRESS LINK
https://blockbook.ranchimall.net/address/F9RgXE1WYxHwvejktpi6EnjnscEkpiWxvS

https://blockbook.ranchimall.net/api/address/F9RgXE1WYxHwvejktpi6EnjnscEkpiWxvS

===
OTHER CHANGES
OLD
https://flosight.ranchimall.net/api/addrs/FBL45szT4jDQmViVirUxPeJjn1tkCDMxeT/txs

https://flosight.ranchimall.net/api/addrs/FBL45szT4jDQmViVirUxPeJjn1tkCDMxeT/txs?from=0&to=359

NEW
https://blockbook.ranchimall.net/api/address/FBL45szT4jDQmViVirUxPeJjn1tkCDMxeT?details=basic

https://blockbook.ranchimall.net/api/address/FBL45szT4jDQmViVirUxPeJjn1tkCDMxeT?details=txs

with paging for address
https://blockbook.ranchimall.net/api/address/FBL45szT4jDQmViVirUxPeJjn1tkCDMxeT?details=txs?details=txs&pageSize=10&page=2


```javascript
//Flosight
var response = ajax("GET",`api/addrs/${addr}/txs?from=0&to=${nRequired}`);
//Blockbook
var response = ajax("GET",`api/address/${addr}&details=txs?pageSize=${nRequired}&page=1`);
```

==

- OLD LINK FOR DIRECT BALANCE QUERY
  https://flosight.duckdns.org/api/addr/FBL45szT4jDQmViVirUxPeJjn1tkCDMxeT/balance

  NO BLOCKBOOK EQUIVALENT
  THIS HAS TO BE FETCHED FROM THIS FORMAT
  https://blockbook.ranchimall.net/api/address/FBL45szT4jDQmViVirUxPeJjn1tkCDMxeT?details=basic

```javascript
// https://blockbook.ranchimall.net/api/address/FBL45szT4jDQmViVirUxPeJjn1tkCDMxeT?details=basic  
// JSON data as a string
var jsonData = '{"addrStr":"FBL45szT4jDQmViVirUxPeJjn1tkCDMxeT","balance":0.1663,"balanceSat":16630000,"totalReceived":29.4465,"totalReceivedSat":2944650000,"totalSent":29.2802,"totalSentSat":2928020000,"unconfirmedBalance":0,"unconfirmedBalanceSat":0,"unconfirmedTxApperances":0,"txApperances":407}';

// Parse the JSON data into a JavaScript object
var dataObject = JSON.parse(jsonData);

// Extract the balance value
var balance = dataObject.balance;

// Output the balance
console.log(balance); // Output: 0.1663
```


# Blockbook

**Blockbook** is back-end service for Trezor wallet. Main features of **Blockbook** are:

-   index of addresses and address balances of the connected block chain
-   fast index search
-   simple blockchain explorer
-   websocket, API and legacy Bitcore Insight compatible socket.io interfaces
-   support of multiple coins (Bitcoin and Ethereum type) with easy extensibility to other coins
-   scripts for easy creation of debian packages for backend and blockbook

## Build and installation instructions

Officially supported platform is **Debian Linux** and **AMD64** architecture.

Memory and disk requirements for initial synchronization of **Bitcoin mainnet** are around 32 GB RAM and over 180 GB of disk space. After initial synchronization, fully synchronized instance uses about 10 GB RAM.
Other coins should have lower requirements, depending on the size of their block chain. Note that fast SSD disks are highly
recommended.

User installation guide is [here](<https://wiki.trezor.io/User_manual:Running_a_local_instance_of_Trezor_Wallet_backend_(Blockbook)>).

Developer build guide is [here](/docs/build.md).

Contribution guide is [here](CONTRIBUTING.md).

## Implemented coins

Blockbook currently supports over 30 coins. The Trezor team implemented

-   Bitcoin, Bitcoin Cash, Zcash, Dash, Litecoin, Bitcoin Gold, Ethereum, Ethereum Classic, Dogecoin, Namecoin, Vertcoin, DigiByte, Liquid

the rest of coins were implemented by the community.

Testnets for some coins are also supported, for example:

-   Bitcoin Testnet, Bitcoin Cash Testnet, ZCash Testnet, Ethereum Testnets (Goerli, Sepolia)

List of all implemented coins is in [the registry of ports](/docs/ports.md).

## Common issues when running Blockbook or implementing additional coins

#### Out of memory when doing initial synchronization

How to reduce memory footprint of the initial sync:

-   disable rocksdb cache by parameter `-dbcache=0`, the default size is 500MB
-   run blockbook with parameter `-workers=1`. This disables bulk import mode, which caches a lot of data in memory (not in rocksdb cache). It will run about twice as slowly but especially for smaller blockchains it is no problem at all.

Please add your experience to this [issue](https://github.com/trezor/blockbook/issues/43).

#### Error `internalState: database is in inconsistent state and cannot be used`

Blockbook was killed during the initial import, most commonly by OOM killer.
By default, Blockbook performs the initial import in bulk import mode, which for performance reasons does not store all data immediately to the database. If Blockbook is killed during this phase, the database is left in an inconsistent state.

See above how to reduce the memory footprint, delete the database files and run the import again.

Check [this](https://github.com/trezor/blockbook/issues/89) or [this](https://github.com/trezor/blockbook/issues/147) issue for more info.

#### Running on Ubuntu

[This issue](https://github.com/trezor/blockbook/issues/45) discusses how to run Blockbook on Ubuntu. If you have some additional experience with Blockbook on Ubuntu, please add it to [this issue](https://github.com/trezor/blockbook/issues/45).

#### My coin implementation is reporting parse errors when importing blockchain

Your coin's block/transaction data may not be compatible with `BitcoinParser` `ParseBlock`/`ParseTx`, which is used by default. In that case, implement your coin in a similar way we used in case of [zcash](https://github.com/trezor/blockbook/tree/master/bchain/coins/zec) and some other coins. The principle is not to parse the block/transaction data in Blockbook but instead to get parsed transactions as json from the backend.

## Data storage in RocksDB

Blockbook stores data the key-value store RocksDB. Database format is described [here](/docs/rocksdb.md).

## API

Blockbook API is described [here](/docs/api.md).
