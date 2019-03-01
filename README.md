# Go authorization proxy for UPS tracking numbers

This API allows you to add authorization to UPS and provide
authorize free API outside (for example for the web pages).

### Authentication data

Provide authorization data as environment variables.

Var name | Description
---|---
UPS_USERNAME | Your UPS user name
UPS_PASSWORD | Your UPS password
UPS_ACCESS_KEY | Access Key generated on UPS devkit page


### How to run?
Generate a self-signed certificate:

```
openssl req -x509 -out localhost.crt -keyout localhost.key \
  -newkey rsa:2048 -nodes -sha256 \
  -subj '/CN=localhost' -extensions EXT -config <( \
   printf "[dn]\nCN=localhost\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:localhost\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")
```

Install Go and run:

```
UPS_USERNAME="login" UPS_PASSWORD="pass" UPS_ACCESS_KEY="access" go run *.go
```

###
Prod push: git push production enable-cors:master
Dev push: git push heroku enable-cors:master