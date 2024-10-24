package main

import (
	"fmt"
	"log"

	"github.com/alecthomas/kong"
	"inet.af/tcpproxy"
)

func main() {

	var cliCfg struct {
		ListenPort string `short:"l" help:"Port to listen on." required:""`
		RemoteAddr string `short:"r" help:"Remote address to forward to." required:""`
	}
	_ = kong.Parse(&cliCfg)

	var p tcpproxy.Proxy
	p.AddRoute(":"+cliCfg.ListenPort, tcpproxy.To(cliCfg.RemoteAddr))
	fmt.Printf("Forwarding from %s -> %s\n", cliCfg.ListenPort, cliCfg.RemoteAddr)

	log.Fatal(p.Run())
}
