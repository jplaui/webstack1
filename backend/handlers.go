package backend

import (
	"bytes"
	"fmt"
	"time"

	ctx "github.com/jplaui/templ-fiber/context"
	"github.com/jplaui/templ-fiber/frontend"
	"github.com/jplaui/templ-fiber/scanner"

	"github.com/gofiber/fiber/v2"
	"github.com/yuin/goldmark"
)

type IndexPage struct {
	StateString string
}

func NewIndexPage(state string) *IndexPage {
	return &IndexPage{StateString: state}
}

func (ip *IndexPage) GetWelcome(c *fiber.Ctx) error {

	fmt.Println("show arg access", ip.StateString)

	ch := HandleComponent(frontend.Index("hallo string"))
	return Render(c, ch)
}

type ScanPage struct {
	StateString string
}

func NewScanPage(state string) *ScanPage {
	return &ScanPage{StateString: state}
}

func (sp *ScanPage) PostScan(c *fiber.Ctx) error {

	fmt.Println("show arg access", sp.StateString)

	// get form params
	nbScans := c.FormValue("numberscans")
	apiLists := c.FormValue("apilistname")
	checkbox := c.FormValue("checkboxValue")

	// call scanner
	scanner.StartScan(apiLists)
	// should return "your html"

	fmt.Println("show form params", nbScans, apiLists, checkbox)

	// ****** replace this logic from here
	// get scan resolts here
	post := getSamplePost()

	// Convert the markdown to HTML, and pass it to the template.
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(post.Content), &buf); err != nil {
		ch := HandleComponent(frontend.Results(apiLists, nil))
		ch = SetRequestCtx(ch, ctx.ErrorStateCtx, "Your scan failed :(")
		return Render(c, ch)
	}
	// ****** replace this logic until here

	// Create an unsafe component containing raw HTML.
	// change to: content := Unsafe("your html" as string)
	content := Unsafe(buf.String())

	ch := HandleComponent(frontend.Results(apiLists, content))
	ch = SetRequestCtx(ch, ctx.SuccessStateCtx, "Your scan was successful :)")
	return Render(c, ch)
}

// example data model
type Post struct {
	Date    string
	Title   string
	Content string
}

func getSamplePost() Post {
	return Post{
		Date:  time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC).Format("02.01.2006"),
		Title: "Happy New Year!",
		Content: `New Year is a widely celebrated occasion in the United Kingdom, marking the end of one year and the beginning of another.

		Top New Year Activities in the UK include:

* Attending a Hogmanay celebration in Scotland
* Taking part in a local First-Foot tradition in Scotland and Northern England
* Setting personal resolutions and goals for the upcoming year
* Going for a New Year's Day walk to enjoy the fresh start
* Visiting a local pub for a celebratory toast and some cheer
`,
	}
}
