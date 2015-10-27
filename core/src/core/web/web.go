package web

//import "fmt" // just for debug

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
    Childs                                []*Object

}

func (o *Object ) Print () {


}

func (o *Object ) Compile ()  ( func ()(*Object) ) { // need to add returning  second argument . *Object

    var child_content string
    //fmt.Printf("\nRunning compile : %s Childs_count :%d \n", o.Name, len(o.ChildsCompilationFunctions))

    for i:= range o.Childs {
        // Collect childs html code
        obj_compile        :=  o.Childs[i].Compile()
        child_content      =   child_content+obj_compile().Content+"\n"
        //fmt.Printf("childs content len %s",child_content)

    }

    var params        string
    for prm_id := range o.Params {
        // Append params to html code
        params+=o.Params[prm_id]+" "
    }
    //fmt.Printf("\nchild_content |%s|\n",child_content)

    return func() *Object {


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
        return o
    }
}

func Append ( parent *Object, child *Object ) () {

    parent.Childs  =  append( parent.Childs , child )
    //fmt.Printf("\nAppend: parent: %s child %s\n",parent.Name,child.Name)

}

func AppendToChild ( parent *Object, existing_parent_child_name string  ,new_child *Object ) () {

  //Example:  parent: html-obj ; existing_parent_child_name: "body" ; child: div-obj

  child , found := ChildDiveSearch( parent, existing_parent_child_name , "")

  if found { child.Childs = append ( child.Childs, new_child ) }

}

func ChildDiveSearch ( parent *Object, existing_parent_child_name string , param string )( child_obj *Object, found bool ) {

    childs     :=  parent.Childs
    child_obj  =   &Object{}

    //fmt.Printf("\n<< parent.Name %s \n",parent.Name)

    for i:= range parent.Childs {

        //fmt.Printf("\n<< child.Name %s \n", childs[i].Name)

        if childs[i].Name == existing_parent_child_name {

            //child_obj  =  childs[i]
            //found      =  true
            //break
            return childs[i], true

        } else {

            //fmt.Printf("\n]] child name: %s childs len: %d \n",childs[i].Name,len(childs[i].Childs))

            if len(childs[i].Childs) > 0 {

                temp_co, found := ChildDiveSearch ( childs[i] , existing_parent_child_name , param )

                if found == true { return temp_co, true }

             }

        }

    }

   //fmt.Printf("\nReturning\n")
   /*if ( child_obj == nil ) { fmt.Printf("\n child_obj is nil\n") }
   fmt.Printf("\nchild_obj.Name: %s Found: %t\n",child_obj.Name,found)*/
   return child_obj, false

}
