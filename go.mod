module github.com/ipfs/go-ipfs-auth/auth-source-eth

go 1.15

require (
	github.com/ethereum/go-ethereum v1.10.13
	github.com/google/martian v2.1.0+incompatible
	github.com/google/uuid v1.1.5
	github.com/ipfs/go-ipfs-auth/standard v0.0.0
	github.com/prometheus/common v0.6.0
)

replace github.com/ipfs/go-ipfs-auth/standard => ../standard
