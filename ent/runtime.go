// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/wolodata/schema/ent/article"
	"github.com/wolodata/schema/ent/html"
	"github.com/wolodata/schema/ent/keywordstrong"
	"github.com/wolodata/schema/ent/promotconfig"
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
	// articleDescHTMLChinese is the schema descriptor for html_chinese field.
	articleDescHTMLChinese := articleFields[7].Descriptor()
	// article.DefaultHTMLChinese holds the default value on creation for the html_chinese field.
	article.DefaultHTMLChinese = articleDescHTMLChinese.Default.(string)
	// articleDescHTMLEnglish is the schema descriptor for html_english field.
	articleDescHTMLEnglish := articleFields[8].Descriptor()
	// article.DefaultHTMLEnglish holds the default value on creation for the html_english field.
	article.DefaultHTMLEnglish = articleDescHTMLEnglish.Default.(string)
	// articleDescTextChinese is the schema descriptor for text_chinese field.
	articleDescTextChinese := articleFields[9].Descriptor()
	// article.DefaultTextChinese holds the default value on creation for the text_chinese field.
	article.DefaultTextChinese = articleDescTextChinese.Default.(string)
	// articleDescTextEnglish is the schema descriptor for text_english field.
	articleDescTextEnglish := articleFields[10].Descriptor()
	// article.DefaultTextEnglish holds the default value on creation for the text_english field.
	article.DefaultTextEnglish = articleDescTextEnglish.Default.(string)
	// articleDescImages is the schema descriptor for images field.
	articleDescImages := articleFields[11].Descriptor()
	// article.DefaultImages holds the default value on creation for the images field.
	article.DefaultImages = articleDescImages.Default.([]string)
	// articleDescImageUploaded is the schema descriptor for image_uploaded field.
	articleDescImageUploaded := articleFields[12].Descriptor()
	// article.DefaultImageUploaded holds the default value on creation for the image_uploaded field.
	article.DefaultImageUploaded = articleDescImageUploaded.Default.(bool)
	// articleDescWeakKeywordProcessed is the schema descriptor for weak_keyword_processed field.
	articleDescWeakKeywordProcessed := articleFields[13].Descriptor()
	// article.DefaultWeakKeywordProcessed holds the default value on creation for the weak_keyword_processed field.
	article.DefaultWeakKeywordProcessed = articleDescWeakKeywordProcessed.Default.(bool)
	// articleDescWeakKeywordRelated is the schema descriptor for weak_keyword_related field.
	articleDescWeakKeywordRelated := articleFields[14].Descriptor()
	// article.DefaultWeakKeywordRelated holds the default value on creation for the weak_keyword_related field.
	article.DefaultWeakKeywordRelated = articleDescWeakKeywordRelated.Default.(bool)
	// articleDescWeakKeywords is the schema descriptor for weak_keywords field.
	articleDescWeakKeywords := articleFields[15].Descriptor()
	// article.DefaultWeakKeywords holds the default value on creation for the weak_keywords field.
	article.DefaultWeakKeywords = articleDescWeakKeywords.Default.([]schema.WeakKeyword)
	// articleDescStrongKeywordProcessed is the schema descriptor for strong_keyword_processed field.
	articleDescStrongKeywordProcessed := articleFields[16].Descriptor()
	// article.DefaultStrongKeywordProcessed holds the default value on creation for the strong_keyword_processed field.
	article.DefaultStrongKeywordProcessed = articleDescStrongKeywordProcessed.Default.(bool)
	// articleDescStrongKeywordRelated is the schema descriptor for strong_keyword_related field.
	articleDescStrongKeywordRelated := articleFields[17].Descriptor()
	// article.DefaultStrongKeywordRelated holds the default value on creation for the strong_keyword_related field.
	article.DefaultStrongKeywordRelated = articleDescStrongKeywordRelated.Default.(bool)
	// articleDescStrongRelatedProcessed is the schema descriptor for strong_related_processed field.
	articleDescStrongRelatedProcessed := articleFields[19].Descriptor()
	// article.DefaultStrongRelatedProcessed holds the default value on creation for the strong_related_processed field.
	article.DefaultStrongRelatedProcessed = articleDescStrongRelatedProcessed.Default.(bool)
	// articleDescStrongRelated is the schema descriptor for strong_related field.
	articleDescStrongRelated := articleFields[20].Descriptor()
	// article.DefaultStrongRelated holds the default value on creation for the strong_related field.
	article.DefaultStrongRelated = articleDescStrongRelated.Default.(bool)
	// articleDescStrongRelatedCategoryProcessed is the schema descriptor for strong_related_category_processed field.
	articleDescStrongRelatedCategoryProcessed := articleFields[21].Descriptor()
	// article.DefaultStrongRelatedCategoryProcessed holds the default value on creation for the strong_related_category_processed field.
	article.DefaultStrongRelatedCategoryProcessed = articleDescStrongRelatedCategoryProcessed.Default.(bool)
	// articleDescStrongRelatedCategory is the schema descriptor for strong_related_category field.
	articleDescStrongRelatedCategory := articleFields[22].Descriptor()
	// article.DefaultStrongRelatedCategory holds the default value on creation for the strong_related_category field.
	article.DefaultStrongRelatedCategory = articleDescStrongRelatedCategory.Default.(string)
	// articleDescSummaryChinese is the schema descriptor for summary_chinese field.
	articleDescSummaryChinese := articleFields[23].Descriptor()
	// article.DefaultSummaryChinese holds the default value on creation for the summary_chinese field.
	article.DefaultSummaryChinese = articleDescSummaryChinese.Default.(string)
	// articleDescStrongRelatedSummaryChineseProcessed is the schema descriptor for strong_related_summary_chinese_processed field.
	articleDescStrongRelatedSummaryChineseProcessed := articleFields[24].Descriptor()
	// article.DefaultStrongRelatedSummaryChineseProcessed holds the default value on creation for the strong_related_summary_chinese_processed field.
	article.DefaultStrongRelatedSummaryChineseProcessed = articleDescStrongRelatedSummaryChineseProcessed.Default.(bool)
	// articleDescStrongRelatedSummaryChinese is the schema descriptor for strong_related_summary_chinese field.
	articleDescStrongRelatedSummaryChinese := articleFields[25].Descriptor()
	// article.DefaultStrongRelatedSummaryChinese holds the default value on creation for the strong_related_summary_chinese field.
	article.DefaultStrongRelatedSummaryChinese = articleDescStrongRelatedSummaryChinese.Default.(string)
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
	promotconfigFields := schema.PromotConfig{}.Fields()
	_ = promotconfigFields
	// promotconfigDescDescription is the schema descriptor for description field.
	promotconfigDescDescription := promotconfigFields[2].Descriptor()
	// promotconfig.DefaultDescription holds the default value on creation for the description field.
	promotconfig.DefaultDescription = promotconfigDescDescription.Default.(string)
	// promotconfigDescAPIModel is the schema descriptor for api_model field.
	promotconfigDescAPIModel := promotconfigFields[3].Descriptor()
	// promotconfig.DefaultAPIModel holds the default value on creation for the api_model field.
	promotconfig.DefaultAPIModel = promotconfigDescAPIModel.Default.(string)
	// promotconfigDescAPIURL is the schema descriptor for api_url field.
	promotconfigDescAPIURL := promotconfigFields[4].Descriptor()
	// promotconfig.DefaultAPIURL holds the default value on creation for the api_url field.
	promotconfig.DefaultAPIURL = promotconfigDescAPIURL.Default.(string)
	// promotconfigDescAPIKey is the schema descriptor for api_key field.
	promotconfigDescAPIKey := promotconfigFields[5].Descriptor()
	// promotconfig.DefaultAPIKey holds the default value on creation for the api_key field.
	promotconfig.DefaultAPIKey = promotconfigDescAPIKey.Default.(string)
	// promotconfigDescPromptSystem is the schema descriptor for prompt_system field.
	promotconfigDescPromptSystem := promotconfigFields[6].Descriptor()
	// promotconfig.DefaultPromptSystem holds the default value on creation for the prompt_system field.
	promotconfig.DefaultPromptSystem = promotconfigDescPromptSystem.Default.(string)
	// promotconfigDescPromptUser is the schema descriptor for prompt_user field.
	promotconfigDescPromptUser := promotconfigFields[7].Descriptor()
	// promotconfig.DefaultPromptUser holds the default value on creation for the prompt_user field.
	promotconfig.DefaultPromptUser = promotconfigDescPromptUser.Default.(string)
	reportFields := schema.Report{}.Fields()
	_ = reportFields
	// reportDescReportType is the schema descriptor for report_type field.
	reportDescReportType := reportFields[1].Descriptor()
	// report.ReportTypeValidator is a validator for the "report_type" field. It is called by the builders before save.
	report.ReportTypeValidator = reportDescReportType.Validators[0].(func(string) error)
	// reportDescSourceIds is the schema descriptor for source_ids field.
	reportDescSourceIds := reportFields[4].Descriptor()
	// report.DefaultSourceIds holds the default value on creation for the source_ids field.
	report.DefaultSourceIds = reportDescSourceIds.Default.([]string)
	// reportDescCategory is the schema descriptor for category field.
	reportDescCategory := reportFields[5].Descriptor()
	// report.DefaultCategory holds the default value on creation for the category field.
	report.DefaultCategory = reportDescCategory.Default.(string)
	// reportDescArticleIds is the schema descriptor for article_ids field.
	reportDescArticleIds := reportFields[6].Descriptor()
	// report.DefaultArticleIds holds the default value on creation for the article_ids field.
	report.DefaultArticleIds = reportDescArticleIds.Default.([]string)
	// reportDescContent is the schema descriptor for content field.
	reportDescContent := reportFields[9].Descriptor()
	// report.DefaultContent holds the default value on creation for the content field.
	report.DefaultContent = reportDescContent.Default.(string)
	// reportDescGeneratedAt is the schema descriptor for generated_at field.
	reportDescGeneratedAt := reportFields[10].Descriptor()
	// report.DefaultGeneratedAt holds the default value on creation for the generated_at field.
	report.DefaultGeneratedAt = reportDescGeneratedAt.Default.(func() time.Time)
	systemconfigFields := schema.SystemConfig{}.Fields()
	_ = systemconfigFields
	// systemconfigDescDescription is the schema descriptor for description field.
	systemconfigDescDescription := systemconfigFields[2].Descriptor()
	// systemconfig.DefaultDescription holds the default value on creation for the description field.
	systemconfig.DefaultDescription = systemconfigDescDescription.Default.(string)
	// systemconfigDescValue is the schema descriptor for value field.
	systemconfigDescValue := systemconfigFields[3].Descriptor()
	// systemconfig.DefaultValue holds the default value on creation for the value field.
	systemconfig.DefaultValue = systemconfigDescValue.Default.(string)
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
