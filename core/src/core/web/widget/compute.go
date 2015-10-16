package widget

import "core/web"

type Widget struct {

    web.Object

}

type Table  struct {
    web.Element
    SelfOpenTag        string  `<table><tbody>`
    SelfCloseTag       string  `</tbody></table>`
}

type Tr struct {
    web.Object
    SelfOpenTag        string  `<tr>`
    SelfCloseTag       string  `</tr>`
}

type Td struct {
    web.Object
    SelfOpenTag        string  `<td>`
    SelfCloseTag       string  `</td>`
}





type Td struct {




}



func NewWidget ()  (w Widget) {

    w.Id = "Hello"


    return
}

