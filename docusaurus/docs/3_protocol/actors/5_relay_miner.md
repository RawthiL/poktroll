---
title: RelayMiner
sidebar_position: 5
---

## RelayMiner <!-- omit in toc -->

- [Overview](#overview)
- [Configuration](#configuration)
- [CLI](#cli)

## Overview

A `RelayMiner` is a specialized operation node (not an onchain actor) designed
for individuals to **offer services** through Pocket Network alongside a staked
`Supplier`. It is responsible for proxying `RelayRequests` between a `PATH Gateway`
and the supplied `Service`.

[Suppliers](6_supplier.md) interested in providing `Service`s on Pocket Network
would need to run a `RelayMiner` in addition to the software that provides the said `Service`.

## Configuration

Configurations and additional documentation related to operating a `RelayMiner`
can be found at [relayminer_config.md](../../1_operate/3_configs/4_relayminer_config.md).

## CLI

All of the operations needed to start and operate a `RelayMiner` can be viewed
by running:

```bash
pocketd relayminer --help
```
