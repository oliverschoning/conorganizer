// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package event

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/Regncon/conorganizer/components"

func Page(name string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Header("Event name").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Cthulhu er nyforelsket</title><link rel=\"stylesheet\" href=\"styles.css\"></head><style>\n\n.container {\n    max-width: 400px;\n    margin: 0 auto;\n    padding: 20px;\n    background-color: #3b4252;\n    border-radius: 10px;\n    box-shadow: 0 0 10px rgba(0, 0, 0, 0.5);\n}\n\n.header h1 {\n    text-align: center;\n    font-size: 24px;\n    color: #eceff4;\n    margin-bottom: 20px;\n}\n\n.details {\n    display: flex;\n    justify-content: space-between;\n    align-items: center;\n    margin-bottom: 20px;\n}\n\n.info p {\n    margin: 5px 0;\n}\n\n.tags {\n    margin-top: 10px;\n}\n\n.tag {\n    display: inline-block;\n    margin-right: 5px;\n    padding: 5px 10px;\n    background-color: #bf616a;\n    border-radius: 5px;\n    font-size: 12px;\n    color: #eceff4;\n}\n\n.profile {\n    text-align: center;\n}\n\n.profile-pic {\n    width: 60px;\n    height: 60px;\n    border-radius: 50%;\n    background-color: #d8dee9;\n    margin-bottom: 5px;\n}\n\n.interest {\n    margin-bottom: 20px;\n}\n\n.interest-level {\n    display: flex;\n    flex-direction: column;\n    align-items: center;\n    margin-bottom: 10px;\n}\n\n.level-bar {\n    background-color: #bf616a;\n    padding: 10px;\n    width: 100%;\n    text-align: center;\n    border-radius: 5px;\n    margin-bottom: 10px;\n}\n\ninput[type=\"range\"] {\n    width: 100%;\n}\n\n.hint {\n    text-align: center;\n    font-size: 14px;\n}\n\n.confused-link {\n    display: block;\n    text-align: center;\n    font-size: 12px;\n    color: #88c0d0;\n    text-decoration: none;\n}\n\n.action {\n    text-align: center;\n    margin-bottom: 20px;\n}\n\n.next-btn, .admin-btn {\n    display: block;\n    width: 100%;\n    margin-bottom: 10px;\n    padding: 10px;\n    background-color: #5e81ac;\n    border: none;\n    color: #eceff4;\n    font-size: 14px;\n    border-radius: 5px;\n    cursor: pointer;\n}\n\n.admin-btn {\n    background-color: #bf616a;\n}\n\n.footer {\n    text-align: center;\n    font-size: 14px;\n    color: #d8dee9;\n}\n</style><body><div class=\"container\"><header class=\"header\"><h1>Cthulhu er nyforelsket</h1></header><section class=\"details\"><div class=\"info\"><p><strong>Gamemaster:</strong> Someone Disturbed</p><p><strong>System:</strong> Cat Call Of Cthulhu!</p><div class=\"tags\"><span class=\"tag\">Can be run in English</span> <span class=\"tag\">Kun for voksne (18+)</span> <span class=\"tag\">6+</span></div></div><div class=\"profile\"><img src=\"/static/images/events/dice-small.webp\" alt=\"Profile\" class=\"profile-pic\"><p><strong>Gerhard</strong></p></div></section><section class=\"interest\"><div class=\"interest-level\"><div class=\"level-bar\"><span>Fredag Kveld Kl 18 - 23</span> <span>Veldig interessert</span></div><input type=\"range\" min=\"1\" max=\"10\" value=\"7\"></div><p class=\"hint\">Dra baren over for å melde din interesse!</p><a href=\"#\" class=\"confused-link\">Forvirret? Les mer om påmeldingssystemet</a></section><section class=\"action\"><button class=\"next-btn\">Neste →</button> <button class=\"admin-btn\">ADMINISTRER ARRANGEMENTET</button></section><footer class=\"footer\"><p>For å ødelegge verden starter Cthulhu på high scool i en magisk verden. Der møter han Mui.</p></footer></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
