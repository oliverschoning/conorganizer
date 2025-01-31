// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func BackButton() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<svg style=\"display: none;\"><symbol id=\"arrow-left\" viewBox=\"0 0 24 24\"><g fill=\"none\" stroke=\"#000\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\"><path stroke-dasharray=\"16\" stroke-dashoffset=\"16\" d=\"M19 12h-13.5\"><animate fill=\"freeze\" attributeName=\"stroke-dashoffset\" dur=\"0.2s\" values=\"16;0\"></animate></path> <path stroke-dasharray=\"10\" stroke-dashoffset=\"10\" d=\"M5 12l5 5M5 12l5 -5\"><animate fill=\"freeze\" attributeName=\"stroke-dashoffset\" begin=\"0.2s\" dur=\"0.2s\" values=\"10;0\"></animate></path></g></symbol></svg><style>\r\n\t.back-btn {\r\n\t\tall: unset;\r\n\t}\r\n\t.svg-back-btn {\r\n\t\tmargin-inline-end: 0.5rem;\r\n\t}\r\n\r\n\t.back-btn {\r\n\t\tposition: relative;\r\n\t\tdisplay: inline-flex;\r\n\t\theight: 3.5rem;\r\n\t\talign-items: center;\r\n\t\tborder-radius: 9999px;\r\n\t\tpadding-left: 0.5rem;\r\n\t\tpadding-right: 0.5rem;\r\n\t\tfont-family: Segoe UI;\r\n\t\tfont-size: 1.2rem;\r\n\t\tfont-weight: 640;\r\n\t\tcolor: #fafaf6;\r\n\t\tletter-spacing: -0.06em;\r\n\t\twidth: 6rem;\r\n\t\tplace-content: center;\r\n\t\tgap: 1rem;\r\n\t}\r\n\r\n\t.back-btn-item {\r\n\t\tbackground-color: transparent;\r\n\t\tcolor: #1d1d1f;\r\n\t}\r\n\r\n\t.back-btn-item .back-btn-bg {\r\n\t\tborder-color: rgba(255, 208, 116);\r\n\t\tbackground-color: rgba(255, 208, 116);\r\n\t}\r\n\r\n\t.back-btn-inner,\r\n\t.back-btn-inner-hover,\r\n\t.back-btn-inner-static {\r\n\t\tpointer-events: none;\r\n\t\tdisplay: block;\r\n\t}\r\n\r\n\t.back-btn-inner-hover,\r\n\t.back-btn-inner-static {\r\n\t\tdisplay: flex;\r\n\t\tplace-items: space-between;\r\n\t}\r\n\r\n\t.back-btn-inner {\r\n\t\tposition: relative;\r\n\t}\r\n\r\n\t.back-btn-inner-hover {\r\n\t\tposition: absolute;\r\n\t\ttop: 0;\r\n\t\tleft: 0;\r\n\t\topacity: 0;\r\n\t\ttransform: translateY(70%);\r\n\t}\r\n\r\n\t.back-btn-bg {\r\n\t\toverflow: hidden;\r\n\t\tborder-radius: 2rem;\r\n\t\tposition: absolute;\r\n\t\ttop: 0;\r\n\t\tleft: 0;\r\n\t\twidth: 100%;\r\n\t\theight: 100%;\r\n\t\ttransform: scale(1);\r\n\t\ttransition: transform 1.8s cubic-bezier(0.19, 1, 0.22, 1);\r\n\t}\r\n\r\n\t.back-btn-bg,\r\n\t.back-btn-bg-layer,\r\n\t.back-btn-bg-layers {\r\n\t\tdisplay: block;\r\n\t}\r\n\r\n\t.back-btn-bg-layers {\r\n\t\tposition: absolute;\r\n\t\tleft: 50%;\r\n\t\ttransform: translate(-50%);\r\n\t\ttop: -60%;\r\n\t\taspect-ratio: 1 / 1;\r\n\t\twidth: max(200%, 10rem);\r\n\t}\r\n\r\n\t.back-btn-bg-layer {\r\n\t\tborder-radius: 9999px;\r\n\t\tposition: absolute;\r\n\t\ttop: 0;\r\n\t\tleft: 0;\r\n\t\twidth: 100%;\r\n\t\theight: 100%;\r\n\t\ttransform: scale(0);\r\n\t}\r\n\r\n\t.back-btn-bg-layer.-purple {\r\n\t\tbackground-color: rgba(163, 116, 255);\r\n\t}\r\n\r\n\t.back-btn-bg-layer.-turquoise {\r\n\t\tbackground-color: rgba(23, 241, 209);\r\n\t}\r\n\r\n\t.back-btn-bg-layer.-yellow {\r\n\t\t--tw-bg-opacity: 1;\r\n\t\tbackground-color: rgba(255, 208, 116, var(--tw-bg-opacity));\r\n\t}\r\n\r\n\t.back-btn:hover .back-btn-inner-static {\r\n\t\topacity: 0;\r\n\t\ttransform: translateY(-70%);\r\n\t\ttransition:\r\n\t\t\ttransform 1.4s cubic-bezier(0.19, 1, 0.22, 1),\r\n\t\t\topacity 0.3s linear;\r\n\t}\r\n\r\n\t.back-btn:hover .back-btn-inner-hover {\r\n\t\topacity: 1;\r\n\t\ttransform: translateY(0);\r\n\t\ttransition:\r\n\t\t\ttransform 1.4s cubic-bezier(0.19, 1, 0.22, 1),\r\n\t\t\topacity 1.4s cubic-bezier(0.19, 1, 0.22, 1);\r\n\t}\r\n\r\n\t.back-btn:hover .back-btn-bg-layer {\r\n\t\ttransition:\r\n\t\t\ttransform 1.3s cubic-bezier(0.19, 1, 0.22, 1),\r\n\t\t\topacity 0.3s linear;\r\n\t}\r\n\r\n\t.back-btn:hover .back-btn-bg-layer-1 {\r\n\t\ttransform: scale(1);\r\n\t}\r\n\r\n\t.back-btn:hover .back-btn-bg-layer-2 {\r\n\t\ttransition-delay: 0.1s;\r\n\t\ttransform: scale(1);\r\n\t}\r\n\r\n\t.back-btn:hover .back-btn-bg-layer-3 {\r\n\t\ttransition-delay: 0.2s;\r\n\t\ttransform: scale(1);\r\n\t}\r\n\r\n\t</style><button class=\"back-btn back-btn-item\" data-on-click=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs("@put('/event/back')")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/backButton.templ`, Line: 175, Col: 77}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "\"><span class=\"back-btn-bg\"><span class=\"back-btn-bg-layers\"><span class=\"back-btn-bg-layer back-btn-bg-layer-1 -purple\"></span> <span class=\"back-btn-bg-layer back-btn-bg-layer-2 -turquoise\"></span> <span class=\"back-btn-bg-layer back-btn-bg-layer-3 -yellow\"></span></span></span> <span class=\"back-btn-inner\"><span class=\"back-btn-inner-static\"><svg width=\"30\" height=\"30\" version=\"1.1\" class=\"svg-back-btn\"><use href=\"#arrow-left\"></use></svg> back</span> <span class=\"back-btn-inner-hover\"><svg width=\"30\" height=\"30\" version=\"1.1\" class=\"svg-back-btn\"><use href=\"#arrow-left\"></use></svg> back</span></span></button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
