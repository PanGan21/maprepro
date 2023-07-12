start-all:
		mapreprod start --home ./data/node-1 &
		mapreprod start --home ./data/node-2 &
		mapreprod start --home ./data/node-3
		
reset-all:
		mapreprod tendermint unsafe-reset-all --home ./data/node-1 &
		mapreprod tendermint unsafe-reset-all --home ./data/node-2 &
		mapreprod tendermint unsafe-reset-all --home ./data/node-3