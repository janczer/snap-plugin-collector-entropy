# Snap Collector Plugin - Entropy

This plugin collects entropy metrics from file `/proc/sys/kernel/random/entropy_avail`.

It's used in the [Snap framework](https://github.com/intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [Installation] (#installation)
  * [Configuration and Usage](#configuration-and-usage)

## Getting Started

### Installation
#### Build the plugin binary:
Fork https://github.com/intelsdi-x/snap-plugin-collector-entropy
Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-collector-entropy.git
```

Build the Snap entropy plugin by running make within the cloned repo:
```
$ make
```
It may take a thile to pull dependencies if you haven't had then already.
This builds the plugin in `./build/`


### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

#### Load the plugin
```
$ snaptel plugin load build/snap-plugin-collector-entropy
$ snaptel metric list
```

### Examples
Create entropy-task.yaml:

```
---
  version: 1
  schedule:
    type: "simple"
    interval: "1s"
  max-failures: 10
  workflow:
    collect:
      metrics:
        /janczer/procfs/entropy: {}
      publish:
        - plugin_name: "file"
          config:
            file: "/tmp/entropy_metrics.log"
```

Create task in Snap:
```
$ snaptel task create -t entropy-task.yaml
