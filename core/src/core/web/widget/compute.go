package widget

import "core/web"

type Widget struct {

    web.Element

}

type Table  struct {

    web.Element
    SelfOpenTag        string  `<table><tbody>`
    SelfCloseTag       string  `</tbody></table>`
    SelfDelimTagOpen   string  `<tr>`
    SelfDelimTagClose  string  `</tr>`





}


func NewWidget ()  (w Widget) {

    w.Id = "Hello"


    return
}

