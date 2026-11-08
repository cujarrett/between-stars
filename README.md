<p align="center"><img src="images/icon.png" width="120" alt="Between Stars icon"></p>

<h1 align="center">Between Stars</h1>

<p align="center">A dark VS Code theme. Purple-black space with a green glow.</p>

![Between Stars in VS Code](images/screenshot.png)

## Install

```sh
code --install-extension cujarrett.between-stars
```

Then run **Preferences: Color Theme** and pick **Between Stars**.

## Credits

Syntax colors come from [Dracula Purple](https://github.com/cbasdev/dracula-purple-theme) by cbasdev (MIT).

<details>
<summary>Publishing</summary>

CI publishes to the Marketplace on every push to `main`. A release is a version bump in `package.json` plus a `CHANGELOG.md` entry.

CI needs a `VSCE_PAT` repo secret. To make one:

1. Open [dev.azure.com/cujarrett/_usersSettings/tokens](https://dev.azure.com/cujarrett/_usersSettings/tokens), signed in as the owner of the `cujarrett` publisher.
2. Click **+ New Token** and fill it in.
   - Name: `between-stars publish`
   - Organization: **All accessible organizations**
   - Expiration: the longest it allows
   - Scopes: click **Show all scopes**, then tick **Marketplace → Manage**
3. Click **Create** and copy the token. Azure shows it only once.
4. Check it works. Paste the token when asked. Success prints `The Personal Access Token verification succeeded for the publisher 'cujarrett'.`

   ```sh
   npx @vscode/vsce verify-pat cujarrett
   ```

5. Save it as the repo secret. Paste the token when asked.

   ```sh
   gh secret set VSCE_PAT -R cujarrett/between-stars
   ```

</details>
