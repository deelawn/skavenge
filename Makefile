.PHONY: all
all: compile

.PHONY: compile
compile:
	solc --abi --bin --base-path eth/node_modules --overwrite eth/skavenge.sol -o eth/build 
	abigen --bin eth/build/Skavenge.bin --abi eth/build/Skavenge.abi --pkg bindings --type Skavenge --out eth/bindings/bindings.go