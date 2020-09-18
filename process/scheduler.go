package process

import (
	"fmt"
	"github.com/adioss/free-tsa-servers-bench/generator"
	"github.com/adioss/free-tsa-servers-bench/utils"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)

func Schedule(scheduling string) {
	log.Print("Create cron")
	cronSpec := "@every" + " " + scheduling
	c := cron.New()
	c.AddFunc(cronSpec, func() {
		tempDirectory := prepareTempFiles()

		for _, tsaServer := range utils.TsaServers {
			outputJarName := fmt.Sprintf("signed_%s.jar", utils.SanitizeServerUrl(tsaServer))
			go signAndCheck(outputJarName, tempDirectory, tsaServer)
		}
		defer os.Remove(tempDirectory)
	})
	c.Start()

	// terminate after 60 days
	time.Sleep(60 * 24 * time.Hour)
	c.Stop()
}

func Loop(tsaServer string) {
	tempDirectory := prepareTempFiles()
	for true {
		outputJarName := fmt.Sprintf("signed_%s.jar", utils.SanitizeServerUrl(tsaServer))
		signAndCheck(outputJarName, tempDirectory, tsaServer)
		os.Remove(tempDirectory)
	}
}

func prepareTempFiles() string {
	uuid, _ := utils.NewUUID()
	tempDirectory := os.TempDir() + "/" + uuid
	os.Mkdir(tempDirectory, 0777)

	// create unique java file
	generator.GenerateUniqueJavaFile(tempDirectory)
	// compile and create jar
	generator.GenerateJarFile(tempDirectory)
	return tempDirectory
}
