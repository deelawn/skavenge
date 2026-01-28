#!/bin/bash
# Setup script for staging deployment host
# Run this on your Digital Ocean droplet as root
#
# Usage: sudo ./setup-staging-host.sh [REPO_URL]
#
# Example:
#   sudo ./setup-staging-host.sh git@github.com:deelawn/skavenge.git

set -e

# Configuration
DEPLOY_USER="deploy"
REPO_URL="${1:-git@github.com:deelawn/skavenge.git}"
APP_DIR="/home/${DEPLOY_USER}/skavenge"

echo "=== Skavenge Staging Host Setup ==="
echo ""

# Check if running as root
if [ "$EUID" -ne 0 ]; then
    echo "Error: Please run as root (sudo ./setup-staging-host.sh)"
    exit 1
fi

# 1. Create deployment user
echo "[1/7] Creating deployment user..."
if ! id "$DEPLOY_USER" &>/dev/null; then
    useradd -m -s /bin/bash "$DEPLOY_USER"
    echo "  User $DEPLOY_USER created"
else
    echo "  User $DEPLOY_USER already exists"
fi

# 2. Install Docker if not present
echo "[2/7] Checking Docker installation..."
if ! command -v docker &>/dev/null; then
    echo "  Installing Docker..."
    curl -fsSL https://get.docker.com | sh
    systemctl enable docker
    systemctl start docker
    echo "  Docker installed"
else
    echo "  Docker already installed"
fi

# 3. Add deploy user to docker group
echo "[3/7] Adding $DEPLOY_USER to docker group..."
usermod -aG docker "$DEPLOY_USER"
echo "  Done"

# 4. Install Docker Compose plugin if not present
echo "[4/7] Checking Docker Compose..."
if ! docker compose version &>/dev/null; then
    echo "  Installing Docker Compose plugin..."
    apt-get update -qq
    apt-get install -y docker-compose-plugin
    echo "  Docker Compose installed"
else
    echo "  Docker Compose already installed"
fi

# 5. Generate SSH key for deploy user (for GitHub access)
echo "[5/7] Setting up SSH keys for GitHub access..."
sudo -u "$DEPLOY_USER" bash -c '
    mkdir -p ~/.ssh
    chmod 700 ~/.ssh
    touch ~/.ssh/authorized_keys
    chmod 600 ~/.ssh/authorized_keys
    if [ ! -f ~/.ssh/id_ed25519 ]; then
        ssh-keygen -t ed25519 -f ~/.ssh/id_ed25519 -N "" -C "deploy@staging"
        echo "  SSH key generated"
    else
        echo "  SSH key already exists"
    fi
'

# Add GitHub to known hosts to prevent interactive prompt
sudo -u "$DEPLOY_USER" bash -c '
    ssh-keyscan -t ed25519 github.com >> ~/.ssh/known_hosts 2>/dev/null
'
echo "  GitHub added to known hosts"

# 6. Create app directory structure
echo "[6/7] Creating application directory..."
sudo -u "$DEPLOY_USER" mkdir -p "$APP_DIR"
echo "  Created $APP_DIR"

# 7. Setup firewall (optional)
echo "[7/7] Firewall configuration (optional)..."
if command -v ufw &>/dev/null; then
    echo "  UFW is available. You may want to run:"
    echo "    sudo ufw allow 22/tcp    # SSH"
    echo "    sudo ufw allow 8080/tcp  # Webapp"
    echo "    sudo ufw allow 4591/tcp  # Gateway"
    echo "    sudo ufw allow 3000/tcp  # Admin Portal"
    echo "    sudo ufw allow 4040/tcp  # Indexer API"
    echo "    sudo ufw enable"
else
    echo "  UFW not installed. Consider installing for firewall protection."
fi

echo ""
echo "=========================================="
echo "  MANUAL STEPS REQUIRED"
echo "=========================================="
echo ""
echo "STEP 1: Add Deploy Key to GitHub"
echo "---------------------------------"
echo "Add this public key as a Deploy Key in your GitHub repository:"
echo "  Repository -> Settings -> Deploy keys -> Add deploy key"
echo ""
echo "Title: Staging Server"
echo "Key:"
sudo -u "$DEPLOY_USER" cat /home/${DEPLOY_USER}/.ssh/id_ed25519.pub
echo ""

echo "STEP 2: Generate SSH Key for GitHub Actions"
echo "--------------------------------------------"
echo "On your LOCAL machine, run:"
echo ""
echo "  ssh-keygen -t ed25519 -f ~/.ssh/skavenge-staging -N '' -C 'github-actions-staging'"
echo ""
echo "Then add the PUBLIC key to this host:"
echo ""
echo "  cat ~/.ssh/skavenge-staging.pub | ssh root@$(hostname -I | awk '{print $1}') 'cat >> /home/${DEPLOY_USER}/.ssh/authorized_keys'"
echo ""

echo "STEP 3: Add GitHub Secrets"
echo "--------------------------"
echo "In your GitHub repository, go to:"
echo "  Settings -> Secrets and variables -> Actions"
echo ""
echo "Add these secrets:"
echo "  STAGING_HOST     = $(hostname -I | awk '{print $1}')"
echo "  STAGING_USER     = ${DEPLOY_USER}"
echo "  STAGING_SSH_KEY  = (contents of ~/.ssh/skavenge-staging from your local machine)"
echo "  STAGING_SSH_PORT = 22"
echo ""

echo "STEP 4: Clone the Repository"
echo "----------------------------"
echo "After adding the deploy key, run:"
echo ""
echo "  sudo -u ${DEPLOY_USER} git clone ${REPO_URL} ${APP_DIR}"
echo ""

echo "STEP 5: Create Configuration Files"
echo "-----------------------------------"
echo "Create the deployment config file:"
echo ""
echo "  sudo -u ${DEPLOY_USER} cat > ${APP_DIR}/deploy-config.json << 'EOF'"
echo '  {'
echo '    "deployerPrivateKey": "YOUR_DEPLOYER_ETHEREUM_PRIVATE_KEY",'
echo '    "hardhatUrl": "http://hardhat:8545"'
echo '  }'
echo "  EOF"
echo ""
echo "Create a placeholder webapp config file (will be updated by deploy-contract):"
echo ""
echo "  sudo -u ${DEPLOY_USER} cat > ${APP_DIR}/webapp/config.json << 'EOF'"
echo '  {'
echo '    "contractAddress": "0x0000000000000000000000000000000000000000",'
echo '    "networkRpcUrl": "http://hardhat:8545",'
echo '    "chainId": 1337,'
echo '    "gatewayUrl": "http://gateway:4591"'
echo '  }'
echo "  EOF"
echo ""
echo "Optionally create mint-config.json for minting clues."
echo ""

echo "STEP 6: Initial Deployment"
echo "--------------------------"
echo "Run the first deployment manually:"
echo ""
echo "  sudo -u ${DEPLOY_USER} bash -c 'cd ~/skavenge && docker compose -f docker-compose.staging.yml up -d'"
echo ""

echo "=========================================="
echo "  Setup script complete!"
echo "=========================================="
