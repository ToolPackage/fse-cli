package internal

import (
	"fmt"
	ini "github.com/Luncert/go-ini"
	"github.com/ToolPackage/fse/tx"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

const (
	actionAuth        = "auth"
	configFileName    = ".fse-cli"
	curConnEnvVarName = "currentConnection"
)

var (
	configFilePath = getConfigFilePath()
)

type FseClient struct {
	cfg     *ini.Config
	channel *tx.Channel
}

func NewClient() *FseClient {
	return &FseClient{cfg: loadConfig()}
}

func (f *FseClient) login() error {
	connName := f.cfg.Section("common").Variable(curConnEnvVarName)
	if connName == nil || connName.Type() != ini.StringType || len(connName.String()) == 0 {
		return fmt.Errorf("current connection not set, use peek command to set one")
	}

	creds, err := lookupCredential(connName.String())
	if err != nil {
		return err
	}
	cred := creds[0]

	if err := f.connect(cred.UserName); err != nil {
		return err
	}

	// send auth packet
	f.channel.NewPacket(actionAuth).Body(string(cred.CredentialBlob)).Emit()
	res := f.channel.RecvPacket()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("auth failed: %s", string(res.Content))
	}

	return nil
}

func (f *FseClient) connect(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	f.channel = tx.NewChannel(conn, conn)
	return nil
}

func (f *FseClient) setConnection(connName string) {
	f.cfg.CreateIfAbsent("common").
		CreateIfAbsent(curConnEnvVarName).
		SetValue(ini.NewStringValue(connName))
	saveConfig(f.cfg)
}

func (f *FseClient) close() {

}

func loadConfig() *ini.Config {
	cfg, err := ini.ReadConfigFile(configFilePath)
	if err != nil {
		panic(err)
	}
	return cfg
}

func saveConfig(cfg *ini.Config) {
	if err := ini.WriteConfigFile(configFilePath, cfg); err != nil {
		panic(err)
	}
}

func getConfigFilePath() string {
	p := filepath.Join(getUserHomeDir(), configFileName)
	if _, err := os.Stat(p); err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}

		err = ioutil.WriteFile(p, []byte{}, 0644)
		if err != nil {
			panic(err)
		}
	}

	return p
}
