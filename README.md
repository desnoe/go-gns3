[![Go Report Card](https://goreportcard.com/badge/github.com/desnoe/go-gns3)](https://goreportcard.com/report/github.com/desnoe/go-gns3)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

# go-gns3

*go-gns3* is a Go library that can be used to interact with a [GNS3 server](https://github.com/GNS3/gns3-server) using its HTTP REST API. A GNS3 server manages emulators such as Dynamips, VirtualBox or Qemu/KVM.

## Disclaimer

This library does not aim at explaining nor documenting in any way the GNS3 server API. Please check the [GNS3 server official documentation](https://gns3-server.readthedocs.io/en/latest/index.html) for further information.

## Getting Started

Tests provide a very good starting point.

### Prerequisites

You'll need a [GNS3 server](https://github.com/GNS3/gns3-server) appliance or virtual machine to use the library. Instructions on how to install a server appliance or virtual machine can be found on the [GNS3 website](https://www.gns3.com/).

### Installing

```
go get -v github.com/desnoe/go-gns3
```

## Running the tests

You'll need a [GNS3 server](https://github.com/GNS3/gns3-server) appliance or virtual machine to test the library. Instructions on how to install a server appliance or virtual machine can be found on the [GNS3 website](https://www.gns3.com/).

Location of this test server is provided via these 2 environment variables:

| Environment variable name | Description                                    | Example        |
|:-------------------------:|------------------------------------------------|:--------------:|
| `GNS3_HOST`               | The IP address or FQDN of the GNS3 test server | 172.16.213.128 |
| `GNS3_PORT`               | The TCP port number of the GNS3 test server    |       3080     |

You then simply need to perform a `go test -v`.

## Limitations

In this version, all [CRUD operations](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) of the following object types have been implemented:
- projects
- nodes:
    * EthernetSwitch
    * VPCS
    * QEMU
- links

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/desnoe/go-gns3/tags). 

## Authors

* **Olivier Desnoë** - *Initial work* - [Albatross Networks](http://albatross-networks.com)

See also the list of [contributors](https://github.com/desnoe/go-gns3/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
