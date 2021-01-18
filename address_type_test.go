package cfxaddress

import (
	"encoding/hex"
	"testing"
)

func TestCalcAddressType(t *testing.T) {
	verifyAddressType(t, "006d49f8505410eb4e671d51f7d96d2c87807b09", BUILTIN)
	verifyAddressType(t, "106d49f8505410eb4e671d51f7d96d2c87807b09", USER)
	verifyAddressType(t, "806d49f8505410eb4e671d51f7d96d2c87807b09", CONTRACT)
	verifyAddressType(t, "0000000000000000000000000000000000000000", NULL)
}

func verifyAddressType(t *testing.T, hexAddress string, expect AddressType) {
	addr, err := hex.DecodeString(hexAddress)
	fatalIfErr(t, err)
	actual, err := CalcAddressType(addr)
	fatalIfErr(t, err)
	if actual != expect {
		t.Fatalf("expect %v actual %v", expect, actual)
	}
}
