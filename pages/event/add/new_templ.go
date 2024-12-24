// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package add

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"database/sql"
	"fmt"
	"github.com/Regncon/conorganizer/components"
	"log"
	"net/http"
)

var validationErrors []string

func EventNew(w http.ResponseWriter, r *http.Request, db *sql.DB) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)

		if r.Method == http.MethodPost {
			validationErrors = nil
			log.Println("EventNew")
			title := r.FormValue("title")
			description := r.FormValue("description")

			if title == "" {
				validationErrors = append(validationErrors, "Title is required.")
			}
			if description == "" {
				validationErrors = append(validationErrors, "Description is required.")
			}
			fmt.Printf("Title: %s, Description: %s\n", title, description)

			if len(validationErrors) == 0 {
				_, err := db.Exec("INSERT INTO events (name, description) VALUES (?, ?)", title, description)
				if err != nil {
					log.Println("Failed to save event to the database.", err)
					validationErrors = append(validationErrors, "Failed to save event to the database.")
				} else {
					log.Println(w, "Event added successfully")
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Header("New Event").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body><h1>New Event</h1><ul>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, error := range validationErrors {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(error)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/event/add/new.templ`, Line: 46, Col: 16}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
