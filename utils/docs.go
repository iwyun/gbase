package utils

import (
	"github.com/iwyun/gbase/log"
	"os/exec"
)

func InitDocs(mainPath string) {
	//本地

	err := exec.Command("swag", "init", "-g", mainPath+"/main.go", "-o", mainPath+"/docs").Run()
	if err != nil {
		log.Errorln(err)
	}

}
