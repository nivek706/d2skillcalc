package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nivek706/d2skillcalc/configs"
	"github.com/nivek706/d2skillcalc/pkg/fileutil"

	"github.com/nivek706/d2skillcalc/internal/structs/skill"
)

func main() {

	c, err := configs.LoadConfig(".")
	if err != nil {
		fmt.Println("fatal")
	}

	//read the Skills.txt file, load into memory
	skillFile := fileutil.ReadFile(fmt.Sprintf("%sSkills.txt", c.TxtDirPath))
	missileFile := fileutil.ReadFile(fmt.Sprintf("%sMissiles.txt", c.TxtDirPath))

	startInputLoop(skillFile, missileFile)
}

func startInputLoop(skillFile *fileutil.File, missileFile *fileutil.File) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter skill name:")
		skillName, _ := reader.ReadString('\n')
		skillName = strings.TrimRight(skillName, "\r\n")
		if skillName == "exit" {
			break
		}
		selectedSkill := skill.NewSkill(skillName, skillFile, missileFile)
		selectedSkill.PrintSkillTable(1, 10)
		selectedSkill.PrintSkillTable(11, 20)
		selectedSkill.PrintSkillTable(21, 30)
		selectedSkill.PrintSkillTable(31, 40)
	}
}
