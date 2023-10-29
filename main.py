# bundles decrypter for a certain game
# this script decrypts all files in the "bundles" folder and saves them in the "decrypted_bundles" folder.
# you can locate the "bundles" folder inside "..\StreamingAssets\Windows".

# author: dromzeh

import os


# this script uses a simple xor cipher for decryption
# the first byte in each file acts as the encryption key, and the second byte is obtained by xoring the key with 0x6e
# the key is then used to decrypt the rest of the file


def decrypt_byte(data, key):
    return data ^ key


def decrypt_data(input_data):
    key = input_data[0] ^ 0x55
    if key != input_data[1] ^ 0x6E:
        raise Exception("Invalid key")
    return bytearray(decrypt_byte(byte, key) for byte in input_data)


def decrypt_file(input_path, output_path):
    with open(input_path, "rb") as input_file, open(output_path, "wb") as output_file:
        bundle_data = input_file.read()
        decrypted_data = decrypt_data(bundle_data)
        output_file.write(bytes(decrypted_data))


def main():
    decrypted = 0
    failed = 0

    if not os.path.exists("decrypted_bundles"):
        os.makedirs("decrypted_bundles")

    if not os.path.exists("bundles"):
        print(
            "bundles folder not found, place the bundles folder in the same directory as this script"
        )
        return

    for file in os.listdir("bundles"):
        try:
            input_path = os.path.join("bundles", file)
            output_path = os.path.join("decrypted_bundles", file)
            decrypt_file(input_path, output_path)

            print(f"decrypted {file}")
            decrypted += 1

        except Exception as e:
            print(f"failed to decrypt {file}: {str(e)}")
            failed += 1

    print(f"decrypted {decrypted} bundles; failed to decrypt {failed} bundles")


if __name__ == "__main__":
    main()
