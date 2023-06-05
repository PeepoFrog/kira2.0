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
Print the application binary version information
Usage:
```
sekaid version [flags]
```
| Flags              | Description                                                                        | Work |
|--------------------|------------------------------------------------------------------------------------|------|
| --help             | help for version                                                                   | yes  |
| --long             | Print long version information                                                     | yes  |
| --output string    | Output format (text\|json) (default "text")                                        | ?no  |

| Global flags         | Description                                                                            | Work |
| -------------------- | -------------------------------------------------------------------------------------- | ---- |
| --home string      | directory for config and data (default "/root/.sekaid")                                  | yes  |
| --log_format       | The logging format (json\|plain) (default "plain")                                       | no   |
| --log_level string | The logging level (trace\|debug\|info\|warn\|error\|fatal\|panic) (default "info")       | ?    |
| --trace            | print out full stack trace on errors                                                     | ?    |
```
#sekaid version  --help
Print the application binary version information

Usage:
  sekaid version [flags]

Flags:
  -h, --help            help for version
      --long            Print long version information
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
#sekaid version
v0.3.15.1
```