package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IP's e Nomes de Servidores na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IP's de endereços na Internet",
			Flags:  flags,
			Action: buscarIPs,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na Internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app
}

func buscarIPs(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for seq, ip := range ips {
		fmt.Printf("(%v) %v\n", seq, ip)
	}

}

func buscarServidores(c *cli.Context) {
	host := c.String("host")
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for seq, servidor := range servidores {
		fmt.Printf("(%v) %v\n", seq, servidor.Host)
	}

}
