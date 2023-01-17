# Hivemapper hdc acl manager


Run this command to store an ACL from tha HDC. That acl and the signature will be stored in the destination folder ...
```bash
acl store {hex_acl} {b85_signature} {destination}
```

Run this command to retrieve the ACL from the HDC. The ACL will be printed as a HEX representation of the ACL json
```bash
acl load {source_path}
```

```json 
{
  "acl": {
    "managers": ["0000000000000000000000000000000000000000", "0000000000000000000000000000000000000000","0000000000000000000000000000000000000000"],
    "driver": ["0000000000000000000000000000000000000000", "0000000000000000000000000000000000000000","0000000000000000000000000000000000000000"] 
  },
  "signature": "0000000000000000000000000000000000000000000000000000000000000000"
}
```

To build a linux binary from a mac, run this command from root folder:
```bash
env GOOS=linux GOARCH=arm64 go build -o acl ./cmd/acl
```

