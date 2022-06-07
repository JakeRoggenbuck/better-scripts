# wallpaperinit
Go implementation of old 'wallpaperinit'

# Notes
Wallpaperinit -random will write the name of a random wallpaper to the file `~/.config/wallpaper/current`. This is what -set reads as it sets the wallpaper.

# Usage
```sh
# Pick and set new random wallpaper
wallpaperinit -random -set
```

```sh
# Set existing wallpaper
wallpaperinit -set
```

```sh
# Pick random wallpaper (without setting)
wallpaperinit -random
```

# Old
```bash
#!/usr/bin/env bash
file=$(ls $HOME/.config/wallpaper/walls | shuf -n 1)

filename=$(basename -- "$file")
extension="${filename##*.}"

rm $HOME/.wallpaper*
cp $HOME/.config/wallpaper/walls/$file $HOME/.wallpaper.$extension
chmod 777 $HOME/.wallpaper.$extension
feh --bg-fill $HOME/.wallpaper.$extension
```

# Issues with old script
- It would sometimes mess up and would not set a wallpaper
- It did not allow user to set same wallpaper repeatedly
- It Would move a wallpaper to the home dir and rename it, rather than save the name and set it

# Future
- Have config and custom path to wallpapers
- Have multiple collections of custom paths
