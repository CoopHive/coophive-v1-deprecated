# Resource Provider Instructions

https://docs.docker.com/engine/install/ubuntu/

https://docs.docker.com/engine/install/linux-postinstall/

https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/latest/install-guide.html


```
curl -fsSL https://nvidia.github.io/libnvidia-container/gpgkey | sudo gpg --dearmor -o /usr/share/keyrings/nvidia-container-toolkit-keyring.gpg \
  && curl -s -L https://nvidia.github.io/libnvidia-container/stable/deb/nvidia-container-toolkit.list | \
    sed 's#deb https://#deb [signed-by=/usr/share/keyrings/nvidia-container-toolkit-keyring.gpg] https://#g' | \
    sudo tee /etc/apt/sources.list.d/nvidia-container-toolkit.list
```

```
sudo apt-get update
```

```
sudo apt-get install -y nvidia-container-toolkit
```

```
sudo nvidia-ctk runtime configure --runtime=docker
```

```
sudo systemctl restart docker
```


https://docs.nvidia.com/cuda/cuda-installation-guide-linux/


```
gcc --version
```

```
sudo apt-get install linux-headers-$(uname -r)
```

```
sudo apt-key del 7fa2af80
```

```
sudo dpkg -i cuda-repo-<distro>_<version>_<architecture>.deb
```

assuming ubuntu 22.04

```
sudo dpkg -i cuda-repo-ubuntu_2204__x86_64.deb

```

```
sudo cp /var/cuda-repo-ubuntu2204-X-Y-local/cuda-*-keyring.gpg /usr/share/keyrings/
```
```
wget https://developer.download.nvidia.com/compute/cuda/repos/ubuntu2204/x86_64/cuda-ubuntu.pin
 sudo mv cuda-ubuntu.pin /etc/apt/preferences.d/cuda-repository-pin-600
```

```
sudo apt-get update

```

```
sudo apt-get install cuda-toolkit
```

```
sudo reboot
```


https://docs.nvidia.com/cuda/cuda-installation-guide-linux/index.html#post-installation-actions

export PATH=/usr/local/cuda-12.4/bin${PATH:+:${PATH}}


https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/latest/sample-workload.html


```
sudo docker run --rm --runtime=nvidia --gpus all ubuntu nvidia-smi
```
