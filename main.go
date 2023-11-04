package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

const (
	decryptionKey     = 0x55
	verificationKey   = 0x6E
	chunkSize         = 4096
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

func decryptByte(data byte, key byte) byte {
	return data ^ key
}

func (fd *FileDecryptor) decryptDataChunk(chunk []byte, key byte) []byte {
	decryptedChunk := make([]byte, len(chunk))
	for i, b := range chunk {
		decryptedChunk[i] = decryptByte(b, key)
	}
	return decryptedChunk
}

func (fd *FileDecryptor) decryptFile(inputPath string, outputPath string) {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", inputPath, err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", outputPath, err)
		return
	}
	defer outputFile.Close()

	initialData := make([]byte, 2)
	_, err = io.ReadFull(inputFile, initialData)
	if err != nil {
		fmt.Printf("Error reading initial data from %s: %v\n", inputPath, err)
		return
	}

	key := initialData[0] ^ decryptionKey
	if key != initialData[1] ^ verificationKey {
		fmt.Println("Invalid key")
		return
	}

	outputFile.Write(initialData)

	buffer := make([]byte, chunkSize)
	for {
		n, err := inputFile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", inputPath, err)
			return
		}
		decryptedChunk := fd.decryptDataChunk(buffer[:n], key)
		outputFile.Write(decryptedChunk)
	}
}

func (fd *FileDecryptor) decryptBundles() (duration time.Duration, filesDecrypted int) {
	startTime := time.Now()
	filepath.Walk(bundlesPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error walking through path %s: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == decryptedDirName {
			return filepath.SkipDir
		}
		if !info.IsDir() && filepath.Ext(path) == ".dat" {
			outputPath := filepath.Join(fd.DecryptedPath, info.Name())
			fd.decryptFile(path, outputPath)
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
	duration, filesDecrypted := decryptor.decryptBundles()

	rps := float64(filesDecrypted) / duration.Seconds()
	fmt.Printf("Decryption completed in %v. Rate: %.2f files/sec\n", duration, rps)
}
