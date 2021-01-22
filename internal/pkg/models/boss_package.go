package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type BossPackage struct {
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	Version      string            `json:"version"`
	Homepage     string            `json:"homepage"`
	MainSrc      string            `json:"mainsrc"`
	Projects     []string          `json:"projects"`
	Scripts      map[string]string `json:"scripts,omitempty"`
	Dependencies map[string]string `json:"dependencies"`
}

func MakeBossPackage() *BossPackage {
	return &BossPackage{
		Scripts:      make(map[string]string),
		Dependencies: make(map[string]string),
		Projects:     []string{},
	}
}

func LoadFile(bossPath string) (*BossPackage, error) {
	buf, err := ioutil.ReadFile(bossPath)
	if err != nil {
		return nil, err
	}

	var bossPackage = MakeBossPackage()

	err = json.Unmarshal(buf, &bossPackage)
	if err != nil {
		return nil, err
	} else {
		return bossPackage, nil
	}
}

func (b *BossPackage) SaveToFile(bossPath string) error {
	buf, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(bossPath, buf, os.ModePerm)
}
