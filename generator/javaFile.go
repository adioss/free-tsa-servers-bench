package generator

import (
	"github.com/adioss/free-tsa-servers-bench/utils"
	"io/ioutil"
	"log"
	"strings"
)

const substr = "TO_REPLACE"
const patternFileName = "Generated_pattern.java"
const GeneratedJavaFileName = "Generated.java"

func GenerateUniqueJavaFile(directory string) {
	input, err := ioutil.ReadFile(patternFileName)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {

		if strings.Contains(line, substr) {
			uuid, _ := utils.NewUUID()
			lines[i] = strings.ReplaceAll(line, substr, uuid)
		}
	}
	output := strings.Join(lines, "\n")

	err = ioutil.WriteFile(directory+"/"+GeneratedJavaFileName, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
