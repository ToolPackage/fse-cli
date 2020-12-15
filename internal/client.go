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
	actionAuth     = "auth"
	configFileName = ".fse-cli"
)

type FseClient struct {
	cfg     *ini.Config
	channel *tx.Channel
}

func NewClient() *FseClient {
	return &FseClient{cfg: loadConfig()}
}

func (f *FseClient) login() error {
	credName := f.cfg.Section("common").Variable(CurConnEnvVarName)
	if credName == nil || credName.Value.Type() != ini.StringType || len(credName.String()) == 0 {
		return fmt.Errorf("current connection not set, use peek command to set one")
	}

	creds, err := lookupCredential(credName.String())
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

func (f *FseClient) Close() {

}

func loadConfig() *ini.Config {
	data, err := ioutil.ReadFile(getConfigFilePath())
	if err != nil {
		panic(err)
	}
	return ini.ParserIni(string(data))
}

func saveConfig() {

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
