package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "time"
)

const (
    decryptionKey     = 0x55
    verificationKey   = 0x6E
    decryptedDirName  = "decrypted_bundles"
    bundlesPath       = "bundles"
)

type FileDecryptor struct {
    DecryptedPath string
}

func NewFileDecryptor(bundlesPath string) *FileDecryptor {
    decryptedPath := filepath.Join(bundlesPath, decryptedDirName)
    if _, err := os.Stat(decryptedPath); os.IsNotExist(err) {
        err := os.MkdirAll(decryptedPath, os.ModePerm)
        if err != nil {
            panic(err)
        }
    }
    return &FileDecryptor{DecryptedPath: decryptedPath}
}

func (fd *FileDecryptor) decryptDataChunk(chunk []byte, key byte) []byte {
    decryptedChunk := make([]byte, len(chunk))
    for i, b := range chunk {
        decryptedChunk[i] = b ^ key
    }
    return decryptedChunk
}

func (fd *FileDecryptor) decryptFile(inputPath string, outputPath string) error {
    inputData, err := ioutil.ReadFile(inputPath)
    if err != nil {
        return fmt.Errorf("error reading file %s: %v", inputPath, err)
    }

    key := inputData[0] ^ decryptionKey
    if key != inputData[1] ^ verificationKey {
        return fmt.Errorf("invalid key")
    }

    decryptedData := fd.decryptDataChunk(inputData[2:], key)
    err = ioutil.WriteFile(outputPath, decryptedData, 0644)
    if err != nil {
        return fmt.Errorf("error writing file %s: %v", outputPath, err)
    }

    return nil
}

func (fd *FileDecryptor) decryptBundles() (duration time.Duration, filesDecrypted int, err error) {
    startTime := time.Now()
    err = filepath.Walk(bundlesPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return fmt.Errorf("error walking through path %s: %v", path, err)
        }
        if info.IsDir() && info.Name() == decryptedDirName {
            return filepath.SkipDir
        }
        if !info.IsDir() && filepath.Ext(path) == ".dat" {
            outputPath := filepath.Join(fd.DecryptedPath, info.Name())
            err := fd.decryptFile(path, outputPath)
            if err != nil {
                return err
            }
            filesDecrypted++
            fmt.Printf("Decrypted %s to %s\n", info.Name(), outputPath)
        }
        return nil
    })
    duration = time.Since(startTime)
    return
}

func main() {
    decryptor := NewFileDecryptor(bundlesPath)
    duration, filesDecrypted, err := decryptor.decryptBundles()
    rps := float64(filesDecrypted) / duration.Seconds()
    fmt.Printf("Decryption completed in %v. Rate: %.2f files/sec\n", duration, rps)

    fmt.Println("Press any key to exit")
    fmt.Scanln()
}
