# Documentation


### handleConnection

In the handleConnection function, after accepting a new connection, we add the new peer to the peer list by calling addPeer.

The handleConnection function then enters a loop where it reads incoming messages from the peer's connection.

When a message is received, it prints the message and calls the broadcastMessage function to forward the message to all other connected peers.

The broadcastMessage function iterates over the peer list and sends the message to all peers except the sender by writing to their connections.

If there's an error sending the message to a peer, it removes that peer from the peer list by calling removePeer.
The removePeer function removes the peer from the peers map.

The peer struct now includes peerX and peerY fields to represent the coordinates of each peer.
The generateCoordinates function is added to generate random coordinates for new peers (currently returning fixed values for simplicity).
The handleConnection function now uses bufio.NewReader to read messages line by line.
The handleMessage function is added to handle different types of messages received from peers.
The BROADCAST message type is added, which allows a peer to broadcast a message to all other peers.
The PEER_LIST message type is added, which requests the peer list from the recipient peer.
The sendPeerList function is added to send the list of connected peers to a specific peer.
In the main function, after connecting to the bootstrap node, we request the peer list from it by sending the PEER_LIST message.