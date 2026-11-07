# Package the theme into a .vsix. This also validates the manifest.
build:
    python3 -m json.tool themes/between-stars.json > /dev/null
    npx --yes @vscode/vsce package

# Install the packaged build into the local editor to try it.
run: build
    code --install-extension $(ls -t *.vsix | head -1) --force
