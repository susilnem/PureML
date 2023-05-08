var cr=r=>typeof r=="boolean"?`${r}`:r===0?"0":r,I=r=>!r||typeof r!="object"||Object.keys(r).length===0,zr=(...r)=>[...r].flat().filter(Boolean),gr=(r,e)=>{let t={};for(let o in r)e!=null&&e.hasOwnProperty(o)?t[o]=typeof r[o]=="object"?gr(r[o],e[o]):e[o]+" "+r[o]:t[o]=r[o];for(let o in e)t.hasOwnProperty(o)||(t[o]=e[o]);return t},ur=r=>!r||typeof r!="string"?r:r.replace(/\s+/g," ").trim();function Gr(){for(var r=0,e,t,o="";r<arguments.length;)(e=arguments[r++])&&(t=mr(e))&&(o&&(o+=" "),o+=t);return o}function mr(r){if(typeof r=="string")return r;for(var e,t="",o=0;o<r.length;o++)r[o]&&(e=mr(r[o]))&&(t&&(t+=" "),t+=e);return t}var sr="-";function Sr(r){var e=Ir(r),t=r.conflictingClassGroups,o=r.conflictingClassGroupModifiers,a=o===void 0?{}:o;function l(n){var f=n.split(sr);return f[0]===""&&f.length!==1&&f.shift(),hr(f,e)||jr(n)}function i(n,f){var p=t[n]||[];return f&&a[n]?[].concat(p,a[n]):p}return{getClassGroupId:l,getConflictingClassGroupIds:i}}function hr(r,e){var i;if(r.length===0)return e.classGroupId;var t=r[0],o=e.nextPart.get(t),a=o?hr(r.slice(1),o):void 0;if(a)return a;if(e.validators.length!==0){var l=r.join(sr);return(i=e.validators.find(function(n){var f=n.validator;return f(l)}))==null?void 0:i.classGroupId}}var pr=/^\[(.+)\]$/;function jr(r){if(pr.test(r)){var e=pr.exec(r)[1],t=e==null?void 0:e.substring(0,e.indexOf(":"));if(t)return"arbitrary.."+t}}function Ir(r){var e=r.theme,t=r.prefix,o={nextPart:new Map,validators:[]},a=Rr(Object.entries(r.classGroups),t);return a.forEach(function(l){var i=l[0],n=l[1];ir(n,o,i,e)}),o}function ir(r,e,t,o){r.forEach(function(a){if(typeof a=="string"){var l=a===""?e:fr(e,a);l.classGroupId=t;return}if(typeof a=="function"){if(Pr(a)){ir(a(o),e,t,o);return}e.validators.push({validator:a,classGroupId:t});return}Object.entries(a).forEach(function(i){var n=i[0],f=i[1];ir(f,fr(e,n),t,o)})})}function fr(r,e){var t=r;return e.split(sr).forEach(function(o){t.nextPart.has(o)||t.nextPart.set(o,{nextPart:new Map,validators:[]}),t=t.nextPart.get(o)}),t}function Pr(r){return r.isThemeGetter}function Rr(r,e){return e?r.map(function(t){var o=t[0],a=t[1],l=a.map(function(i){return typeof i=="string"?e+i:typeof i=="object"?Object.fromEntries(Object.entries(i).map(function(n){var f=n[0],p=n[1];return[e+f,p]})):i});return[o,l]}):r}function Tr(r){if(r<1)return{get:function(){},set:function(){}};var e=0,t=new Map,o=new Map;function a(l,i){t.set(l,i),e++,e>r&&(e=0,o=t,t=new Map)}return{get:function(i){var n=t.get(i);if(n!==void 0)return n;if((n=o.get(i))!==void 0)return a(i,n),n},set:function(i,n){t.has(i)?t.set(i,n):a(i,n)}}}var yr="!";function Er(r){var e=r.separator||":",t=e.length===1,o=e[0],a=e.length;return function(i){for(var n=[],f=0,p=0,b,g=0;g<i.length;g++){var A=i[g];if(f===0){if(A===o&&(t||i.slice(g,g+a)===e)){n.push(i.slice(p,g)),p=g+a;continue}if(A==="/"){b=g;continue}}A==="["?f++:A==="]"&&f--}var v=n.length===0?i:i.substring(p),z=v.startsWith(yr),k=z?v.substring(1):v,S=b&&b>p?b-p:void 0;return{modifiers:n,hasImportantModifier:z,baseClassName:k,maybePostfixModifierPosition:S}}}function Nr(r){if(r.length<=1)return r;var e=[],t=[];return r.forEach(function(o){var a=o[0]==="[";a?(e.push.apply(e,t.sort().concat([o])),t=[]):t.push(o)}),e.push.apply(e,t.sort()),e}function Or(r){return{cache:Tr(r.cacheSize),splitModifiers:Er(r),...Sr(r)}}var Wr=/\s+/;function Lr(r,e){var t=e.splitModifiers,o=e.getClassGroupId,a=e.getConflictingClassGroupIds,l=new Set;return r.trim().split(Wr).map(function(i){var n=t(i),f=n.modifiers,p=n.hasImportantModifier,b=n.baseClassName,g=n.maybePostfixModifierPosition,A=o(g?b.substring(0,g):b),v=!!g;if(!A){if(!g)return{isTailwindClass:!1,originalClassName:i};if(A=o(b),!A)return{isTailwindClass:!1,originalClassName:i};v=!1}var z=Nr(f).join(":"),k=p?z+yr:z;return{isTailwindClass:!0,modifierId:k,classGroupId:A,originalClassName:i,hasPostfixModifier:v}}).reverse().filter(function(i){if(!i.isTailwindClass)return!0;var n=i.modifierId,f=i.classGroupId,p=i.hasPostfixModifier,b=n+f;return l.has(b)?!1:(l.add(b),a(f,p).forEach(function(g){return l.add(n+g)}),!0)}).reverse().map(function(i){return i.originalClassName}).join(" ")}function ar(){for(var r=arguments.length,e=new Array(r),t=0;t<r;t++)e[t]=arguments[t];var o,a,l,i=n;function n(p){var b=e[0],g=e.slice(1),A=g.reduce(function(v,z){return z(v)},b());return o=Or(A),a=o.cache.get,l=o.cache.set,i=f,f(p)}function f(p){var b=a(p);if(b)return b;var g=Lr(p,o);return l(p,g),g}return function(){return i(Gr.apply(null,arguments))}}function m(r){var e=function(o){return o[r]||[]};return e.isThemeGetter=!0,e}var wr=/^\[(?:([a-z-]+):)?(.+)\]$/i,Vr=/^\d+\/\d+$/,Br=new Set(["px","full","screen"]),Ur=/^(\d+(\.\d+)?)?(xs|sm|md|lg|xl)$/,Fr=/\d+(%|px|r?em|[sdl]?v([hwib]|min|max)|pt|pc|in|cm|mm|cap|ch|ex|r?lh|cq(w|h|i|b|min|max))/,qr=/^-?((\d+)?\.?(\d+)[a-z]+|0)_-?((\d+)?\.?(\d+)[a-z]+|0)/;function P(r){return O(r)||Br.has(r)||Vr.test(r)||T(r)}function T(r){return W(r,"length",Hr)}function Zr(r){return W(r,"size",xr)}function Xr(r){return W(r,"position",xr)}function Jr(r){return W(r,"url",_r)}function tr(r){return W(r,"number",O)}function O(r){return!Number.isNaN(Number(r))}function Qr(r){return r.endsWith("%")&&O(r.slice(0,-1))}function H(r){return br(r)||W(r,"number",br)}function w(r){return wr.test(r)}function _(){return!0}function R(r){return Ur.test(r)}function Yr(r){return W(r,"",Dr)}function W(r,e,t){var o=wr.exec(r);return o?o[1]?o[1]===e:t(o[2]):!1}function Hr(r){return Fr.test(r)}function xr(){return!1}function _r(r){return r.startsWith("url(")}function br(r){return Number.isInteger(Number(r))}function Dr(r){return qr.test(r)}function lr(){var r=m("colors"),e=m("spacing"),t=m("blur"),o=m("brightness"),a=m("borderColor"),l=m("borderRadius"),i=m("borderSpacing"),n=m("borderWidth"),f=m("contrast"),p=m("grayscale"),b=m("hueRotate"),g=m("invert"),A=m("gap"),v=m("gradientColorStops"),z=m("gradientColorStopPositions"),k=m("inset"),S=m("margin"),G=m("opacity"),M=m("padding"),K=m("saturate"),C=m("scale"),U=m("sepia"),F=m("skew"),q=m("space"),Z=m("translate"),X=function(){return["auto","contain","none"]},L=function(){return["auto","hidden","clip","visible","scroll"]},$=function(){return["auto",e]},rr=function(){return["",P]},V=function(){return["auto",O,w]},J=function(){return["bottom","center","left","left-bottom","left-top","right","right-bottom","right-top","top"]},E=function(){return["solid","dashed","dotted","double","none"]},Q=function(){return["normal","multiply","screen","overlay","darken","lighten","color-dodge","color-burn","hard-light","soft-light","difference","exclusion","hue","saturation","color","luminosity","plus-lighter"]},Y=function(){return["start","end","center","between","around","evenly","stretch"]},N=function(){return["","0",w]},s=function(){return["auto","avoid","all","avoid-page","page","left","right","column"]},c=function(){return[O,tr]},u=function(){return[O,w]};return{cacheSize:500,theme:{colors:[_],spacing:[P],blur:["none","",R,T],brightness:c(),borderColor:[r],borderRadius:["none","","full",R,T],borderSpacing:[e],borderWidth:rr(),contrast:c(),grayscale:N(),hueRotate:u(),invert:N(),gap:[e],gradientColorStops:[r],gradientColorStopPositions:[Qr,T],inset:$(),margin:$(),opacity:c(),padding:[e],saturate:c(),scale:c(),sepia:N(),skew:u(),space:[e],translate:[e]},classGroups:{aspect:[{aspect:["auto","square","video",w]}],container:["container"],columns:[{columns:[R]}],"break-after":[{"break-after":s()}],"break-before":[{"break-before":s()}],"break-inside":[{"break-inside":["auto","avoid","avoid-page","avoid-column"]}],"box-decoration":[{"box-decoration":["slice","clone"]}],box:[{box:["border","content"]}],display:["block","inline-block","inline","flex","inline-flex","table","inline-table","table-caption","table-cell","table-column","table-column-group","table-footer-group","table-header-group","table-row-group","table-row","flow-root","grid","inline-grid","contents","list-item","hidden"],float:[{float:["right","left","none"]}],clear:[{clear:["left","right","both","none"]}],isolation:["isolate","isolation-auto"],"object-fit":[{object:["contain","cover","fill","none","scale-down"]}],"object-position":[{object:[].concat(J(),[w])}],overflow:[{overflow:L()}],"overflow-x":[{"overflow-x":L()}],"overflow-y":[{"overflow-y":L()}],overscroll:[{overscroll:X()}],"overscroll-x":[{"overscroll-x":X()}],"overscroll-y":[{"overscroll-y":X()}],position:["static","fixed","absolute","relative","sticky"],inset:[{inset:[k]}],"inset-x":[{"inset-x":[k]}],"inset-y":[{"inset-y":[k]}],start:[{start:[k]}],end:[{end:[k]}],top:[{top:[k]}],right:[{right:[k]}],bottom:[{bottom:[k]}],left:[{left:[k]}],visibility:["visible","invisible","collapse"],z:[{z:["auto",H]}],basis:[{basis:[e]}],"flex-direction":[{flex:["row","row-reverse","col","col-reverse"]}],"flex-wrap":[{flex:["wrap","wrap-reverse","nowrap"]}],flex:[{flex:["1","auto","initial","none",w]}],grow:[{grow:N()}],shrink:[{shrink:N()}],order:[{order:["first","last","none",H]}],"grid-cols":[{"grid-cols":[_]}],"col-start-end":[{col:["auto",{span:[H]},w]}],"col-start":[{"col-start":V()}],"col-end":[{"col-end":V()}],"grid-rows":[{"grid-rows":[_]}],"row-start-end":[{row:["auto",{span:[H]},w]}],"row-start":[{"row-start":V()}],"row-end":[{"row-end":V()}],"grid-flow":[{"grid-flow":["row","col","dense","row-dense","col-dense"]}],"auto-cols":[{"auto-cols":["auto","min","max","fr",w]}],"auto-rows":[{"auto-rows":["auto","min","max","fr",w]}],gap:[{gap:[A]}],"gap-x":[{"gap-x":[A]}],"gap-y":[{"gap-y":[A]}],"justify-content":[{justify:["normal"].concat(Y())}],"justify-items":[{"justify-items":["start","end","center","stretch"]}],"justify-self":[{"justify-self":["auto","start","end","center","stretch"]}],"align-content":[{content:["normal"].concat(Y(),["baseline"])}],"align-items":[{items:["start","end","center","baseline","stretch"]}],"align-self":[{self:["auto","start","end","center","stretch","baseline"]}],"place-content":[{"place-content":[].concat(Y(),["baseline"])}],"place-items":[{"place-items":["start","end","center","baseline","stretch"]}],"place-self":[{"place-self":["auto","start","end","center","stretch"]}],p:[{p:[M]}],px:[{px:[M]}],py:[{py:[M]}],ps:[{ps:[M]}],pe:[{pe:[M]}],pt:[{pt:[M]}],pr:[{pr:[M]}],pb:[{pb:[M]}],pl:[{pl:[M]}],m:[{m:[S]}],mx:[{mx:[S]}],my:[{my:[S]}],ms:[{ms:[S]}],me:[{me:[S]}],mt:[{mt:[S]}],mr:[{mr:[S]}],mb:[{mb:[S]}],ml:[{ml:[S]}],"space-x":[{"space-x":[q]}],"space-x-reverse":["space-x-reverse"],"space-y":[{"space-y":[q]}],"space-y-reverse":["space-y-reverse"],w:[{w:["auto","min","max","fit",e]}],"min-w":[{"min-w":["min","max","fit",P]}],"max-w":[{"max-w":["0","none","full","min","max","fit","prose",{screen:[R]},R,T]}],h:[{h:[e,"auto","min","max","fit"]}],"min-h":[{"min-h":["min","max","fit",P]}],"max-h":[{"max-h":[e,"min","max","fit"]}],"font-size":[{text:["base",R,T]}],"font-smoothing":["antialiased","subpixel-antialiased"],"font-style":["italic","not-italic"],"font-weight":[{font:["thin","extralight","light","normal","medium","semibold","bold","extrabold","black",tr]}],"font-family":[{font:[_]}],"fvn-normal":["normal-nums"],"fvn-ordinal":["ordinal"],"fvn-slashed-zero":["slashed-zero"],"fvn-figure":["lining-nums","oldstyle-nums"],"fvn-spacing":["proportional-nums","tabular-nums"],"fvn-fraction":["diagonal-fractions","stacked-fractons"],tracking:[{tracking:["tighter","tight","normal","wide","wider","widest",T]}],"line-clamp":[{"line-clamp":["none",O,tr]}],leading:[{leading:["none","tight","snug","normal","relaxed","loose",P]}],"list-image":[{"list-image":["none",w]}],"list-style-type":[{list:["none","disc","decimal",w]}],"list-style-position":[{list:["inside","outside"]}],"placeholder-color":[{placeholder:[r]}],"placeholder-opacity":[{"placeholder-opacity":[G]}],"text-alignment":[{text:["left","center","right","justify","start","end"]}],"text-color":[{text:[r]}],"text-opacity":[{"text-opacity":[G]}],"text-decoration":["underline","overline","line-through","no-underline"],"text-decoration-style":[{decoration:[].concat(E(),["wavy"])}],"text-decoration-thickness":[{decoration:["auto","from-font",P]}],"underline-offset":[{"underline-offset":["auto",P]}],"text-decoration-color":[{decoration:[r]}],"text-transform":["uppercase","lowercase","capitalize","normal-case"],"text-overflow":["truncate","text-ellipsis","text-clip"],indent:[{indent:[e]}],"vertical-align":[{align:["baseline","top","middle","bottom","text-top","text-bottom","sub","super",T]}],whitespace:[{whitespace:["normal","nowrap","pre","pre-line","pre-wrap","break-spaces"]}],break:[{break:["normal","words","all","keep"]}],hyphens:[{hyphens:["none","manual","auto"]}],content:[{content:["none",w]}],"bg-attachment":[{bg:["fixed","local","scroll"]}],"bg-clip":[{"bg-clip":["border","padding","content","text"]}],"bg-opacity":[{"bg-opacity":[G]}],"bg-origin":[{"bg-origin":["border","padding","content"]}],"bg-position":[{bg:[].concat(J(),[Xr])}],"bg-repeat":[{bg:["no-repeat",{repeat:["","x","y","round","space"]}]}],"bg-size":[{bg:["auto","cover","contain",Zr]}],"bg-image":[{bg:["none",{"gradient-to":["t","tr","r","br","b","bl","l","tl"]},Jr]}],"bg-color":[{bg:[r]}],"gradient-from-pos":[{from:[z]}],"gradient-via-pos":[{via:[z]}],"gradient-to-pos":[{to:[z]}],"gradient-from":[{from:[v]}],"gradient-via":[{via:[v]}],"gradient-to":[{to:[v]}],rounded:[{rounded:[l]}],"rounded-s":[{"rounded-s":[l]}],"rounded-e":[{"rounded-e":[l]}],"rounded-t":[{"rounded-t":[l]}],"rounded-r":[{"rounded-r":[l]}],"rounded-b":[{"rounded-b":[l]}],"rounded-l":[{"rounded-l":[l]}],"rounded-ss":[{"rounded-ss":[l]}],"rounded-se":[{"rounded-se":[l]}],"rounded-ee":[{"rounded-ee":[l]}],"rounded-es":[{"rounded-es":[l]}],"rounded-tl":[{"rounded-tl":[l]}],"rounded-tr":[{"rounded-tr":[l]}],"rounded-br":[{"rounded-br":[l]}],"rounded-bl":[{"rounded-bl":[l]}],"border-w":[{border:[n]}],"border-w-x":[{"border-x":[n]}],"border-w-y":[{"border-y":[n]}],"border-w-s":[{"border-s":[n]}],"border-w-e":[{"border-e":[n]}],"border-w-t":[{"border-t":[n]}],"border-w-r":[{"border-r":[n]}],"border-w-b":[{"border-b":[n]}],"border-w-l":[{"border-l":[n]}],"border-opacity":[{"border-opacity":[G]}],"border-style":[{border:[].concat(E(),["hidden"])}],"divide-x":[{"divide-x":[n]}],"divide-x-reverse":["divide-x-reverse"],"divide-y":[{"divide-y":[n]}],"divide-y-reverse":["divide-y-reverse"],"divide-opacity":[{"divide-opacity":[G]}],"divide-style":[{divide:E()}],"border-color":[{border:[a]}],"border-color-x":[{"border-x":[a]}],"border-color-y":[{"border-y":[a]}],"border-color-t":[{"border-t":[a]}],"border-color-r":[{"border-r":[a]}],"border-color-b":[{"border-b":[a]}],"border-color-l":[{"border-l":[a]}],"divide-color":[{divide:[a]}],"outline-style":[{outline:[""].concat(E())}],"outline-offset":[{"outline-offset":[P]}],"outline-w":[{outline:[P]}],"outline-color":[{outline:[r]}],"ring-w":[{ring:rr()}],"ring-w-inset":["ring-inset"],"ring-color":[{ring:[r]}],"ring-opacity":[{"ring-opacity":[G]}],"ring-offset-w":[{"ring-offset":[P]}],"ring-offset-color":[{"ring-offset":[r]}],shadow:[{shadow:["","inner","none",R,Yr]}],"shadow-color":[{shadow:[_]}],opacity:[{opacity:[G]}],"mix-blend":[{"mix-blend":Q()}],"bg-blend":[{"bg-blend":Q()}],filter:[{filter:["","none"]}],blur:[{blur:[t]}],brightness:[{brightness:[o]}],contrast:[{contrast:[f]}],"drop-shadow":[{"drop-shadow":["","none",R,w]}],grayscale:[{grayscale:[p]}],"hue-rotate":[{"hue-rotate":[b]}],invert:[{invert:[g]}],saturate:[{saturate:[K]}],sepia:[{sepia:[U]}],"backdrop-filter":[{"backdrop-filter":["","none"]}],"backdrop-blur":[{"backdrop-blur":[t]}],"backdrop-brightness":[{"backdrop-brightness":[o]}],"backdrop-contrast":[{"backdrop-contrast":[f]}],"backdrop-grayscale":[{"backdrop-grayscale":[p]}],"backdrop-hue-rotate":[{"backdrop-hue-rotate":[b]}],"backdrop-invert":[{"backdrop-invert":[g]}],"backdrop-opacity":[{"backdrop-opacity":[G]}],"backdrop-saturate":[{"backdrop-saturate":[K]}],"backdrop-sepia":[{"backdrop-sepia":[U]}],"border-collapse":[{border:["collapse","separate"]}],"border-spacing":[{"border-spacing":[i]}],"border-spacing-x":[{"border-spacing-x":[i]}],"border-spacing-y":[{"border-spacing-y":[i]}],"table-layout":[{table:["auto","fixed"]}],caption:[{caption:["top","bottom"]}],transition:[{transition:["none","all","","colors","opacity","shadow","transform",w]}],duration:[{duration:u()}],ease:[{ease:["linear","in","out","in-out",w]}],delay:[{delay:u()}],animate:[{animate:["none","spin","ping","pulse","bounce",w]}],transform:[{transform:["","gpu","none"]}],scale:[{scale:[C]}],"scale-x":[{"scale-x":[C]}],"scale-y":[{"scale-y":[C]}],rotate:[{rotate:[H,w]}],"translate-x":[{"translate-x":[Z]}],"translate-y":[{"translate-y":[Z]}],"skew-x":[{"skew-x":[F]}],"skew-y":[{"skew-y":[F]}],"transform-origin":[{origin:["center","top","top-right","right","bottom-right","bottom","bottom-left","left","top-left",w]}],accent:[{accent:["auto",r]}],appearance:["appearance-none"],cursor:[{cursor:["auto","default","pointer","wait","text","move","help","not-allowed","none","context-menu","progress","cell","crosshair","vertical-text","alias","copy","no-drop","grab","grabbing","all-scroll","col-resize","row-resize","n-resize","e-resize","s-resize","w-resize","ne-resize","nw-resize","se-resize","sw-resize","ew-resize","ns-resize","nesw-resize","nwse-resize","zoom-in","zoom-out",w]}],"caret-color":[{caret:[r]}],"pointer-events":[{"pointer-events":["none","auto"]}],resize:[{resize:["none","y","x",""]}],"scroll-behavior":[{scroll:["auto","smooth"]}],"scroll-m":[{"scroll-m":[e]}],"scroll-mx":[{"scroll-mx":[e]}],"scroll-my":[{"scroll-my":[e]}],"scroll-ms":[{"scroll-ms":[e]}],"scroll-me":[{"scroll-me":[e]}],"scroll-mt":[{"scroll-mt":[e]}],"scroll-mr":[{"scroll-mr":[e]}],"scroll-mb":[{"scroll-mb":[e]}],"scroll-ml":[{"scroll-ml":[e]}],"scroll-p":[{"scroll-p":[e]}],"scroll-px":[{"scroll-px":[e]}],"scroll-py":[{"scroll-py":[e]}],"scroll-ps":[{"scroll-ps":[e]}],"scroll-pe":[{"scroll-pe":[e]}],"scroll-pt":[{"scroll-pt":[e]}],"scroll-pr":[{"scroll-pr":[e]}],"scroll-pb":[{"scroll-pb":[e]}],"scroll-pl":[{"scroll-pl":[e]}],"snap-align":[{snap:["start","end","center","align-none"]}],"snap-stop":[{snap:["normal","always"]}],"snap-type":[{snap:["none","x","y","both"]}],"snap-strictness":[{snap:["mandatory","proximity"]}],touch:[{touch:["auto","none","pinch-zoom","manipulation",{pan:["x","left","right","y","up","down"]}]}],select:[{select:["none","text","all","auto"]}],"will-change":[{"will-change":["auto","scroll","contents","transform",w]}],fill:[{fill:[r,"none"]}],"stroke-w":[{stroke:[P,tr]}],stroke:[{stroke:[r,"none"]}],sr:["sr-only","not-sr-only"]},conflictingClassGroups:{overflow:["overflow-x","overflow-y"],overscroll:["overscroll-x","overscroll-y"],inset:["inset-x","inset-y","start","end","top","right","bottom","left"],"inset-x":["right","left"],"inset-y":["top","bottom"],flex:["basis","grow","shrink"],gap:["gap-x","gap-y"],p:["px","py","ps","pe","pt","pr","pb","pl"],px:["pr","pl"],py:["pt","pb"],m:["mx","my","ms","me","mt","mr","mb","ml"],mx:["mr","ml"],my:["mt","mb"],"font-size":["leading"],"fvn-normal":["fvn-ordinal","fvn-slashed-zero","fvn-figure","fvn-spacing","fvn-fraction"],"fvn-ordinal":["fvn-normal"],"fvn-slashed-zero":["fvn-normal"],"fvn-figure":["fvn-normal"],"fvn-spacing":["fvn-normal"],"fvn-fraction":["fvn-normal"],rounded:["rounded-s","rounded-e","rounded-t","rounded-r","rounded-b","rounded-l","rounded-ss","rounded-se","rounded-ee","rounded-es","rounded-tl","rounded-tr","rounded-br","rounded-bl"],"rounded-s":["rounded-ss","rounded-es"],"rounded-e":["rounded-se","rounded-ee"],"rounded-t":["rounded-tl","rounded-tr"],"rounded-r":["rounded-tr","rounded-br"],"rounded-b":["rounded-br","rounded-bl"],"rounded-l":["rounded-tl","rounded-bl"],"border-spacing":["border-spacing-x","border-spacing-y"],"border-w":["border-w-s","border-w-e","border-w-t","border-w-r","border-w-b","border-w-l"],"border-w-x":["border-w-r","border-w-l"],"border-w-y":["border-w-t","border-w-b"],"border-color":["border-color-t","border-color-r","border-color-b","border-color-l"],"border-color-x":["border-color-r","border-color-l"],"border-color-y":["border-color-t","border-color-b"],"scroll-m":["scroll-mx","scroll-my","scroll-ms","scroll-me","scroll-mt","scroll-mr","scroll-mb","scroll-ml"],"scroll-mx":["scroll-mr","scroll-ml"],"scroll-my":["scroll-mt","scroll-mb"],"scroll-p":["scroll-px","scroll-py","scroll-ps","scroll-pe","scroll-pt","scroll-pr","scroll-pb","scroll-pl"],"scroll-px":["scroll-pr","scroll-pl"],"scroll-py":["scroll-pt","scroll-pb"]},conflictingClassGroupModifiers:{"font-size":["leading"]}}}var Kr=ar(lr);function $r(r,e){for(var t in e)Ar(r,t,e[t]);return r}var re=Object.prototype.hasOwnProperty,ee=new Set(["string","number","boolean"]);function Ar(r,e,t){if(!re.call(r,e)||ee.has(typeof t)||t===null){r[e]=t;return}if(Array.isArray(t)&&Array.isArray(r[e])){r[e]=r[e].concat(t);return}if(typeof t=="object"&&typeof r[e]=="object"){if(r[e]===null){r[e]=t;return}for(var o in t)Ar(r[e],o,t[o])}}function te(r){for(var e=arguments.length,t=new Array(e>1?e-1:0),o=1;o<e;o++)t[o-1]=arguments[o];return typeof r=="function"?ar.apply(void 0,[lr,r].concat(t)):ar.apply(void 0,[function(){return $r(lr(),r)}].concat(t))}var Cr={twMerge:!0,twMergeConfig:{},responsiveVariants:!1},Mr=r=>r||void 0,or=(...r)=>Mr(r.flat(1/0).filter(Boolean).join(" ")),D=(...r)=>(e=Cr)=>{if(!e.twMerge)return or(r);let t=I(e.twMergeConfig)?Kr:te(e.twMergeConfig);return Mr(t(or(r)))},vr=(r,e)=>{let t={...r};for(let o in e)t.hasOwnProperty(o)?t[o]=or(t[o],e[o]):t[o]=e[o];return t},oe=(r,e=Cr)=>{var t,o,a,l,i;let{slots:n={},variants:f={},compoundVariants:p=[],compoundSlots:b=[],defaultVariants:g={}}=r,A=or((t=r==null?void 0:r.extend)==null?void 0:t.base,r==null?void 0:r.base),v=gr(f,(o=r==null?void 0:r.extend)==null?void 0:o.variants),z=Object.assign({},(a=r==null?void 0:r.extend)==null?void 0:a.defaultVariants,g),k=I(n)?{}:{base:r==null?void 0:r.base,...n},S=Array.isArray(e.responsiveVariants)&&e.responsiveVariants.length>0||e.responsiveVariants===!0,G=I((l=r==null?void 0:r.extend)==null?void 0:l.slots)?k:vr((i=r==null?void 0:r.extend)==null?void 0:i.slots,I(k)?{base:r==null?void 0:r.base}:k),M=C=>{var U,F,q,Z;if(I(v)&&I(n)&&I((U=r==null?void 0:r.extend)==null?void 0:U.slots))return D(A,C==null?void 0:C.class,C==null?void 0:C.className)(e);if(p&&!Array.isArray(p))throw new TypeError(`The "compoundVariants" prop must be an array. Received: ${typeof p}`);if(b&&!Array.isArray(b))throw new TypeError(`The "compoundSlots" prop must be an array. Received: ${typeof b}`);let X=(s,c,u=[],d)=>{let h=u;if(typeof c=="string")h.push(ur(c).split(" ").map(y=>`${s}:${y}`));else if(Array.isArray(c))h.push(c.flatMap(y=>`${s}:${y}`));else if(typeof c=="object"&&typeof d=="string"){let y=c==null?void 0:c[d];if(y&&typeof y=="string"){let x=ur(y);h[d]=h[d]?[...h[d],...x.split(" ").map(j=>`${s}:${j}`)]:x.split(" ").map(j=>`${s}:${j}`)}else Array.isArray(y)&&y.length>0&&(h[d]=y.flatMap(x=>`${s}:${x}`))}return h},L=(s,c=v,u=null)=>{let d=c==null?void 0:c[s];if(typeof d!="object"||I(d))return null;let h=C==null?void 0:C[s],y=z==null?void 0:z[s],x=[];if(h===null)return null;let j=cr(h);typeof j=="object"&&S&&(x=Object.keys(j).reduce((nr,er)=>{let dr=j[er],kr=d==null?void 0:d[dr];return er==="initial"?(y=dr,nr):Array.isArray(e.responsiveVariants)&&!e.responsiveVariants.includes(er)?nr:X(er,kr,nr,u)},[]));let B=d[j]||d[cr(y)];return typeof x=="object"&&typeof u=="string"&&x[u]?vr(x,B):x.length>0?[B,...x]:B},$=()=>v?Object.keys(v).map(s=>L(s,v)):null,rr=s=>!v||typeof v!="object"?null:Object.keys(v).map(c=>{let u=L(c,v,s);return s==="base"&&typeof u=="string"?u:u&&u[s]}).filter(Boolean),V=C&&Object.fromEntries(Object.entries(C).filter(([,s])=>s!==void 0)),J=s=>{var c;let u=typeof(C==null?void 0:C[s])=="object"?{[s]:(c=C[s])==null?void 0:c.initial}:{};return{...z,...V,...u}},E=(s=[])=>s==null?void 0:s.filter(({class:c,className:u,...d})=>Object.entries(d).every(([h,y])=>{let x=J(h);return Array.isArray(y)?y.includes(x[h]):x[h]===y})).flatMap(({class:c,className:u})=>[c,u]),Q=()=>{var s;let c=E(p),u=E((s=r==null?void 0:r.extend)==null?void 0:s.compoundVariants);return zr(u,c)},Y=()=>{let s=Q();return Array.isArray(s)?s.reduce((c,u)=>(typeof u=="string"&&(c.base=D(c.base,u)(e)),typeof u=="object"&&Object.entries(u).forEach(([d,h])=>{c[d]=D(c[d],h)(e)}),c),{}):s},N=()=>b.length<1?null:b.reduce((s,c)=>{let{slots:u=[],class:d,className:h,...y}=c;if(!I(y)){let x=Object.keys(y);for(let j of x){let B=J(j)[j];if(!B||B!==y[j])return s}}return u.forEach(x=>{s[x]||(s[x]=[]),s[x].push([d,h])}),s},{});if(!I(n)||!I((F=r==null?void 0:r.extend)==null?void 0:F.slots)){let s=(q=Y())!=null?q:[],c=(Z=N())!=null?Z:[];return{...typeof G=="object"&&!I(G)?Object.keys(G).reduce((u,d)=>(u[d]=h=>D(G[d],rr(d),s==null?void 0:s[d],c==null?void 0:c[d],h==null?void 0:h.class,h==null?void 0:h.className)(e),u),{}):{}}}return D(A,$(),Q(),C==null?void 0:C.class,C==null?void 0:C.className)(e)},K=()=>{if(!(!v||typeof v!="object"))return Object.keys(v)};return M.variantKeys=K(),M.base=A,M.slots=G,M.variants=v,M.defaultVariants=z,M.compoundSlots=b,M.compoundVariants=p,M};export{oe as f};
