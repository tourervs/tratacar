package osutils

import "path/filepath"
import "runtime"
import "os"
import "fmt"

func StaticFilesFinder() (err error) {

//
// Find "static" directory  in current directory and in parent directory.
// Then return all files pathes under "static/js" and "static/css" directories.
//
//
//
//
//
//
//
//

    if err != nil {
        return  err
    }

    _, filename, _, _ := runtime.Caller(1)

    runtime_dir, err := os.Open(filepath.Dir(filename))
    defer runtime_dir.Close()

    dir_content , err := runtime_dir.Readdirnames(-1)

    for i:= range dir_content {

        fmt.Printf("\n| %s |\n",dir_content[i])

    }

    if err != nil {
        return  err
    }


    return nil
}

