package main

import (
	"context"

	"log"

	"testing"

	"fmt"

	"io"

	"google.golang.org/api/option"

	"google.golang.org/api/transport"

	"github.com/googleapis/enterprise-certificate-proxy/darwin"
)

func TestGoExample(t *testing.T) {

	ctx := context.Background()

	hc, _, err := transport.NewHTTPClient(ctx, option.WithQuotaProject("sijunliu-dca-test"))

	if err != nil {

		log.Fatalf("NewHTTPClient: %v", err)

	}

	resp, err := hc.Get("https://pubsub.mtls.googleapis.com/v1/projects/sijunliu-dca-test/topics")

	if err != nil {

		log.Fatalf("Get: %v", err)

	}

	log.Printf("Status: %s", resp.Status)

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	if err != nil {

		log.Fatalln(err)

	}

	fmt.Println(string(b))

}

func TestPublicInClient(t *testing.T) {
	secureKey, err := darwin.NewSecureKey("Google Endpoint Verification")
	if err != nil {
		fmt.Println("Error in getting private key pair")
	}
	fmt.Println(secureKey.Public())
}

// func TestPublicInKeychain(t *testing.T) {
// 	keyPointer, err := keychain.Cred("Google Endpoint Verification")
// 	if err != nil {
// 		t.Errorf("Cred error: %q", err)
// 		return
// 	}
// 	fmt.Println(keyPointer.Public())

// }
