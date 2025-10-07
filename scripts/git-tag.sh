#!/bin/bash

# Đảm bảo tất cả các tags được cập nhật từ remote
git fetch --tags > /dev/null 2>&1

# Lấy tag phiên bản mới nhất từ git (sắp xếp theo phiên bản)
LATEST_GIT_TAG=$(git tag -l "v*" | sort -V | tail -n 1)

# Kiểm tra phiên bản từ GitHub Container Registry (cần có gh CLI)
REPO_OWNER=$(git config --get remote.origin.url | sed -n 's/.*github.com[:/]\([^/]*\)\/[^/]*.*/\1/p')
REPO_NAME=$(basename -s .git "$(git config --get remote.origin.url)")
# Với GHCR, tên package container thường là tên repo
PACKAGE_NAME="$REPO_NAME"

echo "Đang kiểm tra phiên bản từ GitHub Container Registry..."
if command -v gh &> /dev/null; then
  # Sử dụng GitHub CLI để lấy danh sách các package versions nếu có
  if gh auth status &> /dev/null; then
    # Thử truy vấn package theo user (mặc định). Nếu repo thuộc org, người dùng có thể cần đổi sang /orgs/<org>
    GHCR_TAGS=$(gh api "/users/$REPO_OWNER/packages/container/$PACKAGE_NAME/versions" -q '.[] | .metadata.container.tags[]' 2>/dev/null | grep "^v" | sort -V)
    # Fallback endpoint cũ nếu trường hợp trên không khả dụng
    if [ -z "$GHCR_TAGS" ]; then
      GHCR_TAGS=$(gh api "/user/packages/container/$PACKAGE_NAME/versions" -q '.[] | .metadata.container.tags[]' 2>/dev/null | grep "^v" | sort -V)
    fi
    LATEST_GHCR_TAG=$(echo "$GHCR_TAGS" | tail -n 1)
    if [ -n "$LATEST_GHCR_TAG" ]; then
      echo "Phiên bản mới nhất từ GHCR: $LATEST_GHCR_TAG"
    fi
  else
    echo "Bạn chưa đăng nhập GitHub CLI. Bỏ qua kiểm tra từ GHCR."
  fi
else
  echo "GitHub CLI không được cài đặt. Bỏ qua kiểm tra từ GHCR."
fi

# Xác định tag mới nhất từ cả Git và GHCR
if [ -n "$LATEST_GHCR_TAG" ] && [ -n "$LATEST_GIT_TAG" ]; then
  # So sánh phiên bản Git và GHCR để lấy cái mới hơn
  LATEST_VERSION=$(echo -e "$LATEST_GIT_TAG\n$LATEST_GHCR_TAG" | sort -V | tail -n 1)
  echo "Phiên bản mới nhất (từ Git và GHCR): $LATEST_VERSION"
elif [ -n "$LATEST_GIT_TAG" ]; then
  LATEST_VERSION=$LATEST_GIT_TAG
  echo "Phiên bản mới nhất (từ Git): $LATEST_VERSION"
elif [ -n "$LATEST_GHCR_TAG" ]; then
  LATEST_VERSION=$LATEST_GHCR_TAG
  echo "Phiên bản mới nhất (từ GHCR): $LATEST_VERSION"
else
  LATEST_VERSION="v1.0.0"
  echo "Không tìm thấy tag nào. Bắt đầu với $LATEST_VERSION"
fi

## Chuẩn hoá phiên bản và phân tích các thành phần phiên bản
if [[ ! $LATEST_VERSION =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  LATEST_VERSION="v1.0.0"
fi

VERSION_PARTS=(${LATEST_VERSION//./ })
MAJOR=${VERSION_PARTS[0]#v}
MINOR=${VERSION_PARTS[1]:-0}
PATCH=${VERSION_PARTS[2]:-0}

# Đảm bảo giá trị là số nguyên
if ! [[ $MAJOR =~ ^[0-9]+$ ]]; then MAJOR=1; fi
if ! [[ $MINOR =~ ^[0-9]+$ ]]; then MINOR=0; fi
if ! [[ $PATCH =~ ^[0-9]+$ ]]; then PATCH=0; fi

# Tăng phiên bản theo quy tắc: cộng thêm 1 patch; nếu đến 9 thì cuộn lên
if [ "$PATCH" -ge 9 ]; then
  PATCH=0
  if [ "$MINOR" -ge 9 ]; then
    MINOR=0
    MAJOR=$((MAJOR + 1))
  else
    MINOR=$((MINOR + 1))
  fi
else
  PATCH=$((PATCH + 1))
fi

NEW_VERSION="v$MAJOR.$MINOR.$PATCH"

# Kiểm tra tham số --schedule
SCHEDULE_MODE=false
if [ "$1" == "--schedule" ]; then
  SCHEDULE_MODE=true
  echo "🕐 Chế độ scheduled deploy được kích hoạt"
fi

# Tạo tag git mới với annotation phù hợp
echo "Đang tạo tag mới: $NEW_VERSION"
if [ "$SCHEDULE_MODE" == "true" ]; then
  git tag -a "$NEW_VERSION" -m "Release $NEW_VERSION --schedule"
  echo "📅 Tag được tạo với chế độ scheduled deploy"
else
  git tag -a "$NEW_VERSION" -m "Release $NEW_VERSION"
  echo "🚀 Tag được tạo với chế độ immediate deploy"
fi

# Đẩy tags lên remote
echo "Đang đẩy tags lên remote..."
git push --tags

echo "Phiên bản đã được cập nhật thành $NEW_VERSION"
if [ "$SCHEDULE_MODE" == "true" ]; then
  echo "⏰ Sẽ tự động deploy lúc 03:00 UTC hàng ngày"
else
  echo "🚀 Sẽ deploy ngay lập tức"
fi
echo "Quá trình gắn tag phiên bản đã hoàn tất!"