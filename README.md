# wanchain-cli

CLI for a Wanchain RPC endpoint

# Purpose

There is lots of useful information regarding the validators behind Wanchain's
RPC endpoint. However it's quite difficult to read. This tool provides several
commands to inspect the state of the validators and displays the information
in a more readable table layout.

# Usage

It needs an RPC endpoint to talk to. It only reads publicly available information.
In order to be on the safe side don't use an RPC endpoint that can spend your funds.
By default it uses https://mywanwallet.io/api which should work for most users.

You can then execute the commands described below.

# Commands

The following commands are available.

## activity

```console
$ wanchain-cli activity --validator-address 0xfc2730f75330bb75cb28fcff12f0aea5b6e433e1
+----------+------+--------------------------------------------+--------+--------+
| EPOCH ID | ROLE | ADDRESS                                    | ACTIVE | BLOCKS |
+----------+------+--------------------------------------------+--------+--------+
|    18232 | EP   | 0xfc2730f75330bb75cb28fcff12f0aea5b6e433e1 | true   |        |
|    18232 | EP   | 0xfc2730f75330bb75cb28fcff12f0aea5b6e433e1 | true   |        |
|    18232 | RP   | 0xfc2730f75330bb75cb28fcff12f0aea5b6e433e1 | true   |        |
|      ... | ...  | ...                                        | ...    |        |
|    18232 | RP   | 0xfc2730f75330bb75cb28fcff12f0aea5b6e433e1 | true   |        |
|    18232 | RP   | 0xfc2730f75330bb75cb28fcff12f0aea5b6e433e1 | true   |        |
|    18232 | SL   | 0xfc2730f75330bb75cb28fcff12f0aea5b6e433e1 |        | 1703   |
+----------+------+--------------------------------------------+--------+--------+
```

It uses `pos_getActivity` to print Epoch Leader, Random Number Proposer and Slot
Leader activity sorted by Epoch ID. Active shows whether a validator was
considered active and Blocks shows the number of blocks when in Slot leader role.

