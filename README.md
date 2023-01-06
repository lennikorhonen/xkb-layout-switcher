# Keyboard layout switcher
Small program to switch keyboard layout between two languages (fi and us atm)

Didn't bother to look if tiling window managers already handle this problem so I wrote this.

## Table of contents

* [Installing the program](#installing-the-program)
* [Build from source](#build-from-source)
* [Backlog](#backlog)
* [How to use example](#how-to-use-example)

## Installing the program
Download the binary from the releases page and save it to `/usr/bin` directory

You can also curl or wget the program from
```
curl https://github.com/lennikorhonen/xkb-layout-switcher/releases/download/latest-release-number/keyboard-layout-switcher-linux-amd64 /usr/bin/keyboard-layout-switcher
```
```
wget https://github.com/lennikorhonen/xkb-layout-switcher/releases/download/latest-release-number/keyboard-layout-switcher-linux-amd64 -O /usr/bin/keyboard-layout-switcher
```

and then run
```
chmod +x /usr/bin/keyboard-layout-switcher
```
to make it executable

## Build from source
You can also glone this git repo and build the program from source

You need to have `go` installed if you want to do this

```
git clone https://github.com/lennikorhonen/xkb-layout-switcher

cd xkb-layout-switcher

go build
```

and then move it to `/usr/bin`

## How to use example
For example in i3 config file add this line
```
bindsym $mod+space exec "keyboard-layout-switcher"
```

## Backlog

- [ ] More than two languages
- [ ] Support all languages
- [ ] Some test cases
