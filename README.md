# wanchain-cli
CLI for a Wanchain RPC endpoint

# Usage

It currently supports a single command that lists the activity of validator
nodes in a more readable table than the RPC output.

```console
$ wanchain-cli activity
+----------+------+--------------------------------------------+--------+--------+
| EPOCH ID | ROLE | ADDRESS                                    | ACTIVE | BLOCKS |
+----------+------+--------------------------------------------+--------+--------+
|    18143 | EP   | 0x0288c83219701766197373d1149f616c62b52a7d | true   |        |
|    18143 | EP   | 0xfc2730f75330bb75cb28fcff12f0aea5b6e433e1 | true   |        |
|    18143 | EP   | 0x0e83cb3369e065c69b6704f0b16fc2899a6bedb8 | true   |        |
|    18143 | EP   | 0x7212b9e259792879d85ca3227384f1005437e5f5 | true   |        |
|    18143 | EP   | 0x7212b9e259792879d85ca3227384f1005437e5f5 | true   |        |
|    18143 | EP   | 0x7212b9e259792879d85ca3227384f1005437e5f5 | true   |        |
|      ... | ...  | ...                                        | ...    |        |
+----------+------+--------------------------------------------+--------+--------+
```

It uses `pos_getActivity` to print Epoch Leader, Random Number Proposer and Slot
Leader activity sorted by Epoch ID. Active shows whether a validator was
considered active and Blocks shows the number of blocks when in Slot leader role.

# Flags

By default it uses `https://mywanwallet.io/api` for the RPC endpoint, `18143` as
the first Epoch ID to print from and the current Epoch ID as the last Epoch ID.

You can use:
* `--rpc` to change the RPC endpoint
* `--from-epoch-id` to change the starting Epoch ID
* `--to-epoch-id` to change the final Epoch ID
* `--validator-address` to filter by a particular validator

```console
$ wanchain-cli --rpc https://mywanwallet.nl/api activity \
  --validator-address 0x1f45cad3c17ced4d7596a5b40280a3f024b971f4 \
  --from-epoch-id 18148 --to-epoch-id 18150
+----------+------+--------------------------------------------+--------+--------+
| EPOCH ID | ROLE | ADDRESS                                    | ACTIVE | BLOCKS |
+----------+------+--------------------------------------------+--------+--------+
|    18148 | RP   | 0x1f45cad3c17ced4d7596a5b40280a3f024b971f4 | true   |        |
|    18148 | RP   | 0x1f45cad3c17ced4d7596a5b40280a3f024b971f4 | true   |        |
|    18148 | SL   | 0x1f45cad3c17ced4d7596a5b40280a3f024b971f4 |        | 337    |
|    18149 | RP   | 0x1f45cad3c17ced4d7596a5b40280a3f024b971f4 | true   |        |
|    18149 | RP   | 0x1f45cad3c17ced4d7596a5b40280a3f024b971f4 | true   |        |
|    18150 | RP   | 0x1f45cad3c17ced4d7596a5b40280a3f024b971f4 | true   |        |
+----------+------+--------------------------------------------+--------+--------+
```

You can also use a Docker image if you don't like compiling code:

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
