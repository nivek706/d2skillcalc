# d2skillcalc
WIP - a calculator tool for Diablo II skills. 

Disclaimer: a lot of things don't work yet, or do not compute correctly. This is a non-exhaustive list of outstanding issues:
- Damage over time spells don't always calculate correctly (ex. Poison Nova, Fire Wall, etc)
- Physical skills are not supported yet
- Synergies are not considered, even the prerequisite points that would exist to get to allocating a single point in a skill. For example, compare the slight difference in skill damage values for Bone Spear on [Arreat Summit](http://classic.battle.net/diablo2exp/skills/necromancer-poison.shtml) and the results of this tool. That difference is because the Arreat Summit page considers the 7% synergy bonus from a single required skill point in Teeth to get Bone Spear.
- Summon spells are not supported yet

Obviously, there is a long way to go until the tool is fully functional. Please be aware that it's also being used as a learning project, so the project probably won't follow proper golang standards for some time. Contributions, constructive issues, or other advice and discussion is always welcome!


## Usage
Built on golang 1.15.2  
Clone the repository, change to the directory `/d2skillcalc/cmd`, run `go run main.go`.

Currently, the application simply takes the name of a skill, and will print out a skill table for levels 1 through 40. Much more to come in the future.
```
..\code\d2skillcalc\cmd> go run main.go
Enter skill name:Claws of Thunder
+--------------------+------------+------------+------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+
| LEVEL              | 1          | 2          | 3          | 4           | 5           | 6           | 7           | 8           | 9           | 10          |
+--------------------+------------+------------+------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+
| Mana Cost          | 4.0        | 4.0        | 4.0        | 4.0         | 4.0         | 4.0         | 4.0         | 4.0         | 4.0         | 4.0         |
| Lightning Dmg      | 1 - 80     | 1 - 100    | 1 - 120    | 1 - 140     | 1 - 160     | 1 - 180     | 1 - 200     | 1 - 220     | 1 - 260     | 1 - 300     |
| clawsofthundernova | 1.0 - 30.0 | 1.0 - 45.0 | 1.0 - 60.0 | 1.0 - 75.0  | 1.0 - 90.0  | 1.0 - 105.0 | 1.0 - 120.0 | 1.0 - 135.0 | 1.0 - 160.0 | 1.0 - 185.0 |
| clawsofthunderbolt | 1.0 - 40.0 | 1.0 - 60.0 | 1.0 - 80.0 | 1.0 - 100.0 | 1.0 - 120.0 | 1.0 - 140.0 | 1.0 - 160.0 | 1.0 - 180.0 | 1.0 - 220.0 | 1.0 - 260.0 |
+--------------------+------------+------------+------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+
+--------------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+
| LEVEL              | 11          | 12          | 13          | 14          | 15          | 16          | 17          | 18          | 19          | 20          |
+--------------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+
| Mana Cost          | 4.0         | 4.0         | 4.0         | 4.0         | 4.0         | 4.0         | 4.0         | 4.0         | 4.0         | 4.0         |
| Lightning Dmg      | 1 - 340     | 1 - 380     | 1 - 420     | 1 - 460     | 1 - 500     | 1 - 540     | 1 - 600     | 1 - 660     | 1 - 720     | 1 - 780     |
| clawsofthundernova | 1.0 - 210.0 | 1.0 - 235.0 | 1.0 - 260.0 | 1.0 - 285.0 | 1.0 - 310.0 | 1.0 - 335.0 | 1.0 - 370.0 | 1.0 - 405.0 | 1.0 - 440.0 | 1.0 - 475.0 |
| clawsofthunderbolt | 1.0 - 300.0 | 1.0 - 340.0 | 1.0 - 380.0 | 1.0 - 420.0 | 1.0 - 460.0 | 1.0 - 500.0 | 1.0 - 560.0 | 1.0 - 620.0 | 1.0 - 680.0 | 1.0 - 740.0 |
+--------------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+
+--------------------+-------------+-------------+-------------+--------------+--------------+--------------+--------------+--------------+--------------+--------------+
| LEVEL              | 21          | 22          | 23          | 24           | 25           | 26           | 27           | 28           | 29           | 30           |
+--------------------+-------------+-------------+-------------+--------------+--------------+--------------+--------------+--------------+--------------+--------------+
| Mana Cost          | 4.0         | 4.0         | 4.0         | 4.0          | 4.0          | 4.0          | 4.0          | 4.0          | 4.0          | 4.0          |
| Lightning Dmg      | 1 - 840     | 1 - 900     | 1 - 980     | 1 - 1060     | 1 - 1140     | 1 - 1220     | 1 - 1300     | 1 - 1380     | 1 - 1380     | 1 - 1580     |
| clawsofthundernova | 1.0 - 510.0 | 1.0 - 545.0 | 1.0 - 590.0 | 1.0 - 635.0  | 1.0 - 680.0  | 1.0 - 725.0  | 1.0 - 770.0  | 1.0 - 815.0  | 1.0 - 815.0  | 1.0 - 945.0  |
| clawsofthunderbolt | 1.0 - 800.0 | 1.0 - 860.0 | 1.0 - 940.0 | 1.0 - 1020.0 | 1.0 - 1100.0 | 1.0 - 1180.0 | 1.0 - 1260.0 | 1.0 - 1340.0 | 1.0 - 1340.0 | 1.0 - 1540.0 |
+--------------------+-------------+-------------+-------------+--------------+--------------+--------------+--------------+--------------+--------------+--------------+
+--------------------+--------------+--------------+--------------+--------------+--------------+--------------+--------------+--------------+--------------+--------------+
| LEVEL              | 31           | 32           | 33           | 34           | 35           | 36           | 37           | 38           | 39           | 40           |
+--------------------+--------------+--------------+--------------+--------------+--------------+--------------+--------------+--------------+--------------+--------------+
| Mana Cost          | 4.0          | 4.0          | 4.0          | 4.0          | 4.0          | 4.0          | 4.0          | 4.0          | 4.0          | 4.0          |
| Lightning Dmg      | 1 - 1680     | 1 - 1780     | 1 - 1880     | 1 - 1980     | 1 - 2080     | 1 - 2180     | 1 - 2280     | 1 - 2380     | 1 - 2480     | 1 - 2580     |
| clawsofthundernova | 1.0 - 1010.0 | 1.0 - 1075.0 | 1.0 - 1140.0 | 1.0 - 1205.0 | 1.0 - 1270.0 | 1.0 - 1335.0 | 1.0 - 1400.0 | 1.0 - 1465.0 | 1.0 - 1530.0 | 1.0 - 1595.0 |
| clawsofthunderbolt | 1.0 - 1640.0 | 1.0 - 1740.0 | 1.0 - 1840.0 | 1.0 - 1940.0 | 1.0 - 2040.0 | 1.0 - 2140.0 | 1.0 - 2240.0 | 1.0 - 2340.0 | 1.0 - 2440.0 | 1.0 - 2540.0 |
+--------------------+--------------+--------------+--------------+--------------+--------------+--------------+--------------+--------------+--------------+--------------+
Enter skill name:
```


## Credits
Some design decisions and a lot of techniques have been inspired by the d2modmaker repository, find it here: https://github.com/tlentz/d2modmaker