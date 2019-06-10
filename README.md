[![Build Status](https://travis-ci.org/lalamove/nui.svg?branch=master)](https://travis-ci.org/lalamove/nui)
[![Go doc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square
)](https://godoc.org/github.com/lalamove/nui)

# Nui
Nui is package containing common interfaces and tools used across Lalamove's go codebase.
Nui means vast in Maori.

# Packages
- [Context](ncontext/README.md)

Interface to create contexts and default implementation wrapping `context` package.

- [FileSystem](nfs/README.md)

Interface to interact with the operating system and default implentation wrapping `os` package.

- [HTTP](nhttp/README.md)

Interface for http.Client mocking. 

- [io](nio/README.md)

Toolbox for `io`. 

- [Logger](nlogger/README.md)

Interface for a generic logger used across open source packages.

- [Strings](nstrings/README.md)

Strings tool box.

- [Time](ntime/README.md)

Time toolbox. Time interface wrapping `time.Time` package.

- [Getter](ngetter/README.md)

Getter interface to retrieve values.

- [Metrics](nmetrics/README.md)

Interface for a registry in [prometheus](https://github.com/prometheus/client_golang)