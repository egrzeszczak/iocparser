# iocparser

**WARNING!** It is still dangerous to use since the values printed aren't defanged yet. Working on it.

Will detect:
- File Hashes: MD5, SHA-1, SHA-256,
- IPv4 Addresses,
- Domain Names,
- URLs,
- Email Addresses

It will then create links for public CTI services accordingly to the detected IoC type and parse them into a [pretty markdown format](example-output.md).


## Usage

```sh
Usage: go run main.go <file_path>
```
