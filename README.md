

## Known Issues
- Memory leak in the image processing worker (tracking in #142)
- WebSocket disconnects on iOS Safari devices occasionally



### Core Architecture
The system is built using a microservices architecture communicating via generic gRPC endpoints.



### Endpoints
- `GET /api/v1/health` - Health check status code
- `POST /api/v1/users` - Create a new user record

