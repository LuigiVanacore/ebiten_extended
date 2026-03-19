---
description: commit all staged changes with a descriptive message
---

// turbo-all

1. Check git status to see all changed files
```
git -C c:\GoProject\ebiten_extended status --short
```

2. Stage all modified and new files
```
git -C c:\GoProject\ebiten_extended add -A
```

3. Commit with a structured message summarising the changes (use imperative form, include file/package names affected, keep it under 72 chars per line)
```
git -C c:\GoProject\ebiten_extended commit -m "<message>"
```
Replace `<message>` with a concise, imperative summary of what was changed before running this step. Example: `feat(ui): add AnchorLayout and remove deprecated LayoutEx types`
