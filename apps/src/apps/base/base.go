//package base
package main
//
//debug
import "fmt"

//
import "core/web"




func baseHTML ()( func() string )  {

    html            :=  &web.Object{ Name : "html" }
    head            :=  &web.Object{ Name : "head" }
    body            :=  &web.Object{ Name : "body" }
    h1              :=  &web.Object{ Name : "h1", Value : "Hello world!"  }
    h2              :=  &web.Object{ Name : "p", Value : "Hello !<br/>fvdfv<br/>fvdfvf fvdfvf"  }

    web.Append ( html, head )
    web.Append ( html, body )
    web.Append ( body, h1   )
    web.Append ( body, h2   )

    return html.Compile()

}


func main() {

    fmt.Printf("\n====\n")
    fmt.Printf("%s",baseHTML()())

}
