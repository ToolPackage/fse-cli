package internal

import (
	"github.com/ToolPackage/fse/tx"
	"net"
)

type CommandClient struct {
	channel *tx.Channel
}

func NewClient() *CommandClient {
	return &CommandClient{}
}

func (c *CommandClient) login(addr string, token string) {
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
