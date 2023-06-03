# Sekai
ver.: v0.3.16
               
- [Sekai](#sekai)
    - [Context](#context)
      - [3. config](#3-config)
      - [18. status](#18-status)
      - [25. version](#25-version)
### Context
#### 3. config
Create or query an application CLI configuration file

#if send ```sekaid config``` printing current config state of the node 

Usage:
```
sekaid config <key> [value] [flags]
```
| Flags              | Description                                                                        | Work |
|--------------------|------------------------------------------------------------------------------------|------|
| --help             | help for version                                                                   | yes  |
| Global Flags       |                                                                                    |      |
| --home string      | directory for config and data (default "/root/.sekaid")                            | no   |
| --log_format       | The logging format (json\|plain) (default "plain")                                 | no   |
| --log_level string | The logging level (trace\|debug\|info\|warn\|error\|fatal\|panic) (default "info") | ?    |
| --trace            | print out full stack trace on errors                                               | ?    |
```
#sekaid config --help
Create or query an application CLI configuration file

Usage:
  sekaid config <key> [value] [flags]

Flags:
  -h, --help   help for config

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```
#if keyring-backend value will be wrong it will break entire sekaid installation
```
sekaid config keyring-backend test
```



#### 18. status
Query remote node for status

#TODO: the descritpion is not accurate. `sekaid status` queries the node where it was launched by default.

Usage:
```
sekaid status [flags]
```
+-----------------------+-------------------------------------------------------------------------------+-----------+
|   Flags               |                                   Description                                 |   Work    |
+-----------------------+--------------+----------------------------------------------------------------+-----------+
|   --help              |help for status                                                                |   yes     |
|   --node              |Node to connect to (default "tcp://localhost:26657")                           |   yes     |
+-----------------------+-------------------------------------------------------------------------------+-----------+
|   --home              |directory for config and data (default "/root/.sekaid")                        |   yes     |
|   --log_format        |The logging format (json|plain) (default "plain")                              |   ?no     |
|   --log_level string  |The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")   |   ?no     |
|   --trace             |print out full stack trace on errors                                           |   ?no     |
+-----------------------+-------------------------------------------------------------------------------+-----------|

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
# sekaid status -n=tcp://10.1.0.2:26657
{"NodeInfo":{"protocol_version":{"p2p":"8","block":"11","app":"0"},"id":"213267558cc0febde6f1b8149f92fecf16d48b38","listen_addr":"tcp://168.119.228.165:36656","network":"localnet-4","version":"0.34.22","channels":"40202122233038606100","moniker":"KIRA VALIDATOR NODE","other":{"tx_index":"on","rpc_address":"tcp://0.0.0.0:26657"}},"SyncInfo":{"latest_block_hash":"60F7E2C5AAE691F6F709CC8030C6D87D94256F499D122C1FBBA93D8242C8A116","latest_app_hash":"B264A864767D9BE51E73A560008B44B9503E9798C7EBB570FEA7B67BC75BE54F","latest_block_height":"58760","latest_block_time":"2023-06-02T12:47:27.025413009Z","earliest_block_hash":"26AD6395E9AB39AE545D53F64B2DFF4C664F55448A5A96F9963336386CFE2B3C","earliest_app_hash":"E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855","earliest_block_height":"1","earliest_block_time":"2023-05-25T12:56:59.153143882Z","catching_up":false},"ValidatorInfo":{"Address":"AFC8EBD65CE1E7DD38E1E4DD514E9B03A0085E98","PubKey":{"type":"tendermint/PubKeyEd25519","value":"4q2MwShsxSt2UBXmEv/rLXevigs6uln9J74uxA/XbRE="},"VotingPower":"1"}}
```

#### 25. version
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
|                    |                                                                                    |      |
| Global Flags       |                                                                                    |      |
| --home string      | directory for config and data (default "/root/.sekaid")                            | yes  |
| --log_format       | The logging format (json\|plain) (default "plain")                                 | no   |
| --log_level string | The logging level (trace\|debug\|info\|warn\|error\|fatal\|panic) (default "info") | ?    |
| --trace            | print out full stack trace on errors                                               | ?    |
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