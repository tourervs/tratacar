package web

// import "fmt" // just for debug

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
    ChildsCompilationFunctions            []func()( func () string )

}

func (o *Object ) Print () {


}

func (o *Object ) Compile ()  ( func ()(string) ) {

    var child_content string

    //fmt.Printf("\nRunning compile : %s Childs_count :%d \n", o.Name, len(o.ChildsCompilationFunctions))

    for chi_num := range o.ChildsCompilationFunctions {

        content        :=  o.ChildsCompilationFunctions[chi_num]()
        child_content  =   child_content+content()+"\n"
        //fmt.Printf("childs content len %s",child_content)

    }
    //fmt.Printf("\nchild_content |%s|\n",child_content)

    return func() string {


        //fmt.Printf("\n returned function is started %s\n",o.Name)
        if o.SingleTag {
            o.Content = "<"+o.Name+"/>"
        } else {

            var new_line string
            if child_content != "" { new_line="\n" }

            o.Content = "<"+o.Name+">" + new_line + o.Value + child_content + new_line +"</"+o.Name+">"
        }
        return o.Content
    }
}

func Append ( parent *Object, child *Object ) () {

    parent.ChildsCompilationFunctions =   append( parent.ChildsCompilationFunctions, child.Compile )
    //fmt.Printf("\nAppend: parent: %s child %s\n",parent.Name,child.Name)

}

