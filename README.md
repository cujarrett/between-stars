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

CI needs a `VSCE_PAT` repo secret.

1. Create an Azure DevOps token with the scope **Marketplace → Manage**. See [Get a Personal Access Token](https://code.visualstudio.com/api/working-with-extensions/publishing-extension#get-a-personal-access-token).
2. Save it: `gh secret set VSCE_PAT -R cujarrett/between-stars`

</details>
