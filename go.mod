module github.com/zhouxun1995/ethclient

go 1.12

require (
	github.com/allegro/bigcache v1.2.0 // indirect
	github.com/ethereum/go-ethereum v1.8.27
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/zhouxun1995/xlog v0.0.0-20190425111112-92bd6a3bff95
	github.com/zhouxun1995/xutil v0.0.0-20190425112111-e64c53fe6144
)

replace golang.org/x/sys => /opt/go/src/golang.org/x/sys

replace golang.org/x/net => /opt/go/src/golang.org/x/net

replace golang.org/x/text => /opt/go/src/golang.org/x/text

replace golang.org/x/sync => /opt/go/src/golang.org/x/sync

replace golang.org/x/crypto => /opt/go/src/golang.org/x/crypto

replace golang.org/x/tools => /opt/go/src/golang.org/x/tools
