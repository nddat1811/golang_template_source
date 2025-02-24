# FULL_TAG=$(git rev-parse --short HEAD)
PC_NAME=$(hostname)
TAG=$(date +%Y%m%d%H%M%S)
FULL_TAG="${PC_NAME}-${TAG}"

# Information about Docker registry and image
REGISTRY_URL="10.39.125.26:8000"
IMAGE_NAME="mbf_platform_be"
# IMAGE_NAME="mgolf_be"

# Build image Docker without cache
# docker build --no-cache -t ${IMAGE_NAME}:latest .
docker build  -t ${IMAGE_NAME}:latest .

# Attach image with tag (full tag)
docker tag ${IMAGE_NAME}:latest ${REGISTRY_URL}/${IMAGE_NAME}:${FULL_TAG}

# Push image to registry with tag
docker push ${REGISTRY_URL}/${IMAGE_NAME}:${FULL_TAG}

# Attach image with tag -latest
docker tag ${IMAGE_NAME}:latest ${REGISTRY_URL}/${IMAGE_NAME}:latest

# Push image to registry with tag 'latest'
docker push ${REGISTRY_URL}/${IMAGE_NAME}:latest