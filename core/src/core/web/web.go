package web

type Element struct {

    Slug               string
    VerboseName        string
    Id                 string
    Value              string
    Content            string
    //
    SelfOpenTag        string
    SelfCloseTag       string
    //
    //SelfDelimTagOpen   string
    //SelfDelimTagClose  string
    //
    Parent             *Element
    //
    TemplateTag        string
    //
    Classes            []string

}

func (e *Element ) Print () {


}

func (e *Element ) Compile () {


}


