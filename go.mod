module github.com/bdengine/go-ipfs-blockchain-eth

go 1.15

require (
	github.com/bdengine/go-ipfs-blockchain-standard v0.0.2
	github.com/ethereum/go-ethereum v1.10.13
	github.com/google/uuid v1.1.5
	github.com/prometheus/common v0.6.0
)

replace github.com/bdengine/go-ipfs-blockchain-standard => ../go-ipfs-blockchain-standard
