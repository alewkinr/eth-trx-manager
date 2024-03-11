# ethereum-transaction-manager

## Task description
РDevelop a RESTful API service for managing transactions on the Ethereum blockchain (goerli test network), considering high loads.

## Requirements
- [x] The API should support the following operations:
    - [x] Sending a new transaction (POST /transactions)
    - [x] Retrieving the balance of an Ethereum address (GET /balances/{address})
    - [x] Retrieving transaction information by hash (GET /transactions/{hash})
        - [x] hash (transaction hash)
        - [x] from (sender address)
        - [x] to (recipient address)
        - [x] value (amount of Eth sent)
        - [x] timestamp
- [x] Use Ethers.js (for Node.js) or go-ethereum (for Go) to interact with the Ethereum (Goerli) blockchain.
- [x] Implement caching of address balances to reduce RPC load (any caching solution can be used, e.g., Redis or built-in caching).
- [x] Optimize RPC interaction by minimizing the number of requests.
- [x] Provide results of load testing demonstrating the service's ability to handle at least 100 requests per second.
- [x] Publish the source code on GitHub.
- [x] Optional but a plus: Dockerfile + instructions for running in Docker and/or a Docker image.

## Requirements

- [vegeta](https://github.com/tsenart/vegeta) _(for loadtesting)_
- docker _(build&deployment)_

## Build
To build an application you can run a bash-script in `scripts/docker-build.sh`. If you want to customize the image tag, set the `IMAGE_TAG` argument and pass it to the script.
```shell
IMAGE_TAG=v2 ./scripts/docker-build.sh 
```

## How to start
Copy and paste `.env-example` to the `.env` file and change the provided values. After that, you can run docker container using this command:
```shell
docker run -it --rm --env-file .env -p 8080:8080 --name eth-trx-manager ghcr.io/alewkinr/eth-trx-manager:latest
```
> If you have changed the image tag or any other parameters in `scripts/docker-build.sh`, change the image name.

## Load testing
Load testing was proceeded using [vegeta](https://github.com/tsenart/vegeta) CLI. There is a test data file for loadtesting and a `scripts/loadtest/loadtest.sh` bash script to start the process.
> It will loadtest `GET /wallets/{hash}` endpoing using 10k different Ethereum wallets hashes.

The report will be available in the same folder as the run script. Here is the Macbook M2 PRO (16RAM, 512SSD) report:
![](/Users/alewkinr/tech/github.com/alewkinr/eth-trx-manager/scripts/loadtest/hist.png)

_Also, you will be able to see detailed reports after load test run in the same directory as the run-script_ 
