# Keyboard layout switcher
Small program to switch keyboard layout between two languages (fi and us atm)

Didn't bother to look if tiling window managers already handle this problem so I wrote this.

## Table of contents

* [Installing the program](#installing-the-program)
* [Build from source](#build-from-source)
* [How to use example](#how-to-use-example)
* [Backlog](#backlog)

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
You can make a `settings.json` file and store it in `$XDG_CONFIG_HOME/kb_layout_switcher/settings.json`

The contents of `settings.json` should look like this:
```json
{
    "langs": ["fi", "us"]
}
```
Then for example in i3 config file add this line
```
bindsym $mod+space exec "keyboard-layout-switcher"
```

## Backlog

- [x] More than two languages
- [x] Support all languages
- [ ] Some test cases
