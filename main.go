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

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"leetcode-question-today/api"
	"leetcode-question-today/msgpush"
)

var (
	wecom string // wecom 通知链接
	help  bool   // 帮助
)

func init() {
	flag.StringVar(&wecom, "wecom", "", "wecom webhook token")
	flag.BoolVar(&help, "h", false, "帮助")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stdout, `leetcode-question-today - leetcode 每日一题推送
Usage: leetcode-question-today [-h help]
Options:
`)
	flag.PrintDefaults()
}

const msgTemplate = `每日一题(%s)
Title: %s
Difficulty: %s
Tags: %s
Link: %s
LinkCN: %s`

func main() {
	flag.Parse()
	if help {
		flag.PrintDefaults()
		return
	}

	// 获取每日一题，如果有则推送即可
	resp, err := api.GetTodayQuestion(context.TODO())
	if err != nil {
		log.Printf("获取每日一题发生错误: %v\n", err)
		return
	}

	// 获取每日一题，如果有则推送即可
	respAll, err := api.GetALLQuestionsV2()
	if err != nil {
		log.Printf("获取每日一题发生错误: %v\n", err)
		return
	}

	if len(resp.TodayRecord) <= 0 {
		log.Printf("todayRecord 长度为 0,请检查\n")
		return
	}
	if len(respAll.Data.P.Questions) <= 0 {
		log.Printf("todayRecord 长度为 0,请检查\n")
		return
	}

	today := resp.TodayRecord[0]
	date := today.Date
	for i := 0; i < 3; i++ {
		index := rand.Uint32() % 100
		q := respAll.Data.P.Questions[index]
		diff := q.Difficulty
		title := q.TitleCn
		link := fmt.Sprintf("%s/problems/%s", api.Leetcode, q.TitleSlug)
		linkCn := fmt.Sprintf("%s/problems/%s", api.LeetcodeCn, q.TitleSlug)
		tags := make([]string, 0)
		for _, tag := range q.TopicTags {
			tags = append(tags, fmt.Sprintf("%s(%s)", tag.NameTranslated, tag.Name))
		}
		tagsValue := strings.Join(tags, "、")
		content := fmt.Sprintf(msgTemplate, date, title, diff, tagsValue, link, linkCn)

		log.Println(content)

		if wecom != "" {
			w := msgpush.NewWeCom(wecom)
			_ = w.SendText(content, []string{"hongyu.an", "quan.ren", "binyuan.rong", "ming.wen"})
		}
	}

	difficulty := today.Question.Difficulty
	title := fmt.Sprintf("%s(%s)", today.Question.TitleCn, today.Question.Title)
	tags := make([]string, 0)
	for _, tag := range today.Question.TopicTags {
		tags = append(tags, fmt.Sprintf("%s(%s)", tag.NameTranslated, tag.Name))
	}
	tagsValue := strings.Join(tags, "、")
	link := fmt.Sprintf("%s/problems/%s", api.Leetcode, today.Question.TitleSlug)
	linkCn := fmt.Sprintf("%s/problems/%s", api.LeetcodeCn, today.Question.TitleSlug)

	content := fmt.Sprintf(msgTemplate, date, title, difficulty, tagsValue, link, linkCn)

	log.Println(content)

	if wecom != "" {
		w := msgpush.NewWeCom(wecom)
		_ = w.SendText(content, []string{"bin.zhang"})
	}

	return
}
