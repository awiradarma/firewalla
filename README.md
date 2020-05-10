# firewalla

## Compile to local (Mac OS X) 
* Compile
> go build -o native tcpdump_parser.go
* Check file format
> file native  
> native: Mach-O 64-bit executable x86_64
* try it!
> cat tcpdump.out | ./native | head -3  
> 192.168.1.18:61963 ( 02:01:e4:xx:xx:xx ) ==> 52.41.xxx.xxx:443 ( 10:da:43:xx:xx:xx ) 135 bytes  
> 192.168.1.24:59240 ( 02:01:e4:xx:xx:xx ) ==> 167.24.xxx.xxx:80 ( 10:da:43:xx:xx:xx ) 78 bytes  
> 192.168.1.7:56459 ( 02:01:e4:xx:xx:xx ) ==> 52.21.xxx.xxx:80 ( 10:da:43:xx:xx:xx ) 74 bytes  
  
## Compile to ARM64 - Linux for Firewalla device 
* Compile
> GOOS=linux GOARCH=arm64 go build -o firewalla_binary tcpdump_parser.go
* Check file format
> file firewalla_binary  
> firewalla_binary: ELF 64-bit LSB executable, ARM aarch64, version 1 (SYSV), statically linked, not stripped
