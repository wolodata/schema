// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/wolodata/schema/ent/article"
	"github.com/wolodata/schema/ent/html"
	"github.com/wolodata/schema/ent/keywordstrong"
	"github.com/wolodata/schema/ent/report"
	"github.com/wolodata/schema/ent/schema"
	"github.com/wolodata/schema/ent/systemconfig"
	"github.com/wolodata/schema/ent/topic"
	"github.com/wolodata/schema/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	articleFields := schema.Article{}.Fields()
	_ = articleFields
	// articleDescOriginShortID is the schema descriptor for origin_short_id field.
	articleDescOriginShortID := articleFields[1].Descriptor()
	// article.OriginShortIDValidator is a validator for the "origin_short_id" field. It is called by the builders before save.
	article.OriginShortIDValidator = articleDescOriginShortID.Validators[0].(func(string) error)
	// articleDescIsChinese is the schema descriptor for is_chinese field.
	articleDescIsChinese := articleFields[2].Descriptor()
	// article.DefaultIsChinese holds the default value on creation for the is_chinese field.
	article.DefaultIsChinese = articleDescIsChinese.Default.(bool)
	// articleDescURL is the schema descriptor for url field.
	articleDescURL := articleFields[3].Descriptor()
	// article.URLValidator is a validator for the "url" field. It is called by the builders before save.
	article.URLValidator = func() func(string) error {
		validators := articleDescURL.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(url string) error {
			for _, fn := range fns {
				if err := fn(url); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// articleDescTitleChinese is the schema descriptor for title_chinese field.
	articleDescTitleChinese := articleFields[4].Descriptor()
	// article.DefaultTitleChinese holds the default value on creation for the title_chinese field.
	article.DefaultTitleChinese = articleDescTitleChinese.Default.(string)
	// articleDescTitleEnglish is the schema descriptor for title_english field.
	articleDescTitleEnglish := articleFields[5].Descriptor()
	// article.DefaultTitleEnglish holds the default value on creation for the title_english field.
	article.DefaultTitleEnglish = articleDescTitleEnglish.Default.(string)
	// articleDescAuthor is the schema descriptor for author field.
	articleDescAuthor := articleFields[6].Descriptor()
	// article.DefaultAuthor holds the default value on creation for the author field.
	article.DefaultAuthor = articleDescAuthor.Default.([]string)
	// articleDescHTMLChinese is the schema descriptor for html_chinese field.
	articleDescHTMLChinese := articleFields[8].Descriptor()
	// article.DefaultHTMLChinese holds the default value on creation for the html_chinese field.
	article.DefaultHTMLChinese = articleDescHTMLChinese.Default.(string)
	// articleDescHTMLEnglish is the schema descriptor for html_english field.
	articleDescHTMLEnglish := articleFields[9].Descriptor()
	// article.DefaultHTMLEnglish holds the default value on creation for the html_english field.
	article.DefaultHTMLEnglish = articleDescHTMLEnglish.Default.(string)
	// articleDescTextChinese is the schema descriptor for text_chinese field.
	articleDescTextChinese := articleFields[10].Descriptor()
	// article.DefaultTextChinese holds the default value on creation for the text_chinese field.
	article.DefaultTextChinese = articleDescTextChinese.Default.(string)
	// articleDescTextEnglish is the schema descriptor for text_english field.
	articleDescTextEnglish := articleFields[11].Descriptor()
	// article.DefaultTextEnglish holds the default value on creation for the text_english field.
	article.DefaultTextEnglish = articleDescTextEnglish.Default.(string)
	// articleDescImages is the schema descriptor for images field.
	articleDescImages := articleFields[12].Descriptor()
	// article.DefaultImages holds the default value on creation for the images field.
	article.DefaultImages = articleDescImages.Default.([]string)
	// articleDescWeakKeywords is the schema descriptor for weak_keywords field.
	articleDescWeakKeywords := articleFields[13].Descriptor()
	// article.DefaultWeakKeywords holds the default value on creation for the weak_keywords field.
	article.DefaultWeakKeywords = articleDescWeakKeywords.Default.([]schema.WeakKeyword)
	// articleDescStrongRelatedCategory is the schema descriptor for strong_related_category field.
	articleDescStrongRelatedCategory := articleFields[15].Descriptor()
	// article.DefaultStrongRelatedCategory holds the default value on creation for the strong_related_category field.
	article.DefaultStrongRelatedCategory = articleDescStrongRelatedCategory.Default.(string)
	// articleDescSummaryChinese is the schema descriptor for summary_chinese field.
	articleDescSummaryChinese := articleFields[16].Descriptor()
	// article.DefaultSummaryChinese holds the default value on creation for the summary_chinese field.
	article.DefaultSummaryChinese = articleDescSummaryChinese.Default.(string)
	htmlFields := schema.Html{}.Fields()
	_ = htmlFields
	// htmlDescOriginShortID is the schema descriptor for origin_short_id field.
	htmlDescOriginShortID := htmlFields[1].Descriptor()
	// html.OriginShortIDValidator is a validator for the "origin_short_id" field. It is called by the builders before save.
	html.OriginShortIDValidator = htmlDescOriginShortID.Validators[0].(func(string) error)
	// htmlDescIsChinese is the schema descriptor for is_chinese field.
	htmlDescIsChinese := htmlFields[2].Descriptor()
	// html.DefaultIsChinese holds the default value on creation for the is_chinese field.
	html.DefaultIsChinese = htmlDescIsChinese.Default.(bool)
	// htmlDescURL is the schema descriptor for url field.
	htmlDescURL := htmlFields[3].Descriptor()
	// html.URLValidator is a validator for the "url" field. It is called by the builders before save.
	html.URLValidator = func() func(string) error {
		validators := htmlDescURL.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(url string) error {
			for _, fn := range fns {
				if err := fn(url); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// htmlDescHTML is the schema descriptor for html field.
	htmlDescHTML := htmlFields[4].Descriptor()
	// html.HTMLValidator is a validator for the "html" field. It is called by the builders before save.
	html.HTMLValidator = htmlDescHTML.Validators[0].(func(string) error)
	// htmlDescCrawledAt is the schema descriptor for crawled_at field.
	htmlDescCrawledAt := htmlFields[5].Descriptor()
	// html.DefaultCrawledAt holds the default value on creation for the crawled_at field.
	html.DefaultCrawledAt = htmlDescCrawledAt.Default.(func() time.Time)
	// htmlDescReason is the schema descriptor for reason field.
	htmlDescReason := htmlFields[8].Descriptor()
	// html.DefaultReason holds the default value on creation for the reason field.
	html.DefaultReason = htmlDescReason.Default.(string)
	keywordstrongFields := schema.KeywordStrong{}.Fields()
	_ = keywordstrongFields
	// keywordstrongDescSub is the schema descriptor for sub field.
	keywordstrongDescSub := keywordstrongFields[4].Descriptor()
	// keywordstrong.DefaultSub holds the default value on creation for the sub field.
	keywordstrong.DefaultSub = keywordstrongDescSub.Default.(string)
	// keywordstrongDescSubCount is the schema descriptor for sub_count field.
	keywordstrongDescSubCount := keywordstrongFields[5].Descriptor()
	// keywordstrong.DefaultSubCount holds the default value on creation for the sub_count field.
	keywordstrong.DefaultSubCount = keywordstrongDescSubCount.Default.(uint64)
	reportFields := schema.Report{}.Fields()
	_ = reportFields
	// reportDescReportType is the schema descriptor for report_type field.
	reportDescReportType := reportFields[1].Descriptor()
	// report.ReportTypeValidator is a validator for the "report_type" field. It is called by the builders before save.
	report.ReportTypeValidator = reportDescReportType.Validators[0].(func(string) error)
	// reportDescContent is the schema descriptor for content field.
	reportDescContent := reportFields[6].Descriptor()
	// report.DefaultContent holds the default value on creation for the content field.
	report.DefaultContent = reportDescContent.Default.(string)
	// reportDescGeneratedAt is the schema descriptor for generated_at field.
	reportDescGeneratedAt := reportFields[7].Descriptor()
	// report.DefaultGeneratedAt holds the default value on creation for the generated_at field.
	report.DefaultGeneratedAt = reportDescGeneratedAt.Default.(func() time.Time)
	systemconfigFields := schema.SystemConfig{}.Fields()
	_ = systemconfigFields
	// systemconfigDescDescription is the schema descriptor for description field.
	systemconfigDescDescription := systemconfigFields[2].Descriptor()
	// systemconfig.DefaultDescription holds the default value on creation for the description field.
	systemconfig.DefaultDescription = systemconfigDescDescription.Default.(string)
	// systemconfigDescAPIModel is the schema descriptor for api_model field.
	systemconfigDescAPIModel := systemconfigFields[3].Descriptor()
	// systemconfig.DefaultAPIModel holds the default value on creation for the api_model field.
	systemconfig.DefaultAPIModel = systemconfigDescAPIModel.Default.(string)
	// systemconfigDescAPIURL is the schema descriptor for api_url field.
	systemconfigDescAPIURL := systemconfigFields[4].Descriptor()
	// systemconfig.DefaultAPIURL holds the default value on creation for the api_url field.
	systemconfig.DefaultAPIURL = systemconfigDescAPIURL.Default.(string)
	// systemconfigDescAPIKey is the schema descriptor for api_key field.
	systemconfigDescAPIKey := systemconfigFields[5].Descriptor()
	// systemconfig.DefaultAPIKey holds the default value on creation for the api_key field.
	systemconfig.DefaultAPIKey = systemconfigDescAPIKey.Default.(string)
	// systemconfigDescPromptSystem is the schema descriptor for prompt_system field.
	systemconfigDescPromptSystem := systemconfigFields[6].Descriptor()
	// systemconfig.DefaultPromptSystem holds the default value on creation for the prompt_system field.
	systemconfig.DefaultPromptSystem = systemconfigDescPromptSystem.Default.(string)
	// systemconfigDescPromptUser is the schema descriptor for prompt_user field.
	systemconfigDescPromptUser := systemconfigFields[7].Descriptor()
	// systemconfig.DefaultPromptUser holds the default value on creation for the prompt_user field.
	systemconfig.DefaultPromptUser = systemconfigDescPromptUser.Default.(string)
	topicFields := schema.Topic{}.Fields()
	_ = topicFields
	// topicDescKeyword is the schema descriptor for keyword field.
	topicDescKeyword := topicFields[2].Descriptor()
	// topic.KeywordValidator is a validator for the "keyword" field. It is called by the builders before save.
	topic.KeywordValidator = topicDescKeyword.Validators[0].(func(string) error)
	// topicDescFollowTitle is the schema descriptor for follow_title field.
	topicDescFollowTitle := topicFields[3].Descriptor()
	// topic.DefaultFollowTitle holds the default value on creation for the follow_title field.
	topic.DefaultFollowTitle = topicDescFollowTitle.Default.(bool)
	// topicDescFollowContent is the schema descriptor for follow_content field.
	topicDescFollowContent := topicFields[4].Descriptor()
	// topic.DefaultFollowContent holds the default value on creation for the follow_content field.
	topic.DefaultFollowContent = topicDescFollowContent.Default.(bool)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescIsAdmin is the schema descriptor for is_admin field.
	userDescIsAdmin := userFields[3].Descriptor()
	// user.DefaultIsAdmin holds the default value on creation for the is_admin field.
	user.DefaultIsAdmin = userDescIsAdmin.Default.(bool)
}
