package configs

import (
	"context"
	"fmt"
	"os"

	"github.com/lestrrat-go/jwx/v3/jwk"
)

// SupabaseConfig holds Supabase configuration
type SupabaseConfig struct {
	URL            string
	AnonKey        string
	ServiceRoleKey string
	JWKSURL        string
}

var supabase *SupabaseConfig

// InitSupabase initializes Supabase configuration from environment variables
func InitSupabase() error {
	url := os.Getenv("SUPABASE_URL")
	if url == "" {
		return fmt.Errorf("SUPABASE_URL is required")
	}

	jwksURL := fmt.Sprintf("%s/auth/v1/.well-known/jwks.json", url)

	// Verify we can fetch the JWKS endpoint
	ctx := context.Background()
	_, err := jwk.Fetch(ctx, jwksURL)
	if err != nil {
		return fmt.Errorf("failed to fetch JWKS: %w", err)
	}

	supabase = &SupabaseConfig{
		URL:            url,
		AnonKey:        os.Getenv("SUPABASE_ANON_KEY"),
		ServiceRoleKey: os.Getenv("SUPABASE_SERVICE_ROLE_KEY"),
		JWKSURL:        jwksURL,
	}

	fmt.Println("Done configuring Supabase with JWK verification.")
	return nil
}

// Supabase returns the Supabase configuration
func Supabase() *SupabaseConfig {
	return supabase
}
