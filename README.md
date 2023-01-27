# Hivemapper hdc acl manager

## ACL json representation
```json 
{
  "acl": {
    "managers": ["0000000000000000000000000000000000000000", "0000000000000000000000000000000000000000","0000000000000000000000000000000000000000"],
    "drivers": ["0000000000000000000000000000000000000000", "0000000000000000000000000000000000000000","0000000000000000000000000000000000000000"] 
  },
  "signature": "0000000000000000000000000000000000000000000000000000000000000000"
}
```

Run this command to store an ACL on the HDC.
- `{hex_acl}` is the HEX representation of the ACL
- `{b85_signature}` is the base85 representation of the signature of the ACL JSON
- `{destination}` is the path where ACL will be stored 
```bash
acl store {hex_acl} {b85_signature} {destination}
```

Run this command to retrieve the ACL from the HDC. The ACL will be printed as a HEX representation of the ACL JSON
- `{destination}` is the path where the ACL is stored
```bash
acl load {source_path}
```

## Load result json
```json 
{
  "acl": {
    "managers": ["0000000000000000000000000000000000000000", "0000000000000000000000000000000000000000","0000000000000000000000000000000000000000"],
    "drivers": ["0000000000000000000000000000000000000000", "0000000000000000000000000000000000000000","0000000000000000000000000000000000000000"] 
  },
  "signature": "0000000000000000000000000000000000000000000000000000000000000000"
}
```

To build a linux binary from a Mac, run this command from root folder:
```bash
env GOOS=linux GOARCH=arm64 go build -o acl ./cmd/acl
```

