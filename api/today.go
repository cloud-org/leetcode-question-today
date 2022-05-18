package api

import (
	"context"
	"github.com/machinebox/graphql"
)

func GetTodayQuestion(ctx context.Context) (*QuestionTodayResp, error) {
	// create a client (safe to share across requests)
	client := graphql.NewClient("https://leetcode.cn/graphql/")

	// make a request
	req := graphql.NewRequest(QuestionQuery)

	// set header fields
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.67 Safari/537.36")
	req.Header.Set("referer", "https://leetcode.cn/problemset/all/")
	req.Header.Set("origin", "https://leetcode.cn")

	// run it and capture the response
	var resp QuestionTodayResp
	if err := client.Run(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
