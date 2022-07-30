package main

import (
    "fmt"
    "strings"
    atomparser "github.com/wbernest/atom-parser"
)

func main() {
	processatomsubscription()
}

func processatomsubscription()  {
	// config := p.getConfiguration()

	// get new rss feed string from url
	newFeed,_ , _ := atomparser.ParseURL("https://www.google.co.jp/alerts/feeds/07908382636288630076/17396211672342771178")

	oldFeed, _ := atomparser.ParseString(`<feed xmlns="http://www.w3.org/2005/Atom" xmlns:idx="urn:atom-extension:indexing"> <id>tag:google.com,2005:reader/user/07908382636288630076/state/com.google/alerts/17396211672342771178</id> <title>Google アラート - あ</title> <link href="https://www.google.com/alerts/feeds/07908382636288630076/17396211672342771178" rel="self"/> <updated>2022-07-30T07:10:36Z</updated></feed>`)
	// fmt.Println(oldFeed)

	items := atomparser.CompareItemsBetweenOldAndNew(oldFeed, newFeed)

	// if this is a new subscription only post the latest
	// and not spam the channel
	if len(oldFeed.Entry) == 0 && len(items) > 0 {
		items = items[:1]
	}

	for _, item := range items {
		post := ""

		// if config.FormatTitle {
			post = post + "##### "
		// }
		// post = post + newFeed.Title + "\n" // 削除
		post = post + "ラクスアラート (ラクト)\n" // 削除

		// if config.ShowAtomItemTitle {
			// if config.FormatTitle {
				post = post + "###### "
			// }
			post = post + item.Title + "\n"
		// }

		// if config.ShowAtomLink {
			for _, link := range item.Link {
				// if link.Rel == "alternate" {
					post = post + strings.TrimSpace(link.Href) + "\n"
				// }
			}
		// }
			for _, content := range item.Content {
				// if link.Rel == "alternate" {

		// fmt.Println(content)
					// post = post + content + "\n"
				// }
			}
			// post = post + item.Link + "\n"
		// fmt.Println(item.Content)
		// fmt.Println(item.Link)

			if !tryParseRichNode(item.Content, &post) {
				p.API.LogInfo("Missing content in atom feed item",
					"subscription_url", subscription.URL,
					"item_title", item.Title)
				post = post + "\n"
			}
		fmt.Println(post)

		// if config.ShowSummary {
			// if !tryParseRichNode(item.Summary, &post) {
			// 	p.API.LogInfo("Missing summary in atom feed item",
			// 		"subscription_url", subscription.URL,
			// 		"item_title", item.Title)
			// 	post = post + "\n"
			// }
		// }

		// if config.ShowContent {
		// 	if !tryParseRichNode(item.Content, &post) {
		// 		p.API.LogInfo("Missing content in atom feed item",
		// 			"subscription_url", subscription.URL,
		// 			"item_title", item.Title)
		// 		post = post + "\n"
		// 	}
		// }

		// p.createBotPost(subscription.ChannelID, post, "custom_git_pr")
	}
}
func tryParseRichNode(node *atom.Text, post *string) bool {
	if node != nil {
		if node.Type != "text" {
			*post = *post + html2md.Convert(strings.TrimSpace(node.Body)) + "\n"
		} else {
			*post = *post + node.Body + "\n"
		}
		return true
	} else {
		return false
	}
}

