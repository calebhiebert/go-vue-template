package main

import "github.com/calebhiebert/go-vue-template/models/modelcrud"

func ConfigAdminCrudModels() {
	modelcrud.AccessLogsModelConfig.ID.ShowOnGraph = false

	modelcrud.AccessLogsModelConfig.ResponseBody.ShowOnGraph = false
	modelcrud.AccessLogsModelConfig.ResponseHeaders.ShowOnGraph = false

	modelcrud.AccessLogsModelConfig.RequestBody.ShowOnGraph = false
	modelcrud.AccessLogsModelConfig.RequestHeaders.ShowOnGraph = false

	modelcrud.AccessLogsModelConfig.Apply()
}
