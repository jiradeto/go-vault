package secret

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"vault/encrypt"
)

type Vault struct {
	VaultKey  string
	KeyValues map[string]string
	FilePath  string
	mutex     sync.Mutex
}

func (v *Vault) loadKeyValues() error {
	f, err := os.Open(v.FilePath)
	if err != nil {
		v.KeyValues = make(map[string]string)
		return nil
	}
	defer f.Close()
	var sb strings.Builder
	_, err = io.Copy(&sb, f)
	if err != nil {
		return err
	}
	decryptedJson, err := encrypt.Decrypt(v.VaultKey, sb.String())
	if err != nil {
		return err
	}

	r := strings.NewReader(decryptedJson)
	dec := json.NewDecoder(r)
	err = dec.Decode(&v.KeyValues)
	if err != nil {
		return err
	}
	return nil
}

func (v *Vault) saveKeyValues() error {
	var sb strings.Builder
	enc := json.NewEncoder(&sb)
	err := enc.Encode(v.KeyValues)
	if err != nil {
		return err
	}
	encryptedJSON, err := encrypt.Encrypt(v.VaultKey, sb.String())
	if err != nil {
		return err
	}

	f, err := os.OpenFile(v.FilePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = fmt.Fprint(f, encryptedJSON)
	if err != nil {
		return err
	}
	return nil
}

func NewVault(vaultKey string, filePath string) *Vault {
	return &Vault{
		VaultKey: vaultKey,
		FilePath: filePath,
	}
}

func (v *Vault) Get(key string) (string, error) {
	v.loadKeyValues()
	val, exist := v.KeyValues[key]
	if !exist {
		return "nil", errors.New("key not exist")
	}
	return val, nil
}

func (v *Vault) Set(key, val string) error {
	v.loadKeyValues()
	v.mutex.Lock()
	defer v.mutex.Unlock()
	v.KeyValues[key] = val
	fmt.Print(v.KeyValues)
	err := v.saveKeyValues()
	if err != nil {
		return err
	}
	return nil
}
