## Golang Lib for DevOps
  - metrics
  - monitoring
  - servers
  - global locking / paxos / raft
  - Cloud API clients





## Distributed Computing with Golang
  - https://code.google.com/p/go-wiki/wiki/Courses
    - [15-440: Distributed Systems Syllabus](http://www.cs.cmu.edu/~dga/15-440/F12/syllabus.html) - very-very-very good course!
    - [6.824 Lab 1: Lock Server](http://pdos.csail.mit.edu/6.824/labs/lab-1.html)

    - [Chinese Courses](http://www1.i2r.a-star.edu.sg/~jlang/golang.html)

## Locking Server/PubSub/Service Discovery
  - GNats/Nats:
    - http://www.reddit.com/r/golang/comments/1oqqx7/gnatsd_from_apcera_a_high_performance_nats_server/
    - NATS does not have persistence, or transactions. It is more like a nervous system, and it will protect itself at all costs and does not have SPOFs. It does publish/subscribe, and distributed queues.
    - http://www.quora.com/Cloud-Foundry/Why-does-CloudFoundry-use-NATS-a-specially-written-messaging-system-whereas-OpenStack-uses-AMQP
    AMQP, and implementations like RabbitMQ, are enterprise messaging systems built with things like durability, transactions, and formal queues. NATS was designed and built to be like a dial-tone publish-subscribe service, something that is always on and available. However, NATS does not provide durability or transactions, and its queuing model is interest-based only. It also protects itself, the NATS service, at all costs, so that it can always be available. This forms a great base platform for building scalable and reliable distributed systems, but is probably not a good fit for the typical enterprise application.

  - Serf:
    - http://www.serfdom.io/
    - https://news.ycombinator.com/item?id=6600063


## Distributed Execution
  - [Исполнение SSH-команд на сотнях серверов с помощью Golang - 2014.02](http://habrahabr.ru/post/215111/)


## Why Golang for DevOps?
  - [Go Language for Ops and Site Reliability Engineering](http://talks.golang.org/2013/go-sreops.slide)



This repository is supposed to work with [DirEnv](https://github.com/zimbatm/direnv). It will set the GOPATH to current directory and append the ./bin folder to your PATH variable.


<!-- PROJECTS_LIST_START -->
    *** GENERATED BY https://github.com/mindreframer/techwatcher (ruby _sh/pull golang-devops-stuff) *** 

    abh/geodns:
      DNS server with per-client targeted responses
       241 commits, last change: 2014-01-08 05:50:49, 280 stars, 24 forks

    abhishekkr/goshare:
      Go Share your TimeSeries/NameSpace/KeyVal DataStore (using leveldb) over HTTP /or ZeroMQ
       57 commits, last change: 2013-10-29 21:25:09, 20 stars, 3 forks

    adnaan/hamster:
      A back end as a service based on MongoDB
       45 commits, last change: 2014-02-20 18:47:38, 51 stars, 3 forks

    apcera/gnatsd:
      High Performance NATS Server
       289 commits, last change: 2014-02-06 19:34:12, 311 stars, 35 forks

    apcera/nats:
      NATS client for Go
       219 commits, last change: 2014-02-22 10:49:47, 70 stars, 9 forks

    ARolek/lilpinger:
      A small site pinging application with email and SMS notifications written in Golang
       5 commits, last change: , 17 stars, 3 forks

    benbjohnson/go-raft:
      A Go implementation of the Raft distributed consensus protocol.
       496 commits, last change: 2014-03-03 22:36:58, 745 stars, 88 forks

    bitly/google_auth_proxy:
      A reverse proxy that provides authentication using Google OAuth2
       19 commits, last change: 2013-10-24 10:42:28, 159 stars, 24 forks

    bitly/nsq:
      A realtime distributed messaging platform
       1,059 commits, last change: 2014-03-05 10:58:33, 2,154 stars, 185 forks

    buger/gor:
      HTTP traffic replay in real-time. Replay traffic from production to staging and dev environments.
       248 commits, last change: 2014-01-29 07:29:46, 1,131 stars, 60 forks

    BurntSushi/cmail:
      cmail runs a command and sends the output to your email address at certain intervals.
       11 commits, last change: , 3 stars, 0 forks

    ccding/go-stun:
      a go implementation of the STUN client (RFC 3489 and RFC 5389)
       4 commits, last change: 2013-08-18 01:10:34, 12 stars, 0 forks

    cloudflare/redoctober:
      Go server for two-man rule style file encryption and decryption.
       72 commits, last change: 2014-03-02 09:15:51, 521 stars, 36 forks

    cloudfoundry/gorouter:

       415 commits, last change: 2014-03-03 15:39:37, 114 stars, 42 forks

    cloudfoundry/gosigar:

       15 commits, last change: 2013-08-06 16:12:49, 49 stars, 19 forks

    cloudfoundry/hm9000:

       316 commits, last change: 2014-03-03 11:08:44, 18 stars, 6 forks

    cloudfoundry/yagnats:
      Yet Another Go NATS client
       61 commits, last change: , 6 stars, 2 forks

    coreos/etcd:
      A highly-available key value store for shared configuration and service discovery
       1,553 commits, last change: 2014-03-06 14:12:51, 2,576 stars, 240 forks

    crosbymichael/skydock:
      Service discovery via DNS for docker
       78 commits, last change: 2014-03-07 12:10:02, 248 stars, 18 forks

    crowdmob/goamz:
      Fork of the GOAMZ version developed within Canonical with additional functionality with DynamoDB
       579 commits, last change: 2014-03-07 08:49:37, 113 stars, 69 forks

    dotcloud/docker:
      Docker - the open-source application container engine
       6,724 commits, last change: 2014-03-07 13:20:02, 10,277 stars, 1,596 forks

    efficient/epaxos:

       26 commits, last change: 2014-01-20 05:21:47, 151 stars, 8 forks

    erikh/gollector:
      json-based metrics collector
       164 commits, last change: 2014-03-01 20:34:42, 90 stars, 5 forks

    errplane/errplane-go:
      Go library for metrics for Errplane
       52 commits, last change: , 13 stars, 0 forks

    flynn/go-crypto-ssh:
      Forked from go.crypto as Flynn working copy until changes merged upstream
       5 commits, last change: 2013-10-09 20:09:38, 2 stars, 1 forks

    flynn/go-discover:
      Service discovery system written in Go
       104 commits, last change: 2014-03-04 14:57:25, 144 stars, 12 forks

    flynn/rpcplus:
      Go RPC plus streaming responses (forked from vitess)
       18 commits, last change: 2014-02-02 16:37:57, 11 stars, 1 forks

    globocom/gandalf:
      Gandalf is an API to manage git repositories.
       464 commits, last change: 2014-02-01 23:55:04, 112 stars, 15 forks

    globocom/tsuru:
      Open source Platform as a Service.
       6,620 commits, last change: 2014-03-06 15:51:20, 809 stars, 74 forks

    golang/groupcache:
      groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases.
       21 commits, last change: 2013-10-30 09:55:26, 2,608 stars, 226 forks

    happypancake/fsd:

        commits, last change: ,  stars,  forks

    hashicorp/serf:
      Service orchestration and management tool.
       950 commits, last change: 2014-03-06 17:16:50, 1,677 stars, 92 forks

    inconshreveable/muxado:
      Stream multiplexing for Go
       13 commits, last change: 2014-01-10 01:38:28, 123 stars, 6 forks

    inconshreveable/ngrok:
      Introspected tunnels to localhost
       262 commits, last change: 2014-02-27 11:56:40, 1,572 stars, 99 forks

    inconshreveable/slt:
      A TLS reverse proxy with SNI multiplexing in Go
       9 commits, last change: 2014-03-01 16:44:01, 79 stars, 1 forks

    influxdb/influxdb:
      Scalable datastore for metrics, events, and real-time analytics
       975 commits, last change: 2014-03-07 08:22:31, 1,528 stars, 72 forks

    jingweno/gotask:
      Idiomatic build tool in Go
       123 commits, last change: 2014-01-12 10:05:21, 105 stars, 5 forks

    joewalnes/websocketd:
      Like inetd, but for WebSockets. Turn any application that uses STDIN/STDOUT into a WebSocket server.
       121 commits, last change: 2014-03-03 14:42:29, 2,976 stars, 134 forks

    jondot/groundcontrol:
      Manage and monitor your Raspberry Pi with ease
       50 commits, last change: 2013-08-22 07:19:32, 645 stars, 42 forks

    jordansissel/lumberjack:
      An experiment to cut logs in preparation for processing elsewhere.
       536 commits, last change: 2014-02-05 02:37:24, 428 stars, 103 forks

    jyap808/jaeger:
      Jaeger is a JSON encoded GPG encrypted key value store. It is useful for generating and keeping configuration files secure. Jaeger is written in Go.
       32 commits, last change: , 7 stars, 2 forks

    kelseyhightower/confd:
      Manage local application configuration files using templates and data from etcd
       199 commits, last change: 2014-03-05 07:21:20, 279 stars, 16 forks

    mailgun/vulcan:
      Programmatic HTTP rate limiting proxy for JSON based APIs.
       221 commits, last change: 2013-12-27 15:28:29, 168 stars, 6 forks

    mindreframer/emtail:
      extract whitebox monitoring data from logs and insert into a timeseries database - mirror for https://code.google.com/p/emtail/
       273 commits, last change: , 1 stars, 0 forks

    mitchellh/packer:
      Packer is a tool for creating identical machine images for multiple platforms from a single source configuration.
       2,720 commits, last change: 2014-03-01 07:20:35, 2,101 stars, 308 forks

    mozilla-services/heka:
      Data collection and processing made easy.
       2,213 commits, last change: 2014-03-06 09:19:33, 970 stars, 90 forks

    necrogami/watchdog:
      Watchdog
       7 commits, last change: , 2 stars, 1 forks

    nf/gohttptun:
      A tool to tunnel TCP over HTTP, written in Go
       22 commits, last change: 2014-01-22 13:29:05, 76 stars, 15 forks

    oleiade/trousseau:
      Networked and encrypted key-value database
       224 commits, last change: 2014-03-06 13:35:05, 518 stars, 19 forks

    petar/GoTeleport:
      Teleport Transport: End-to-end resilience to network outages
       6 commits, last change: , 100 stars, 1 forks

    pubsubsql/pubsubsql:
      An open-source distributed in-memory database integrated with Publish-Subscribe
       433 commits, last change: , 28 stars, 5 forks

    PuerkitoBio/throttled:
      throttling strategies for HTTP handlers
       39 commits, last change: 2014-02-25 06:02:27, 103 stars, 4 forks

    rackspace/gophercloud:
      The Go SDK for the Rackspace Cloud
       187 commits, last change: 2014-02-27 13:09:50, 184 stars, 22 forks

    rcrowley/go-metrics:
      Go port of Coda Hale's Metrics library
       186 commits, last change: 2014-01-27 10:15:57, 301 stars, 36 forks

    Sendhub/shipbuilder:
      The Open-source self-hosted Platform-as-a-Service written in Go
       139 commits, last change: 2013-08-29 02:59:20, 2 stars, 19 forks

    skydb/sky:
      Sky is an open source, behavioral analytics database.
       557 commits, last change: 2014-01-31 08:39:53, 417 stars, 36 forks

    skynetservices/skydns:
      DNS for skynet or any other service discovery
       343 commits, last change: 2014-02-25 11:23:18, 342 stars, 24 forks

    smira/aptly:
      aptly - Debian repository management tool
       470 commits, last change: 2014-03-07 13:24:45, 149 stars, 9 forks

    spf13/nitro:
      Quick and easy performance analyzer library for golang
       7 commits, last change: 2013-10-03 16:43:07, 86 stars, 4 forks

    stripe-ctf/octopus:
      Many-armed network simulator
       17 commits, last change: 2014-03-05 19:58:36, 43 stars, 3 forks

    tsenart/tb:
      A generic lock-free implementation of the "Token-Bucket" algorithm
       49 commits, last change: 2014-02-11 08:14:50, 26 stars, 0 forks

    tsenart/vegeta:
      HTTP load testing tool and library. It's over 9000!
       155 commits, last change: 2014-02-02 16:01:46, 1,432 stars, 69 forks

    uniqush/uniqush-push:
      Uniqush is a free and open source software which provides a unified push service for server-side notification to apps on mobile devices.
       406 commits, last change: 2013-11-07 12:23:32, 399 stars, 51 forks

    vektra/tachyon:
      An experimental configuration management system inspired by ansible
       57 commits, last change: 2014-03-04 23:03:21, 17 stars, 3 forks

    vito/garden:
      Go Warden
       287 commits, last change: 2014-02-09 00:16:48, 0 stars, 4 forks

    VividCortex/robustly:
      Run functions resiliently
       33 commits, last change: , 49 stars, 2 forks

    xtaci/gonet:
      a game server skeleton with golang
       918 commits, last change: 2014-03-04 22:34:18, 166 stars, 76 forks

    YuriyNasretdinov/GoSSHa:
      Ssh client that supports multiple servers
       11 commits, last change: 2014-03-08 20:28:07, 7 stars, 0 forks

    zimbatm/socketmaster:
      Zero downtime restarts for your apps
       51 commits, last change: 2014-01-31 16:54:23, 397 stars, 18 forks
<!-- PROJECTS_LIST_END -->
