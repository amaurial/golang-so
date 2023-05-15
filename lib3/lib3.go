package main

import(	
	"fmt"
)

type libname string

func (n libname)GetLibName(){
	fmt.Println("LIB3.SO")
}

var LibName libname
