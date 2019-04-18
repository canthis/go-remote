## Generating app icon

1. Copy your *.ico (for win) and *.png (for unix) files to this folder (```scripts```)
2. Launch scripts:

Windows
```
make_icon.bat your_icon.ico
```

Unix
```
make_icon.sh your_icon.png
```

Those scripts will produce ```iconwin.go``` and ```iconunix.go``` files and place them into ```icon``` folder