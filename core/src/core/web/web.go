package web

import "fmt"

type Object struct {

    Slug                                 string
    Name                                 string
    Id                                   string
    Value                                string
    Content                              string // All element content
    //
    //SelfOpenTag        string
    //SelfCloseTag       string
    SingleTag                            bool
    //
    //SelfDelimTagOpen    string
    //SelfDelimTagClose   string
    //
    Parent                                *Object
    //
    templateTag                           string
    //
    Classes                               []string
    Style                                 string
    ChildsCompilationFunctions            []func() string

}

func (o *Object ) Print () {


}

func (o *Object ) Compile ()  ( func ()(string) ) {

    var child_content string

    for chi_num := range o.ChildsCompilationFunctions {

        content        :=  o.ChildsCompilationFunctions[chi_num]()
        child_content  =   child_content+content

    }
    fmt.Printf("\nchild_content |%s|\n",child_content)

    return func() string {


        fmt.Printf("\n returned function is started %s\n",o.Name)
        if o.SingleTag {
            o.Content = "<"+o.Name+"/>"
        } else {
            o.Content = "<"+o.Name+">"+ o.Value + child_content + "</"+o.Name+">"
        }
        return o.Content
    }
}

func Append ( parent *Object,child *Object ) (err error) {

    parent.ChildsCompilationFunctions =   append( parent.ChildsCompilationFunctions, child.Compile() )
    return nil

}

