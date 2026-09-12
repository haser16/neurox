package users_service

import (
	"context"
	"fmt"
	"io"

	"github.com/google/uuid"
)

type Avatar struct {
	File        io.Reader
	ContentType string
	Size        int64
}

func (s *UsersService) UploadAvatar(
	ctx context.Context,
	userID int64,
	avatar Avatar,
) error {
	if avatar.Size > 5<<20 {
		return fmt.Errorf("upload avatar too big")
	}

	switch avatar.ContentType {
	case "image/png", "image/jpeg", "image/gif", "image/webp":
	default:
		return fmt.Errorf("upload avatar content type not supported")
	}
	key := fmt.Sprintf("avatars/%d/%s", userID, uuid.NewString())

	if err := s.s3.Upload(
		ctx,
		key,
		avatar.File,
		avatar.ContentType,
	); err != nil {
		return fmt.Errorf("upload avatar to storage: %w", err)
	}

	if err := s.usersRepository.UpdateAvatarKey(
		ctx,
		userID,
		key,
	); err != nil {
		_ = s.s3.Delete(ctx, key)

		return fmt.Errorf("update avatar key: %w", err)
	}

	return nil
}
