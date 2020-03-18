package utils

import (
	"github.com/iwyun/gbase/log"
	"os/exec"
)

func InitDocs(mainPath string) {
	//todo check os
	err := exec.Command("swag", "init", "-g", mainPath+"/main.go", "-o", mainPath+"/docs").Run()
	if err != nil {
		log.Errorln(err)
	}

}
