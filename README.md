# ci-template

### Get latest (CI/CD)
```bash
curl -LO https://github.com/hashmap-kz/ci-template/releases/latest/download/ci-template_linux_amd64.apk
apk add ci-template_linux_amd64.apk --allow-untrusted
```

### Helm
```bash
helm repo add ci-template https://hashmap-kz.github.io/ci-template
helm repo update
helm search repo ci-template
```

