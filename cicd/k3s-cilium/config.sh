#!/bin/bash

source ../common.sh

echo "#########################################"
echo "Spawning all hosts"
echo "#########################################"

spawn_docker_host --dock-type loxilb --dock-name llb1 --with-bgp yes --bgp-config $(pwd)/llb1_gobgp_config --with-ka in --ka-config $(pwd)/keepalived_config1
spawn_docker_host --dock-type loxilb --dock-name llb2 --with-bgp yes --bgp-config $(pwd)/llb2_gobgp_config --with-ka in --ka-config $(pwd)/keepalived_config2
spawn_docker_host --dock-type host --dock-name ep1
spawn_docker_host --dock-type host --dock-name ep2
spawn_docker_host --dock-type host --dock-name ep3
spawn_docker_host --dock-type host --dock-name r1 --with-bgp yes --bgp-config $(pwd)/r1_config
spawn_docker_host --dock-type host --dock-name r2 --with-bgp yes --bgp-config $(pwd)/r2_config
spawn_docker_host --dock-type host --dock-name user

echo "#########################################"
echo "Connecting and configuring  hosts"
echo "#########################################"

connect_docker_hosts user r1
connect_docker_hosts r1 llb1
connect_docker_hosts r1 llb2
connect_docker_hosts llb1 r2
connect_docker_hosts llb2 r2
connect_docker_hosts r2 ep1
connect_docker_hosts r2 ep2
connect_docker_hosts r2 ep3

## Make network for k3s connectivity
sudo ip link add ellb1sys type veth peer name esysllb1
sleep 3
sudo ip link set ellb1sys netns llb1
sleep 3
sudo ip -n llb1 link set ellb1sys up
sudo ip -n llb1 addr add 12.12.12.1/24 dev ellb1sys
sudo ip link set esysllb1 up
sudo ip addr add 12.12.12.254/24 dev esysllb1

sudo ip link add ellb2sys type veth peer name esysllb2
sleep 3
sudo ip link set ellb2sys netns llb2
sleep 3
sudo ip -n llb2 link set ellb2sys up
sudo ip -n llb2 addr add 14.14.14.1/24 dev ellb2sys
sudo ip link set esysllb2 up
sudo ip addr add 14.14.14.254/24 dev esysllb2

#node1 config
config_docker_host --host1 user --host2 r1 --ptype phy --addr 1.1.1.1/24 --gw 1.1.1.254
config_docker_host --host1 r1 --host2 user --ptype phy --addr 1.1.1.254/24

create_docker_host_vlan --host1 r1 --host2 llb1 --id 11 --ptype untagged
create_docker_host_vlan --host1 llb1 --host2 r1 --id 11 --ptype untagged

config_docker_host --host1 r1 --host2 llb1 --ptype vlan --id 11 --addr 11.11.11.254/24 --gw 11.11.11.11
config_docker_host --host1 llb1 --host2 r1 --ptype vlan --id 11 --addr 11.11.11.1/24

create_docker_host_vlan --host1 r1 --host2 llb2 --id 11 --ptype untagged
create_docker_host_vlan --host1 llb2 --host2 r1 --id 11 --ptype untagged
config_docker_host --host1 llb2 --host2 r1 --ptype vlan --id 11 --addr 11.11.11.2/24


create_docker_host_vlan --host1 llb1 --host2 r2 --id 10 --ptype untagged
config_docker_host --host1 llb1 --host2 r2 --ptype vlan --id 10 --addr 10.10.10.1/24
create_docker_host_vlan --host1 llb2 --host2 r2 --id 10 --ptype untagged
config_docker_host --host1 llb2 --host2 r2 --ptype vlan --id 10 --addr 10.10.10.2/24

create_docker_host_vlan --host1 r2 --host2 llb1 --id 10 --ptype untagged
create_docker_host_vlan --host1 r2 --host2 llb2 --id 10 --ptype untagged
create_docker_host_vlan --host1 r2 --host2 ep1 --id 10 --ptype untagged
create_docker_host_vlan --host1 r2 --host2 ep2 --id 10 --ptype untagged
create_docker_host_vlan --host1 r2 --host2 ep3 --id 10 --ptype untagged
config_docker_host --host1 r2 --host2 llb1 --ptype vlan --id 10 --addr 10.10.10.254/24

create_docker_host_vlan --host1 r2 --host2 ep1 --id 31 --ptype untagged
config_docker_host --host1 r2 --host2 ep1 --ptype vlan --id 31 --addr 31.31.31.254/24

create_docker_host_vlan --host1 ep1 --host2 r2 --id 31 --ptype untagged
config_docker_host --host1 ep1 --host2 r2 --ptype vlan --id 31 --addr 31.31.31.1/24 --gw 31.31.31.254

create_docker_host_vlan --host1 r2 --host2 ep2 --id 32 --ptype untagged
config_docker_host --host1 r2 --host2 ep2 --ptype vlan --id 32 --addr 32.32.32.254/24

