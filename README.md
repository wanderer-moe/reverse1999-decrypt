# Reverse: 1999 Bundles Decryptor

- This utility is a simple decryption tool written in Go or Python that processes & decrypts `.dat` files located in the `bundles` directory. 
- The decrypted files are then saved into a `decrypted_bundles` directory. It uses a basic XOR cipher for decryption, with the first byte of each file acting as the encryption key.

## Usage

### Go

To use this tool, ensure that you have Go installed on your system. Then, follow these steps:

1. Clone the repository or download the source code to your local machine.
2. Place your `.dat` files inside a directory named `bundles` located at `../StreamingAssets/Windows`.
3. Run the program with the command `go run main.go` from the root directory where `bundles` is located.

The decrypted files will be available in the `decrypted_bundles` directory.

### Python

To use this tool, ensure that you have Go installed on your system. Then, follow these steps:

1. Clone the repository or download the source code to your local machine.
2. Place your `.dat` files inside a directory named `bundles` located at `../StreamingAssets/Windows`.
3. Run the program with the command `py run main.py` from the root directory where `bundles` is located.

## Credits

Python: [dromzeh](https://github.com/dromzeh)

Go: [okpuud](https://github.com/okpuud)

## Performance Comparison

The Go version of this decrypter significantly outperforms the Python version, making it a preferable choice for larger datasets or when performance is a critical factor.

Here are the performance metrics:

### Go
![Go Speed](https://cdn.signed.host/6542c2ad433ded5b2a172c16/iUeXe.png)
  
  The Go decrypter processed the files at a rate of approximately 355.15 files/sec.

### Python
  ![Python Speed](https://cdn.signed.host/6542c2ad433ded5b2a172c16/A9zjG.png)
  
  The Python decrypter processed at a slower rate, with the exact metrics visible in the linked image.
