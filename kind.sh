#!/usr/bin/env bash
set -Eeuo pipefail
IFS=$'\n\t'
header() {
    printf '\n'
    printf '%*s\n' 60 '' | tr ' ' '='
    printf ' %s\n' "$1"
    printf '%*s\n' 60 '' | tr ' ' '='
}

create_cluster() {

    header "Kind create cluster"
    kind create cluster --wait 30s --name default-cluster
    header "Kind get clusters"
    kind get clusters
    
   
}
load_images() {
    kind load postgres:15.18-alpine3.24 busybox nginx:trixie-perl redis:trixie
    echo "$!"
}
pull_images() {
    if docker info &>/dev/null; then 
        echo "docker is not running or is not installed"
        exit 1
    header "Installing Docker Images"
    docker pull postgres:15.18-alpine3.24 & 
    pid1=$!
    docker pull busybox &
    pid2=$!
    docker pull nginx:trixie-perl &
    pid3=$!
    docker pull redis:trixie &
    pid4=$!

    wait

    if pid1 then;
        echo "postgres installed successfully"
    fi

    if pid2 ; then
        echo "busybox installed successfully"
    fi

    if pid3 ; then
        echo "nginx insalled successfully"
    fi

    if pid4 ; then
        echo "redis installed successfully"
    fi

    if pid1 && pid2 && pid3 && pid4; then
        echo "all images pulled successfully"
    else 
        echo "There was errors installing one or more of the images: (redis, nginx, busybox or postgres)"
    
}
install_kubernetes() {
    header "installing kubectl"
    sudo dnf install kubectl
    
    if dnf info kubectl &>/dev/null; then
        echo "kubectl installed successfully"
    else 
        echo "kubectl is not installed halting execution"
        exit 1
}

probe_clusters() {
    header "Get info for default cluster" 
    kubectl cluster-info --context kind-default-cluster
}

main() {
    create_cluster 
    install_kubernetes
    probe_clusters
    pull_images
}

main "$@"
