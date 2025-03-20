# GithubStore

Proposes a db backed cache for github, as an extra option for the bridge.

## Background

Githubridge is a gRPC bridge for github, designed to run inside Kubernetes. It
has some limited caching of elements but this is incomplete. We propose to add a cache
layer to githubridge called githubstore that will cache objects  as they are updated
and ensure that our requests out to github are reduced to just the updates.

## Functionality

The caching layer is a new job that sits between clients and githubridge, which implements
the full githubridge API but runs as a pass through. It either serves out of the cache or reaches
out to githubridge to fulfil requests.

## Implementation

1. Write a shell githubstore that has no caching
1. Bring up the shell in prod
1. Dashboards tracking outbound and inbound request
1. Dashbaords show percentage of requests served from cache
1. Implement get labels caching
1. Validate that requests are served our of cache
