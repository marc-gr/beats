// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build windows

package amcache

import (
	"fmt"

	"github.com/elastic/beats/v7/x-pack/osquerybeat/ext/osquery-extension/pkg/amcache/tables"
	"github.com/elastic/beats/v7/x-pack/osquerybeat/ext/osquery-extension/pkg/hooks"
	"github.com/elastic/beats/v7/x-pack/osquerybeat/ext/osquery-extension/pkg/logger"
	elasticamcacheapplicationsview "github.com/elastic/beats/v7/x-pack/osquerybeat/ext/osquery-extension/pkg/views/generated/elastic_amcache_applications_view"
)

func init() {
	// Register the hooks function with the generated view
	elasticamcacheapplicationsview.RegisterHooksFunc(registerHooks)
}

func CreateViewHook(socket *string, log *logger.Logger, hookData any) error {
	view, ok := hookData.(*hooks.View)
	if !ok {
		return fmt.Errorf("hook data is not a view")
	}
	return view.Create(socket, log)
}

func DeleteViewHook(socket *string, log *logger.Logger, hookData any) error {
	view, ok := hookData.(*hooks.View)
	if !ok {
		return fmt.Errorf("hook data is not a view")
	}
	return view.Delete(socket, log)
}

func CleanupInstanceHook(socket *string, log *logger.Logger, hookData any) error {
	amcacheState := tables.GetAmcacheState()
	amcacheState.Close()
	return nil
}

func registerHooks(hm *hooks.HookManager) {
	hm.Register(hooks.NewHook("CreateAmcacheApplicationsView", CreateViewHook, DeleteViewHook, elasticamcacheapplicationsview.View()))
	hm.Register(hooks.NewHook("CleanupAmcacheInstance", nil, CleanupInstanceHook, nil))
}
