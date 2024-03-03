## Golang DoS CVE proof of concept

HTTPS/2 Server with vulnerable version of `golang.org/x/net`, rapid reset attack got to be known last
year, impacting multiple big cloud providers and CDNs, with the record of 201MI RPS.

**WARNING: Do not use the example without FIXING the version**

More on:

https://blog.cloudflare.com/technical-breakdown-http2-rapid-reset-ddos-attack

Scanner results:
```
┌──────────────────┬────────────────┬──────────┬────────┬───────────────────┬───────────────┬──────────────────────────────────────────────────────────────┐
│     Library      │ Vulnerability  │ Severity │ Status │ Installed Version │ Fixed Version │                            Title                             │
├──────────────────┼────────────────┼──────────┼────────┼───────────────────┼───────────────┼──────────────────────────────────────────────────────────────┤
│ golang.org/x/net │ CVE-2022-41723 │ HIGH     │ fixed  │ v0.6.0            │ 0.7.0         │ net/http, golang.org/x/net/http2: avoid quadratic complexity │
│                  │                │          │        │                   │               │ in HPACK decoding                                            │
│                  │                │          │        │                   │               │ https://avd.aquasec.com/nvd/cve-2022-41723                   │
│                  ├────────────────┤          │        │                   ├───────────────┼──────────────────────────────────────────────────────────────┤
│                  │ CVE-2023-39325 │          │        │                   │ 0.17.0        │ golang: net/http, x/net/http2: rapid stream resets can cause │
│                  │                │          │        │                   │               │ excessive work (CVE-2023-44487)                              │
│                  │                │          │        │                   │               │ https://avd.aquasec.com/nvd/cve-2023-39325                   │
│                  ├────────────────┼──────────┤        │                   ├───────────────┼──────────────────────────────────────────────────────────────┤
│                  │ CVE-2023-3978  │ MEDIUM   │        │                   │ 0.13.0        │ golang.org/x/net/html: Cross site scripting                  │
│                  │                │          │        │                   │               │ https://avd.aquasec.com/nvd/cve-2023-3978                    │
│                  ├────────────────┤          │        │                   ├───────────────┼──────────────────────────────────────────────────────────────┤
│                  │ CVE-2023-44487 │          │        │                   │ 0.17.0        │ HTTP/2: Multiple HTTP/2 enabled web servers are vulnerable   │
│                  │                │          │        │                   │               │ to a DDoS attack...                                          │
│                  │                │          │        │                   │               │ https://avd.aquasec.com/nvd/cve-2023-44487                   │
└──────────────────┴────────────────┴──────────┴────────┴───────────────────┴───────────────┴──────────────────────────────────────────────────────────────┘
```

### Running

```
$ go mod tidy
$ go run main.go

2024/03/03 09:39:26 Starting to listen HTTPS server on :6443
```