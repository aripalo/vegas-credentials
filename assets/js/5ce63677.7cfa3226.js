"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[9733],{3905:function(e,t,n){n.d(t,{Zo:function(){return c},kt:function(){return d}});var r=n(7294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function u(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var s=r.createContext({}),p=function(e){var t=r.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},c=function(e){var t=p(e.components);return r.createElement(s.Provider,{value:t},e.children)},l={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},y=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,i=e.originalType,s=e.parentName,c=u(e,["components","mdxType","originalType","parentName"]),y=p(n),d=o,m=y["".concat(s,".").concat(d)]||y[d]||l[d]||i;return n?r.createElement(m,a(a({ref:t},c),{},{components:n})):r.createElement(m,a({ref:t},c))}));function d(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var i=n.length,a=new Array(i);a[0]=y;var u={};for(var s in t)hasOwnProperty.call(t,s)&&(u[s]=t[s]);u.originalType=e,u.mdxType="string"==typeof e?e:o,a[1]=u;for(var p=2;p<i;p++)a[p]=n[p];return r.createElement.apply(null,a)}return r.createElement.apply(null,n)}y.displayName="MDXCreateElement"},4994:function(e,t,n){n.r(t),n.d(t,{assets:function(){return c},contentTitle:function(){return s},default:function(){return d},frontMatter:function(){return u},metadata:function(){return p},toc:function(){return l}});var r=n(7462),o=n(3366),i=(n(7294),n(3905)),a=["components"],u={sidebar_position:1},s="Yubikey Setup",p={unversionedId:"yubikeys/setup",id:"yubikeys/setup",title:"Yubikey Setup",description:"This section covers how you can configure and use Yubikeys as one of the MFA authentication methods with vegas-credentials. The nice thing is the \u201cone of the methods\u201d part: You can (and should) set your MFA to multiple devices, for example into one Yubikey and into your mobile Authenticator App; This way if you don't have your Yubikey device with you at all times, you'd still be able to authenticate with your mobile device!",source:"@site/docs/yubikeys/setup.md",sourceDirName:"yubikeys",slug:"/yubikeys/setup",permalink:"/docs/yubikeys/setup",editUrl:"https://github.com/aripalo/vegas-credentials/tree/main/packages/create-docusaurus/templates/shared/docs/yubikeys/setup.md",tags:[],version:"current",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"Setup",permalink:"/docs/setup"},next:{title:"AWS MFA Setup",permalink:"/docs/yubikeys/aws-mfa"}},c={},l=[{value:"Prerequsites",id:"prerequsites",level:2}],y={toc:l};function d(e){var t=e.components,n=(0,o.Z)(e,a);return(0,i.kt)("wrapper",(0,r.Z)({},y,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"yubikey-setup"},"Yubikey Setup"),(0,i.kt)("p",null,"This section covers how you can configure and use Yubikeys as one of the MFA authentication methods with ",(0,i.kt)("inlineCode",{parentName:"p"},"vegas-credentials"),". The nice thing is the ",(0,i.kt)("em",{parentName:"p"},"\u201cone of the methods\u201d")," part: You can (and should) set your MFA to multiple devices, for example into one Yubikey and into your mobile Authenticator App; This way if you don't have your Yubikey device with you at all times, you'd still be able to authenticate with your mobile device!"),(0,i.kt)("h2",{id:"prerequsites"},"Prerequsites"),(0,i.kt)("ol",null,(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("p",{parentName:"li"},"You MUST have at least one ",(0,i.kt)("a",{parentName:"p",href:"https://www.yubico.com/products/yubikey-5-overview/"},"Yubikey Touch device")," with ",(0,i.kt)("a",{parentName:"p",href:"https://en.wikipedia.org/wiki/Time-based_One-Time_Password"},"OATH TOTP")," support (Yubikey 5 or 5C recommended).")),(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("p",{parentName:"li"},"You MUST have Yubikey Manager CLI ",(0,i.kt)("a",{parentName:"p",href:"https://developers.yubico.com/yubikey-manager/"},(0,i.kt)("inlineCode",{parentName:"a"},"ykman")," CLI")," installed in your machine and available in your ",(0,i.kt)("inlineCode",{parentName:"p"},"$PATH"),".")),(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("p",{parentName:"li"},"You MAY also install the ",(0,i.kt)("a",{parentName:"p",href:"https://www.yubico.com/products/yubico-authenticator/"},"Yubico Authenticator")," GUI which helps you when you're setting up new OATH accounts into your Yubikey.")),(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("p",{parentName:"li"},"You SHOULD consider password protecting OATH application on your Yubikey device. You can do that with ",(0,i.kt)("inlineCode",{parentName:"p"},"ykman"),":"),(0,i.kt)("pre",{parentName:"li"},(0,i.kt)("code",{parentName:"pre",className:"language-sh"},"ykman oath access change\n")),(0,i.kt)("ol",{parentName:"li"},(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("p",{parentName:"li"},"For extra protection, you SHOULD NOT use ",(0,i.kt)("inlineCode",{parentName:"p"},"ykman oath access remember")," or click \u201cRemember password\u201d with the Yubico Authenticator GUI.")),(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("p",{parentName:"li"},"For password protected devices, ",(0,i.kt)("inlineCode",{parentName:"p"},"vegas-credentials")," will prompt you for the device OATH application password and cache it for 12 hours: Which usually is enought to get you through the work day without having to retype the password."))))))}d.isMDXComponent=!0}}]);