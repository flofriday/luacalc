package main

// #cgo CFLAGS:
// #cgo LDFLAGS: -L./lua/ -llua
// #include "./lua/lua.h"
// #include "./lua/lualib.h"
// #include "./lua/lauxlib.h"
//
// int helper_dostring(lua_State *L, const char *s) {return luaL_dostring(L, s);}
// int helper_pcall(lua_State *L, int n, int r, int f ) {return lua_pcallk(L, (n), (r), (f), 0, NULL);}
// int helper_tonumber(lua_State *L, int i) {return lua_tonumber(L, (i));}
//
import "C"
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s IMPLEMENTATION\n", os.Args[0])
		os.Exit(1)
	}

	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	input := string(bytes)

	L := C.luaL_newstate()
	C.luaL_openlibs(L)
	C.helper_dostring(L, C.CString(input))

	fmt.Println("LuaCalc - A calculator with plugable implementation")
	fmt.Println("Enter two numbers like 5 + 3")

	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, err := reader.ReadString('\n')
		if err != nil || strings.TrimSpace(text) == "" {
			break
		}

		parts := strings.Split(text, "+")
		if len(parts) != 2 {
			log.Fatal("Only two numbers are supported")
		}

		first, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			log.Fatal(err)
		}
		second, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			log.Fatal(err)
		}

		C.lua_getglobal(L, C.CString("sum"))
		C.lua_pushinteger(L, C.longlong(first))
		C.lua_pushinteger(L, C.longlong(second))

		C.helper_pcall(L, 2, 1, 0)
		res := C.helper_tonumber(L, -1)
		fmt.Printf("%v\n", res)
	}

	C.lua_close(L)
}
