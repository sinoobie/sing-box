package provider

import (
	"fmt"
	"os"
	"strings"
	"time"

	E "github.com/sagernet/sing/common/exceptions"
)

func saveCache(file string, c *fileContent) error {
	w, err := os.Create(file)
	if err != nil {
		return err
	}
	defer w.Close()
	_, err = w.WriteString(fmt.Sprintf(
		"# Updated: %s\n# Links Hash: %s\n",
		c.updated.Format(time.RFC3339), c.linksHash,
	))
	if err != nil {
		return err
	}
	_, err = w.WriteString(c.links)
	if err != nil {
		return err
	}
	return nil
}

func saveCacheIfNeed(file string, content *fileContent) error {
	if content.links == "" {
		return nil
	}
	saved, _ := loadCache(file)
	if saved == nil || saved.linksHash != content.linksHash {
		return saveCache(file, content)
	}
	return nil
}

func loadCache(file string) (*fileContent, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	fc := &fileContent{}
	lines := strings.Split(string(content), "\n")
	links := make([]string, 0, len(lines))
	for _, line := range lines {
		if !strings.HasPrefix(line, "#") {
			links = append(links, line)
			continue
		}
		var name, value string
		parts := strings.SplitN(line[1:], ":", 2)
		if len(parts) != 2 {
			return nil, E.New("invalid header line: ", line)
		}
		name = strings.ToLower(strings.TrimSpace(parts[0]))
		value = strings.TrimSpace(parts[1])
		switch name {
		case "updated":
			fc.updated, err = time.Parse(time.RFC3339, value)
			if err != nil {
				return nil, err
			}
		case "links hash":
			fc.linksHash = value
		}
	}
	if fc.linksHash == "" {
		return nil, E.New("invalid cache file")
	}
	fc.links = strings.Join(links, "\n")
	return fc, nil
}
