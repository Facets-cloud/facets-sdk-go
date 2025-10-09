# GitHub Secrets Configuration

This document explains how to configure GitHub secrets for the SDK automation workflows.

## Required Secrets

The following secrets need to be configured in your GitHub repository for the automated workflows to function properly.

### 1. FACETS_API_HOST (Optional)

- **Description**: The hostname of your Facets API endpoint
- **Default**: `facetsdemo.console.facets.cloud`
- **Required**: No (uses default if not set)
- **Example**: `facetsdemo.console.facets.cloud` or `production.facets.cloud`

### 2. FACETS_USERNAME (Optional)

- **Description**: Username for Facets API authentication
- **Required**: No (but required for integration tests)
- **Example**: `your-username` or `service-account@facets.cloud`
- **Note**: If not set, integration tests will be skipped

### 3. FACETS_API_TOKEN (Optional)

- **Description**: API token/password for Facets API authentication
- **Required**: No (but required for integration tests)
- **Example**: `your-api-token-here`
- **Note**: If not set, integration tests will be skipped

## How to Add Secrets

### Via GitHub Web UI

1. Go to your repository on GitHub
2. Click **Settings** → **Secrets and variables** → **Actions**
3. Click **New repository secret**
4. Add each secret:

   **FACETS_API_HOST** (optional)
   ```
   Name: FACETS_API_HOST
   Value: facetsdemo.console.facets.cloud
   ```

   **FACETS_USERNAME** (optional, for tests)
   ```
   Name: FACETS_USERNAME
   Value: your-username
   ```

   **FACETS_API_TOKEN** (optional, for tests)
   ```
   Name: FACETS_API_TOKEN
   Value: your-api-token
   ```

5. Click **Add secret**

### Via GitHub CLI

```bash
# Set the API host (optional)
gh secret set FACETS_API_HOST --body "facetsdemo.console.facets.cloud"

# Set username for integration tests (optional)
gh secret set FACETS_USERNAME --body "your-username"

# Set API token for integration tests (optional)
gh secret set FACETS_API_TOKEN --body "your-api-token"
```

## Secret Usage in Workflows

### Weekly Automated Update (`update-sdk.yml`)

```yaml
env:
  FACETS_API_HOST: ${{ secrets.FACETS_API_HOST || 'facetsdemo.console.facets.cloud' }}
  FACETS_USERNAME: ${{ secrets.FACETS_USERNAME }}
  FACETS_API_TOKEN: ${{ secrets.FACETS_API_TOKEN }}
```

- Downloads spec from: `https://${FACETS_API_HOST}/v3/api-docs`
- Runs integration test if username and token are provided
- Works without secrets (uses defaults, skips tests)

### Manual Update (`manual-update.yml`)

Same configuration as weekly update. Allows you to trigger updates on-demand.

## Integration Tests

Integration tests will **only run** if both `FACETS_USERNAME` and `FACETS_API_TOKEN` are configured.

If these secrets are not set:
- Spec download still works (uses public endpoint)
- SDK generation still works
- Integration test is skipped (with `continue-on-error: true`)
- Workflow completes successfully

## Security Best Practices

1. **Never commit secrets** to the repository
2. **Use service accounts** for automation (not personal accounts)
3. **Rotate tokens regularly** (update GitHub secret when you do)
4. **Minimum permissions**: Use API tokens with minimum required permissions
5. **Monitor usage**: Check GitHub Actions logs for unauthorized access attempts

## Updating Secrets

To update a secret:

### Via GitHub Web UI
1. Go to **Settings** → **Secrets and variables** → **Actions**
2. Click the secret name
3. Click **Update secret**
4. Enter new value and click **Update secret**

### Via GitHub CLI
```bash
gh secret set FACETS_API_TOKEN --body "new-token-value"
```

## Troubleshooting

### Workflow fails with "unauthorized" or "forbidden"

- Check that `FACETS_API_TOKEN` is valid and not expired
- Verify `FACETS_USERNAME` is correct
- Ensure the token has required permissions

### Different API endpoint needed

Update `FACETS_API_HOST` secret to point to your endpoint:
```bash
gh secret set FACETS_API_HOST --body "your-custom-host.facets.cloud"
```

### Integration tests not running

This is expected if `FACETS_USERNAME` or `FACETS_API_TOKEN` are not set. The workflow will still succeed and generate the SDK.

To enable integration tests, add both secrets:
```bash
gh secret set FACETS_USERNAME --body "your-username"
gh secret set FACETS_API_TOKEN --body "your-token"
```

## Local Testing

To test locally with the same configuration:

```bash
# Create .env file (do not commit!)
cat > .env << EOF
FACETS_API_HOST=facetsdemo.console.facets.cloud
FACETS_USERNAME=your-username
FACETS_API_TOKEN=your-token
EOF

# Source the variables
export $(cat .env | xargs)

# Run the workflow steps
make clean-spec
make generate-client
make integration-test
```

## Environment-Specific Configurations

If you have multiple environments (staging, production), you can:

1. Use **GitHub Environments** to have different secrets per environment
2. Create separate workflows for each environment
3. Use branch-specific secrets

Example with GitHub Environments:

```yaml
jobs:
  update-sdk:
    runs-on: ubuntu-latest
    environment: production  # or 'staging'

    env:
      FACETS_API_HOST: ${{ secrets.FACETS_API_HOST }}
      # ... other secrets
```

Then configure different secrets for each environment in:
**Settings** → **Environments** → **production** → **Add secret**
