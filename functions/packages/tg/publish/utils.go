package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func rawJSONFromDO(in Request) ([]byte, error) {
	// если OpenWhisk положил тело в __ow_body (обычный случай для web=true)
	if in.OwBody != "" {
		// пробуем как обычный текст
		if strings.HasPrefix(strings.TrimSpace(in.OwBody), "{") {
			return []byte(in.OwBody), nil
		}
		// иначе это, скорее всего, base64
		b, err := base64.StdEncoding.DecodeString(in.OwBody)
		if err == nil {
			return b, nil
		}
		return nil, fmt.Errorf("cannot decode __ow_body")
	}
	// fallback: если вдруг body уже пришёл как JSON-объект
	if len(in.Body) > 0 {
		return in.Body, nil
	}
	return nil, fmt.Errorf("empty body")
}

func parseOwnerIDs(raw string) map[int64]bool {
	m := map[int64]bool{}
	for _, p := range strings.Split(raw, ",") {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		if id, err := strconv.ParseInt(p, 10, 64); err == nil {
			m[id] = true
		}
	}
	return m
}

func hasCmd(s, cmd string) bool {
	// допускаем "/cmd" или "/cmd@botname"
	if !strings.HasPrefix(s, cmd) {
		return false
	}
	rest := s[len(cmd):]
	return rest == "" || rest[0] == ' ' || rest[0] == '@'
}

func extractPayload(s string) string {
	// поддерживает "/cmd payload" и "/cmd\npayload"
	if i := strings.IndexAny(s, " \n"); i >= 0 && i+1 < len(s) {
		return strings.TrimSpace(s[i+1:])
	}
	return ""
}

func sendMessage(chatID int64, text string) error {
	return sendMessageStr(strconv.FormatInt(chatID, 10), text)
}

func sendMessageStr(chatIDStr string, text string) error {
	id, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		return fmt.Errorf("bad chat id")
	}

	req := TgSendMessage{
		ChatID:    id,
		Text:      text,
		ParseMode: "Markdown",
	}
	b, _ := json.Marshal(req)
	resp, err := http.Post(apiURL+"/sendMessage", "application/json", bytes.NewReader(b))

	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("failed to close response body:", err)
		}
	}(resp.Body)

	if resp.StatusCode >= 300 {
		return fmt.Errorf("telegram status %d", resp.StatusCode)
	}
	return nil
}
