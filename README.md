# Reverse: 1999 Bundles Decryptor

- Fast decryption tool for Reverse: 1999 bundles written in Go that processes & decrypts `.dat` files located in the `bundles` directory. 
- The decrypted files are then saved into a `decrypted_bundles` directory. It uses a basic XOR cipher for decryption, with the first byte of each file acting as the encryption key.

![Go Speed](https://cdn.signed.host/6542c2ad433ded5b2a172c16/iUeXe.png)

## Usage

### Building from source
To use this tool, ensure that you have [Go installed on your system](https://go.dev/doc/install). Then, follow these steps:

1. Clone the repository or download the source code.
2. Copy your `bundles` folder located at `../StreamingAssets/Windows` to the root directory where `main.go` is located.
3. Run the program with the command `go run main.go`.

The decrypted files will be available in the `decrypted_bundles` directory.

## Credits

Original Implementation: [dromzeh](https://github.com/dromzeh)

Go Rewrite: [alluding](https://github.com/alluding)

