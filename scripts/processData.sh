#!/bin/bash
curl -X POST -H "Content-Type: application/json" -d '{"board":"board1", "soil":30}' http://localhost:8090/processData
