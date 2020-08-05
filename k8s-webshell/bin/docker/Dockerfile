FROM alpine:3.12

LABEL maintainer="k8a-container-terminal by huary"

COPY k8s-container-terminal_linux_amd64 /opt/kube/k8s-container-terminal_linux_amd64
COPY static /opt/kube/static
COPY views /opt/kube/views
COPY k8s-config /opt/kube/k8s-config
COPY conf /opt/kube/conf

RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.12/main/ > /etc/apk/repositories && \
    apk add -U tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime && \
    chmod a+x /opt/kube/k8s-container-terminal_linux_amd64

WORKDIR /opt/kube
CMD ./k8s-container-terminal_linux_amd64