## Go Flatpak CLI

![preview image](img.png)

A simple command line tool to search and install flatpaks.
Using "flatpak search org.XYZ.123" is not only inconvenient, its woefully slow.
With go-flatpak you can just run it and install what you wan't too!

<i>Please note that go-flatpak is still in an early phase!</i>

Run:

```sh
  ./go-flatpak
```

Use the Fzf-like interface to browse through the flatpak database and hit ENTER to install a package
or mark multiple packages with TAB to install more than one at a time.

# Export flatpak apps to the Command line

## Link Flatpak apps to their usual command line names

```sh
./go-flatpak --link

```

Creates a "bin" folder in the current directory with executable files that link
the flatpak app with their shorthand names that you would usually use on the command line.

Example:

- `org.Blender.blender` becomes `blender`
- `io.github.giantpinkrobots.flatsweep` becomes `flatsweep`
- `com.google.AndroidStudio` becomes `android_studio`
