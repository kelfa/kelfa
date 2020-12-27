# Kelfa Extended Log Format Analytics

Analyse logs in ELF format.

## SQLite in podman

    podman run --rm --privileged -v $(pwd):/data:Z -it quay.io/fale/sqlite sqlite3 database.db

[![Build Status](https://travis-ci.org/kelfa/kelfa.svg?branch=master)](https://travis-ci.org/kelfa/kelfa)

[![Go Report Card](https://goreportcard.com/badge/go.kelfa.io/kelfa)](https://goreportcard.com/report/go.kelfa.io/kelfa)
