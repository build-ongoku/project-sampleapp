package mthd

import (
	"context"

	app_client "sampleapp/backend/.goku/generated/client"
	svcdefault_entclient_typ "sampleapp/backend/.goku/generated/service/default/entity/client/typ"
)

func HookReadPre(ctx context.Context, c app_client.Client, req svcdefault_entclient_typ.Client) (svcdefault_entclient_typ.Client, error) {

	return req, nil
}

func HookInit(ctx context.Context, c app_client.Client) error {

	// Add an entity

	// Fetch an entity

	return nil
}
