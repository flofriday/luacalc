# LuaCalc

An excuse to play with the CLua runtime and play with code isolation.

## Build

```bash
cd lua
make all
cd ..
go run main.go add.lua
```

Note: If you change the C code you will need to rebuild lua and you might need
to force a rebuild in Go with `go run -a main.go`

## Features

At the moment it only let's the user select an implementation but there are
no security mitigation to ensure the provided code doesn't crash the 
calculator.

The following features are planned
- [x] Swapable implementation
- [ ] Memory limit
- [ ] Time limit
- [ ] No I/O
- [ ] No imports
