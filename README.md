# Roublard

![Roublard screen shot](./.meta/assets/sshot.png)

[roublard.low.webm](https://github.com/user-attachments/assets/6a3987a4-523d-4038-a618-0758c64e9afe)

[Better quality video](https://youtu.be/C0HtNDdppLU)

> Roublard is the french name of the Rogue class in DnD

Rogue-like game based on the [tutorial series from IdiotCoder/FatOldYeti](https://web.archive.org/web/20220818041908/https://www.fatoldyeti.com/posts/roguelike-tutorial-0/).

As Ebiten didn't work on my fresh setup of Ubuntu 24.04
(see [here](https://github.com/hajimehoshi/ebiten/issues/3304). (EDIT: It has been fixed since.)), I decided to 
use http://g3n.rocks since I'm following the project for a while and wanted to
play with it for years.

So here is my attempt to have a playable base for a Rogue-like game, in Go, with
g3n and ECS.

It includes assets from PolyHaven and @mz4250.


## How to run

- install [Go](https://go.dev): https://go.dev/doc/install
- install all the dependencies for [g3n](http://g3n.rocks): https://github.com/g3n/engine#dependencies
- run the project

```
git clone https://github.com/dolanor/roublard
cd roublard
go run .
```

## Compare with original

This repository olds the original upstream code from @fatoldyeti in a non-accessible branch by default (master).
It was used to keep my modification to use g3n instead of ebiten as small and separate as possible to have the minimalist difference between the code base.
First, as to see how much work is needed to have the 3D version of the game, and second, for me to check if I didn't implement some of the logic correctly (which happened at some point)

You can compare ebiten version versus g3n with

```
git diff v0.X.0 upstream-v0.0.X
```

Just replace X by the number of the tutorial on the blog https://web.archive.org/web/20220818041908/https://www.fatoldyeti.com/posts/.
