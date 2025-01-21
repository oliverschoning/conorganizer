// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package index

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"database/sql"
	"fmt"
	"github.com/Regncon/conorganizer/components"
	"github.com/Regncon/conorganizer/models"
	"github.com/Regncon/conorganizer/pages/root"
	"log/slog"
)

type EventMVC struct {
	Events    []models.Event `json:"event"`
	Idx       int            `json:"editingIdx"`
	IsEditing bool           `json:"isEditing"`
}

func AllEventMVCView(mvc *EventMVC, db *sql.DB, logger *slog.Logger) templ.Component {
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
		events, err := root.GetEvents(db, logger)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div id=\"events-container\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if err != nil {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<p>Error fetching events: ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(err.Error())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/index/allEvents.templ`, Line: 22, Col: 42}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "</p>return")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<div class=\"events-grid\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ShowEvents(mvc, events, db, logger).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</div><code><pre data-text=\"ctx.signals.JSON()\"></pre></code></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func ShowEvents(mvc *EventMVC, events []models.Event, db *sql.DB, logger *slog.Logger) templ.Component {
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
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		fmt.Println("Id: ", mvc.Idx, mvc.Idx == -1)
		if mvc.Idx == -1 {
			for _, event := range events {
				templ_7745c5c3_Err = root.EventCard(event, mvc.Idx).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		} else {
			event, err := GetEvent(db, mvc.Idx, logger)
			if err != nil {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<p>Error fetching events: ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(err.Error())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/index/allEvents.templ`, Line: 43, Col: 42}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "</p>return")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, " ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = EventCard(event, mvc).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return nil
	})
}

func EventCard(event models.Event, mvc *EventMVC) templ.Component {
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
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "<style>\n\t\t.events-grid {\n\t\t\t--events-grid: 1fr;\n\t\t}\n\t\t.back-btn {\n\t\t\twidth: 100%;\n\t\t}\n\t\t/* Reset and base styles */\n\t\t* {\n\t\t\tmargin: 0;\n\t\t\tpadding: 0;\n\t\t\tbox-sizing: border-box;\n\t\t}\n\t\t/* Card styles */\n\t\t.card {\n\t\t\tborder-radius: 0.5rem;\n\t\t\tbox-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);\n\t\t\twidth: 100%;\n\t\t\toverflow: hidden;\n\t\t\tcursor: pointer;\n\t\t}\n\n\t\t.card-header {\n\t\t\tposition: relative;\n\t\t\theight: 16rem;\n\t\t}\n\n\t\t.card-image {\n\t\t\twidth: 100%;\n\t\t\theight: 100%;\n\t\t\tobject-fit: cover;\n\t\t}\n\n\t\t.card-title {\n\t\t\tposition: absolute;\n\t\t\ttop: 50%;\n\t\t\tleft: 50%;\n\t\t\ttransform: translate(-50%, -50%);\n\t\t\tcolor: white;\n\t\t\tfont-size: 3.5rem;\n\t\t\tfont-weight: bold;\n\t\t\ttext-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);\n\t\t}\n\n\t\t.card-content {\n\t\t\tpadding: 1.5rem;\n\t\t\tbackground-color: #21273d;\n\t\t}\n\n\t\t.card-grid {\n\t\t\tdisplay: grid;\n\t\t\tgrid-template-columns: repeat(2, 1fr);\n\t\t\tgap: 1.5rem;\n\t\t\tmargin-bottom: 1.5rem;\n\t\t}\n\n\t\t.card-field {\n\t\t\tdisplay: flex;\n\t\t\tflex-direction: column;\n\t\t\tgap: 0.5rem;\n\t\t}\n\n\t\t.field-label {\n\t\t\tcolor: #a9b4ca;\n\t\t\tfont-size: 0.875rem;\n\t\t\tfont-weight: 500;\n\t\t}\n\n\t\t.field-value {\n\t\t\tfont-size: 1.125rem;\n\t\t\tfont-weight: 500;\n\t\t}\n\n\t\t/* Modal styles */\n\t\t.modal {\n\t\t\tposition: fixed;\n\t\t\ttop: 0;\n\t\t\tleft: 0;\n\t\t\twidth: 100%;\n\t\t\theight: 100%;\n\t\t\tbackground: rgba(0, 0, 0, 0.5);\n\t\t\tdisplay: none;\n\t\t\tplace-items: center;\n\t\t\tpadding: 1rem;\n\t\t}\n\n\t\t.modal-content {\n\t\t\tbackground: white;\n\t\t\tpadding: 2rem;\n\t\t\tborder-radius: 0.5rem;\n\t\t\twidth: 100%;\n\t\t\tmax-width: 36rem;\n\t\t\tposition: relative;\n\t\t}\n\n\t\t.modal-close {\n\t\t\tposition: absolute;\n\t\t\ttop: 1rem;\n\t\t\tright: 1rem;\n\t\t\tbackground: none;\n\t\t\tborder: none;\n\t\t\tfont-size: 1.5rem;\n\t\t\tcursor: pointer;\n\t\t\tcolor: #6b7280;\n\t\t}\n\n\t\t/* Modal toggle functionality */\n\t\t#modal-toggle {\n\t\t\tdisplay: none;\n\t\t}\n\n\t\t#modal-toggle:checked + .modal {\n\t\t\tdisplay: grid;\n\t\t}\n\n\t\t/* Responsive adjustments */\n\t\t@media (max-width: 640px) {\n\t\t\t.card-grid {\n\t\t\t\tgrid-template-columns: 1fr;\n\t\t\t}\n\n\t\t\t.card-title {\n\t\t\t\tfont-size: 2.5rem;\n\t\t\t}\n\t\t}\n\t</style><div style=\"display: flex; place-content: space-between; margin-inline: 1rem;\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.BackButton().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.EditButton(event.ID).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "</div><!-- Card --><label class=\"card\"><div class=\"card-header\"><img src=\"/static/images/events/dice-small.webp\" alt=\"Colorful dice collection\" class=\"card-image\"> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if mvc.IsEditing {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "<input type=\"text\" class=\"card-title\" placeholder=\"Enter a title\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(event.Title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/index/allEvents.templ`, Line: 194, Col: 24}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "<h1 class=\"card-title\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(event.Title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/index/allEvents.templ`, Line: 197, Col: 40}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "</h1>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "</div><div class=\"card-content\"><div class=\"card-grid\"><div class=\"card-field\"><span class=\"field-label\">Game Master</span> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if mvc.IsEditing {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "<input type=\"text\" class=\"field-value\" placeholder=\"Enter a game master\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(event.GameMaster)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/index/allEvents.templ`, Line: 209, Col: 31}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "<p class=\"field-value\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(event.GameMaster)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/index/allEvents.templ`, Line: 212, Col: 47}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "</p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "</div><div class=\"card-field\"><span class=\"field-label\">System</span> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if mvc.IsEditing {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "<input type=\"text\" class=\"field-value\" placeholder=\"Enter a system\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(event.System)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/index/allEvents.templ`, Line: 222, Col: 27}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, "\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, "<p class=\"field-value\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var11 string
			templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(event.System)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/index/allEvents.templ`, Line: 225, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 24, "</p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 25, "</div></div><div class=\"card-field\"><span class=\"field-label\">Description</span> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if mvc.IsEditing {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 26, "<textarea class=\"field-value\" rows=\"4\" cols=\"50\" placeholder=\"Enter a description\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var12 string
			templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(event.ShortDescription)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/index/allEvents.templ`, Line: 237, Col: 30}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 27, "</textarea>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 28, "<p class=\"field-value\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(event.ShortDescription)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/index/allEvents.templ`, Line: 239, Col: 52}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 29, "</p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 30, "</div></div></label> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if mvc.IsEditing {
			templ_7745c5c3_Err = components.EditEvent(event).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
