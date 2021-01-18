package cfxaddress

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// CfxAddress ...
type CfxAddress struct {
	NetworkType
	AddressType
	Body
	Checksum
}

func (c CfxAddress) String() string {
	return fmt.Sprintf("%v:%v%v", c.NetworkType, c.Body, c.Checksum)
}

// VerboseString ...
func (c CfxAddress) VerboseString() string {
	return fmt.Sprintf("%v:%v:%v%v", c.NetworkType, c.AddressType, c.Body, c.Checksum)
}

// NewCfxAddressByBase32String ...
func NewCfxAddressByBase32String(base32Str string) (cfxAddress CfxAddress, err error) {
	if strings.ToLower(base32Str) != base32Str || strings.ToUpper(base32Str) != base32Str {
		err = errors.Errorf("not support base32 string with mix lowercase and uppercase %v", base32Str)
		return
	}
	base32Str = strings.ToLower(base32Str)

	parts := strings.Split(base32Str, ":")
	if len(parts) < 2 || len(parts) > 3 {
		err = errors.New("invalid format")
	}

	cfxAddress.NetworkType, err = NewNetowrkType(parts[0])
	if err != nil {
		return
	}

	bodyWithChecksum := parts[len(parts)-1]
	if len(bodyWithChecksum) < 8 {
		err = errors.New("invalid length")
		return
	}
	bodyStr := bodyWithChecksum[0 : len(bodyWithChecksum)-8]

	cfxAddress.Body, err = NewBodyByString(bodyStr)
	if err != nil {
		return
	}

	_, hexAddress, err := cfxAddress.Body.Encode()
	if err != nil {
		return
	}

	cfxAddress.AddressType, err = CalcAddressType(hexAddress)
	if len(parts) == 3 && strings.ToLower(parts[1]) != cfxAddress.AddressType.String() {
		err = errors.Errorf("invalid address type, expect %v actual %v", cfxAddress.AddressType, parts[1])
		return
	}

	cfxAddress.Checksum, err = CalcChecksum(cfxAddress.NetworkType, cfxAddress.Body)
	if err != nil {
		return
	}

	expectChk := cfxAddress.Checksum.String()
	actualChk := bodyWithChecksum[len(bodyWithChecksum)-8:]
	if expectChk != actualChk {
		err = errors.Errorf("invalid checksum, expect %v actual %v", expectChk, actualChk)
	}
	return
}

// NewCfxAddressByHexAddress encode hex address with networkID to base32 address according to CIP37
// INPUT: an addr (20-byte conflux-hex-address), a network-id (4 bytes)
// OUTPUT: a conflux-base32-address
func NewCfxAddressByHexAddress(hexAddress []byte, networkID uint32) (CfxAddress, error) {
	val := CfxAddress{}
	var err error
	val.NetworkType.EncodeFromID(networkID)
	val.AddressType, err = CalcAddressType(hexAddress)
	if err != nil {
		return val, errors.Wrap(err, "failed to calculate address type")
	}
	versionType, err := CalcVersionType(hexAddress)
	if err != nil {
		return val, errors.Wrap(err, "failed to calculate version type")
	}
	err = val.Body.Decode(versionType, hexAddress)
	if err != nil {
		return val, errors.Wrap(err, "failed to decode to body")
	}
	val.Checksum, err = CalcChecksum(val.NetworkType, val.Body)
	if err != nil {
		return val, errors.Wrap(err, "failed to calc checksum")
	}
	return val, nil
}

func (c CfxAddress) Encode() (hexAddress []byte, networkID uint32, err error) {
	// verify checksum
	var actualChecksum Checksum
	actualChecksum, err = CalcChecksum(c.NetworkType, c.Body)
	if err != nil {
		return
	}
	if actualChecksum != c.Checksum {
		err = errors.New("invalid checksum")
		return
	}

	_, hexAddress, err = c.Body.Encode()
	if err != nil {
		return
	}

	networkID, err = c.NetworkType.Decode()
	if err != nil {
		return
	}
	return
}
