package fse

import (
	"github.com/ToolPackage/fse/tx"
	"github.com/jroimartin/gocui"
	"net"
)

type Client struct {
	channel *tx.Channel
	gui     *gocui.Gui
}

func NewClient(gui *gocui.Gui) *Client {
	return &Client{gui: gui}
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
