package message

import "errors"

func rawMessage2MessageContent(rawContent string) (*MessageContent, error) {
	// todo
	return nil, errors.New("unSupport")
}

func messageContent2RawMessage(content *MessageContent) (string, error) {
	if content == nil {
		return "", nil
	}
	// todo
	return "", errors.New("unSupport")
}
