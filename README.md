# nginx-reloader

![GitHub last commit](https://img.shields.io/github/last-commit/amirhnajafiz/nginx-reloader)
![GitHub issues](https://img.shields.io/github/issues/amirhnajafiz/nginx-reloader)
![GitHub forks](https://img.shields.io/github/forks/amirhnajafiz/nginx-reloader)
![GitHub stars](https://img.shields.io/github/stars/amirhnajafiz/nginx-reloader)
![GitHub Workflow Status](https://github.com/amirhnajafiz/nginx-reloader/actions/workflows/release.yaml/badge.svg)

## Overview

`nginx-reloader` is an init-container designed to streamline the deployment of frontend applications using nginx by automating the process of replacing default nginx HTML files with those from your own web projects. This tool is perfect for deploying `react.js`, `vue.js`, `angular.js`, or even static HTML applications.

## Features

- **Easy Deployment**: Set up your frontend apps with nginx without writing complex Dockerfiles.
- **Flexibility**: Supports various frameworks like React.js, Vue.js, Angular.js, and more.
- **Automation**: Automatically places your built app into the nginx HTML directory.

## Getting Started

### Prerequisites

- Docker installed on your machine.
- Basic knowledge of Docker and containerization.

### Usage

Pull the following image:

```bash
docker pull ghcr.io/amirhnajafiz/nginx-reloader:latest
```

A docker-compose example:

```yaml
version: '3'
services:
  nginx-reloader:
    image: ghcr.io/amirhnajafiz/nginx-reloader:latest
    environment:
      - NR_TYPE=clone
      - NR_ADDRESS=https://github.com/research-camp/profile.git
      - NR_NGINX_HTML_DIR=/etc/nginx-reloader/tmp
      - NR_TMP_LOCAL_DIR=/etc/nginx-reloader/nginx
    volumes:
      - nginx-reloader-data:/etc/nginx-reloader/tmp
      - nginx-var-dir:/etc/nginx-reloader/nginx

  nginx:
    image: nginx:latest
    ports:
      - "8080:80"
    volumes:
      - nginx-var-dir:/usr/share/nginx/html
    depends_on:
      - nginx-reloader

volumes:
  nginx-reloader-data:
  nginx-var-dir:
```

### Env variables

- `NR_TYPE` : can be clone (for git) or fetch (for normal file download from an address)
- `NR_ADDRESS` : the address of your web-app build content
- `NR_DOWNLOAD_FILENAME` : can be used when you are fetching a file with an address that does not have filename in it
- `NR_NGINX_HTML_DIR` : by default is set to /usr/share/nginx/html
- `NR_TMP_LOCAL_DIR` :  by default is set to /etc/nginx-reloader/tmp
