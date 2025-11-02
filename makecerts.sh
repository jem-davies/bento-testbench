#!/bin/bash
# call this script with an email address (valid or not).
# like:
# ./makecerts.sh foo@foo.com

if [ "$1" == "" ]; then
    echo "Need email as argument"
    exit 1
fi

EMAIL=$1

rm -rf certs
mkdir certs
cd certs

echo "make CA"
PRIVKEY="test"
openssl req -new -x509 -days 365 -keyout ca.key -out ca.pem \
    -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=www.random.com/emailAddress=KryptoKings@random.com" \
    -passout pass:$PRIVKEY

echo "make server cert"
# Generate server key
openssl genrsa -out server.key 2048

# Create server CSR with SAN
cat > server_ext.cnf << EOF
[req]
distinguished_name = req_distinguished_name
req_extensions = v3_req

[req_distinguished_name]

[v3_req]
subjectAltName = @alt_names

[alt_names]
DNS.1 = localhost
IP.1 = 127.0.0.1
EOF

# Create server certificate request
openssl req -new -key server.key -out server.csr \
    -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=localhost/emailAddress=${EMAIL}"

# Sign server cert with CA
openssl x509 -req -days 365 -in server.csr \
    -CA ca.pem -CAkey ca.key -CAcreateserial \
    -out server.pem \
    -extfile server_ext.cnf -extensions v3_req \
    -passin pass:$PRIVKEY

echo "make client cert"
# Generate client key
openssl genrsa -out client.key 2048

# Create client certificate request
openssl req -new -key client.key -out client.csr \
    -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=client.com/emailAddress=${EMAIL}"

# Sign client cert with CA
openssl x509 -req -days 365 -in client.csr \
    -CA ca.pem -CAkey ca.key -CAcreateserial \
    -out client.pem \
    -passin pass:$PRIVKEY

# Cleanup
rm -f server.csr client.csr server_ext.cnf

echo "Done! Certificates created:"
echo "  CA: ca.pem"
echo "  Server: server.pem, server.key"
echo "  Client: client.pem, client.key"