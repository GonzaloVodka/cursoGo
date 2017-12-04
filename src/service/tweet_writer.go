package service

import (
	"os"

	"gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/domain"
)

type TweetWriter interface {
	WriteTweet(domain.Tweet)
}

type MemoryTweetWriter struct {
	Tweets []domain.Tweet
}

func NewMemoryTweetWriter() *MemoryTweetWriter {
	return new(MemoryTweetWriter)
}

func (writer *MemoryTweetWriter) WriteTweet(tweet domain.Tweet) {
	writer.Tweets = append(writer.Tweets, tweet)
}

type FileTweetWriter struct {
	file *os.File
}

func NewFileTweetWriter() *FileTweetWriter {

	file, _ := os.OpenFile(
		"tweets.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)

	writer := new(FileTweetWriter)
	writer.file = file

	return writer
}

func (writer *FileTweetWriter) WriteTweet(tweet domain.Tweet) {

	if writer.file != nil {
		byteSlice := []byte(tweet.PrintableTweet() + "\n")
		writer.file.Write(byteSlice)
	}
}

type ChannelTweetWriter struct {
	writer TweetWriter
}

func NewChannelTweetWriter(writer TweetWriter) *ChannelTweetWriter {
	channelTweetWriter := new(ChannelTweetWriter)
	channelTweetWriter.writer = writer
	return channelTweetWriter
}

func (channelWriter *ChannelTweetWriter) WriteTweet(tweetsToWrite chan domain.Tweet, quit chan bool) {

	tweet, open := <-tweetsToWrite

	for open {

		channelWriter.writer.WriteTweet(tweet)

		tweet, open = <-tweetsToWrite
	}

	quit <- true
}
