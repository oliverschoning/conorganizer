// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Menu() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<style>\n\t\t.main-menu-link {\n\t\t\tdisplay: flex;\n\t\t\talign-items: center;\n\t\t\tcolor: var(--color-secondary);\n\t\t\ttext-decoration: none;\n\t\t\ttext-transform: uppercase;\n\t\t\tmargin-left: 1rem;\n\t\t\tpadding: 0.5rem;\n\t\t}\n\t\t.main-menu-link:hover {\n\t\t\tbackground-color: rgba(255, 255, 255, 0.1);\n\t\t}\n\n\t\ta.disabled {\n\t\t\tpointer-events: none;\n\t\t\tcolor: #999;\n\t\t\tcursor: not-allowed;\n\t\t\ttext-decoration: none;\n\t\t\topacity: 0.6;\n\t\t}\n\n\t\tdetails {\n\t\t\tposition: relative;\n\t\t\tdisplay: inline-block;\n\t\t\tpadding-right: 2rem;\n\t\t}\n\n\t\tsummary {\n\t\t\tcursor: pointer;\n\t\t\tlist-style: none; /* hide the default disclosure triangle in some browsers */\n\t\t\tcolor: var(--color-secondary);\n\t\t\ttext-transform: uppercase;\n\t\t\ttext-decoration: none;\n\t\t\tmargin-left: 1rem;\n\t\t\tpadding: 0.5rem;\n\t\t}\n\t\tsummary::-webkit-details-marker {\n\t\t\tdisplay: none; /* hide default marker in Chrome/Safari */\n\t\t}\n\n\t\t.dropdown-panel {\n\t\t\tposition: absolute;\n\t\t\tright: 0;  \n\t\t\tleft: auto;\n\t\t\tbackground-color: #2d3748;\n\t\t\tmin-width: 160px;\n\t\t\tbox-shadow: 0 8px 16px rgba(0,0,0,0.2);\n\t\t\tmargin-top: 0.5rem;\n\t\t\tz-index: 999;\n\t\t}\n\t\t.dropdown-panel a {\n\t\t\tdisplay: block;\n\t\t\tcolor: var(--color-primary-text);\n\t\t\tpadding: 0.5rem 1rem;\n\t\t\ttext-decoration: none;\n\t\t\ttext-transform: uppercase;\n\t\t}\n\t\t.dropdown-panel a:hover {\n\t\t\tbackground-color: rgba(255, 255, 255, 0.1);\n\t\t}\n\t</style><nav style=\"\n\t\tdisplay: flex;\n\t\tjustify-content: space-between;\n\t\tbackground-color: #2d3748;\n\t\twidth: 100%;\n\t\tpadding: 0.25rem;\n\t\tmargin-bottom: 1rem;\"><div style=\"\n\t\t\tdisplay:flex;\n\t\t\talign-items: center;\"><a href=\"/\"><img src=\"/static/assets/RegnconLogo.svg\" alt=\"logo\" style=\"width: 2rem; height: 2rem;\"></a> <a href=\"/\" class=\"main-menu-link\"><img src=\"/static/assets/icons/house-solid.svg\" alt=\"hjem\" style=\"width: 1.25rem; height: 1.25rem; margin-right: 0.25rem;\"> Hjem</a></div><details><summary><img src=\"/static/assets/icons/bars-solid.svg\" alt=\"menu\" style=\"width: 1rem; height: 1rem; margin-right: 0.25rem;\"></summary><div class=\"dropdown-panel\"><a href=\"/logout\" class=\"disabled\">Logg ut</a> <a href=\"/mine\" class=\"disabled\">Mine arrangementer</a> <a href=\"/profil\" class=\"disabled\">Min profil</a> <a href=\"/admin\" class=\"disabled\">Admin</a> <a href=\"/arrangementer/edit\" class=\"disabled\">Rediger arrangement</a></div></details></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
