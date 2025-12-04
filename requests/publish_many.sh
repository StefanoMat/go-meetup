#!/bin/bash
for i in {1..20}; do
  curl -s -X POST localhost:8080/publish -H "Content-Type: application/json" -d "{\"body\":\"msg $i\"}" > /dev/null
done
echo "Enviado!"


