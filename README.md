# Inspektor

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/AdrianWR/inspektor/build-go-app.yml?style=flat-square&logo=github)
![Go Mod Version](https://img.shields.io/github/go-mod/go-version/AdrianWR/inspektor/main?style=flat-square&logo=go&label=go.mod&labelColor=white&color=darkblue)
<object>
<img alt="Docker Image Version (latest semver)" src="https://img.shields.io/docker/v/awroque/inspektor/v1?style=flat-square&logo=docker&label=latest&labelColor=white&link=https%3A%2F%2Fhub.docker.com%2Frepository%2Fdocker%2Fawroque%2Finspektor">
</object>

Inspektor is a simple API tool that I use to test container and Kubernetes deplomeyents wheneven I'm doing infrastructure tests; it's modeled as a Golang API server that can extract useful information about it own environment. If you hit the `/v1/inspect` endpoint, you should get a response similar to the following payload example:

```json
{
  "name": "inspektor",
  "version": "v1",
  "pod": {
    "name": "inspektor-7d4b7f7b5f-8q9qj",
    "namespace": "dev",
    "ip": "172.10.5.24"
  }
}
```

# Development

## How to Build and Run

The development workflow requires a working Golang environment and `make` installed. To build and run the application, just run the following commands:

```bash
$ make build
$ ./bin/inspektor

# or just run the following command
$ make run
```

# License

MIT License
