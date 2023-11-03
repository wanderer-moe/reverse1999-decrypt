# Game Bundles Decrypter

This utility is a simple decryption tool written in Go that processes and decrypts `.dat` files located in a `bundles` directory. The decrypted files are then saved into a `decrypted_bundles` directory. It uses a basic XOR cipher for decryption, with the first byte of each file acting as the encryption key.

## Usage

To use this tool, ensure that you have Go installed on your system. Then, follow these steps:

1. Clone the repository or download the source code to your local machine.
2. Place your `.dat` files inside a directory named `bundles` located at `../StreamingAssets/Windows`.
3. Run the program with the command `go run main.go` from the root directory of the source code.

The decrypted files will be available in the `decrypted_bundles` directory.

## Credits

Original author: [dromzeh](https://github.com/dromzeh)
Go rewrite: [lolpuud](https://github.com/lolpuud)

## Performance Comparison

The Go version of this decrypter significantly outperforms the Python version, making it a preferable choice for larger datasets or when performance is a critical factor.

Here are the performance metrics:

- **Go Decrypter**:
  ![Go Speed](https://cdn.signed.host/6542c2ad433ded5b2a172c16/iUeXe.png)
  The Go decrypter processed the files at a rate of approximately 355.15 files/sec (0.7s).

- **Python Decrypter**:
  ![Python Speed](https://cdn.signed.host/6542c2ad433ded5b2a172c16/A9zjG.png)
  The Python decrypter processed at a slower rate, with the exact metrics visible in the linked image.
