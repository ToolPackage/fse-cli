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

type CommandClient struct {
	cfg     *ini.Config
	channel *tx.Channel
}

func NewClient() *CommandClient {
	return &CommandClient{cfg: loadConfig()}
}

func (c *CommandClient) login() error {
	credName := c.cfg.Section("common").Variable(CurConnEnvVarName)
	if credName == nil || credName.Value.Type() != ini.StringType || len(credName.String()) == 0 {
		return fmt.Errorf("current connection not set, use peek command to set one")
	}

	creds, err := lookupCredential(credName.String())
	if err != nil {
		return err
	}
	cred := creds[0]

	if err := c.connect(cred.UserName); err != nil {
		return err
	}

	// send auth packet
	c.channel.NewPacket(actionAuth).Body(string(cred.CredentialBlob)).Emit()
	res := c.channel.RecvPacket()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("auth failed: %s", string(res.Content))
	}

	return nil
}

func (c *CommandClient) connect(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	c.channel = tx.NewChannel(conn, conn)
	return nil
}

func (c *CommandClient) Close() {

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
