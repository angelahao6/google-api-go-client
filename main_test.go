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

func TestDarwin(t *testing.T) {
	darwin.TestHi()
}
