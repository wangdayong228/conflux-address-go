package cfxaddress

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
)

func TestCfxAddress(t *testing.T) {
	verify(t, "106d49f8505410eb4e671d51f7d96d2c87807b09", 1029, "cfx:0086ujfsa1a11uuecwen3xytdmp8f03v140ypk3mxc")
	verify(t, "106d49f8505410eb4e671d51f7d96d2c87807b09", 1029, "CFX:TYPE.USER:0086UJFSA1A11UUECWEN3XYTDMP8F03V140YPK3MXC")
	verify(t, "806d49f8505410eb4e671d51f7d96d2c87807b09", 1029, "cfx:0206ujfsa1a11uuecwen3xytdmp8f03v14ksvfyh2z")
	verify(t, "806d49f8505410eb4e671d51f7d96d2c87807b09", 1029, "CFX:TYPE.CONTRACT:0206UJFSA1A11UUECWEN3XYTDMP8F03V14KSVFYH2Z")
	verify(t, "006d49f8505410eb4e671d51f7d96d2c87807b09", 1029, "cfx:0006ujfsa1a11uuecwen3xytdmp8f03v1400dt9usz")
	verify(t, "006d49f8505410eb4e671d51f7d96d2c87807b09", 1029, "CFX:TYPE.BUILTIN:0006UJFSA1A11UUECWEN3XYTDMP8F03V1400DT9USZ")
}

func verify(t *testing.T, hexAddressStr string, networkID uint32, base32Address string) {
	hexAddress, err := hex.DecodeString(hexAddressStr)
	fmt.Printf("hexAddress:%x\n", hexAddress)
	fatalIfErr(t, err)

	cfxAddressFromHex, err := NewCfxAddressByHexAddress(hexAddress, networkID)
	fatalIfErr(t, err)

	fmt.Printf("cfxAddressFromHex %v\n", cfxAddressFromHex)
	cfxAddressFromBase32, err := NewCfxAddressByBase32String(base32Address)
	fatalIfErr(t, err)

	if !reflect.DeepEqual(cfxAddressFromHex, cfxAddressFromBase32) {
		t.Fatalf("expect %#v, actual %#v", cfxAddressFromHex, cfxAddressFromBase32)
	}
}

func fatalIfErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
