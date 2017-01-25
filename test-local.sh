#!/bin/bash
go test \
  . \
  ./model/... \
  ./modules/... \
  ./router/...
