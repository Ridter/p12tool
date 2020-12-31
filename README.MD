# P12tool
学习golang写的一个用来解析和爆破p12证书的小工具，代码写的有点戳，能用就行。

## Usage
```

██████╗  ██╗██████╗ ████████╗ ██████╗  ██████╗ ██╗
██╔══██╗███║╚════██╗╚══██╔══╝██╔═══██╗██╔═══██╗██║
██████╔╝╚██║ █████╔╝   ██║   ██║   ██║██║   ██║██║
██╔═══╝  ██║██╔═══╝    ██║   ██║   ██║██║   ██║██║
██║      ██║███████╗   ██║   ╚██████╔╝╚██████╔╝███████╗
╚═╝      ╚═╝╚══════╝   ╚═╝    ╚═════╝  ╚═════╝ ╚══════╝

Version: dev (n/a) - 12/31/20 - Evi1cg

NAME:
   p12tool - A tool to parse p12 cert file or bruteforce attacks against cert password

USAGE:
   main [global options] command [command options] [arguments...]

AUTHOR:
   Evi1cg

COMMANDS:
   parse    Parse p12 file and print cert info
   crack    Crack p12 file password.
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --cert value, -c value  The cert file you choice.
   --debug, -d             Debug mode. (default: false)
   --help, -h              show help (default: false)
```

## crack
```bash
go run cmd/main.go crack -c file.pfx -f passwords.txt
```
>可选: -t 指定线程，-d 开启debug模式，-o 将破解成功的密码输出至指定文件。

![](https://blogpics-1251691280.file.myqcloud.com/imgs/20201231232938.png)


## parse
```bash
go run cmd/main.go parse -c file.pfx -p password
```

![](https://blogpics-1251691280.file.myqcloud.com/imgs/20201231233107.png)

对证书进行解析，输出证书信息，顺便输出是否可对程序进行签名，免去手动验证的烦恼。 ~~

2020年推得最后一个小工具，祝大家新年快乐！！

## 参考
* 1、https://github.com/allyomalley/p12CrackerGo
* 2、https://golang.org/pkg/crypto/x509
* 3、https://godoc.org/golang.org/x/crypto/pkcs12