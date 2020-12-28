package common

import (
	"github.com/nivek706/d2skillcalc/internal/structs/damage"
)

func RemoveDuplicateDamageRows(damageArray [][]damage.Damage) [][]damage.Damage {
	tempDmgArray := make([][]damage.Damage, 0)

	for _, row := range damageArray {
		//check if tempDmgArray contains row
		//if not, add row to tempDmgArray
		//if yes, move to next iteration
		if checkDmgArrayContainsRow(tempDmgArray, row) == false {
			tempDmgArray = append(tempDmgArray, row)
		}
	}

	return tempDmgArray

}

func checkDmgArrayContainsRow(dmgArray [][]damage.Damage, rowToCheck []damage.Damage) bool {
	containsRow := false
	for _, row := range dmgArray {
		if len(row) == len(rowToCheck) {
			rowsMatch := true
			for i, value := range row {
				if value != rowToCheck[i] {
					rowsMatch = false
				}

			}
			if rowsMatch == true {
				containsRow = true
			}
		}

	}

	return containsRow
}
