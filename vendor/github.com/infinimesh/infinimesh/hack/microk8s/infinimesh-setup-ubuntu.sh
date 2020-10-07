#!/bin/sh
clear

# create infinimesh directory
mkdir -p ~/infinimesh-local && cd ~/infinimesh-local

# check if we on Linux
 if [[ "$OSTYPE" == "linux-gnu" ]]; then
	 echo "Linux OS found, proceed ..."
     else
	 echo "no Linux environment, aborting ..."
       	 exit 0
fi

# check if kubectl is installed
if which kubectl >/dev/null; then
        echo " kubectl found, proceed .."
    else
	echo " kubectl not found, please install:" \
	echo " for Ubuntu: sudo snap install kubectl --classic"
        exit 0
 fi
# creating RND file for secure cert creation
dd if=/dev/urandom of=randfile bs=256 count=1

echo "=> everything ready, let's start"
printf '\n'
# install microk8s
sudo snap install microk8s --classic --channel=1.13/stable
sleep 30

# set sudo for accessing k8s > 1.14
sudo usermod -a -G microk8s multipass

sudo iptables -P FORWARD ACCEPT
sleep 15

/snap/bin/microk8s.enable dns dashboard ingress registry
sleep 5

/snap/bin/microk8s.enable storage
sleep 10

mkdir -p ${HOME}/.kube
/snap/bin/microk8s.config > ${HOME}/.kube/config

# retrieve token
token=$(kubectl -n kube-system get secret | grep default-token | cut -d " " -f1)
echo `kubectl -n kube-system describe secret $token` > .k8stoken
kubectl -n kube-system describe secret $token

# check if we can reach kubernetes and abort if not
 if kubectl cluster-info; then
	echo "kubernetes running ..."
    else
	echo "something went wrong, check the logs, aborting "
	exit 0
 fi

# add standard storage class for postgres op
echo " enable standard storage class and patch to non-default"
kubectl apply -f https://raw.githubusercontent.com/infinimesh/infinimesh/master/hack/microk8s/storage.yaml
sleep 15

kubectl patch storageclass standard -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"false"}}}'
sleep 10

echo " creating self - signed certificates "
printf '\n'
mkdir -p certs && cd certs
openssl genrsa -out ca.key 4096
openssl req -subj '/CN=infinimesh.local/O=Infinimesh' -new -x509 -sha256 -key ca.key -out ca.crt -days 3650
openssl genrsa -out apiserver_grpc.key 4096
openssl req -subj /CN=grpc.api.infinimesh.local -out apiserver_grpc.csr -key apiserver_grpc.key -new
openssl x509 -req -days 3650 -in apiserver_grpc.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out apiserver_grpc.crt -sha256
openssl genrsa -out apiserver_rest.key 4096
openssl req -subj /CN=api.infinimesh.local -out apiserver_rest.csr -key apiserver_rest.key -new
openssl x509 -req -days 3650 -in apiserver_rest.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out apiserver_rest.crt -sha256 
openssl genrsa -out mqtt_bridge.key 4096
openssl req -subj /CN=mqtt.api.infinimesh.local -out mqtt_bridge.csr -key mqtt_bridge.key -new
openssl x509 -req -days 3650 -in mqtt_bridge.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out mqtt_bridge.crt -sha256 
openssl genrsa -out app.key 4096
openssl req -subj /CN=app.infinimesh.local -out app.csr -key app.key -new
openssl x509 -req -days 3650 -in app.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out app.crt -sha256
sleep 3

echo " install the certificates"
printf '\n'
kubectl create secret tls apiserver-grpc-tls --cert apiserver_grpc.crt --key apiserver_grpc.key 
kubectl create secret tls apiserver-rest-tls --cert apiserver_rest.crt --key apiserver_rest.key 
kubectl create secret tls mqtt-bridge-tls --cert mqtt_bridge.crt --key mqtt_bridge.key 
kubectl create secret tls app-tls --cert app.crt --key app.key 
cd -

