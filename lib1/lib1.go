package main

import(	
	"fmt"
)

type libname string

func (n libname)GetLibName(){
	fmt.Println("LIB1.SO")
}

var LibName libname
