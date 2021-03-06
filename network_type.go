package cfxaddress

import (
	"errors"
	"fmt"
	"strconv"
)

/*
Network-prefix:
    match network-id:
        case 1029: "cfx"
        case 1:    "cfxtest"
        case n:    "net[n]"
Examples of valid network-prefixes: "cfx", "cfxtest", "net17"
Examples of invalid network-prefixes: "bch", "conflux", "net1", "net1029"
*/

type NetworkType string

func (n NetworkType) String() string {
	return string(n)
}

const (
	MAINNET_PREFIX NetworkType = "cfx"
	TESTNET_PREFIX NetworkType = "cfxtest"

	MAINNET_ID uint32 = 1029
	TESTNET_ID uint32 = 1
)

func NewNetowrkType(nt string) (NetworkType, error) {
	if nt == MAINNET_PREFIX.String() || nt == TESTNET_PREFIX.String() {
		return NetworkType(nt), nil
	}
	return "", errors.New("invalid network type")
}

func NewNetworkTypeByID(networkID uint32) NetworkType {
	var nt NetworkType
	switch networkID {
	case MAINNET_ID:
		nt = MAINNET_PREFIX
	case TESTNET_ID:
		nt = TESTNET_PREFIX
	default:
		nt = NetworkType(fmt.Sprintf("net%v", networkID))
	}
	return nt
}

func (n NetworkType) ToNetworkID() (uint32, error) {
	switch n {
	case MAINNET_PREFIX:
		return MAINNET_ID, nil
	case TESTNET_PREFIX:
		return TESTNET_ID, nil
	default:
		if n[0:3] == "net" {
			netID, err := strconv.Atoi(string(n[3:]))
			if err != nil {
				return 0, err
			}
			if netID >= (1 >> 32) {
				return 0, errors.New("NetworkID must in range 0~0xffffffff")
			}
			return uint32(netID), nil
		}
		return 0, errors.New("Invalid network")
	}
}
