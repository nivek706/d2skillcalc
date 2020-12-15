package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nivek706/d2skillcalc/pkg/fileutil"

	"github.com/nivek706/d2skillcalc/internal/structs/skill"
)

func main() {

	//read the Skills.txt file, load into memory
	skillFile := fileutil.ReadFile("../assets/113c-data/Skills.txt")
	missileFile := fileutil.ReadFile("../assets/113c-data/Missiles.txt")

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
		// selectedSkill.PopulateMissileDamage(1, 5)
		selectedSkill.PrintSkillTable(1, 5)
		// selectedSkill.PrintSkillTable(11, 20)
		// selectedSkill.PrintSkillTable(21, 30)
		// selectedSkill.PrintSkillTable(31, 40)
	}
}
