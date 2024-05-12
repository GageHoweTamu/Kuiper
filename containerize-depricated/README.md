# Notes for the "Sednoid" project

This is a software project that aims to democratize cloud computing. Users can pay for compute credits on the website. They can use these credits to pay other users to perform computations. Like a decentralized cryptocurrency, Sednoid will be controlled and maintained primarily by its users.

## Tasks

1. Task one is the cornerstone of the project, and requires a method to safely perform remote code execution (RCE.)

2. Task two will be creating a suitable user interface. Performance is not critical and we can probably get away with using Electron, though I'd prefer something more modern.

3. Networking. How do software clients find other software clients? How are packets of code sent? How do we ensure user payment and computation is successful?

4. Website. This will include a database for users and payments for compute credits.

## Challenges

1. Safe but arbitrary code execution is a huge challenge. Even solutions like Docker or Proxmox don't completely account for the dangers of remote code execution (RCE.) If this isn't possible, ask clients to run inside of a virtual machine.

2. The second challenge will be networking and security. This will likely require concepts from blockchain technology.

## Technologies to be used (flexible)

* Docker, Proxmox, etc - for containerizing code
* Golang - Code for blockchain and networking

## Roadmap

With all that said, here's a potential roadmap for the project:

* write code to spin up a docker image, run some bash script, and shut down
* research blockchain and implement currency
* 