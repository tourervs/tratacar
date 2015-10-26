package web

// import "fmt" // just for debug

import "core/web/tag"

type HtmlReturner func() string

type Object struct {

    Slug                                 string
    Name                                 string
    Id                                   string
    Value                                string
    Content                              string // All element content
    //
    //SelfOpenTag        string
    //SelfCloseTag       string
    DoubleTag                            bool
    //
    //SelfDelimTagOpen    string
    //SelfDelimTagClose   string
    //
    Parent                                *Object
    Params                                []string
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
        // Collect childs html code
        content        :=  o.ChildsCompilationFunctions[chi_num]()
        child_content  =   child_content+content()+"\n"
        //fmt.Printf("childs content len %s",child_content)

    }

    var params        string
    for prm_id := range o.Params {
        // Append params to html code
        params+=o.Params[prm_id]+" "
    }
    //fmt.Printf("\nchild_content |%s|\n",child_content)

    return func() string {


        //fmt.Printf("\n returned function is started %s\n",o.Name)

        if tag.IsDouble(o.Name) == false  {

            var value string

            if o.Value != "" { value="value="+o.Value } else { value="" }

            o.Content = "<" + o.Name + " " + params + " " + value + " />"

        } else {

            var new_line string
            var space    string

            if child_content != "" { new_line="\n" }
            if params        != "" { space=" " }

            o.Content = "<" + o.Name + space + params + space +">" + new_line + o.Value + child_content + new_line + "</" + o.Name + ">"
        }
        return o.Content
    }
}

func Append ( parent *Object, child *Object ) () {

    parent.ChildsCompilationFunctions =   append( parent.ChildsCompilationFunctions, child.Compile )
    //fmt.Printf("\nAppend: parent: %s child %s\n",parent.Name,child.Name)

}

