package web

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
    ChildsCompilationFunctions            []func()

}

func (o *Object ) Print () {


}

func (o *Object ) Compile ()  ( func()) {

    for chi_num := range o.ChildsCompilationFunctions {

        o.ChildsCompilationFunctions[chi_num]()

    }

    return func() {

        if o.SingleTag {
            o.Content = "<"+o.Name+"/>"
        } else {
            o.Content = "<"+o.Name+">"+o.Value+o.Content+"</"+o.Name+">"

        }
    }
}

func Append ( parent *Object,child *Object ) (err error) {


  return nil

}

