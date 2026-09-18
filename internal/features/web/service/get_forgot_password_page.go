package web_service

import (
	"fmt"
	"os"
	"path"
)

func (s *WebService) GetForgotPasswordPage() ([]byte, error) {
	htmlFilePath := path.Join(
		os.Getenv("PROJECT_ROOT"),
		"/public/forgot-password.html",
	)
	html, err := s.webRepository.GetFile(htmlFilePath)
	if err != nil {
		return nil, fmt.Errorf("get file from repository: %w", err)
	}
	return html, nil
}
