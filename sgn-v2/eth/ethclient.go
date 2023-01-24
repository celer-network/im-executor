package eth

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

// if ksfile is like awskms:us-west-2:alias/mytestkey, use KmsSigner
// passphrase will be awsKey:awsSec or if empty, will use aws auto search env variable etc
// otherwise normal ks json file based signer
const awskmsPre = "awskms"

// parse ksfile string if valid awskms, return region and alias. otherwise return empty strings
func ParseAwsKms(ksfile, passphrase string) (region, alias, awsKey, awsSec string) {
	if !strings.HasPrefix(ksfile, awskmsPre) {
		return "", "", "", ""
	}
	kmskeyinfo := strings.SplitN(ksfile, ":", 3)
	if len(kmskeyinfo) != 3 {
		return "", "", "", ""
	}
	awskeysec := []string{"", ""}
	if passphrase != "" {
		awskeysec = strings.SplitN(passphrase, ":", 2)
		if len(awskeysec) != 2 {
			// if awskeysec is invalid, return region and let aws scan local cre file seems better.
			return kmskeyinfo[1], kmskeyinfo[2], "", ""
		}
	}
	return kmskeyinfo[1], kmskeyinfo[2], awskeysec[0], awskeysec[1]
}

// return signer, address
func CreateSigner(ksfile, passphrase string, chainid *big.Int) (ethutils.Signer, Addr, error) {
	if strings.HasPrefix(ksfile, awskmsPre) {
		kmskeyinfo := strings.SplitN(ksfile, ":", 3)
		if len(kmskeyinfo) != 3 {
			return nil, ZeroAddr, fmt.Errorf("%s has wrong format", ksfile)
		}
		awskeysec := []string{"", ""}
		if passphrase != "" {
			awskeysec = strings.SplitN(passphrase, ":", 2)
			if len(awskeysec) != 2 {
				return nil, ZeroAddr, fmt.Errorf("%s has wrong format", passphrase)
			}
		}
		kmsSigner, err := ethutils.NewKmsSigner(kmskeyinfo[1], kmskeyinfo[2], awskeysec[0], awskeysec[1], chainid)
		if err != nil {
			return nil, ZeroAddr, err
		}
		return kmsSigner, kmsSigner.Addr, nil
	}
	ksBytes, err := ioutil.ReadFile(ksfile)
	if err != nil {
		return nil, ZeroAddr, err
	}
	key, err := keystore.DecryptKey(ksBytes, passphrase)
	if err != nil {
		return nil, ZeroAddr, err
	}
	signer, err := ethutils.NewSigner(hex.EncodeToString(crypto.FromECDSA(key.PrivateKey)), chainid)
	return signer, key.Address, err
}
