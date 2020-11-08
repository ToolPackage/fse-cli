package fse

//
//import (
//	"github.com/ToolPackage/fse-cli/fse/parser"
//	"github.com/ToolPackage/fse/tx"
//	"net"
//)
//
//type CommandClient struct {
//	channel *tx.Channel
//}
//
//func NewClient() *CommandClient {
//	return &CommandClient{}
//}
//
//func (c *CommandClient) Execute(input string) {
//	ts := parser.Parse(input)
//	token, _ := ts.Next()
//	switch token {
//	case "login":
//		addr, ok := ts.Next()
//		if !ok {
//			_ = UI.ErrorOutput("error: login address token")
//			return
//		}
//		token, ok := ts.Next()
//		if !ok {
//			_ = UI.ErrorOutput("error: login address token")
//			return
//		}
//		_, ok = ts.Next()
//		if ok {
//			_ = UI.ErrorOutput("error: login address token")
//			return
//		}
//		c.login(addr, token)
//	default:
//		_ = UI.ErrorOutput("error: unknown command " + token)
//	}
//}
//
//func (c *CommandClient) login(addr string, token string) {
//	_ = UI.SuccessOutput("connected to " + addr)
//}
//
//func (c *CommandClient) connect(addr string) error {
//	conn, err := net.Dial("tcp", addr)
//	if err != nil {
//		return err
//	}
//
//	c.channel = tx.NewChannel(conn, conn)
//	return nil
//}
//
//func (c *CommandClient) Close() {
//
//}
