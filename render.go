package main

import (
	"bytes"
	"log"
	"os"
	"text/template"
)

// Render 将朋友信息和最近的文章渲染成Markdown格式
func Render(friends []*Friend, posts []Post) []byte {
	markdownTemplate, err := os.ReadFile(Config.TemplatePath)
	if err != nil {
		log.Fatalf("Failed to read template file: %v", err)
	}

	tmpl, err := template.New("markdown").Parse(string(markdownTemplate))
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	var buffer bytes.Buffer
	data := struct {
		Friends []*Friend
		Posts   []Post
	}{
		Friends: friends,
		Posts:   posts,
	}

	if err := tmpl.Execute(&buffer, data); err != nil {
		log.Fatalf("Failed to render template: %v", err)
	}

	return buffer.Bytes()
}
