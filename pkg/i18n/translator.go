package translator

import (
	"encoding/json"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/language"
)

type Translator struct {
	bundle *i18n.Bundle
}

func NewTranslator() *Translator {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	dirPath := "resources/locales/"
	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(d.Name(), ".json") {
			_, err = bundle.LoadMessageFile(path)
			return err
		}
		return nil
	})
	if err != nil {
		log.Errorf("failed to load i18n files with error: %s", err)
		return nil
	}
	return &Translator{bundle: bundle}
}

func (tr Translator) Localize(msg string, langs ...string) string {
	lang := language.English.String()
	if len(langs) > 0 {
		for _, t := range tr.bundle.LanguageTags() {
			if t.String() == langs[0] {
				lang = t.String()
				break
			}
		}
	}
	localize := i18n.NewLocalizer(tr.bundle, lang)
	message, err := localize.LocalizeMessage(&i18n.Message{ID: msg})
	if err != nil {
		return msg
	}
	return message
}
