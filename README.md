Requirements:

1. `libvirt-dev` package. `apt install libvirt-dev`
2. Go Package Manager: [Glide](https://github.com/Masterminds/glide)


How to Run: 
1. `go build .`
2. `./libvirt-go-examples`

How to prepare images:

1. Base Image: `wget https://cloud-images.ubuntu.com/bionic/20190703/bionic-server-cloudimg-amd64.img`
2. Prepare User Data Image:
   a. Create user-data file with following content:
     ```
#cloud-config
password: asdfqwer
chpasswd: {expire: False}
ssh_pwauth: True

runcmd:
  - echo "127.0.0.1 kube-cp-01" >> /etc/hosts
  - kubeadm init --pod-network-cidr 10.40.0.0/16 --kubernetes-version 1.14.12. Prepare User Data Image:
     ```
  b. Create meta-data filw with folowing content:
    ```
instance-id: kube-cp-01
local-hostname: kube-cp-01
    ```
