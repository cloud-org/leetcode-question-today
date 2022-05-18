/*
 * MIT License
 *
 * Copyright (c) 2021 ashing
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

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
