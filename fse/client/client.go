package client

import (
	"github.com/ToolPackage/fse/tx"
	"net"
)

type Client struct {
	channel *tx.Channel
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Connect(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	c.channel = tx.NewChannel(conn, conn)
	return nil
}

func (c *Client) Close() {

}
