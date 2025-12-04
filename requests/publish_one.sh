#!/bin/bash
curl -X POST localhost:8080/publish -H "Content-Type: application/json" -d '{"body":"mensagem única"}'


