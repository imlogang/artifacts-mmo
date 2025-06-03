package main

import (
	artifactclient "artifacts-mmo/client"
	"artifacts-mmo/resource"
	"context"
	"fmt"
	"github.com/circleci/ex/httpclient"
	"log"
	"os"
	"time"
)

func main() {
	tokenBytes, err := os.ReadFile("token")
	if err != nil {
		log.Fatal(err)
	}

	token := string(tokenBytes)

	ctx := context.Background()
	artifactHTTPClient := NewArtifactClient(ctx, token)
	//resp, err := MoveCharacter(ctx, artifactHTTPClient, 0, 1)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("Character moved to %d, %d.\n", resp.Destination.X, resp.Destination.Y)

	fightResp, err := Fight(ctx, artifactHTTPClient)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Current HP after Fight: %d\nCurrent cooldown is: %d\nFight Logs: %s\n", fightResp.Character.Hp, fightResp.Cooldown.TotalSeconds, fightResp.Fight.Logs)
	cooldown := time.Duration(fightResp.Cooldown.TotalSeconds)
	time.Sleep(cooldown * time.Second)
	//if fightResp.Cooldown.TotalSeconds == 0 {
	//	fmt.Println("We're resting now!")
	//	restResp, err := Rest(ctx, artifactHTTPClient)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Printf("Current HP: %v\n", restResp.Character.Hp)
	//}
	restResp, err := Rest(ctx, artifactHTTPClient)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Current HP: %v\n", restResp.Character.Hp)
	fmt.Printf("Curernt cooldown is: %d\n", restResp.Character.Cooldown)

	//getCharacter(ctx, artifactHTTPClient)
}

type CharacterResponse struct {
	Data struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"data"`
}

func NewArtifactClient(ctx context.Context, token string) *httpclient.Client {
	httpclientConfig := httpclient.Config{
		Name:       "artifactHTTPClient",
		BaseURL:    "https://api.artifactsmmo.com",
		AuthToken:  token,
		AcceptType: "application/json",
	}

	return httpclient.New(httpclientConfig)
}

func authTest(token string) {
	ctx := context.Background()

	client := NewArtifactClient(ctx, token)
	req := httpclient.NewRequest("GET", "/my/details")
	resp := CharacterResponse{}
	httpclient.JSONDecoder(&resp)(&req)
	httpclient.Body(req)
	err := client.Call(ctx, req)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(resp.Data)
}

func createCharacter(client *httpclient.Client, ctx context.Context) {
	req := httpclient.NewRequest("POST", "/characters/create")
	resp := struct {
		Data resource.Character `json:"data"`
	}{}
	responseJSONDecoder := httpclient.JSONDecoder(&resp)
	responseJSONDecoder(&req)
	requestBodySetter := httpclient.Body(struct {
		Name string `json:"name"`
		Skin string `json:"skin"`
	}{
		Name: "Logan",
		Skin: "men1",
	})

	requestBodySetter(&req)

	err := client.Call(ctx, req)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(resp)
}

func getCharacter(ctx context.Context, client *httpclient.Client) {
	req := httpclient.NewRequest("GET", "/characters/Logan")
	resp := resource.Response{}
	httpclient.JSONDecoder(&resp)(&req)
	httpclient.Body(req)
	err := client.Call(ctx, req)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("\nCurrent HP: %d", resp.Data.Character.Hp)
}

func MoveCharacter(ctx context.Context, client *httpclient.Client, x, y int) (resource.Data, error) {
	req := httpclient.NewRequest("POST", "/my/Logan/action/move")
	resp := struct {
		Data resource.Data `json:"data"`
	}{}
	responseJSONDecoder := httpclient.JSONDecoder(&resp)
	responseJSONDecoder(&req)
	requestBodySetter := httpclient.Body(struct {
		X int `json:"x"`
		Y int `json:"y"`
	}{
		X: x,
		Y: y,
	})

	requestBodySetter(&req)
	err := client.Call(ctx, req)
	if err != nil {
		if isError := httpclient.HasStatusCode(err, 404, 486, 490, 498, 499); isError {
			switch err.(*httpclient.HTTPError).Code() {
			case 404:
				return resp.Data, artifactclient.MapNotFound{}
			case 486:
				return resp.Data, artifactclient.ActionInProgress{}
			case 490:
				return resp.Data, artifactclient.CharacterAtDestinationError{}
			case 498:
				return resp.Data, artifactclient.CharacterNotFound{}
			case 499:
				return resp.Data, artifactclient.CharacterInCooldown{}
			default:
				fmt.Printf("Unknown error: %s", err)
			}
		}
		return resp.Data, err
	}
	return resp.Data, nil
}

func Fight(ctx context.Context, client *httpclient.Client) (resource.FightResponse, error) {
	req := httpclient.NewRequest("POST", "/my/Logan/action/fight")
	resp := struct {
		Data resource.FightResponse `json:"data"`
	}{}
	httpclient.JSONDecoder(&resp)(&req) //I need a refresher on what this does and why it's needed here.
	err := client.Call(ctx, req)
	if err != nil {
		if isError := httpclient.HasStatusCode(err, 486, 497, 498, 499, 598); isError {
			switch err.(*httpclient.HTTPError).Code() {
			case 486:
				return resp.Data, artifactclient.ActionInProgress{}
			case 490:
				return resp.Data, artifactclient.CharacterInvFull{}
			case 498:
				return resp.Data, artifactclient.CharacterNotFound{}
			case 499:
				return resp.Data, artifactclient.CharacterInCooldown{}
			case 598:
				return resp.Data, artifactclient.MonsterNotFound{}
			default:
				fmt.Printf("Unknown error: %s", err)
			}
		}
		return resp.Data, err
	}
	return resp.Data, nil
}

func Rest(ctx context.Context, client *httpclient.Client) (resource.RestResponse, error) {
	req := httpclient.NewRequest("POST", "/my/Logan/action/rest")
	resp := struct {
		Data resource.RestResponse `json:"data"`
	}{}
	responseDecoderSetter := httpclient.JSONDecoder(&resp)
	responseDecoderSetter(&req)
	err := client.Call(ctx, req)
	if err != nil {
		if isError := httpclient.HasStatusCode(err, 486, 498, 499); isError {
			switch err.(*httpclient.HTTPError).Code() {
			case 486:
				return resp.Data, artifactclient.ActionInProgress{}
			case 498:
				return resp.Data, artifactclient.CharacterNotFound{}
			case 499:
				return resp.Data, artifactclient.CharacterInCooldown{}
			default:
				fmt.Printf("Unknown error: %s", err)
			}
		}
		return resp.Data, err
	}
	return resp.Data, nil
}
