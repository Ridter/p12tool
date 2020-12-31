package work

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
	"os"
	"p12tool/util"
	"p12tool/vars"
	"strconv"
)

func P12FileBrute(ctx *cli.Context) (err error)  {
	Parse(ctx)
	vars.Logger = util.NewLogger(vars.DebugMode, "")
	p12Bytes, err := ioutil.ReadFile(vars.Cert)
	if err != nil {
		vars.Logger.Log.Error("[-] Please input cert file")
		return nil
	}
	targetFile, err := os.Open(vars.File)
	if err != nil {
		vars.Logger.Log.Error("[-] Please input pass list file")
		return nil
	}
	scanner := bufio.NewScanner(targetFile)
	scanner.Split(bufio.ScanLines)
	vars.Logger.Log.Info("[*] Brute forcing...")
	crack(scanner, p12Bytes, vars.Threads)
	if vars.CrackedPassword !=""{
		success := fmt.Sprintf("[+] Password found ==> %s", vars.CrackedPassword)
		vars.Logger.Log.Noticef(success)
		if vars.OutFile != ""{
			fi, err := os.Create(vars.OutFile)
			if err != nil{
				vars.Logger.Log.Errorf("[!] Can't create file %s",vars.OutFile)
			}
			defer fi.Close()
			_, err2 := fi.WriteString(success)
			if err2 != nil {
				vars.Logger.Log.Errorf("[!] Write file error")
			}
		}
		vars.Logger.Log.Infof("[*] Successfully cracked password after " + strconv.Itoa(vars.Attempts) + " attempts!")
	}
	return err
}

func crack(scanner *bufio.Scanner, p12Bytes []byte, threads int) {
	vars.Logger.Log.Infof("[*] Start thread num %d",vars.Threads)
	semaphore := make(chan bool, threads)
	lineNo := 0
	for scanner.Scan() {
		lineNo = lineNo + 1
		semaphore <- true

		go func(password string, line int) {
			decrypted := checkPass(p12Bytes, password)
			if decrypted != "" {
				vars.ResultsLock.Lock()
				vars.Attempts = line
				vars.CrackedPassword = password
				vars.ResultsLock.Unlock()
			}
			<-semaphore
		}(scanner.Text(), lineNo)
	}

	for i := 0; i < cap(semaphore); i++ {
		semaphore <- true
	}
}
func checkPass(p12Bytes []byte, password string) string {
	_,  err := pkcs12.ToPEM(p12Bytes, password)
	if err != nil{
		if err == pkcs12.ErrIncorrectPassword{
			if vars.DebugMode{
				vars.Logger.Log.Debugf("[-] Password [%s] incorrect",password)
			}
		}else{
			vars.Logger.Log.Error("[-] Check your file is P12 cert file !!")
			os.Exit(1)
		}

	}else{
		return password
	}
	return ""
}