# d2skillcalc
Highly WIP - a calculator tool for Diablo II skills. 

Disclaimer: a lot of things don't work yet, or do not compute correctly. This is a non-exhaustive list of outstanding issues:
- Damage over time spells don't calculate correctly (ex. Poison Nova, Fire Wall, etc)
- Physical skills are not supported yet
- Synergies are not considered, even the prerequisite points that would exist to get to allocating a single point in a skill. For example, compare the slight difference in skill damage values for Bone Spear on [Arreat Summit](http://classic.battle.net/diablo2exp/skills/necromancer-poison.shtml) and the results of this tool. That difference is because the Arreat Summit page considers the 7% synergy bonus from a single required skill point in Teeth to get Bone Spear.
- Summon spells are not supported yet

Obviously, there is a long way to go until the tool is fully functional. Please be aware that it's also being used as a learning project, so the project probably won't follow proper golang standards for some time. Contributions, constructive issues, or other advice and discussion is always welcome!


## Usage
Built on golang 1.15.2  
Clone the repository, change to the directory `/d2skillcalc/internal`, run `go run main.go`.

Currently, the application simply takes the name of a skill, and will print out a skill table for levels 1 through 40. Much more to come in the future.
```
PS E:\code\d2skillcalc\internal> go run main.go
Enter skill name:Teeth
+-----------+-------+-------+-------+-------+-------+-------+--------+--------+---------+---------+
| LEVEL     | 1     | 2     | 3     | 4     | 5     | 6     | 7      | 8      | 9       | 10      |
+-----------+-------+-------+-------+-------+-------+-------+--------+--------+---------+---------+
| Mana Cost | 3.0   | 3.5   | 4.0   | 4.5   | 5.0   | 5.5   | 6.0    | 6.5    | 7.0     | 7.5     |
| Ele Dmg   | 2 - 4 | 3 - 5 | 4 - 6 | 5 - 7 | 6 - 8 | 7 - 9 | 8 - 10 | 9 - 11 | 10 - 12 | 11 - 14 |
+-----------+-------+-------+-------+-------+-------+-------+--------+--------+---------+---------+
+-----------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+
| LEVEL     | 11      | 12      | 13      | 14      | 15      | 16      | 17      | 18      | 19      | 20      |
+-----------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+
| Mana Cost | 8.0     | 8.5     | 9.0     | 9.5     | 10.0    | 10.5    | 11.0    | 11.5    | 12.0    | 12.5    |
| Ele Dmg   | 12 - 16 | 13 - 17 | 14 - 18 | 15 - 20 | 16 - 22 | 17 - 23 | 18 - 25 | 20 - 27 | 22 - 29 | 23 - 31 |
+-----------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+
+-----------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+
| LEVEL     | 21      | 22      | 23      | 24      | 25      | 26      | 27      | 28      | 29      | 30      |
+-----------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+
| Mana Cost | 13.0    | 13.5    | 14.0    | 14.5    | 15.0    | 15.5    | 16.0    | 16.5    | 17.0    | 17.5    |
| Ele Dmg   | 24 - 33 | 26 - 35 | 28 - 38 | 30 - 40 | 32 - 42 | 34 - 45 | 36 - 48 | 38 - 50 | 38 - 50 | 43 - 56 |
+-----------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+
+-----------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+
| LEVEL     | 31      | 32      | 33      | 34      | 35      | 36      | 37      | 38      | 39      | 40      |
+-----------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+
| Mana Cost | 18.0    | 18.5    | 19.0    | 19.5    | 20.0    | 20.5    | 21.0    | 21.5    | 22.0    | 22.5    |
| Ele Dmg   | 46 - 59 | 48 - 62 | 50 - 65 | 53 - 68 | 56 - 71 | 58 - 74 | 60 - 77 | 63 - 80 | 66 - 83 | 68 - 86 |
+-----------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+
Enter skill name:
```


## Credits
Some design decisions and a lot of techniques have been inspired by the d2modmaker repository, find it here: https://github.com/tlentz/d2modmaker