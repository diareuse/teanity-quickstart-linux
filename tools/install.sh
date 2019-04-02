INSTALL_DIR="~"
TEANITY_DIR=".teanity-quickstart"

main() {
  ensureGit
  ensureCorrectDirectory

  cloneRepo

  ensureLinks
  ensureInstalled
}

ensureGit() {
  command -v git >/dev/null 2>&1 || {
    echo "Error: install git first"
    exit 1
  }
}

ensureCorrectDirectory() {
  cd $INSTALL_DIR
}

cloneRepo() {
  git clone --depth=1 https://github.com/diareuse/teanity-quickstart-linux.git $TEANITY_DIR
}

ensureLinks() {
  cd "$INSTALL_DIR/$TEANITY_DIR"
  ln -s quickstart /usr/local/bin/quickstart
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
