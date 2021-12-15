package main


import (
    "archive/tar"
    "errors"
    "path/filepath"
    "fmt"
    "io"
    "os"
)


func CreateTarball(tarballFilePath string, packageDir string, filePaths []string) error {
    file, err := os.Create(tarballFilePath)
    if err != nil {
        return errors.New(fmt.Sprintf("Could not create tarball file '%s', got error '%s'", tarballFilePath, err.Error()))
    }
    defer file.Close()

    tarWriter := tar.NewWriter(file)
    defer tarWriter.Close()

    for _, filePath := range filePaths {
        err := addFileToTarWriter(filePath, packageDir, tarWriter)
        if err != nil {
            return errors.New(fmt.Sprintf("Could not add file '%s', to tarball, got error '%s'", filePath, err.Error()))
        }
    }

    return nil
}


func addFileToTarWriter(filePath string, packageDir string, tarWriter *tar.Writer) error {
    file, err := os.Open(filePath)
    if err != nil {
        return errors.New(fmt.Sprintf("Could not open file '%s', got error '%s'", filePath, err.Error()))
    }
    defer file.Close()

    stat, err := file.Stat()
    if err != nil {
        return errors.New(fmt.Sprintf("Could not get stat for file '%s', got error '%s'", filePath, err.Error()))
    }

    header := &tar.Header{
        Name:    filepath.Join(packageDir, filePath),
        Size:    stat.Size(),
        Mode:    int64(stat.Mode()),
        ModTime: stat.ModTime(),
    }

    err = tarWriter.WriteHeader(header)
    if err != nil {
        return errors.New(fmt.Sprintf("Could not write header for file '%s', got error '%s'", filePath, err.Error()))
    }

    _, err = io.Copy(tarWriter, file)
    if err != nil {
        return errors.New(fmt.Sprintf("Could not copy the file '%s' data to the tarball, got error '%s'", filePath, err.Error()))
    }

    return nil
}


func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: mytar output-file package-dir [input-files...]")
        os.Exit(1)
    }
    CreateTarball(os.Args[1], os.Args[2], os.Args[3:])
}
