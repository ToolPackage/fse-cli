package internal

import (
	"github.com/ToolPackage/fse/tx"
	"net"
)

type FseClient struct {
	channel *tx.Channel
}

func newClient() *FseClient {
	return &FseClient{}
}

func (c *FseClient) login(addr string, token string) {
}

func (c *FseClient) connect(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	c.channel = tx.NewChannel(conn, conn)
	return nil
}

func (c *FseClient) Close() {

}
