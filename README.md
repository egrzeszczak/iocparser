# iocparser

**WORK IN PROGRESS**

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

- URLs
- IPv4 Addresses,
- MD5, SHA1, SHA256 hashes (TODO: Work in progress)
- Domain Names, (TODO: Work in progress)
- Email Addresses, (TODO: Work in progress)

## 2. Enriching IoCs

`iocparser` will enrich detected IoCs with:

- Links to public reputation services (such as Talos Intelligence, VirusTotal, AbuseIPDB, IBM XForce) (TODO: Work in progress)
- Create RSA NetWitness queries (TODO: Work in progress)

## 3. Printing and formating IoCs

`iocparser` will print detected and enriched IoCs:

- In a Markdown format (usage: `-f markdown`) (TODO: Work in progress)
- As a [Jira wiki text](https://jira.atlassian.com/secure/WikiRendererHelpAction.jspa?section=all) (TODO: Work in progress)
- As a CSV (TODO: Work in progress)
