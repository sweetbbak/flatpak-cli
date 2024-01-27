default:
    go build -o go-flatpak
build:
    go build -o go-flatpak
install:
    #!/usr/bin/env bash
    [ -d "$HOME/bin" ] && cp go-flatpak "$HOME/bin"
reinstall:
    #!/usr/bin/env bash
    [ -e "$HOME/bin/go-flatpak" ] && rm "$HOME/bin/go-flatpak"
    [ -d "$HOME/bin" ] && cp go-flatpak "$HOME/bin"
