---
sidebar_position: 1
title: Service Creation (~ 5 min)
---

## Service Creation Cheat Sheet <!-- omit in toc -->

- [Pre-Requisites](#pre-requisites)
- [How do I query for all existing onchain Services?](#how-do-i-query-for-all-existing-onchain-services)
- [How do I create a new service?](#how-do-i-create-a-new-service)
- [How do I learn more about interacting with Services?](#how-do-i-learn-more-about-interacting-with-services)
  - [Service Transactions](#service-transactions)
  - [Service Queries](#service-queries)

### Pre-Requisites

1. Make sure to [install the `pocketd` CLI](../../tools/user_guide/pocketd_cli.md).
2. Make sure you know how to [create and fund a new account](../../tools/user_guide/create-new-wallet.md).

### How do I query for all existing onchain Services?

You can query for all services like so:

```bash
pocketd query service all-services --node https://shannon-testnet-grove-rpc.beta.poktroll.com --output json | jq
```

Here is an example of the output on Beta TestNet as of writing this document:

```json
{
  "service": [
    {
      "id": "svc_8ymf38",
      "name": "name for svc_8ymf38",
      "compute_units_per_relay": "7",
      "owner_address": "pokt1aqsr8ejvwwnjwx3ppp234l586kl06cvas7ag6w"
    },
    {
      "id": "svc_drce83",
      "name": "name for svc_drce83",
      "compute_units_per_relay": "7",
      "owner_address": "pokt1mgtf9k4k3pze57gwp3qsne88jmvqkc37t7vd9g"
    },
    {
      "id": "svc_jk07qh",
      "name": "name for svc_jk07qh",
      "compute_units_per_relay": "7",
      "owner_address": "pokt1mwynfsnzesc38f98zrk08pttjn48tu7crc2p09"
    }
  ],
  "pagination": {
    "total": "3"
  }
}
```

### How do I create a new service?

You can use the `add-service` command to create a new service like so:

```bash
pocketd tx service add-service ${SERVICE_ID} "${SERVICE_NAME_OR_DESCRIPTION}" ${COMPUTE_UNITS_PER_RELAY} \
  --fees 1upokt --from ${SERVICE_OWNER} --chain-id ${CHAIN_ID}
```

Here is a concrete copy-pasta assuming you have created and funded a new account called `$USER`:

```bash
pocketd tx service add-service "svc-$USER" "service description for $USER" 13 \
    --node https://shannon-testnet-grove-rpc.beta.poktroll.com \
    --fees 1upokt --from $USER --chain-id pocket-beta
```

Optionally, you can add some more flags to be ultra-verbose about your local environment:

```bash
pocketd tx service add-service "svc-$USER" "service description for $USER" 13 \
    --node https://shannon-testnet-grove-rpc.beta.poktroll.com \
    --fees 1upokt --from $USER --chain-id pocket-beta \
    --home ~/.pocket --keyring-backend test \
    --yes --output json
```

### How do I learn more about interacting with Services?

#### Service Transactions

```bash
pocketd tx service -h
```

#### Service Queries

```bash
pocketd query service -h
```
