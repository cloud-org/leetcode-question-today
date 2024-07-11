package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/machinebox/graphql"
)

func GetALLQuestions(ctx context.Context) (*problemsetQuestionList, error) {
	// create a client (safe to share across requests)
	client := graphql.NewClient("https://leetcode.cn/graphql/")

	// make a request
	req := graphql.NewRequest(QuestionsQuery)

	// set header fields
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.67 Safari/537.36")
	req.Header.Set("referer", "https://leetcode.cn/problemset/")
	req.Header.Set("origin", "https://leetcode.cn")

	// run it and capture the response
	var resp problemsetQuestionList
	if err := client.Run(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func GetALLQuestionsV2() (*Data, error) {
	url := "https://leetcode.cn/graphql/"
	method := "POST"

	payload := strings.NewReader(`{"query":"\n    query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {\n  problemsetQuestionList(\n    categorySlug: $categorySlug\n    limit: $limit\n    skip: $skip\n    filters: $filters\n  ) {\n    hasMore\n    total\n    questions {\n      acRate\n      difficulty\n      freqBar\n      frontendQuestionId\n      isFavor\n      paidOnly\n      solutionNum\n      status\n      title\n      titleCn\n      titleSlug\n      topicTags {\n        name\n        nameTranslated\n        id\n        slug\n      }\n      extra {\n        hasVideoSolution\n        topCompanyTags {\n          imgUrl\n          slug\n          numSubscribed\n        }\n      }\n    }\n  }\n}\n    ","variables":{"categorySlug":"all-code-essentials","skip":0,"limit":200,"filters":{}},"operationName":"problemsetQuestionList"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("authorization", "")
	req.Header.Add("baggage", "sentry-environment=production,sentry-release=c2ab2648,sentry-transaction=%2Fproblemset%2F%5B%5B...slug%5D%5D,sentry-public_key=1595090ae2f831f9e65978be5851f865,sentry-trace_id=1c4ed2f9d5c14d3fac4106618a94f8db,sentry-sample_rate=0.03")
	req.Header.Add("random-uuid", "8b625142-a311-dd83-cbe2-d8a6be23c4be")
	req.Header.Add("sentry-trace", "1c4ed2f9d5c14d3fac4106618a94f8db-9a256ea4aa22ad04-0")
	req.Header.Add("x-csrftoken", "krzx7Gy6Wm8fUWdhvWs0oUPuPKEOQRKWNqit0tTWlrbMSxJqCbMo5BP4fmF0nuX3")
	req.Header.Add("Cookie", "aliyungf_tc=9cf57447268a6a45bf77ac4fdff01a0390003bd40bbdb4fc4b900cf25a63f8dd; csrftoken=krzx7Gy6Wm8fUWdhvWs0oUPuPKEOQRKWNqit0tTWlrbMSxJqCbMo5BP4fmF0nuX3; _ga=GA1.1.1219614871.1720702589; _ga_PDVPZYN3CW=GS1.1.1720702589.1.1.1720702884.60.0.0")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "leetcode.cn")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(string(body))
	resp := Data{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &resp, nil
}
