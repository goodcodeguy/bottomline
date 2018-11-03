# Process Configuration

POST /process-configuration/
PUT /process-configuration/:id
DELETE /process-configuration/:id
GET /process-configuration/:id

POST /process-configuration/search

GET /process-configuration/:id/processes

# Step configuration

POST /process-configuration/:id/step
PUT /process-configuration/:id/step/:id
DELETE /process-configuration/:id/step/:id
GET /process-configuration/:id/step/:id

GET /process-configuration/:id/steps

# Users

POST /user/
PUT /user/:id
DELETE /user/:id
GET /user/:id

POST /user/search

GET /user/:id/workspaces

# Workspaces

POST /workspace/
PUT /workspace/:id
DELETE /workspace/:id
GET /workspace/:id

POST /workspaces/search

GET /workspace/:id/process-configurations
GET /workspace/:id/users

# Running Processes

POST /process/
PUT /process/:id
DELETE /process/:id
GET /process/:id

POST /process/search

GET /process/:id/configuration
GET /process/:id/status
