// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/wolodata/schema/ent/article"
	"github.com/wolodata/schema/ent/schema"
	"github.com/wolodata/schema/ent/tag"
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
	articleDescURL := articleFields[4].Descriptor()
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
	articleDescTitleChinese := articleFields[5].Descriptor()
	// article.DefaultTitleChinese holds the default value on creation for the title_chinese field.
	article.DefaultTitleChinese = articleDescTitleChinese.Default.(string)
	// articleDescTitleEnglish is the schema descriptor for title_english field.
	articleDescTitleEnglish := articleFields[6].Descriptor()
	// article.DefaultTitleEnglish holds the default value on creation for the title_english field.
	article.DefaultTitleEnglish = articleDescTitleEnglish.Default.(string)
	// articleDescAuthor is the schema descriptor for author field.
	articleDescAuthor := articleFields[7].Descriptor()
	// article.DefaultAuthor holds the default value on creation for the author field.
	article.DefaultAuthor = articleDescAuthor.Default.(string)
	// articleDescTags is the schema descriptor for tags field.
	articleDescTags := articleFields[8].Descriptor()
	// article.DefaultTags holds the default value on creation for the tags field.
	article.DefaultTags = articleDescTags.Default.([]string)
	// articleDescHTMLChinese is the schema descriptor for html_chinese field.
	articleDescHTMLChinese := articleFields[10].Descriptor()
	// article.DefaultHTMLChinese holds the default value on creation for the html_chinese field.
	article.DefaultHTMLChinese = articleDescHTMLChinese.Default.(string)
	// articleDescHTMLEnglish is the schema descriptor for html_english field.
	articleDescHTMLEnglish := articleFields[11].Descriptor()
	// article.DefaultHTMLEnglish holds the default value on creation for the html_english field.
	article.DefaultHTMLEnglish = articleDescHTMLEnglish.Default.(string)
	// articleDescTextChinese is the schema descriptor for text_chinese field.
	articleDescTextChinese := articleFields[12].Descriptor()
	// article.DefaultTextChinese holds the default value on creation for the text_chinese field.
	article.DefaultTextChinese = articleDescTextChinese.Default.(string)
	// articleDescTextEnglish is the schema descriptor for text_english field.
	articleDescTextEnglish := articleFields[13].Descriptor()
	// article.DefaultTextEnglish holds the default value on creation for the text_english field.
	article.DefaultTextEnglish = articleDescTextEnglish.Default.(string)
	// articleDescCrawledAt is the schema descriptor for crawled_at field.
	articleDescCrawledAt := articleFields[14].Descriptor()
	// article.DefaultCrawledAt holds the default value on creation for the crawled_at field.
	article.DefaultCrawledAt = articleDescCrawledAt.Default.(func() time.Time)
	// articleDescSummaryChinese is the schema descriptor for summary_chinese field.
	articleDescSummaryChinese := articleFields[15].Descriptor()
	// article.DefaultSummaryChinese holds the default value on creation for the summary_chinese field.
	article.DefaultSummaryChinese = articleDescSummaryChinese.Default.(string)
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescEnglish is the schema descriptor for english field.
	tagDescEnglish := tagFields[0].Descriptor()
	// tag.EnglishValidator is a validator for the "english" field. It is called by the builders before save.
	tag.EnglishValidator = tagDescEnglish.Validators[0].(func(string) error)
	// tagDescChinese is the schema descriptor for chinese field.
	tagDescChinese := tagFields[1].Descriptor()
	// tag.DefaultChinese holds the default value on creation for the chinese field.
	tag.DefaultChinese = tagDescChinese.Default.(string)
}
