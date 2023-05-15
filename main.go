package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"   
    "plugin" 
)

type LibName interface{
    GetLibName()
}

// Users struct which contains
// an array of users
type Libs struct {
    Libs []Lib `json:"libs"`
}

// User struct which contains a name
// a type and a list of social links
type Lib struct {
    Name   string `json:"name"`
}


func main() {
    // Open our jsonFile
    jsonFile, err := os.Open("libs.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened libs.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened xmlFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    // we initialize our Users array
    var libs Libs

    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'users' which we defined above
    json.Unmarshal(byteValue, &libs)

    // we iterate through every user within our users array and
    // print out the user Type, their name, and their facebook url
    // as just an example
    for i := 0; i < len(libs.Libs); i++ { 
        
        // Loading the libraries
        // load module
        // 1. open the so file to load the symbols
        fmt.Println("Lib Name: " + libs.Libs[i].Name)   
        lib := "./libs/" + libs.Libs[i].Name
        plug, err := plugin.Open(lib)
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        // 2. look up a symbol (an exported function or variable)
        // in this case, variable LibName
        symLibName, err := plug.Lookup("LibName")
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        // 3. Assert that loaded symbol is of a desired type
        // in this case interface type LibName (defined above)
        var libname LibName
        libname, ok := symLibName.(LibName)
        if !ok {
            fmt.Println("unexpected type from module symbol")
            os.Exit(1)
        }

        // 4. use the module
        libname.GetLibName()
             
    }

}
