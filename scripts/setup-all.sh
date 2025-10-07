#!/bin/bash

# Script hoàn chỉnh: Từ thêm SSH key đến import vào GitHub
# Sử dụng: ./scripts/setup-all.sh

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
NC='\033[0m'

print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

print_step() {
    echo -e "${PURPLE}[STEP]${NC} $1"
}

# Kiểm tra xem có đang ở trong git repository không
if [ ! -d ".git" ]; then
    print_error "Không tìm thấy .git directory. Hãy chạy script này từ root của git repository."
    exit 1
fi

# Kiểm tra file secret.txt
SECRET_FILE="scripts/secret.txt"
if [ ! -f "$SECRET_FILE" ]; then
    print_error "Không tìm thấy file $SECRET_FILE"
    exit 1
fi

print_step "🚀 Bắt đầu quy trình hoàn chỉnh: SSH Key → GitHub Secrets"

# Function để tìm SSH keys trong ~/.ssh
find_ssh_keys() {
    local ssh_dir="$HOME/.ssh"
    local keys=()
    
    if [ ! -d "$ssh_dir" ]; then
        print_warning "Thư mục ~/.ssh không tồn tại"
        return 1
    fi
    
    # Tìm tất cả private keys (bao gồm cả custom names)
    for key in "$ssh_dir"/*; do
        if [ -f "$key" ] && [[ "$key" != *".pub" ]] && [[ "$key" != *"known_hosts" ]] && [[ "$key" != *"config" ]]; then
            # Kiểm tra xem có phải SSH key không
            if ssh-keygen -l -f "$key" &>/dev/null; then
                keys+=("$key")
            fi
        fi
    done
    
    echo "${keys[@]}"
}

# Function để hiển thị nội dung SSH key
show_ssh_key() {
    local key_file="$1"
    echo ""
    print_info "Nội dung file: $key_file"
    echo "----------------------------------------"
    cat "$key_file"
    echo "----------------------------------------"
    echo ""
}

# Function để tạo SSH key mới
create_ssh_key() {
    local key_name="$1"
    local ssh_dir="$HOME/.ssh"
    
    print_info "Tạo SSH key mới: $key_name"
    
    # Tạo thư mục .ssh nếu chưa có
    mkdir -p "$ssh_dir"
    chmod 700 "$ssh_dir"
    
    # Tạo key
    ssh-keygen -t rsa -b 4096 -f "$ssh_dir/$key_name" -N ""
    
    # Set permissions
    chmod 600 "$ssh_dir/$key_name"
    chmod 644 "$ssh_dir/$key_name.pub"
    
    print_success "Đã tạo SSH key: $ssh_dir/$key_name"
}

# Function để thêm secret vào GitHub
add_secret() {
    local secret_name=$1
    local secret_value=$2
    
    if [ -z "$secret_value" ]; then
        print_warning "Secret $secret_name trống, bỏ qua..."
        return
    fi
    
    print_info "Thêm secret: $secret_name"
    echo "$secret_value" | gh secret set "$secret_name" --repo "$FULL_REPO_NAME" --body -
    print_success "Đã thêm secret: $secret_name"
}

# Bước 1: Tìm SSH keys hiện có
print_step "1. Tìm SSH keys hiện có trong ~/.ssh"

ssh_keys=($(find_ssh_keys))
if [ ${#ssh_keys[@]} -eq 0 ]; then
    print_warning "Không tìm thấy SSH keys nào trong ~/.ssh"
    ssh_keys_found=false
else
    print_info "Tìm thấy ${#ssh_keys[@]} SSH key(s):"
    for i in "${!ssh_keys[@]}"; do
        echo "  $((i+1)). ${ssh_keys[$i]}"
    done
    ssh_keys_found=true
fi

# Bước 2: Cho phép user chọn SSH key
selected_key=""
if [ "$ssh_keys_found" = true ]; then
    echo ""
    print_info "Bạn có muốn sử dụng một trong các SSH keys trên không?"
    for i in "${!ssh_keys[@]}"; do
        echo "  $((i+1)). ${ssh_keys[$i]}"
        read -p "    Hiển thị nội dung? (y/n): " show_content
        if [[ $show_content =~ ^[Yy]$ ]]; then
            show_ssh_key "${ssh_keys[$i]}"
        fi
    done
    
    echo ""
    read -p "Chọn SSH key để sử dụng (1-${#ssh_keys[@]}) hoặc Enter để bỏ qua: " choice
    
    if [[ $choice =~ ^[0-9]+$ ]] && [ $choice -ge 1 ] && [ $choice -le ${#ssh_keys[@]} ]; then
        selected_key="${ssh_keys[$((choice-1))]}"
        print_success "Đã chọn: $selected_key"
    else
        print_warning "Bỏ qua tất cả SSH keys hiện có"
    fi
fi

# Bước 3: Hỏi có muốn tạo SSH key mới không
if [ -z "$selected_key" ]; then
    echo ""
    read -p "Bạn có muốn tạo SSH key mới không? (y/n): " create_new
    
    if [[ $create_new =~ ^[Yy]$ ]]; then
        read -p "Nhập tên file SSH key (mặc định: id_rsa): " key_name
        key_name=${key_name:-id_rsa}
        
        create_ssh_key "$key_name"
        selected_key="$HOME/.ssh/$key_name"
        
        # Hiển thị nội dung key mới tạo
        show_ssh_key "$selected_key"
        show_ssh_key "$selected_key.pub"
        
        print_info "Public key (để thêm vào server):"
        cat "$selected_key.pub"
        echo ""
        print_warning "Hãy thêm public key vào ~/.ssh/authorized_keys trên server!"
    fi
fi

# Bước 4: Cập nhật file secret.txt
if [ -n "$selected_key" ]; then
    print_step "2. Cập nhật file secret.txt"
    
    # Đọc nội dung SSH key
    ssh_key_content=$(cat "$selected_key")
    
    # Cập nhật file secret.txt
    print_info "Cập nhật SSH_PRIVATE_KEY trong $SECRET_FILE"
    
    # Tạo file tạm thời
    temp_file=$(mktemp)
    
    # Đọc file secret.txt và cập nhật SSH_PRIVATE_KEY
    while IFS= read -r line; do
        if [[ $line == SSH_PRIVATE_KEY=* ]]; then
            echo "SSH_PRIVATE_KEY=\"$ssh_key_content\"" >> "$temp_file"
        else
            echo "$line" >> "$temp_file"
        fi
    done < "$SECRET_FILE"
    
    # Thay thế file gốc
    mv "$temp_file" "$SECRET_FILE"
    
    print_success "Đã cập nhật SSH_PRIVATE_KEY trong $SECRET_FILE"
else
    print_warning "Không có SSH key được chọn, giữ nguyên SSH_PRIVATE_KEY trong $SECRET_FILE"
fi

# Bước 5: Import vào GitHub
print_step "3. Import secrets vào GitHub"

# Kiểm tra GitHub CLI
if ! command -v gh &> /dev/null; then
    print_error "GitHub CLI (gh) chưa được cài đặt. Vui lòng cài đặt trước:"
    echo "  macOS: brew install gh"
    echo "  Ubuntu/Debian: sudo apt install gh"
    exit 1
fi

# Kiểm tra đăng nhập GitHub
if ! gh auth status &> /dev/null; then
    print_error "Chưa đăng nhập GitHub CLI. Vui lòng chạy: gh auth login"
    exit 1
fi

# Lấy repository name
REPO_OWNER=$(git config --get remote.origin.url | sed -n 's/.*github.com[:/]\([^/]*\)\/[^/]*.*/\1/p')
REPO_NAME=$(basename -s .git $(git config --get remote.origin.url))
FULL_REPO_NAME="$REPO_OWNER/$REPO_NAME"

print_info "Repository: $FULL_REPO_NAME"
print_info "Bắt đầu import secrets từ $SECRET_FILE..."

# Đọc file secret.txt và import secrets
SSH_PRIVATE_KEY=""
SSH_HOST=""
SSH_USERNAME=""
CONTAINER_REGISTRY_PAT=""

while IFS= read -r line; do
    if [[ $line == SSH_PRIVATE_KEY=* ]]; then
        SSH_PRIVATE_KEY="${line#SSH_PRIVATE_KEY=}"
        # Loại bỏ dấu ngoặc kép nếu có
        SSH_PRIVATE_KEY="${SSH_PRIVATE_KEY%\"}"
        SSH_PRIVATE_KEY="${SSH_PRIVATE_KEY#\"}"
    elif [[ $line == SSH_HOST=* ]]; then
        SSH_HOST="${line#SSH_HOST=}"
    elif [[ $line == SSH_USERNAME=* ]]; then
        SSH_USERNAME="${line#SSH_USERNAME=}"
    elif [[ $line == CONTAINER_REGISTRY_PAT=* ]]; then
        CONTAINER_REGISTRY_PAT="${line#CONTAINER_REGISTRY_PAT=}"
    fi
done < "$SECRET_FILE"

# Thêm SSH_PRIVATE_KEY
if [ -n "$SSH_PRIVATE_KEY" ]; then
    add_secret "SSH_PRIVATE_KEY" "$SSH_PRIVATE_KEY"
fi

# Thêm SSH_HOST
if [ -n "$SSH_HOST" ]; then
    add_secret "SSH_HOST" "$SSH_HOST"
fi

# Thêm SSH_USERNAME
if [ -n "$SSH_USERNAME" ]; then
    add_secret "SSH_USERNAME" "$SSH_USERNAME"
fi

# Thêm CONTAINER_REGISTRY_PAT
if [ -n "$CONTAINER_REGISTRY_PAT" ]; then
    add_secret "CONTAINER_REGISTRY_PAT" "$CONTAINER_REGISTRY_PAT"
fi

# Bước 6: Hiển thị kết quả
print_step "4. Kiểm tra kết quả"

print_success "Hoàn tất! Tất cả secrets đã được import vào GitHub."
print_info "Bạn có thể kiểm tra tại: https://github.com/$FULL_REPO_NAME/settings/secrets/actions"

# Hiển thị public key nếu có
if [ -n "$selected_key" ]; then
    public_key_file="$selected_key.pub"
    if [ -f "$public_key_file" ]; then
        echo ""
        print_info "📋 SSH Public Key (để thêm vào server):"
        echo "----------------------------------------"
        cat "$public_key_file"
        echo "----------------------------------------"
        echo ""
        print_warning "⚠️  Hãy thêm public key vào ~/.ssh/authorized_keys trên server!"
    fi
fi

echo ""
print_info "🎯 Bây giờ bạn có thể:"
echo "  - Chạy ./scripts/git-tag.sh để deploy"
echo "  - Chạy ./scripts/check-secrets.sh để kiểm tra secrets"
echo ""
print_info "📝 Lưu ý:"
echo "  - Secrets sẽ được mã hóa và chỉ hiển thị dưới dạng ***"
echo "  - Chỉ có quyền admin mới có thể xem/sửa secrets"
echo "  - Secrets sẽ được sử dụng trong workflow CI/CD" 