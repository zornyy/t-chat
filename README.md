```
- Server listens on a port
- Clients connect via TCP
- Simple text message broadcasting
- No encryption yet
```

**Phase 2: Add structure**
```
- Define a message protocol (binary or JSON)
- Add usernames/identity
- Implement channels for message routing
- Add rooms/channels
```

**Phase 3: Encryption**
```
- TLS for transport (easiest with crypto/tls)
- End-to-end encryption for messages
- Key exchange between clients
```

**Phase 4: Terminal UI**
```
- Use libraries like bubbletea or termui
- Split screen for messages/input
- User lists, etc.
```

## Quick Architecture Sketch
```
Server:
- Main goroutine: Accept incoming connections
- Per-client goroutine: Read messages from client
- Broadcast goroutine: Fan-out messages to all clients
- Channels to coordinate everything

Client:
- Connection goroutine: Maintain TCP connection
- Read goroutine: Receive messages from server
- Input goroutine: Read from terminal
- Display goroutine: Render to terminal