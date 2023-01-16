package hivemapper_hdc_acl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/streamingfast/solana-go"
)

const AclFileName = "acl.data"
const AclSignatureFileName = "acl.signature"

type Acl struct {
	Managers []string `json:"managers"`
	Drivers  []string `json:"drivers"`
}

func NewAclFromFile(sourceFolder string) (*Acl, solana.Signature, error) {
	aclFile, err := os.Open(path.Join(sourceFolder, AclFileName))
	if err != nil {
		return nil, solana.Signature{}, fmt.Errorf("opening acl aclFile: %s", err)
	}

	alData, err := io.ReadAll(aclFile)
	if err != nil {
		return nil, solana.Signature{}, fmt.Errorf("reading acl aclFile: %s", err)
	}

	acl, err := NewAclFromData(alData)
	if err != nil {
		return nil, solana.Signature{}, fmt.Errorf("creating acl: %s", err)
	}

	signatureFile, err := os.Open(path.Join(sourceFolder, AclSignatureFileName))
	if err != nil {
		return nil, solana.Signature{}, fmt.Errorf("opening acl signature aclFile: %s", err)
	}

	signatureData, err := io.ReadAll(signatureFile)
	if err != nil {
		return nil, solana.Signature{}, fmt.Errorf("reading acl signature aclFile: %s", err)
	}

	signature, err := solana.NewSignatureFromBytes(signatureData)
	if err != nil {
		return nil, solana.Signature{}, fmt.Errorf("creating signature: %s", err)
	}

	return acl, signature, nil
}

func NewAclFromData(data []byte) (*Acl, error) {
	var acl *Acl

	err := json.Unmarshal(data, &acl)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling acl data: %s", err)
	}

	return acl, nil
}

func AclExistOnDevice(sourceFolder string) bool {
	if _, err := os.Stat(path.Join(sourceFolder, AclFileName)); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func (a *Acl) Store(destinationFolder string, signature solana.Signature) error {
	data, err := json.Marshal(a)
	if err != nil {
		return fmt.Errorf("marshalling acl: %s", err)
	}

	aclFile := path.Join(destinationFolder, AclFileName)
	err = os.WriteFile(aclFile, data, 0644)
	if err != nil {
		return fmt.Errorf("writing acl file: %s", err)
	}

	signFile := path.Join(destinationFolder, AclSignatureFileName)
	err = os.WriteFile(signFile, signature.ToSlice(), 0644)
	if err != nil {
		return fmt.Errorf("writing acl file: %s", err)
	}

	return nil
}
