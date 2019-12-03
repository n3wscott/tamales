package update

import (
	"bufio"
	"fmt"
	"github.com/n3wscott/tomles/pkg/queue"
	"log"
	"os"
	"strings"
)

type Update struct {
	// TODO
	Filename string

	DryRun  bool
	Verbose bool

	Name     string
	Branch   string
	Version  string
	Revision string
}

type state struct {
	inConstraint bool
	match        bool
}

func (u *Update) constraint(indent string) string {
	key := ""
	value := ""
	if u.Branch != "" {
		key = "branch"
		value = u.Branch
	}
	if u.Version != "" {
		key = "version"
		value = u.Version
	}
	if u.Revision != "" {
		key = "revision"
		value = u.Revision
	}

	return fmt.Sprintf("%s%s = %q", indent, key, value)
}

func (u *Update) Do() error {
	logger := log.New(os.Stderr, "", 0)

	out := log.New(os.Stdout, "", 0)

	q := queue.New()

	// Wrap the import path
	if !strings.HasPrefix(u.Name, `"`) {
		u.Name = fmt.Sprintf(`"%s"`, u.Name)
	}
	u.Name = strings.ToLower(strings.TrimSpace(u.Name))

	if u.Verbose {
		logger.Printf("Params: \n%v\n", u)
	}

	file, err := os.Open(u.Filename)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	now := state{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if u.Verbose {
			logger.Println(line)
		}
		q.Enqueue(line)

		text := strings.TrimSpace(line)

		if len(text) == 0 {
			continue
		}

		if strings.HasPrefix(text, "[[") {
			// Drain
			for q.Len() > 0 {
				out.Print(q.Dequeue())
			}

			if text == "[[override]]" || text == "[[constraint]]" {
				now = state{
					inConstraint: true,
				}
			} else {
				now = state{
					inConstraint: false,
				}
			}

		} else if strings.HasPrefix(text, "[") {
			// Nothing...
		} else if strings.HasPrefix(text, "#") {
			// Nothing...
		} else {
			parts := strings.Split(text, " = ")
			if len(parts) != 2 {
				continue
			}
			key := strings.ToLower(strings.TrimSpace(parts[0]))
			value := strings.ToLower(strings.TrimSpace(parts[1]))
			switch key {
			case "name":
				if now.inConstraint && value == u.Name {
					now.match = true
				}
			case "branch", "version", "revision":
				if now.match {
					// Pop off old constraint.
					q.Pop()
					if u.Verbose {
						logger.Printf("replacing %s with %s", line, u.constraint(""))
					}
					indent := strings.Split(line, key)
					// Push back new constraint.
					q.Enqueue(u.constraint(indent[0]))
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Drain
	for q.Len() > 0 {
		out.Print(q.Dequeue())
	}

	return nil
}
