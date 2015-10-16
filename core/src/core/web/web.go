package web

type Object struct {

    Slug               string
    Name               string
    Id                 string
    Value              string
    Content            string // All element content
    //
    //SelfOpenTag        string
    //SelfCloseTag       string
    SingleTag          bool
    //
    //SelfDelimTagOpen   string
    //SelfDelimTagClose  string
    //
    Parent             *Element
    //
    templateTag        string
    //
    Classes            []string
    Style              string

}

func (o *Object ) Print () {


}

func (e *Object ) Compile () {


}


