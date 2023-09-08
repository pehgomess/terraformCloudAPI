package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var (
	workspaceName    = os.Getenv("TF_WORKSPACE_NAME")
	organizationName = os.Getenv("TF_ORGANIZATION_NAME")
	token            = os.Getenv("TF_API_TOKEN")
	projectName      = os.Getenv("TF_PROJECT_NAME")
)

func main() {
	errorMsg := `Certifique-se de definir as variáveis de ambiente 
	TF_WORKSPACE_NAME, TF_ORGANIZATION_NAME, TF_API_TOKEN e TF_PROJECT_NAME`

	if workspaceName == "" || organizationName == "" || projectName == "" || token == "" {
		fmt.Println(errorMsg)
		return
	}

	requestBody := map[string]interface{}{
		"data": map[string]interface{}{
			"type": "workspaces",
			"attributes": map[string]interface{}{
				"name":                  workspaceName,
				"terraform-version":     "0.13.5",
				"speculative-enabled":   true,
				"auto-apply":            false,
				"operations":            true,
				"file-triggers-enabled": true,
				"queue-all-runs":        false,
			},
			"relationships": map[string]interface{}{
				"organization": map[string]interface{}{
					"data": map[string]interface{}{
						"type": "organizations",
						"id":   organizationName,
					},
				},

				"workspace": map[string]interface{}{
					"data": map[string]interface{}{
						"type": "projects",
						"id":   projectName,
					},
				},
			},
		},
	}

	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}

	url := "https://app.terraform.io/api/v2/organizations/" + organizationName + "/workspaces"
	req, err := http.NewRequest("POST", url, strings.NewReader(string(requestJSON)))
	if err != nil {
		fmt.Println("Erro ao criar solicitação HTTP:", err)
		return
	}

	req.Header.Set("Content-Type", "application/vnd.api+json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a solicitação HTTP:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Println("Erro ao criar o espaço de trabalho. Status:", resp.Status)
		return
	}

	fmt.Println("Espaço de trabalho criado com sucesso.")
}
