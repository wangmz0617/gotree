# Gotree: toolkit and api for phylogenetic tree manipulation

## Commands

### reroot

This command reroots a tree in two ways:
1. `gotree reroot outgroup` : Using an outgroup. If the outgroup is not monophyletic, 2 possibilities: 1) By default (`--strict=false`) it takes the lca of given tips to reroot the tree, and print a warning, 2) if `--strict` is given, it exits with an error.
2. `gotree reroot midpoint`: At midpoint.

#### Usage

General command
```
Usage:
  gotree reroot [command]

Available Commands:
  midpoint    Reroot trees at midpoint
  outgroup    Reroot trees using an outgroup

Flags:
  -i, --input string    Input Tree (default "stdin")
  -o, --output string   Rerooted output tree file (default "stdout")
```

#### Examples

* Reroot a random tree using an outgroup in a file

outgroup.txt
```
Tip2
Tip4
Tip7
```

```
gotree generate yuletree --seed 10 -o outtree1.nw
gotree reroot outgroup -l outgroup.txt -i outtree1.nw -o outtree2.nw
gotree draw svg -w 200 -H 200  -i outtree1.nw -o commands/reroot_1.svg
gotree draw svg -w 200 -H 200  -i outtree2.nw --with-branch-support --support-cutoff 0.5 -o commands/reroot_2.svg
```

Initial random Tree            | Rerooted Tree
-------------------------------|---------------------------------------
![Random Tree 1](reroot_1.svg) | ![Rerooted](reroot_2.svg)

* Reroot a random tree using an outgroup in commandline

```
gotree generate yuletree --seed 10 -o outtree1.nw
gotree reroot outgroup -i outtree1.nw -o outtree2.nw Tip2 Tip4 Tip7
gotree draw svg -w 200 -H 200  -i outtree1.nw -o commands/reroot_1.svg
gotree draw svg -w 200 -H 200  -i outtree2.nw --with-branch-support --support-cutoff 0.5 -o commands/reroot_2.svg
```


* Reroot a random tree using an outgroup in commandline, and remove the outgroup

```
gotree generate yuletree --seed 10 -o outtree1.nw
gotree reroot outgroup --remove-outgroup -i outtree1.nw -o outtree2.nw Tip2 Tip4 Tip7
gotree draw svg -w 200 -H 200  -i outtree1.nw -o commands/reroot_5.svg
gotree draw svg -w 200 -H 200  -i outtree2.nw --with-branch-support --support-cutoff 0.5 -o commands/reroot_6.svg
```

Initial random Tree            | Rerooted Tree without the outgroup
-------------------------------|---------------------------------------
![Random Tree 1](reroot_5.svg) | ![Rerooted](reroot_6.svg)


* Reroot a random tree at midpoint

```
gotree generate yuletree --seed 10 -o outtree1.nw
gotree reroot midpoint -i outtree1.nw -o outtree2.nw 
gotree draw svg -w 200 -H 200  -i outtree1.nw -o commands/reroot_3.svg
gotree draw svg -w 200 -H 200  -i outtree2.nw --with-branch-support --support-cutoff 0.5 -o commands/reroot_4.svg
```

Initial random Tree            | Rerooted Tree at Midpoint
-------------------------------|---------------------------------------
![Random Tree 1](reroot_3.svg) | ![Rerooted](reroot_4.svg)
