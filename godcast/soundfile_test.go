package godcast

import (
	"log"
	"testing"
)

func TestGetTagMp3(t *testing.T) {
	filename := "./test/test.mp3"
	tag, err := GetTag(filename)
	if err != nil {
		log.Fatal("Error on get mp3 tag.")
	}
	if tag.Title != "テストtitle" {
		log.Fatal("Invalid mp3 title tag.")
	}
	if tag.Description != "no description" {
		log.Fatal("Invalid mp3 description tag.")
	}
}

func TestGetTagMp4(t *testing.T) {
	filename := "./test/test.mp4"
	tag, err := GetTag(filename)
	if err != nil {
		log.Fatal("Error on get mp3 tag.")
	}
	if tag.Title != "テストtitle" {
		log.Fatal("Invalid mp3 title tag.")
	}
}

func TestGetTagsInDir(t *testing.T) {
	dirname := "./test/"
	tags, err := GetTagsInDir(dirname)
	if err != nil {
		log.Fatal("Error on walking sound directory.", err)
	}
	if len(tags) != 2 {
		log.Fatal("Error files of sound directory.")
	}
	if tags[0].Title != "テストtitle" {
		log.Fatal("Error files of sound Title.")
	}
}
