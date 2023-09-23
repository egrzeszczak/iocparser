# iocparser

A tool for detecting and enriching IoCs.

**Usage:**

```sh
go run main.go -i file-with-iocs.txt
```

```sh
go run main.go -i file-with-iocs.txt -o iocs.md -f markdown
```

```sh
go run main.go -i file-with-iocs.txt -o iocs.jira -f jira
```

## 1. Detecting IoCs

`iocparser` will detect:

- MD5, SHA1, SHA256 hashes
- IPv4 Addresses,
- Domain Names,
- Email Addresses,
- URLs

## 2. Enriching IoCs

`iocparser` will enrich detected IoCs with:

- Links to public reputation services (such as Talos Intelligence, VirusTotal, AbuseIPDB, IBM XForce) 

## 3. Printing and formating IoCs

`iocparser` will print detected and enriched IoCs:

- In a Markdown format (usage: `-f markdown`)
