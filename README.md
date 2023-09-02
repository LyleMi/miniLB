# MiniLoadBalancer

An extremely simplified LoadBalancer for Kubernetes.

## Disclaimer

Please note that this tool is designed for scenarios where you want to *access services inside a Kubernetes cluster from outside the cluster*. It was developed because existing solutions did not meet the requirements. Before using it, please consider if your scenario is similar:

- Avoiding the use of NodePort + affinity configuration due to the lack of fixed nodes.
- Binary services cannot use configurations like Ingress.
- Not wanting to use a paid LoadBalancer provided by a cloud provider.
- Do not want to use the unrestricted ``kubectl proxy``.
- ``kubectl port-forward`` cannot handle service restarts.

## Prerequisites

The binary program needs to be deployed in a location that can access the Kubernetes service (e.g., a Kubernetes node with resolver configured).

If you are not familiar with the configuration, you can follow the steps below to configure it:

```bash
# Get the DNS resolver server address

# For Kubernetes clusters using kube-dns (newer version)
DNS=`kubectl get svc -n kube-system kube-dns -o jsonpath='{.spec.clusterIP}'`

# For Kubernetes clusters using coredns
DNS=`kubectl get svc -n kube-system coredns -o jsonpath='{.spec.clusterIP}'`

# Configure the resolver server
echo "nameserver $DNS" | sudo tee -a /etc/resolv.conf
```

## Usage

```bash
git clone https://github.com/lylemi/minilb
cd minilb
make
cp config-example.yaml config.yaml
./minilb
```

## Contribution

If you have any ideas or suggestions, please feel free to submit a pull request. We appreciate any contributions.

## Contact

If you have any questions or suggestions, please feel free to contact us.
