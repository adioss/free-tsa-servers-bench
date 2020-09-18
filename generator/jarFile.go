package generator

import (
	"fmt"
	"github.com/adioss/free-tsa-servers-bench/utils"
)

func GenerateJarFile(directory string) {
	utils.RunCommand("javac Generated.java", directory)
	utils.RunCommand(fmt.Sprintf("jar cvfe %s Generated Generated.class", utils.JarName), directory)
}
