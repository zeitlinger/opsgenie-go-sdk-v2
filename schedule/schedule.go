package schedule

import (
	"context"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
)

type Client struct {
	ogClient client.OpsGenieClient
}

type Schedule struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description,omitempty"`
	Timezone    string        `json:"timezone,omitempty"`
	Enabled     bool          `json:"enabled"`
	OwnerTeam   *og.OwnerTeam `json:"ownerTeam,omitempty"`
	Rotations   []og.Rotation `json:"rotations,omitempty"`
}

func NewClient(config *client.Config) (*Client, error) {
	opsgenieClient, err := client.NewOpsGenieClient(config)
	if err != nil {
		return nil, err
	}
	client := &Client{}
	client.ogClient = *opsgenieClient
	return client, nil
}

func (client *Client) Create(context context.Context, request CreateRequest) (*CreateResult, error) {
	result := &CreateResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Get(context context.Context, request GetRequest) (*GetResult, error) {
	result := &GetResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Update(context context.Context, request UpdateRequest) (*UpdateResult, error) {
	result := &UpdateResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Delete(context context.Context, request DeleteRequest) (*DeleteResult, error) {
	result := &DeleteResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) List(context context.Context, request ListRequest) (*ListResult, error) {
	result := &ListResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) GetTimeline(context context.Context, request GetTimelineRequest) (*TimelineResult, error) {
	result := &TimelineResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
