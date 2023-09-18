# iocenhancer

Will detect:
- File Hashes: MD5, SHA-1, SHA-256 hashes of files known to be malicious.
- File Names: Suspicious or known malicious filenames.
- IP Addresses: IP addresses known to be associated with malicious activities.
- Domain Names: Suspicious or malicious domain names.
- URLs: Suspicious or known malicious URLs.
- Email Addresses: Suspicious or known malicious email addresses.
from a file

It will then create links for public CTI services accordingly to the detected IoC type

## Usage

```sh
go run main.go iocs.txt
```
