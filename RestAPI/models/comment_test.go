package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateComment(t *testing.T) {
	testCase := []struct {
		name    string
		comment *Comment
		isValid bool
	}{
		{
			name: "default",
			comment: &Comment{
				Text:     "testText",
				UserID:   "1111111111111111111111111111",
				UserName: "testUser",
			},
			isValid: true,
		},
		{
			name: "invalid userID",
			comment: &Comment{
				Text:     "testText",
				UserID:   "11111111111",
				UserName: "testUser",
			},
			isValid: false,
		},
		{
			name: "invalid text",
			comment: &Comment{
				Text:     "",
				UserID:   "11111111111",
				UserName: "testUser",
			},
			isValid: false,
		},
		{
			name: "invalid user name",
			comment: &Comment{
				Text:     "testText",
				UserID:   "11111111111",
				UserName: "",
			},
			isValid: false,
		},
	}
	for _, item := range testCase {
		t.Run(item.name, func(t *testing.T) {
			if item.isValid {
				assert.NoError(t, item.comment.Validate())
			} else {
				assert.Error(t, item.comment.Validate())
			}
		})
	}
}

func TestBeforeCreateComment(t *testing.T) {
	comment := &Comment{
		Text:     "testText",
		UserID:   "1111111111111111111111111111",
		UserName: "testUser",
	}
	assert.NoError(t, comment.BeforeCreate())
	assert.NotNil(t, comment.UpdateAt)
	assert.NotNil(t, comment.CreateAt)
}
