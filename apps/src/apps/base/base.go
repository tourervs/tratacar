//package base
package main
//
//debug
import "fmt"

//
import "core/web"





func baseHTML ()(string)  {

    html            :=  &web.Object{ Name : "html" }
    head            :=  &web.Object{ Name : "head" }
    body            :=  &web.Object{ Name : "body" }
    h1              :=  &web.Object{ Name : "h1", Value : "Hello world!"  }

    web.Append ( html, head )
    web.Append ( html, body )
    web.Append ( body, h1   )

    html.Compile()()

    return html.Content

}


func main() {

    fmt.Printf("%s",baseHTML())

}
