FROM golang:1.21-alpine AS builder
ENV CGO_ENABLED=0
RUN apk add --update make
WORKDIR /backend
COPY go.* .
ARG TARGETARCH
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go mod download
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  make bin

FROM node:14.17-alpine3.13 AS client-builder
WORKDIR /ui
# cache packages in layer
COPY ui/package.json /ui/package.json
COPY ui/package-lock.json /ui/package-lock.json
ARG TARGETARCH
RUN --mount=type=cache,target=/usr/src/app/.npm-${TARGETARCH} \
  npm set cache /usr/src/app/.npm-${TARGETARCH} && \
  npm ci
# install
COPY ui /ui
RUN npm run build

FROM alpine
ARG GIT_STRIPPED_VERSION
ARG GIT_VERSION
ARG RELEASE_CHANNEL
LABEL org.opencontainers.image.title="Meshplay" \
  org.opencontainers.image.description="Meshplay, the collaborative Kubernetes manager. Go multi-player as you design and operate your cloud native infrastructure with teammates." \
  org.opencontainers.image.vendor="Layer5, Inc." \
  com.docker.desktop.extension.api.version=">= 0.2.0" \
   com.docker.extension.screenshots="[ \
      { \
        \"alt\": \"Meshplay Docker Extension\", \
        \"url\": \"https://raw.githubusercontent.com/meshplay/meshplay/master/install/docker-extension/docs/img/meshmap-docker-extension-for-meshplay.png\" \
      },{ \
        \"alt\": \"Meshplay Docker Extension\", \
        \"url\": \"https://raw.githubusercontent.com/meshplay/meshplay/master/install/docker-extension/docs/img/Docker-extension-meshplay.png\" \
      } \
    ]" \
  com.docker.extension.detailed-description="\
  <h2>Visually and collaboratively design and operate your Kubernetes clusters (<a href='https://meshplay.io/catalog'>video</a>).</h2> \
  <p>The Meshplay Docker Extension is your cloud native infrastructure designer, complete with multi-cluster Kubernetes management. Meshplay provides cloud native engineers with visual and collaborative interface to designing and operating cloud native infrastructure.</p> \
  <ul> \
    <li><b>Discovery of your Kubernetes environments</b> - Meshplay is a multi-cluster manager and will scan your kubeconfig, allowing you to select which of your Kubernetes contexts to connect. Switch from one K8s cluster to the next or manage multiple concurrently.</li> \
    <li><b>Support for your Docker Compose apps -</b> Import your Docker Compose apps. Configure and deploy them to Kubernetes.</li> \
    <li><b>Visual designer for Docker Compose apps, Helm charts, and Kubernetes manifests -</b> No code, visual topology for designing Docker Compose applications, operating Kubernetes and their workloads.</li> \
    <li><b>Save time with design patterns - </b>Turbo-charge your infrastructure with best practice cloud native design patterns from the <a href='https://meshplay.io/catalog'>Meshplay Catalog</a>.</li> \
    <li><b>Single-click deployment of all cloud native infrastructure -</b> Support for hundreds of different cloud native infrastructure tools right at your fingertips.</li> \
  </ul>" \
  com.docker.desktop.extension.icon="https://raw.githubusercontent.com/meshplay/meshplay/master/install/docker-extension/docs/img/meshplay-logo-light.svg" \
  com.docker.extension.publisher-url="https://meshplay.io" \
  com.docker.extension.additional-urls="[{\"title\":\"Documentation\",\"url\":\"https://docs.meshplay.io\"},{\"title\":\"Project\",\"url\":\"https://layer5.io/meshplay\"},{\"title\":\"Slack\",\"url\":\"https://slack.meshplay.io\"},{\"title\":\"Discussion Forum\",\"url\":\"http://discuss.meshplay.io\"}]"
COPY --from=builder /backend/bin/service /
COPY docker-compose.yaml .
COPY metadata.json .
RUN mkdir -p /meshplayctl/darwin && mkdir -p /meshplayctl/windows && mkdir -p /meshplayctl/linux
ADD https://github.com/meshplay/meshplay/releases/download/${GIT_VERSION}/meshplayctl_${GIT_STRIPPED_VERSION}_Darwin_x86_64.tar.gz /meshplayctl/darwin/
ADD https://github.com/meshplay/meshplay/releases/download/${GIT_VERSION}/meshplayctl_${GIT_STRIPPED_VERSION}_Windows_x86_64.zip /meshplayctl/windows/
ADD https://github.com/meshplay/meshplay/releases/download/${GIT_VERSION}/meshplayctl_${GIT_STRIPPED_VERSION}_Linux_x86_64.tar.gz /meshplayctl/linux/
RUN tar -xzf /meshplayctl/darwin/meshplayctl_${GIT_STRIPPED_VERSION}_Darwin_x86_64.tar.gz -C /meshplayctl/darwin/
RUN unzip /meshplayctl/windows/meshplayctl_${GIT_STRIPPED_VERSION}_Windows_x86_64.zip -d /meshplayctl/windows/
RUN tar -xzf /meshplayctl/linux/meshplayctl_${GIT_STRIPPED_VERSION}_Linux_x86_64.tar.gz -C /meshplayctl/linux/
COPY meshplay-logo-light.svg .
COPY --from=client-builder /ui/build ui
EXPOSE 7877/tcp
CMD ./service
