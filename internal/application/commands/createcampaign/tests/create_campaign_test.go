package createcampaign_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/ybalcin/ecommerce-study/internal/application"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/createcampaign"
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"github.com/ybalcin/ecommerce-study/internal/domain/repositories"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		t    string
		h    *createcampaign.Handler
		c    *createcampaign.Command
		fail bool
	}{
		{
			"handler nil",
			nil,
			new(createcampaign.Command),
			true,
		},
		{
			"command nil",
			new(createcampaign.Handler),
			nil,
			true,
		},
		{
			"repo nil",
			createcampaign.NewHandler(nil, new(application.SystemTime)),
			new(createcampaign.Command),
			true,
		},
		{
			"sysTime nil",
			createcampaign.NewHandler(new(repositories.MockCampaignRepository), nil),
			new(createcampaign.Command),
			true,
		},
		{
			"repo returns error",
			createcampaign.NewHandler(repositories.MockCampaignRepository{
				AddCampaignInvoked: true,
				AddCampaignFn: func(ctx context.Context, campaign *domain.Campaign) error {
					return errors.New("")
				}}, new(application.SystemTime)),
			new(createcampaign.Command),
			true,
		},
		{
			"command invalid",
			createcampaign.NewHandler(repositories.MockCampaignRepository{
				AddCampaignInvoked: true,
				AddCampaignFn: func(ctx context.Context, campaign *domain.Campaign) error {
					return nil
				}}, new(application.SystemTime)),
			createcampaign.Build("create_campaign C11 P11 10 20"),
			true,
		},
		{
			"success",
			createcampaign.NewHandler(repositories.MockCampaignRepository{
				AddCampaignInvoked: true,
				AddCampaignFn: func(ctx context.Context, campaign *domain.Campaign) error {
					return nil
				}}, new(application.SystemTime)),
			createcampaign.Build("create_campaign C11 P11 10 20 100"),
			false,
		},
	}

	for _, c := range testCases {
		t.Run(c.t, func(t *testing.T) {
			resp, err := c.h.Handle(ctx, c.c)
			if c.fail {
				assert.Nil(t, resp)
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}
