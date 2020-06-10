#!/bin/bash

INSTALL_DIR=`eval echo ~$USER`
TEANITY_DIR=".teanity-quickstart"

main() {
  ensureGit
  ensureCorrectDirectory

  echo "Cloning repository..."
  cloneRepo

  echo "Creating links..."
  ensureLinks
  ensureExecutable
  ensureInstalled
}

ensureGit() {
  command -v git >/dev/null 2>&1 || {
    echo "Error: install git first"
    exit 1
  }
}

ensureCorrectDirectory() {
  cd "$INSTALL_DIR"
}

cloneRepo() {
  git clone --depth=1 https://github.com/skoumalcz/teanity-quickstart.git $TEANITY_DIR --quiet
}

ensureLinks() {
  CURRENT_DIR="$INSTALL_DIR/$TEANITY_DIR"
  cd $CURRENT_DIR
  rm -rf /usr/local/bin/quickstart
  ln quickstart /usr/local/bin/quickstart
}

ensureExecutable() {
  chmod +x quickstart
}

ensureInstalled() {
  if ! command -v quickstart >/dev/null 2>&1; then
    echo "Failed to install. There's literally nothing you can do."
    exit
  else
    echo "Successfully installed!"
    echo "Use quickstart command to create new project"
  fi
}

main