The command takes several flags, use `--help` and [see below](#flags) for a description
of the flags.

## incentives

```console
$ wanchain-cli incentives
+----------+-----------+--------------------------------------------+--------------+
| EPOCH ID | TYPE      | ADDRESS                                    | INCENTIVE    |
+----------+-----------+--------------------------------------------+--------------+
|    18143 | validator | 0x0288c83219701766197373d1149f616c62b52a7d |  10.20666436 |
|    18143 | validator | 0xfc2730f75330bb75cb28fcff12f0aea5b6e433e1 |  29.99589130 |
|    18143 | validator | 0x0e83cb3369e065c69b6704f0b16fc2899a6bedb8 |  15.51228138 |
|    18143 | validator | 0x7212b9e259792879d85ca3227384f1005437e5f5 |  14.39512195 |
|    18143 | validator | 0x7212b9e259792879d85ca3227384f1005437e5f5 |  14.39512195 |
|    18143 | validator | 0x7212b9e259792879d85ca3227384f1005437e5f5 |  14.39775959 |
|      ... | ...       | ...                                        |          ... |
+----------+-----------+--------------------------------------------+--------------+
|          | VALIDATOR |                                            |  97.65604884 |
|          | DELEGATOR |                                            | 461.46278191 |
+----------+-----------+--------------------------------------------+--------------+
```

This one uses `pos_getEpochIncentivePayDetail` to print rewards paid out to validators
and delegators sorted by Epoch ID. The incentive is displayed in WAN.

The command takes several flags, use `--help` and [see below](#flags) for a description
of the flags.

## selected

```console
$ wanchain-cli selected
+----------+------+--------------------------------------------+
| EPOCH ID | ROLE | ADDRESS                                    |
+----------+------+--------------------------------------------+
|    18333 | EP   | 0xd4a0f5df136f65345b9d6fcd306fdb018213005a |
|    18333 | EP   | 0x0288c83219701766197373d1149f616c62b52a7d |
|      ... | ...  |                                        ... |
+----------+------+--------------------------------------------+
```

Checks which validator is selected for a given Epoch ID. If no Epoch ID is provided
the next Epoch ID is used. If used after roughly 16:00 of Epoch ID x it will display
the selected validators for Epoch ID x+1.

The command takes several flags, use `--help` and [see below](#flags) for a description
of the flags.

## validators

```console
$ wanchain-cli validators --validator-address 0x7212b9e259792879d85ca3227384f1005437e5f5
+---------------------+--------------------------------------------+
| Address             | 0x7212B9e259792879d85Ca3227384f1005437E5f5 |
| PubSec256           | 0x04b6883...d1433f2c8f                     |
| PubBn256            | 0x1a63987...8c91c0db0                      |
| Amount              | 110149.00                                  |
| VotingPower         | 165223500.00 (1.50x)                       |
| LockEpochs          | 90                                         |
| NextLockEpochs      | 90                                         |
| From                | 0x6cb37EF52F184d631F425b18969D854F2E2E9077 |
| StakingEpoch        | 18321                                      |
| FeeRate             | 8.00%                                      |
| # Delegators        | 634 (5409873.70)                           |
| # Partners          | 5 (993559.00)                              |
| ValidatorAmount     | 6513581.70 (of 26402532.82)                |
| StakeWeight         | 24.67%                                     |
| PowerWeight         | 24.54%                                     |
| MaxFeeRate          | 10.00%                                     |
| FeeRateChangedEpoch | 18263                                      |
+---------------------+--------------------------------------------+
```

Given a validator address it prints some information about that validator.

The command takes several flags, use `--help` and [see below](#flags) for a description
of the flags.

## current

```console
$ wanchain-cli current
+--------------+---------+
| Epoch ID     |   18332 |
| Slot ID      |   12813 |
| Block Height | 7325035 |
+--------------+---------+
```

Displays the current Epoch ID, Slot ID and block height.

# Flags

Flags allow you to tailor your request to what you need. You can change the RPC
endpoint to connect to, configure the Epoch IDs for requests that displays data
across a range of Epochs, filter by a particular validator address and so on.

| Name                | Description                                                                             | Default value              |
|---------------------|-----------------------------------------------------------------------------------------|----------------------------|
| --rpc               | Specify the RPC endpoint to connect to                                                  | https://mywanwallet.io/api |
| --format            | Specify the output format. Possible values are: text, csv, html, markdown               | text                       |
| --from-epoch-id     | Specify the beginning Epoch ID. Can either be absolute or relative to --to-epoch-id     | -3 (current Epoch ID - 3)  |
| --to-epoch-id       | Specify the ending Epoch ID. Can either be absolute or relative to the current Epoch ID | 0 (current Epoch ID)       |
| --validator-address | Specify the address of the validator to filter the output by                            | ""                         |
| --delegator-address | Specify the address of the delegator to filter the output by                            | ""                         |
| --filter-type       | Filter results by incentive type, either blank, validator or delegator                  | ""                         |
| --block-height      | Specify the block height at which to take the data from                                 | 0 (current block height)   |

For example, if you want to see the activity of the GerWAN validator up to Epoch ID 18331
beginning from two Epochs before while using a different RPC endpoint, you can use the
following command.

```console
$ wanchain-cli --rpc https://mywanwallet.nl/api activity \
  --validator-address 0x83a1b91Ade06BCd2137D953DbDFAbE60f78Fd157 \
  --to-epoch-id 18331 --from-epoch-id -2
+----------+------+--------------------------------------------+--------+--------+
| EPOCH ID | ROLE | ADDRESS                                    | ACTIVE | BLOCKS |
+----------+------+--------------------------------------------+--------+--------+
|    18329 | EP   | 0x83a1b91ade06bcd2137d953dbdfabe60f78fd157 | true   |        |
|    18329 | EP   | 0x83a1b91ade06bcd2137d953dbdfabe60f78fd157 | true   |        |
|    18329 | SL   | 0x83a1b91ade06bcd2137d953dbdfabe60f78fd157 |        |    692 |
|    18330 | EP   | 0x83a1b91ade06bcd2137d953dbdfabe60f78fd157 | true   |        |
|    18330 | SL   | 0x83a1b91ade06bcd2137d953dbdfabe60f78fd157 |        |    698 |
|    18331 | EP   | 0x83a1b91ade06bcd2137d953dbdfabe60f78fd157 | true   |        |
|    18331 | EP   | 0x83a1b91ade06bcd2137d953dbdfabe60f78fd157 | true   |        |
|    18331 | RP   | 0x83a1b91ade06bcd2137d953dbdfabe60f78fd157 | true   |        |
|    18331 | SL   | 0x83a1b91ade06bcd2137d953dbdfabe60f78fd157 |        |    316 |
+----------+------+--------------------------------------------+--------+--------+
```

# Releases

You can find find automatically compiled binaries for several operating systems
on [the releases page](./releases).

You can also use an automatically built Docker image:

```console
$ docker run -it quay.io/linki/wanchain-cli activity
+----------+------+--------------------------------------------+--------+--------+
| EPOCH ID | ROLE | ADDRESS                                    | ACTIVE | BLOCKS |
+----------+------+--------------------------------------------+--------+--------+
|    18143 | EP   | 0x0288c83219701766197373d1149f616c62b52a7d | true   |        |
|    18143 | EP   | 0xfc2730f75330bb75cb28fcff12f0aea5b6e433e1 | true   |        |
|      ... | ...  | ...                                        | ...    |        |
+----------+------+--------------------------------------------+--------+--------+
```
