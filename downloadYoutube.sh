#!/bin/bash
# Downloads a video from Youtube
# Usage: downloadYoutube.sh <url>

url="$1"

if [ -z "url" ]; then
  echo "Usage: downloadYoutube.sh <url>"
  exit 1
fi

if ! type "brew" >/dev/null 2>&1; then
  echo "Downloading Homebrew"
  NONINTERACTIVE=1 /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
  [ $? -ne 0 ] && echo "Failed to install Homebrew" && exit 1
fi

if ! type "ffmpeg" >/dev/null 2>&1; then
  echo "Downloading ffmpeg"
  brew install ffmpeg
  [ $? -ne 0 ] && echo "Failed to install ffmpeg" && exit 1
fi

if ! type "yt-dlp" >/dev/null 2>&1; then
  echo "Downloading yt-dlp"
  brew install yt-dlp
  [ $? -ne 0 ] && echo "Failed to install yt-dlp" && exit 1
fi

echo "Updating yt-dlp"
brew upgrade yt-dlp

pushd ~/Downloads

yt-dlp -o "%(title)s.%(ext)s" --restrict-filenames -f mp4 "$url"
[ $? -ne 0 ] && echo "Failed to download video" && exit 1

echo "Download is complete. Check your download folder"
open ~/Downloads

popd
