package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func splitPort(s string) (string, string) {
	x := strings.Split(s,".")
	x_ip_address := x[0]+"."+x[1]+"."+x[2]+"."+x[3]
	x_port := x[4]
	if (strings.IndexAny(x[4],":") > 0) {
		x_port = strings.TrimRight(x[4],":")
	} 
	return x_ip_address, x_port
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		//x_date := fields[0]
		//x_time := fields[1]
		x_mac_src := fields[2]
		x_length := strings.TrimRight(fields[9],":")
		x_mac_dst := strings.TrimRight(fields[4],",")
		//x_src := strings.Split(fields[10],".")
		//x_ip_src := x_src[0]+"."+x_src[1]+"."+x_src[2]+"."+x_src[3]
		//x_port_src := x_src[4]
		x_ip_src, x_port_src := splitPort(fields[10])

		//x_dst := strings.Split(fields[12],".")
		//x_ip_dst := x_dst[0]+"."+x_dst[1]+"."+x_dst[2]+"."+x_dst[3]
		//x_port_dst := strings.TrimRight(x_dst[4],":")
		x_ip_dst, x_port_dst := splitPort(fields[12])

    	fmt.Println(x_ip_src+":"+x_port_src+" ( "+x_mac_src+" ) ==> "+x_ip_dst+":"+x_port_dst+" ( "+x_mac_dst+" ) "+x_length+" bytes" )
	}
}

