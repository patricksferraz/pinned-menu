# Kubernetes Deployment Instructions 🚀

This document provides detailed instructions for deploying the Pinned Menu application to Kubernetes.

## 📋 Prerequisites

- Kubernetes cluster up and running
- `kubectl` configured and connected to your cluster
- Docker registry access (if using private images)
- Basic understanding of Kubernetes concepts

## 🔐 Secrets Management

### 1. Application Secrets

1. Create a `.env` file in the `k8s` directory using the provided `.env.example` as a template:
   ```bash
   cp .env.example k8s/.env
   ```

2. Fill in all required environment variables in `k8s/.env`

3. Create the Kubernetes secret:
   ```bash
   kubectl create secret generic menu-secret --from-env-file k8s/.env
   ```

### 2. Docker Registry Secret

If you're using a private Docker registry, create a registry secret:

```bash
kubectl create secret docker-registry regsecret \
  --docker-server=$DOCKER_REGISTRY_SERVER \
  --docker-username=$DOCKER_USER \
  --docker-password=$DOCKER_PASSWORD \
  --docker-email=$DOCKER_EMAIL
```

Required variables:
- `$DOCKER_REGISTRY_SERVER`: Docker registry URL (e.g., docker.io, gcr.io)
- `$DOCKER_USER`: Registry username
- `$DOCKER_PASSWORD`: Registry password
- `$DOCKER_EMAIL`: (Optional) Email address

## 🚀 Deployment

### 1. Deploy All Resources

To deploy all Kubernetes resources at once:
```bash
kubectl apply -f ./k8s
```

### 2. Verify Deployment

Check the status of your deployments:
```bash
# Check pods
kubectl get pods

# Check services
kubectl get services

# Check deployments
kubectl get deployments
```

### 3. Access the Application

After deployment, you can access the application through the configured service:
```bash
# Get the service URL
kubectl get service pinned-menu-service
```

## 🔧 Troubleshooting

### Common Issues

1. **Pods not starting**
   ```bash
   # Check pod logs
   kubectl logs <pod-name>

   # Describe pod for more details
   kubectl describe pod <pod-name>
   ```

2. **Secrets not found**
   ```bash
   # Verify secrets exist
   kubectl get secrets

   # Check secret details
   kubectl describe secret menu-secret
   ```

3. **Image pull errors**
   ```bash
   # Check if registry secret is properly configured
   kubectl get secret regsecret --output=yaml
   ```

## 🔄 Maintenance

### Updating the Deployment

1. Update your application code
2. Build and push new Docker image
3. Update the image tag in deployment files
4. Apply changes:
   ```bash
   kubectl apply -f ./k8s
   ```

### Scaling

To scale your application:
```bash
# Scale deployment
kubectl scale deployment pinned-menu --replicas=3
```

## 🧹 Cleanup

To remove all deployed resources:
```bash
kubectl delete -f ./k8s
```

To remove secrets:
```bash
kubectl delete secret menu-secret
kubectl delete secret regsecret
```

## 📚 Additional Resources

- [Kubernetes Documentation](https://kubernetes.io/docs/)
- [Kubectl Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)
- [Kubernetes Best Practices](https://kubernetes.io/docs/concepts/configuration/overview/)

---

Remember to always backup your data and configurations before making significant changes! 🔒
