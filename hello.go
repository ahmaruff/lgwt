package main

import "fmt"

const (
    spanish = "es";
    french = "fr";


    englishHelloPrefix = "Hello, "
    spanishHelloPrefix = "Hola, "
    frenchHelloPrefix = "Bonjour, "
)

func Hello(name string, lang string) string {
    if(name == "") {
        name = "World"
    }
    
    return greetingPrefix(lang) + name
}


func greetingPrefix(lang string) (prefix string){
    switch lang {
        case spanish:
            prefix = spanishHelloPrefix
        case french:
            prefix = frenchHelloPrefix
        default:
            prefix = englishHelloPrefix
    }

    return 
}


func main(){
    fmt.Println(Hello("Chris", "en"))
}


