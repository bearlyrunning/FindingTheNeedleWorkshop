# Finding the Needle: An Introduction to Detection Engineering

This repository contains course materials for the 'Finding the Needle: An Introduction to Detection Engineering" workshop, delivered at [DEF CON 32](https://defcon.org/html/defcon-32/dc-32-workshops.html#54217) and [The Diana Initiative 2024](https://www.dianainitiative.org/workshops-2024/finding-the-needle/).

This workshop can be completed in either Golang or Python (3), or both! (Note, however, that we do not support Python 2.7!)

## Prerequisites

It is strongly advised that workshop attendees are comfortable implementing code in either Python3 or Golang, and in using the execution/compilation tooling of their chosen language.

Attendees are also advised to use a laptop (or similar) with 50GB of hard drive space, and 16GB+ of RAM, with your preferred IDE installed.

### Recommended Readings

#### Protocol Buffers

We will be making heavy use of Protocol Buffers to serialize / deserialize log and detection data. 

If you have not previously worked with proto messages before, we recommend going through the following tutorials:
* Golang: https://protobuf.dev/getting-started/gotutorial/
* Python: https://protobuf.dev/getting-started/pythontutorial/

We also recommend adding the following language guide to your bookmark:
* https://protobuf.dev/programming-guides/proto3/

#### gRPC

A preliminary understanding of gRPC (especially how to write client side code) would be beneficial for completing part of the workshop; we recommend going through the following tutorials:

* Golang: https://grpc.io/docs/languages/go/quickstart/
* Python: https://grpc.io/docs/languages/python/quickstart/

## Getting Started

To complete the workshop, you will need to:

* Install [Golang](https://go.dev/doc/install) or [Python 3](https://www.python.org/downloads/).
* Clone this repository locally (in a file location of your choice: `git clone https://github.com/bearlyrunning/FindingTheNeedleWorkshop.git`).
* Implement the missing code (and unit tests where appropriate), indicated by `<TODO: Implement me!>`.

### (Golang) Setup and Installing Requirements

Install Protocol Buffer Compiler `protoc`: https://grpc.io/docs/protoc-installation/

Install Go plugins for the protocol compiler: https://grpc.io/docs/protoc-installation/

Once the code repository is shared with you (we'll be sharing it closer to the workshop date!), ensure all dependencies are set up.
```
$ cd go/
$ go mod tidy
```

### (Python) Setup and Installing Requirements

We recommend using a Python virtual environment (`venv`) to avoid breaking your own Python installation, and to avoid dependency collisions. To do so, navigate to the `python` directory and run the following command to create a virtual environment directory named `venv`:

```
$ python3 -m venv venv
```

Next, activate your virtual environment as below.

For those using Mac/Linux:

```
$ source venv/bin/activate
```

And for Windows:

```
> venv/bin/Activate.ps1
```

Once the code repository is shared with you (we'll be sharing it closer to the workshop date!), ensure all dependencies are set up:

```
(venv) $ pip3 install -r requirements.txt
```

### IDE Configuration

If applicable, don't forget to configure your IDE to add linter / debugger support for your chosen langugage. 

For example, if your are using Visual Studio Code, the following extensions could be handy:
* https://marketplace.visualstudio.com/items?itemName=ms-python.python
* https://marketplace.visualstudio.com/items?itemName=golang.Go
* https://marketplace.visualstudio.com/items?itemName=zxh404.vscode-proto3
