package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

const filePrefix = "provider_cmd_"
const fileSuffix = ".go"
const packageCmdPath = "cmd"

func main() {
	provider := os.Args[1]
	provider = strings.Replace(provider,"{","",-1)
	provider = strings.Replace(provider,"}","",-1)
	log.Println("Build terraformer with " + provider + " provider...")
	files, err := ioutil.ReadDir(packageCmdPath)
	if err != nil {
		log.Println(err)
	}
	deletedProvider := []string{}
	for _, f := range files {
		if strings.HasPrefix(f.Name(), filePrefix) {
			if !strings.HasPrefix(f.Name(), filePrefix+provider+fileSuffix) {
				providerName := strings.Replace(f.Name(), filePrefix, "", -1)
				providerName = strings.Replace(providerName, fileSuffix, "", -1)
				deletedProvider = append(deletedProvider, providerName)
			}
		}
	}
	// move files for deleted providers
	os.MkdirAll(packageCmdPath+"/tmp", os.ModePerm)
	for _, provider := range deletedProvider {
		err := os.Rename(packageCmdPath+"/"+filePrefix+provider+fileSuffix, packageCmdPath+"/tmp/"+filePrefix+provider+fileSuffix)
		if err != nil {
			log.Println(err)
		}
	}

	// comment deleted providers in code
	rootCode, err := ioutil.ReadFile(packageCmdPath + "/root.go")
	lines := strings.Split(string(rootCode), "\n")
	newRootCodeLines := make([]string, len(lines))
	for i, line := range lines {
		for _, provider := range deletedProvider {
			if strings.Contains(strings.ToLower(line), "newcmd"+provider+"importer") {
				line = "// " + line
			}
			if strings.Contains(strings.ToLower(line), "new"+provider+"provider") {
				line = "// " + line
			}
		}
		newRootCodeLines[i] = line
	}
	newRootCode := strings.Join(newRootCodeLines, "\n")
	ioutil.WriteFile(packageCmdPath+"/root.go", []byte(newRootCode), os.ModePerm)

	// build....
	cmd := exec.Command("go", "build", "-o", "terraformer-"+provider)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err = cmd.Run()
	if err != nil {
		log.Fatal("err:", errb.String())
	}
	fmt.Println(outb.String())

	//revert code and files
	ioutil.WriteFile(packageCmdPath+"/root.go", []byte(rootCode), os.ModePerm)
	for _, provider := range deletedProvider {
		err := os.Rename(packageCmdPath+"/tmp/"+filePrefix+provider+fileSuffix, "cmd/"+filePrefix+provider+fileSuffix)
		if err != nil {
			log.Println(err)
		}
	}
}
