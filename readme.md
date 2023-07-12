# maprepro

**maprepro** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

`ignite chain build`

### Run

In three terminals start the nodes:
<br/>
`mapreprod start --home ./data/node-1`
<br/>
`mapreprod start --home ./data/node-2`
<br/>
`mapreprod start --home ./data/node-2`

### Send transaction

In a forth terminal send a tx:
<br/>
`mapreprod tx repro create-mymap 1 --from node-1 --home ./data/node-1`

### Result

This will result into error:
<br/>
`ERR prevote step: ProposalBlock is invalid err="wrong Block.Header.AppHash.  Expected 2BB849D1FA2775CC876FC56D8797FF931DA832ABDB2FC686A1AC97C6616A514F, got 032E76DE5259009B79E80579CC923073D0D2C729B2821B314F12010C239E54BF" height=10 module=consensus round=0`
<br/>
The reason is the non deterministic way of serializing a map in go with protobufs
