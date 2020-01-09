# Architecture

Announcer System have following architecture:
![Architecture Diagram](architecture.jpg?raw=true "Architecture overview")

### Content Service
Content service is application that running on spesific time to update resources inside git repository.

### Announcement Orchestrator
Announcement Orchestrator is long running service and automatically pull data from git repository.

### Dispatcher
Service send announcement to spesific platform like Telegram and Twitter.