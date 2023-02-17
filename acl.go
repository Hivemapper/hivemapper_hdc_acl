package hivemapper_hdc_acl

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/streamingfast/solana-go"
)

const AclFileName = "acl.data"
const AclSignatureFileName = "acl.signatureB58"

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
		return nil, solana.Signature{}, fmt.Errorf("opening acl signatureB58 aclFile: %s", err)
	}

	signatureData, err := io.ReadAll(signatureFile)
	if err != nil {
		return nil, solana.Signature{}, fmt.Errorf("reading acl signatureB58 aclFile: %s", err)
	}

	signature, err := solana.NewSignatureFromBytes(signatureData)
	if err != nil {
		return nil, solana.Signature{}, fmt.Errorf("creating signatureB58: %s", err)
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

func AclClearFromDevice(sourceFolder string) error {
	aclFile := path.Join(sourceFolder, AclFileName)
	if _, err := os.Stat(aclFile); err == nil {
		if err := os.Remove(aclFile); err != nil {
			return fmt.Errorf("removing acl file: %s", err)
		}
	}

	signatureFile := path.Join(sourceFolder, AclSignatureFileName)
	if _, err := os.Stat(signatureFile); err == nil {
		if err := os.Remove(signatureFile); err != nil {
			return fmt.Errorf("removing acl file: %s", err)
		}
	}

	return nil
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

func (a *Acl) legacyMessageToSign() ([]byte, error) {

	data, err := json.Marshal(a)
	if err != nil {
		return nil, fmt.Errorf("marshalling acl: %s", err)
	}

	return data, nil
}
func (a *Acl) messageToSign() ([]byte, error) {

	data, err := json.Marshal(a)
	if err != nil {
		return nil, fmt.Errorf("marshalling acl: %s", err)
	}

	h := md5.New()
	io.WriteString(h, string(data))

	hash := h.Sum(nil)
	hexHash := hex.EncodeToString(hash)

	message := fmt.Sprintf("Lock Camera with Access Control List with %d manager(s) and %d driver(s). Hash: %s", len(a.Managers), len(a.Drivers), hexHash)

	return []byte(message), nil
}

func (a *Acl) ValidateSignature(signature solana.Signature) bool {
	data, err := a.messageToSign()
	if err != nil {
		return false
	}
	valid := a.validateSignature(data, signature)
	if valid {
		return true
	}

	data, err = a.legacyMessageToSign()
	return a.validateSignature(data, signature)
}

func (a *Acl) validateSignature(data []byte, signature solana.Signature) bool {
	for _, managerAddress := range a.Managers {
		pubKey, err := solana.PublicKeyFromBase58(managerAddress)
		if err != nil {
			return false
		}
		if signature.Verify(pubKey, data) {
			return true
		}
	}
	return false
}
