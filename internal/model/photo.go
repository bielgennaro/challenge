package model

import (
	"time"

	"github.com/bielgennaro/v3-challenge-cloud/internal/errors"
	"github.com/google/uuid"
)

type Photo struct {
	ID         uuid.UUID
	MacAddress string
	FileURL    string
	IsMatch    bool
	Timestamp  time.Time
	CreatedAt  time.Time
}

type PhotoBuilder struct {
	macAddress string
	fileUrl    string
	isMatch    bool
	timestamp  time.Time
	err        error
}

func NewPhotoBuilder() *PhotoBuilder {
	return &PhotoBuilder{}
}

func (b *PhotoBuilder) WithMacAddress(macAddress string) *PhotoBuilder {
	if b.err != nil {
		return b
	}

	if macAddress == "" {
		b.err = errors.NewErrorBadRequest(
			"missing_device_id",
			"device identifier (MAC address) is required",
		)
		return b
	}

	b.macAddress = macAddress
	return b
}

func (b *PhotoBuilder) WithFileUrl(fileUrl string) *PhotoBuilder {
	if b.err != nil {
		return b
	}

	if fileUrl == "" {
		b.err = errors.NewErrorBadRequest(
			"missing_file_path",
			"file path for the photo is required",
		)
		return b
	}

	b.fileUrl = fileUrl
	return b
}

func (b *PhotoBuilder) WithRecognitionStatus(isMatch bool) *PhotoBuilder {
	if b.err != nil {
		return b
	}

	b.isMatch = isMatch
	return b
}

func (b *PhotoBuilder) WithTimestamp(timestamp time.Time) *PhotoBuilder {
	if b.err != nil {
		return b
	}

	if timestamp.IsZero() {
		b.err = errors.NewErrorBadRequest(
			"invalid_timestamp",
			"timestamp is required and must be valid",
		)
		return b
	}

	b.timestamp = timestamp
	return b
}

func (b *PhotoBuilder) Build() (*Photo, error) {
	if b.err != nil {
		return nil, b.err
	}

	return &Photo{
		ID:         uuid.New(),
		MacAddress: b.macAddress,
		FileURL:    b.fileUrl,
		IsMatch:    b.isMatch,
		Timestamp:  b.timestamp,
		CreatedAt:  time.Now(),
	}, nil
}
