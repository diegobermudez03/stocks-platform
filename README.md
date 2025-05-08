# Stocks Application
---

## 🚀 DEPLOYED SITE 🚀

**Access the live application here:** [**http://stocks-frontend-bucket.s3-website-us-east-1.amazonaws.com/**](http://stocks-frontend-bucket.s3-website-us-east-1.amazonaws.com/)

---

This project is a full-stack application designed to display or manage stock information. It consists of a backend service and a frontend web application, both automatically deployed using GitHub Actions.

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Technology Stack](#technology-stack)
- [CI/CD](#cicd)
- [Prerequisites](#prerequisites)
- [Local Development](#local-development)
- [Deployment](#deployment)
- [Configuration](#configuration)

## Overview

The application is divided into two main components:

*   **Backend (`stocks-backend`)**: A service responsible for business logic and data management. It is containerized using Docker and deployed to Amazon ECR (Elastic Container Registry).
*   **Frontend (`stocks-frontend`)**: A web application, likely built with Vue.js (using Vite), that provides the user interface. It is deployed as a static site to Amazon S3.

## Project Structure

```
platform/
├── .github/
│   └── workflows/
│       ├── backend.yml     # CI/CD workflow for the backend
│       └── frontend.yml    # CI/CD workflow for the frontend
├── stocks-backend/         # Contains all backend-related code and Dockerfile
└── stocks-frontend/        # Contains all frontend-related code (Vue.js, Vite)
└── README.md               # This file
```

## Technology Stack

*   **Backend**:
    *   Docker
*   **Frontend**:
    *   Node.js (v20)
    *   npm
    *   Vite (assumed, based on `VITE_API_URL`)
    *   Vue.js (assumed)
*   **CI/CD**:
    *   GitHub Actions
*   **Cloud Provider**:
    *   AWS (ECR for backend, S3 for frontend)
*   **Live URL:** [**http://stocks-frontend-bucket.s3-website-us-east-1.amazonaws.com/**](http://stocks-frontend-bucket.s3-website-us-east-1.amazonaws.com/)


## CI/CD

Continuous Integration and Continuous Deployment (CI/CD) are managed using GitHub Actions.

*   **Backend CI/CD (`.github/workflows/backend.yml`)**:
    *   Triggered on pushes to the `main` branch if changes occur in `stocks-backend/` or the workflow file itself.
    *   Builds a Docker image from `./stocks-backend`.
    *   Pushes the Docker image to Amazon ECR (`stocks-backend` repository with `latest` tag).

*   **Frontend CI/CD (`.github/workflows/frontend.yml`)**:
    *   Triggered on pushes to the `main` branch if changes occur in `stocks-frontend/` or the workflow file itself.
    *   Sets up Node.js v20.
    *   Installs dependencies (`npm ci`) in `stocks-frontend/`.
    *   Builds the Vue application (`npm run build`) using the `BACKEND_URL` secret.
    *   Syncs the built assets from `./stocks-frontend/dist` to an Amazon S3 bucket (`stocks-frontend-bucket`).

## Prerequisites

For local development or contributing, you might need:

*   Node.js (version 20 recommended for frontend)
*   npm (comes with Node.js)
*   Docker
*   AWS CLI (optional, for interacting with AWS services manually)
*   Git

## Local Development

### Backend

1.  Navigate to the `stocks-backend` directory: `cd stocks-backend`
2.  Build the Docker image: `docker build -t stocks-backend .`
3.  Run the Docker container: `docker run -p <host_port>:<container_port> stocks-backend` (replace ports as needed)

### Frontend

1.  Navigate to the `stocks-frontend` directory: `cd stocks-frontend`
2.  Install dependencies: `npm install`
3.  Set up your local environment variables, especially `VITE_API_URL` (e.g., in a `.env.local` file).
4.  Run the development server: `npm run dev`

## Deployment

Deployment is automated via GitHub Actions upon pushing to the `main` branch.

*   **Backend**: Deployed to Amazon ECR.
*   **Frontend**: Deployed to Amazon S3.

## Configuration

The CI/CD pipelines rely on several environment variables and secrets:

### Environment Variables (defined in workflow files):
*   `AWS_REGION`: `us-east-1` (for both backend and frontend)
*   `ECR_REPOSITORY`: `stocks-backend` (for backend)
*   `IMAGE_TAG`: `latest` (for backend)
*   `S3_BUCKET_NAME`: `stocks-frontend-bucket` (for frontend)

### GitHub Secrets (need to be configured in repository settings):
*   `AWS_ACCESS_KEY_ID`: AWS access key for authentication.
*   `AWS_SECRET_ACCESS_KEY`: AWS secret key for authentication.
*   `BACKEND_URL`: The URL of the deployed backend API, used by the frontend during its build process.

