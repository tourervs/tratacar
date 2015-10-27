//package base
package main
//
//debug
import "fmt"

//
import "core/web"
//import "core/osutils"

const APP_NAME = "base"


func TestTableExample()( *web.Object ) {

    table  :=  &web.Object{ Name : "table" }
    tbody  :=  &web.Object{ Name : "tbody" }
    tr     :=  &web.Object{ Name : "tr" }
    td1    :=  &web.Object{ Name : "td" , Value:"Hello world !" }
    td2    :=  &web.Object{ Name : "td" , Value:"Hello Russia !" }

    web.Append(table,tbody)
    web.Append(tbody,tr)
    web.Append(tr,td1)
    web.Append(tr,td2)

    return table

}

func BaseHtml ()( *web.Object )  {

    html            :=  &web.Object{ Name : "html" }
    head            :=  &web.Object{ Name : "head" }
    static_angular  :=  &web.Object{ Name : "script", Params : []string{`src="/angular.min.js"`} }
    body            :=  &web.Object{ Name : "body" }
    h1              :=  &web.Object{ Name : "h1", Value : "Hello world!"  }
    p               :=  &web.Object{ Name : "p" }
    input           :=  &web.Object{ Name : "input" , Value : "ok" }


    web.Append ( head, static_angular     )
    web.Append ( html, head               )
    web.Append ( html, body               )
    web.Append ( body, h1                 )
    web.Append ( body, p                  )
    web.Append ( p, input                 )
    web.Append ( body, TestTableExample() )

    return html

}


func main() {

    html_code := BaseHtml()

    content   := html_code.Compile()().Content

    //osutils.StaticFilesFinder()

    fmt.Printf( "%s", content )

}
