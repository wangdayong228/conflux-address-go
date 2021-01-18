package cfxaddress

import (
	"encoding/hex"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

/*
[OPTIONAL] Address-type:
    match addr[0] & 0xf0
        case b00000000: "type=builtin"
        case b00010000: "type=user"
        case b10000000: "type=contract"
Implementations can choose to use "type=null" for the null address (0x0000000000000000000000000000000000000000).
*/

type AddressType string

const (
	BUILTIN  AddressType = "builtin"
	USER     AddressType = "user"
	CONTRACT AddressType = "contract"
	NULL     AddressType = "null"
)

const (
	ErrInvalidAddressType string = "Invalid address type"
)

func CalcAddressType(hexAddress []byte) (AddressType, error) {
	nullAddr, err := hex.DecodeString("0000000000000000000000000000000000000000")
	if err != nil {
		return "", err
	}
	if reflect.DeepEqual(nullAddr, hexAddress) {
		return NULL, nil
	}
	var addressType AddressType
	addressType.Decode(hexAddress[0])
	return addressType, nil
}

func (a AddressType) Encode() (byte, error) {
	switch a {
	case NULL:
		return 0x00, nil
	case BUILTIN:
		return 0x00, nil
	case USER:
		return 0x10, nil
	case CONTRACT:
		return 0x80, nil
	}
	return 0, errors.Errorf("Invalid address type %v", a)
}

func (a *AddressType) Decode(b byte) error {
	var addressType AddressType
	switch b & 0xf0 {
	case 0x00:
		addressType = BUILTIN
	case 0x10:
		addressType = USER
	case 0x80:
		addressType = CONTRACT
	default:
		return errors.Errorf("Invalid address type byte %v", b)
	}
	*a = addressType
	return nil
}

func (a AddressType) String() string {
	return fmt.Sprintf("type.%v", string(a))
}