create_docker_host_vlan --host1 ep2 --host2 r2 --id 32 --ptype untagged
config_docker_host --host1 ep2 --host2 r2 --ptype vlan --id 32 --addr 32.32.32.1/24 --gw 32.32.32.254

create_docker_host_vlan --host1 r2 --host2 ep3 --id 33 --ptype untagged
config_docker_host --host1 r2 --host2 ep3 --ptype vlan --id 33 --addr 33.33.33.254/24

create_docker_host_vlan --host1 ep3 --host2 r2 --id 33 --ptype untagged
config_docker_host --host1 ep3 --host2 r2 --ptype vlan --id 33 --addr 33.33.33.1/24 --gw 33.33.33.254

##Pod networks
$hexec r1 ip route add 20.20.20.1/32 via 11.11.11.11
#add_route llb1 1.1.1.0/24 11.11.11.254
#add_route llb2 1.1.1.0/24 11.11.11.254

# Route back to user
sudo ip route add 1.1.1.1/32 via 12.12.12.1

# Change default route in llb1
$hexec llb1 ip route del default
$hexec llb1 ip route add default via 12.12.12.254

# Change default route in llb2
$hexec llb2 ip route del default
$hexec llb2 ip route add default via 14.14.14.254

sleep 1
##Create LB rule
create_lb_rule llb1 20.20.20.1 --tcp=2020:8080 --endpoints=31.31.31.1:1,32.32.32.1:1,33.33.33.1:1 --mode=fullnat --bgp
create_lb_rule llb2 20.20.20.1 --tcp=2020:8080 --endpoints=31.31.31.1:1,32.32.32.1:1,33.33.33.1:1 --mode=fullnat --bgp

# keepalive will take few seconds to be UP and running with valid states
sleep 60

# K3s setup

if [ "$1" ]; then
  KUBECONFIG="$1"
fi

# If k3s setup exists, skip installation
if [[ -f "/usr/local/bin/k3s-uninstall.sh" ]]; then
  echo "K3s exists"
  sleep 10
else
  echo "Start K3s installation"

  # Install k3s without external cloud-manager and disabled servicelb
  #curl -sfL https://get.k3s.io | INSTALL_K3S_VERSION=v1.22.9+k3s1 INSTALL_K3S_EXEC="server --disable traefik --disable servicelb --disable-cloud-controller --kubelet-arg cloud-provider=external" K3S_KUBECONFIG_MODE="644" sh -
  curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="server --disable traefik --disable servicelb --disable-cloud-controller  --disable "metrics-server" --kubelet-arg cloud-provider=external --flannel-backend=none --cluster-cidr=10.42.0.0/16" K3S_KUBECONFIG_MODE="644" sh -

  sleep 10

  # Install Cilium
  CILIUM_CLI_VERSION=$(curl -s https://raw.githubusercontent.com/cilium/cilium-cli/master/stable.txt)
  CLI_ARCH=amd64
  if [ "$(uname -m)" = "aarch64" ]; then CLI_ARCH=arm64; fi
  curl -L --fail --remote-name-all https://github.com/cilium/cilium-cli/releases/download/${CILIUM_CLI_VERSION}/cilium-linux-${CLI_ARCH}.tar.gz{,.sha256sum}
  sha256sum --check cilium-linux-${CLI_ARCH}.tar.gz.sha256sum
  sudo tar xzvfC cilium-linux-${CLI_ARCH}.tar.gz /usr/local/bin
  rm cilium-linux-${CLI_ARCH}.tar.gz{,.sha256sum}

  sudo cat /etc/rancher/k3s/k3s.yaml > ~/.kube/config
  cilium install
  #kubectl create -f https://raw.githubusercontent.com/cilium/cilium/v1.9/install/kubernetes/quick-install.yaml
  sleep 60

  # Check kubectl works
  kubectl $KUBECONFIG get pods -A

  # Remove taints in k3s if any (usually happens if started without cloud-manager)
  kubectl $KUBECONFIG taint nodes --all node.cloudprovider.kubernetes.io/uninitialized=false:NoSchedule-

  echo "End K3s installation"
fi

# Install Bird to work with k3s
sudo apt install bird2 --yes

sleep 5

sudo cp -f bird_config/bird.conf /etc/bird/bird.conf
if [ ! -f  /var/log/bird.log ]; then
  sudo touch /var/log/bird.log
fi
sudo chown bird:bird /var/log/bird.log
sudo systemctl restart bird

sleep 10

# Start nginx pods and services for test(using kube-loxilb)
kubectl $KUBECONFIG apply -f kube-loxilb.yml
sleep 15
kubectl $KUBECONFIG apply -f nginx-svc-lb1.yml
sleep 10
kubectl $KUBECONFIG apply -f udp-svc-lb.yml
sleep 10
#kubectl $KUBECONFIG apply -f sctp-svc-lb.yml
#sleep 30

# External LB service must be created by now
kubectl $KUBECONFIG get svc


