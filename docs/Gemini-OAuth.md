# Use Gemini with Google OAuth

You can use Gemini models in Fabric without a Gemini API key. To do this, use the **VertexAI** vendor. It signs in to Google Cloud with your Google account through OAuth.

## Subscriptions are not API access

A Google AI Pro or Google AI Ultra subscription does not give you access to the Gemini API. The subscription is for the Gemini app and for Google tools such as Gemini CLI.

Some third-party tools use the OAuth client of Gemini CLI or Antigravity to send requests on your subscription. Google suspends accounts that do this. For this reason, Fabric does not have this feature, and it will not get it.

Fabric gives you two methods to use Gemini:

| Vendor | Authentication | Billing |
| --- | --- | --- |
| `Gemini` | API key from [Google AI Studio](https://aistudio.google.com/apikey) | Gemini API (free tier is available) |
| `VertexAI` | Google OAuth (Application Default Credentials) | Your Google Cloud project |

## Set up the VertexAI vendor

1. Install the [Google Cloud CLI](https://cloud.google.com/sdk/docs/install).
2. Make or select a Google Cloud project. Make sure that billing is on.
3. Turn on the Vertex AI API:

   ```bash
   gcloud services enable aiplatform.googleapis.com --project YOUR_PROJECT_ID
   ```

4. Sign in. This step opens your browser for the OAuth flow:

   ```bash
   gcloud auth application-default login
   ```

   The Google libraries keep the credentials and refresh the tokens. Fabric does not keep a token file.

5. Run `fabric --setup`, select **VertexAI**, and type your project ID. The region is `global` by default.

   You can also set these values in `~/.config/fabric/.env`:

   ```bash
   VERTEXAI_PROJECT_ID=your-project-id
   VERTEXAI_REGION=global
   ```

## Use a Gemini model

```bash
fabric --listmodels --vendor VertexAI
echo "Tell me about OAuth" | fabric --vendor VertexAI --model gemini-2.5-flash
```

The `Gemini` vendor and the `VertexAI` vendor can both be configured at the same time. Use `--vendor` to select one.

## Troubleshooting

- **`could not find default credentials`**: Do step 4 again.
- **`PERMISSION_DENIED`**: Make sure that the Vertex AI API is on for the project, and that your account has the `Vertex AI User` role.
- **A preview model is not found**: Fabric sends preview models to the `global` region. Make sure that the model is available for your project.