echo " installing infinimesh operator "
printf '\n'
kubectl apply -f https://raw.githubusercontent.com/infinimesh/operator/master/manifests/crd.yaml
kubectl apply -f https://raw.githubusercontent.com/infinimesh/operator/master/manifests/operator.yaml
sleep 2

echo " installing kubeDB from https://github.com/kubedb "
printf '\n'
curl -fsSL https://github.com/kubedb/installer/raw/v0.13.0-rc.0/deploy/kubedb.sh | bash
sleep 5

echo " installing Kafka from https://github.com/strimzi "
printf '\n'
kubectl create namespace kafka &&
curl -L https://github.com/strimzi/strimzi-kafka-operator/releases/download/0.11.1/strimzi-cluster-operator-0.11.1.yaml \
  | sed 's/namespace: .*/namespace: kafka/' \
  | kubectl -n kafka apply -f -
printf '\n'
echo "=> now we install infinimesh ..."
printf '\n'

kubectl apply -f https://raw.githubusercontent.com/InfiniteDevices/infinimesh/master/hack/microk8s/postgres-dev.yaml
sleep 30
kubectl apply -f https://raw.githubusercontent.com/infinimesh/infinimesh/master/hack/microk8s/infinimesh-platform.yaml
kubectl apply -f https://raw.githubusercontent.com/infinimesh/infinimesh/master/hack/microk8s/infinimesh-kafka.yaml -n kafka

# getting IP and add hosts entries
echo " checking for host entries"
printf '\n'

IP=`sudo hostname -i|grep -E -o "([0-9]{1,3}[\.]){3}[0-9]{1,3}"`
if ! grep -q infinimesh.local "/etc/hosts"; then
echo "=> please add host entries into /etc/hosts: "
printf '\n'
cat <<EOL
# infinimesh-dev
$IP grpc.api.infinimesh.local
$IP api.infinimesh.local
$IP app.infinimesh.local
$IP mqtt.api.infinimesh.local
EOL
   else
        echo "=> host entries already present"
fi
printf '\n'

echo "=> installing inf (infinimesh CLI) and point to the local setup:"
curl -L https://bit.ly/2CNKWzJ | BINDIR=$HOME/bin bash  
echo "inf CLI installed in" $HOME/bin

# check if inf is already in $PATH
if ! grep -q HOME/bin ~/.profile; then
echo 'export PATH=$HOME/bin:$PATH' >> ~/.profile
else
        echo "=> PATH already set"
fi
printf '\n'

echo "=> set CLI context"
~/bin/inf config set-context local --apiserver grpc.api.infinimesh.local:443 --tls=true --ca-file ~/infinimesh-local/certs/ca.crt

printf '\n'
echo "wait a few secs to get all things in place ...."
secs=$((90))
while [ $secs -gt 0 ]; do
   echo -ne "$secs\033[0K\r"
   sleep 1
   : $((secs--))
done
printf '\n'

# export kubeconfig again, to be sure
export KUBECONFIG=$KUBECONFIG:~/kubeconfig

echo "infinimesh is now ready, point your browser to app.infinimesh.local or use our CLI"
printf '\n'
echo "your master user credentials are: "
echo +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
echo "root" 
kubectl get secret my-infinimesh-root-account -o=jsonpath='{.data.password}' | base64 -d
printf '\n'
echo +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
printf '\n'

echo "To trust the root certificate, you must go to your browser settings and add the file ca.crt as an certificate Authority." 
echo "This works best with Firefox or Safari, we encounter some issues with Chrome."
echo "In Firefox open "about:preferences#privacy" and import ca.crt from ~/infinimesh-local/certs/ into Certificates." 
echo "Please read our documention under https://infinimesh.github.io/infinimesh/docs/#/ to proceed with adding devices and sending data. Happy IoTing ..."
printf '\n'
echo "starting infinimesh CLI:"
printf '\n'
~/bin/inf login
