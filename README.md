# jumpto

Personal little CLI to quickly jump to folders. Only made for macOS. Will work unexpectedly and weirdly on any other OS.

```shell
# Create new route
jp2 -m mc /Volumes/Media/Minecraft

# Jump to route
jp2 mc

# Delete route
jp2 -d mc
```

## Installation

In your shell config (~/.zshrc, ~/.bashrc, etc.):

```shell
jp2() {
    target=$(command jp2-bin "$@")
    if [ -d "$target" ]; then
        cd "$target"
    else
        echo "$target"
    fi
}
```