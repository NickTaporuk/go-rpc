//go:build !windows
// +build darwin

package rpc

func RpcParse(resp []byte) error {
	return nil
}
