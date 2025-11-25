# csd2goalgen

**csd2goalgen** is a reimplementation of my [Keymaster's Keep module](https://github.com/mathstar/KeymastersKeepGames) 
for the game *Cook, Serve, Delicious! 2!!* as a standalone goal generator.

Goals are generated for both Chef 4 Hire and Cook, Serve, Delicious modes
taking the following parameters:

* **-n --num-goals** Number of goals to generate (default: 1)
* **-y --yum** Max yum level that will be required (default: 50)
* **-l --include-locked** Whether to include locked foods which are only available as random unlocks (default: false)
* **-s --stress** Whether to include stress mode goals (default: false)
* **-d --display-yum** Whether to display yum level requirements (default: false)
