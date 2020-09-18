package process

import (
	"fmt"
	"github.com/adioss/free-tsa-servers-bench/utils"
	"log"
	"strings"
)

const keyAlias string = "adioss"
const keystore string = "adioss.p12"
const storePass string = "changeit"

func signAndCheck(outputJarName string, directory string, tsaServer string) {
	log.Printf("-------------------------------- With '%s' ------------------------------------------------", tsaServer)
	outputJarPath := directory + "/" + outputJarName

	cmd := fmt.Sprintf("jarsigner -tsa %s -keystore %s -storePass %s -keypass %s -signedjar %s %s %s",
		tsaServer, keystore, storePass, storePass, outputJarPath, directory+"/"+utils.JarName, keyAlias)
	commandOutput, err := utils.RunCommand(cmd, "")
	if err != nil {
		log.Println("################################################################")
		log.Printf("Signature failed for TSA sever '%s' with %v\n", tsaServer, err)
		log.Println("################################################################")
		return
	}
	log.Printf("Signature done? %v with '%s'\n", strings.Contains(commandOutput, "jar signed."), tsaServer)

	cmd = fmt.Sprintf("jarsigner -verify -certs %s", outputJarPath)
	commandOutput, _ = utils.RunCommand(cmd, "")
	log.Printf("Verification ok? %v with '%s'\n", strings.Contains(commandOutput, "jar verified."), tsaServer)
}
