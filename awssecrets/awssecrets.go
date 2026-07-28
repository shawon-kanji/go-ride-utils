// Package awssecrets fetches JSON-valued AWS Secrets Manager secrets for
// services running in staging/production, where secrets are created by
// Terraform directly in Secrets Manager (not as Kubernetes Secret objects)
// and read by the application at startup via IRSA credentials.
package awssecrets

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// FetchJSON fetches the named secret and unmarshals its string value as a
// flat JSON object, e.g. {"DB_PASSWORD": "...", "DB_USER": "..."}. Uses the
// AWS SDK's default credential chain, which IRSA populates automatically on
// the pod (AWS_ROLE_ARN / AWS_WEB_IDENTITY_TOKEN_FILE) — no explicit
// credentials or region need to be wired up by the caller.
func FetchJSON(ctx context.Context, secretName string) (map[string]string, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("load aws config: %w", err)
	}

	client := secretsmanager.NewFromConfig(cfg)
	out, err := client.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		return nil, fmt.Errorf("get secret value %q: %w", secretName, err)
	}
	if out.SecretString == nil {
		return nil, fmt.Errorf("secret %q has no string value", secretName)
	}

	values := make(map[string]string)
	if err := json.Unmarshal([]byte(*out.SecretString), &values); err != nil {
		return nil, fmt.Errorf("unmarshal secret %q: %w", secretName, err)
	}
	return values, nil
}
