//package base
package main
//
//debug
import "fmt"

//
import "core/web"


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
    body            :=  &web.Object{ Name : "body" }
    h1              :=  &web.Object{ Name : "h1", Value : "Hello world!"  }
    h2              :=  &web.Object{ Name : "p", Value : "Hello !<br/>fvdfv<br/>fvdfvf fvdfvf"  }

    web.Append ( html, head )
    web.Append ( html, body )
    web.Append ( body, h1   )
    web.Append ( body, h2   )
    web.Append ( body, TestTableExample() )

    return html

}


func main() {

    html_code := baseHTML()

    html_code.Compile()()

    fmt.Printf( "%s",html_code.Content )

}
