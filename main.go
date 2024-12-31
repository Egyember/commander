package main

import (
	"espcommander/commands"
	"espcommander/webgui"
	"flag"
	"fmt"
	"net"
	"slices"
	"strings"
)


var ADDRES = flag.String("i", "autodetect", "target ip addres")
var cli = flag.Bool("c", false, "use cli interface")

func getip() (string, error){
	pc, err := net.ListenPacket("udp", ":40000")
	if err != nil {
		return "", err
	}
	defer pc.Close()
	addres := ""
	for {
		buff := make([]byte, 10)
		_, addr, err := pc.ReadFrom(buff)
		if err != nil {
			return "", err
		}
		if slices.Compare(buff, []byte{104, 101, 108, 108, 111, 32, 115, 105, 114, 0}) == 0{
			parts := strings.Split(addr.String(), ":")
			addres = parts[0]+":40001"
			break
		}
	}
	return addres, nil
}

func ConecttoDev() net.Conn{
	addres := ""
	var err error
	if *ADDRES == "autodetect" {
		fmt.Println("searching")
		addres, err = getip()
		if err != nil {
			panic(err)
		}
	} else{
		addres = *ADDRES
	}
	fmt.Println("connecting to: ",addres)
	conn, err := net.Dial("tcp", addres)
	/* to file
	conn, err := os.Create("./output.bin")
	*/
	if err != nil {
		panic(err)
	}

	return conn
}

func main() {
	flag.Parse()
	fmt.Println("connecting")
	conn := ConecttoDev()
	defer conn.Close() 
	if *cli {
		fmt.Println("cli")
		com := commands.Command{
			Command: commands.PLAY,
		}
		argnum := 0
		fmt.Printf("enter argnum:")
		fmt.Scanf("%d", &argnum)
		if argnum<=0{
			return
		}
		nodes := make([]commands.Node, argnum)
		for k := range nodes{
			fmt.Printf("enter args: (freq;time)\n")
			fmt.Scanf("%d;%d\n", &nodes[k].Freq, &nodes[k].Time)
		}
		com.Args = make([]interface{Encode() []byte}, len(nodes))
		for k := range nodes{
			com.Args[k] = nodes[k]
		}
		conn.Write(com.Encode())
	} else {
		webgui.Start(conn)
	}
}
