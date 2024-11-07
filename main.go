package main

import (
	"espcommander/commands"
	"flag"
	"fmt"
	"net"
	"strings"
)


var ADDRES = flag.String("i", "autodetect", "target ip addres")

func getip() (string, error){
	pc, err := net.ListenPacket("udp", ":40000")
	if err != nil {
		return "", err
	}
	defer pc.Close()
	buff := make([]byte, 32)
	addres := ""
	for {
		_, addr, err := pc.ReadFrom(buff)
		if err != nil {
			return "", err
		}
		if string(buff) == "Hello sir"{
			parts := strings.Split(addr.String(), ":")
			addres = parts[0]+"40001"
			break
		}
	}
	return addres, nil
}


func main() {
	flag.Parse()
	addres := ""
	var err error
	if *ADDRES == "autodetect" {
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
	
	defer conn.Close() 
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
	for k,_ := range nodes{
		fmt.Printf("enter args: (freq;time)\n")
		fmt.Scanf("%d;%d\n", &nodes[k].Freq, &nodes[k].Time)
	}
	com.Args = make([]interface{Encode() []byte}, len(nodes))
	for k := range nodes{
		com.Args[k] = nodes[k]
	}
	conn.Write(com.Encode())
}
