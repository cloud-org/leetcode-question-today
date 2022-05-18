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

const (
	LeetcodeCn = "https://leetcode.cn"
	Leetcode   = "https://leetcode.com"
)

// QuestionQuery graphql query
const QuestionQuery = `
	query questionOfToday {
  todayRecord {
    date
    userStatus
    question {
      questionId
      frontendQuestionId: questionFrontendId
      difficulty
      title
      titleCn: translatedTitle
      titleSlug
      paidOnly: isPaidOnly
      freqBar
      isFavor
      acRate
      status
      solutionNum
      hasVideoSolution
      topicTags {
        name
        nameTranslated: translatedName
        id
      }
      extra {
        topCompanyTags {
          imgUrl
          slug
          numSubscribed
        }
      }
    }
    lastSubmission {
      id
    }
  }
}
`

// QuestionTodayResp 注意与官网直接 restful api 请求返回的少一个 data 字段嵌套
type QuestionTodayResp struct {
	TodayRecord []struct {
		Date       string `json:"date"`
		UserStatus string `json:"userStatus"`
		Question   struct {
			QuestionID         string      `json:"questionId"`
			FrontendQuestionID string      `json:"frontendQuestionId"`
			Difficulty         string      `json:"difficulty"`
			Title              string      `json:"title"`
			TitleCn            string      `json:"titleCn"`
			TitleSlug          string      `json:"titleSlug"`
			PaidOnly           bool        `json:"paidOnly"`
			FreqBar            interface{} `json:"freqBar"`
			IsFavor            bool        `json:"isFavor"`
			AcRate             float64     `json:"acRate"`
			Status             interface{} `json:"status"`
			SolutionNum        int         `json:"solutionNum"`
			HasVideoSolution   bool        `json:"hasVideoSolution"`
			TopicTags          []struct {
				Name           string `json:"name"`
				NameTranslated string `json:"nameTranslated"`
				ID             string `json:"id"`
			} `json:"topicTags"`
			Extra struct {
				TopCompanyTags []struct {
					ImgURL        string `json:"imgUrl"`
					Slug          string `json:"slug"`
					NumSubscribed int    `json:"numSubscribed"`
				} `json:"topCompanyTags"`
			} `json:"extra"`
		} `json:"question"`
		LastSubmission interface{} `json:"lastSubmission"`
	} `json:"todayRecord"`
}
