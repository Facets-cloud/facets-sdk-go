# Quick Start: GitHub Actions Setup

## 1. Configure Secrets (5 minutes)

Go to your repository **Settings** → **Secrets and variables** → **Actions**, then add:

| Secret Name | Required? | Purpose | Example Value |
|------------|-----------|---------|---------------|
| `FACETS_API_HOST` | No | API endpoint hostname | `facetsdemo.console.facets.cloud` |
| `FACETS_USERNAME` | No* | Username for auth | `your-username` |
| `FACETS_API_TOKEN` | No* | API token | `your-secret-token` |

*Required only if you want to run integration tests

**Using GitHub CLI:**
```bash
gh secret set FACETS_API_HOST --body "facetsdemo.console.facets.cloud"
gh secret set FACETS_USERNAME --body "your-username"
gh secret set FACETS_API_TOKEN --body "your-token"
```

## 2. Enable GitHub Actions

1. Go to **Settings** → **Actions** → **General**
2. Under **Actions permissions**, select:
   - ✅ Allow all actions and reusable workflows
3. Under **Workflow permissions**, select:
   - ✅ Read and write permissions
4. Click **Save**

## 3. Test the Setup

### Option A: Manual Trigger
1. Go to **Actions** tab
2. Click **Manual SDK Update** workflow
3. Click **Run workflow**
4. Select `patch` and click **Run workflow**

### Option B: Wait for Weekly Run
- Runs automatically every Monday at 9 AM UTC
- Check **Actions** tab to see the run

## 4. Verify Success

After the workflow runs:

1. ✅ Check **Actions** tab for green checkmark
2. ✅ Check **Releases** for new version
3. ✅ Check **Tags** for new version tag
4. ✅ Try: `go get github.com/Facets-cloud/facets-sdk-go@latest`

## Common Issues

### "Permission denied" error
- Fix: Enable write permissions in **Settings** → **Actions** → **General**

### Integration test fails
- Expected if `FACETS_USERNAME` or `FACETS_API_TOKEN` not set
- Workflow continues anyway (integration test is optional)

### No new release created
- Expected if OpenAPI spec hasn't changed
- Check workflow logs: "No changes detected in OpenAPI spec"

## What Happens Automatically?

```
Every Monday at 9 AM UTC:
├─ Download latest OpenAPI spec
├─ Check for changes
├─ If changed:
│  ├─ Generate new SDK code
│  ├─ Run integration test (if credentials provided)
│  ├─ Commit to main branch
│  ├─ Create new version tag (auto-increment)
│  └─ Publish GitHub release
└─ If no changes: Skip release
```

## Need Help?

- 📖 Full docs: [AUTOMATION.md](../AUTOMATION.md)
- 🔐 Secrets guide: [SECRETS.md](SECRETS.md)
- 🐛 Issues: [GitHub Issues](../../issues)
