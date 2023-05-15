package main

import(	
	"fmt"
)

type libname string

func (n libname)GetLibName(){
	fmt.Println("LIB2.SO")
}

var LibName libname
