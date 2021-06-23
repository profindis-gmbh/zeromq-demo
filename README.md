# ZeroMQ
ZeroMQ is a messagging library. build using C and it suports many languages including GoLang anfd Java etc.
ZeroMQ is brokerless: unlike other messaging system such as kafak or rabbitMQ, zeroMQ does not rely on a centrelized broker.
[source](http://api.zeromq.org/4-1:zmq-socket#toc2)


## Supported message patterns 
4 Messaging pattern: 
1. Request-reply pattern
2. Publish-subscribe pattern
3. Pipeline pattern
4. Exclusive pair pattern

### Request-reply pattern
The request-reply pattern is used for sending requests from a ``ZMQ_REQ`` client to one or more ``ZMQ_REP`` services, and receiving subsequent replies to each request sent.
[Example](./zmq4-demo/request-reply)

### Publish-subscribe pattern
The publish-subscribe pattern is used for one-to-many distribution of data from a single publisher to multiple subscribers in a fan out fashion.
[Example](./zmq4-demo/pub-sub)

### Pipeline pattern
The pipeline pattern is used for distributing data to nodes arranged in a pipeline. Data always flows down the pipeline, and each stage of the pipeline is connected to at least one node. When a pipeline stage is connected to multiple nodes data is round-robined among all connected nodes.
[Example](./zmq4-demo/push-pull)

## Exclusive pair pattern
The exclusive pair pattern is used to connect a peer to precisely one other peer. This pattern is used for inter-thread communication across the inproc transport.
[Example](./zmq4-demo/inproc)

## Socket types
* REQ: used by a client to send a request and then get a reply; the peer socket (server side) shoudl be either REP or ROUTER
If no server is available to receive the receive the requet the send operation on the socket blocks until there is at least one service is available. 
  
  
* REP: listens for the received requests and sends back the response

* PUB: (send only) used by a publisher to distribute messages among the existing subscribers.

* SUB: (receive only) used to subscribe to a publisher 

* PUSH: (send only) used to send messages to a downstream pipeline nodes

* PULL: (receive only) used by a pipeline node to receive messages from upstream pipeline nodes

* ROUTER: like REp but it envellops the reply with the destination ... 

* DEALER:

### Advatange : 
 * Simple
 * No pre-configuration 
 * Easy to impelement
 * Different pattern can be implemented

### Cons
* No much of examples are provided 
* The documentation is not simple
* Many Socket types look similar (
  router <-> REP ; Dealer <-> (REP/REQ)
  XPUB <-> PUB ; XSUB <-> SUB
)
