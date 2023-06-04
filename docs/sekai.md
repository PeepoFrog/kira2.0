# Sekai
`ver.: v0.3.16`
               
* [Sekai](#sekai)
  * [Context](#context)
    * [1. add-genesis-account](#1-add-genesis-account)
    * [2. collect-gentxs](#2-collect-gentxs)
    * [3. config](#3-config)
    * [4. debug](#4-debug)
      * [4.1 addr](#41-addr)
      * [4.2 pubkey](#42-pubkey)
      * [4.3 raw-bytes](#43-raw-bytes)
    * [5. export](#5-export)
    * [6. export-metadata](#6-export-metadata)
    * [7. export-minimized-genesis](#7-export-minimized-genesis)
    * [8. gentx](#8-gentx)
    * [9. gentx-claim](#9-gentx-claim)
    * [10. help](#10-help)
    * [11. init](#11-init)
    * [12. keys](#12-keys)
    * [13. new-genesis-from-exported](#13-new-genesis-from-exported)
    * [14. query](#14-query)
    * [15. rollback](#15-rollback)
    * [16. rosetta](#16-rosetta)
    * [17. start](#17-start)
    * [18. status](#18-status)
    * [19. tendermint](#19-tendermint)
    * [20. testnet](#20-testnet)
    * [21. tx](#21-tx)
    * [22. val-address](#22-val-address)
    * [23. valcons-address](#23-valcons-address)
    * [24. validate-genesis](#24-validate-genesis)
    * [25. version](#25-version)

## Context

### 1. add-genesis-account

### 2. collect-gentxs

### 3. config

### 4. debug

#### 4.1 addr

#### 4.2 pubkey

#### 4.3 raw-bytes

### 5. export

### 6. export-metadata

### 7. export-minimized-genesis

### 8. gentx

### 9. gentx-claim

### 10. help

### 11. init

### 12. keys

### 13. new-genesis-from-exported

### 14. query

### 15. rollback

### 16. rosetta

### 17. start
Run the full node application with Tendermint in or out of process. By
default, the application will run with Tendermint in process.

Pruning options can be provided via the '--pruning' flag or alternatively with '--pruning-keep-recent', 'pruning-keep-every', and 'pruning-interval' together.

For '--pruning' the options are as follows:

default: the last 100 states are kept in addition to every 500th state; pruning at 10 block intervals
nothing: all historic states will be saved, nothing will be deleted (i.e. archiving node)
everything: all saved states will be deleted, storing only the current and previous state; pruning at 10 block intervals
custom: allow pruning options to be manually specified through 'pruning-keep-recent', 'pruning-keep-every', and 'pruning-interval'

Node halting configurations exist in the form of two flags: '--halt-height' and '--halt-time'. During the ABCI Commit phase, the node will check if the current block height is greater than or equal to
the halt-height or if the current block time is greater than or equal to the halt-time. If so, the node will attempt to gracefully shutdown and the block will not be committed. In addition, the node
will not be able to commit subsequent blocks.

For profiling and benchmarking purposes, CPU profiling can be enabled via the '--cpu-profile' flag which accepts a path for the resulting pprof file.

The node may be started in a 'query only' mode where only the gRPC and JSON HTTP API services are enabled via the 'grpc-only' flag. In this mode, Tendermint is bypassed and can be used when legacy queries are needed after an on-chain upgrade is performed. Note, when enabled, gRPC will also be automatically enabled.

Usage:
```
sekaid start [flags]
```

| Flags                                           | Description                                                                                                                                    | Work |
|-------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------|------|
| --abci string                                   | genesis file chain-id, if left blank will be randomly created                                                                                  | ?    |
| --address string                                | Listen address (default "tcp://0.0.0.0:26658")                                                                                                 | ?    |
| --consensus.create_empty_blocks                 | set this to false to only produce blocks when there are txs or when the AppHash changes (default true)                                         | ?    |
| --consensus.create_empty_blocks_interval string | the possible interval between empty blocks (default "0s")                                                                                      | ?    |
| --consensus.double_sign_check_height int        | how many blocks to look back to check existence of the node's consensus votes before joining consensus                                         | ?    |
| --cpu-profile string                            | Enable CPU profiling and write to the provided file                                                                                            | ?    |
| --db_backend string                             | database backend: goleveldb \| cleveldb \| boltdb \| rocksdb \| badgerdb (default "goleveldb")                                                 | ?    |
| --db_dir string                                 | database directory (default "data")                                                                                                            | ?    |
| --fast_sync                                     | fast blockchain syncing (default true)                                                                                                         | ?    |
| --genesis_hash bytesHex                         | optional SHA-256 hash of the genesis file                                                                                                      | ?    |
| --grpc-only                                     | Start the node in gRPC query only mode (no Tendermint process is started)                                                                      | ?    |
| --grpc-web.address string                       | The gRPC-Web server address to listen on (default "0.0.0.0:9091")                                                                              | ?    |
| --grpc-web.enable                               | Define if the gRPC-Web server should be enabled. (Note: gRPC must also be enabled.) (default true)                                             | ?    |
| --grpc.address string                           | the gRPC server address to listen on (default "0.0.0.0:9090")                                                                                  | ?    |
| --grpc.enable                                   | Define if the gRPC server should be enabled (default true)                                                                                     | ?    |
| --halt-height uint                              | Block height at which to gracefully halt the chain and shutdown the node                                                                       | ?    |
| --halt-time uint                                | Minimum block time (in Unix seconds) at which to gracefully halt the chain and shutdown the node                                               | ?    |
| -h, --help                                      | help for start                                                                                                                                 | ?    |
| --iavl-disable-fastnode                         | Enable fast node for IAVL tree (default true)                                                                                                  | ?    |
| --inter-block-cache                             | Enable inter-block caching (default true)                                                                                                      | ?    |
| --inv-check-period uint                         | Assert registered invariants every N blocks                                                                                                    | ?    |
| --min-retain-blocks uint                        | Minimum block height offset during ABCI commit to prune Tendermint blocks                                                                      | ?    |
| --minimum-gas-prices string                     | Minimum gas prices to accept for transactions; Any fee in a tx must meet this minimum (e.g. 0.01photino;0.0001stake)                           | ?    |
| --moniker string                                | node name (default "validator.local")                                                                                                          | ?    |
| --p2p.external-address string                   | ip:port address to advertise to peers for them to dial                                                                                         | ?    |
| --p2p.laddr string                              | node listen address. (0.0.0.0:0 means any interface, any port) (default "tcp://0.0.0.0:26656")                                                 | ?    |
| --p2p.persistent_peers string                   | comma-delimited ID@host:port persistent peers                                                                                                  | ?    |
| --p2p.pex                                       | enable/disable Peer-Exchange (default true)                                                                                                    | ?    |
| --p2p.private_peer_ids string                   | comma-delimited private peer IDs                                                                                                               | ?    |
| --p2p.seed_mode                                 | enable/disable seed mode                                                                                                                       | ?    |
| --p2p.seeds string                              | comma-delimited ID@host:port seed nodes                                                                                                        | ?    |
| --p2p.unconditional_peer_ids string             | comma-delimited IDs of unconditional peers                                                                                                     | ?    |
| --p2p.upnp                                      | enable/disable UPNP port forwarding                                                                                                            | ?    |
| --priv_validator_laddr string                   | socket address to listen on for connections from external priv_validator process                                                               | ?    |
| --proxy_app string                              | proxy app address, or one of: 'kvstore', 'persistent_kvstore', 'counter', 'e2e' or 'noop' for local testing. (default "tcp://127.0.0.1:26658") | ?    |
| --pruning string                                | Pruning strategy (default\|nothing\|everything\|custom) (default "default")                                                                    | ?    |
| --pruning-interval uint                         | Height interval at which pruned heights are removed from disk (ignored if pruning is not 'custom')                                             | ?    |
| --pruning-keep-every uint                       | Offset heights to keep on disk after 'keep-every' (ignored if pruning is not 'custom')                                                         | ?    |
| --pruning-keep-recent uint                      | Number of recent heights to keep on disk (ignored if pruning is not 'custom')                                                                  | ?    |
| --rpc.grpc_laddr string                         | GRPC listen address (BroadcastTx only). Port required                                                                                          | ?    |
| --rpc.laddr string                              | RPC listen address. Port required (default "tcp://127.0.0.1:26657")                                                                            | yes  |
| --rpc.pprof_laddr string                        | pprof listen address (https://golang.org/pkg/net/http/pprof)                                                                                   | ?    |
| --rpc.unsafe                                    | enabled unsafe rpc methods                                                                                                                     | ?    |
| --state-sync.snapshot-interval uint             | State sync snapshot interval                                                                                                                   | ?    |
| --state-sync.snapshot-keep-recent uint32        | State sync snapshot to keep (default 2)                                                                                                        | ?    |
| --trace-store string                            | Enable KVStore tracing to an output file                                                                                                       | ?    |
| --transport string                              | Transport protocol: socket, grpc (default "socket")                                                                                            | ?    |
| --unsafe-skip-upgrades ints                     | Skip a set of upgrade heights to continue the old binary                                                                                       | ?    |
| --with-tendermint                               | Run abci app embedded in-process with tendermint (default true)                                                                                | ?    |
| --x-crisis-skip-assert-invariants               | Skip x/crisis invariants check on startup                                                                                                      | ?    |

| Global flags         | Description                                                                            | Work |
| -------------------- | -------------------------------------------------------------------------------------- | ---- |
| `--home`             | directory for config and data (default `"/root/.sekaid"`)                              | yes  |
| `--log_format`       | The logging format (`json\|plain`) (default `"plain"`)                                 | ?no  |
| `--log_level string` | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ?no  |
| `--trace`            | Print out full stack trace on errors                                                   | ?no  |

```
sekaid start --rpc.laddr "tcp://0.0.0.0:26657" 
```
### 18. status
Query remote node for status

#TODO: the descritpion is not accurate. `sekaid status` queries the node where it was launched by default.

Usage:
```
sekaid status [flags]
```

| Flags    | Description                                            | Work |
| -------- | ------------------------------------------------------ | ---- |
| `--help` | Help for status                                        | yes  |
| `--node` | Node to connect to (default `"tcp://localhost:26657"`) | yes  |

| Global flags         | Description                                                                            | Work |
| -------------------- | -------------------------------------------------------------------------------------- | ---- |
| `--home`             | directory for config and data (default `"/root/.sekaid"`)                              | yes  |
| `--log_format`       | The logging format (`json\|plain`) (default `"plain"`)                                 | ?no  |
| `--log_level string` | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ?no  |
| `--trace`            | Print out full stack trace on errors                                                   | ?no  |

```
/# sekaid status --help
Query remote node for status

Usage:
  sekaid status [flags]

Flags:
  -h, --help          help for status
  -n, --node string   Node to connect to (default "tcp://localhost:26657")

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
# sekaid status -n=tcp://10.1.0.2:26657 | jq
{
  "NodeInfo": {
    "protocol_version": {
      "p2p": "8",
      "block": "11",
      "app": "0"
    },
    "id": "213267558cc0febde6f1b8149f92fecf16d48b38",
    "listen_addr": "tcp://168.119.228.165:36656",
    "network": "localnet-4",
    "version": "0.34.22",
    "channels": "40202122233038606100",
    "moniker": "KIRA VALIDATOR NODE",
    "other": {
      "tx_index": "on",
      "rpc_address": "tcp://0.0.0.0:26657"
    }
  },
  "SyncInfo": {
    "latest_block_hash": "EF98E1777F8D9477AED36B0BB45238445D3CED2C5E3AA7EF0EAAB984BEED382D",
    "latest_app_hash": "37B64A94D872F0DD081EACE2ECCAD5B397121C7DD6F9E6D3B419299E94D74BD6",
    "latest_block_height": "74627",
    "latest_block_time": "2023-06-04T10:14:45.658686623Z",
    "earliest_block_hash": "26AD6395E9AB39AE545D53F64B2DFF4C664F55448A5A96F9963336386CFE2B3C",
    "earliest_app_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
    "earliest_block_height": "1",
    "earliest_block_time": "2023-05-25T12:56:59.153143882Z",
    "catching_up": false
  },
  "ValidatorInfo": {
    "Address": "AFC8EBD65CE1E7DD38E1E4DD514E9B03A0085E98",
    "PubKey": {
      "type": "tendermint/PubKeyEd25519",
      "value": "4q2MwShsxSt2UBXmEv/rLXevigs6uln9J74uxA/XbRE="
    },
    "VotingPower": "1"
  }
}
```

### 19. tendermint

### 20. testnet

### 21. tx

### 22. val-address

### 23. valcons-address

### 24. validate-genesis

### 25. version
