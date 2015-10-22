//package base
package main
//
//debug
import "fmt"

//
import "core/web"

func baseHTML ()(string)  {

    objHtml            :=  &web.Object{}
    objHead            :=  &web.Object{}
    objBody            :=  &web.Object{}

    objHtml.Name       =   "html"
    objHead.Name       =   "head"
    objBody.Name       =   "body"

    objHead.Parent     =   objHtml
    objBody.Parent     =   objHtml

    _                  =   web.Append ( objHead, objHead )
    _                  =   web.Append ( objHead, objBody )

    objBody.Value = "<h1>Hello</h1>"

    objHtml.Compile()()
    if objHtml.Parent == nil { fmt.Printf("HTML parent is nil")}
    return objHtml.Content

}


func main() {

    fmt.Printf("%s",baseHTML())

}
