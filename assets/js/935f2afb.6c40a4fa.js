"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[53],{1109:function(e){e.exports=JSON.parse('{"pluginId":"default","version":"current","label":"Next","banner":null,"badge":false,"className":"docs-version-current","isLast":true,"docsSidebars":{"tutorialSidebar":[{"type":"link","label":"Install","href":"/docs/install","docId":"install"},{"type":"link","label":"Setup","href":"/docs/setup","docId":"setup"},{"type":"category","label":"Yubikeys","collapsible":true,"collapsed":true,"items":[{"type":"link","label":"Yubikey Setup","href":"/docs/yubikeys/setup","docId":"yubikeys/setup"},{"type":"link","label":"AWS MFA Setup","href":"/docs/yubikeys/aws-mfa","docId":"yubikeys/aws-mfa"},{"type":"link","label":"Yubikey(s) with Vegas","href":"/docs/yubikeys/vegas","docId":"yubikeys/vegas"}]},{"type":"link","label":"CLI","href":"/docs/cli","docId":"cli"},{"type":"link","label":"Logs","href":"/docs/logs","docId":"logs"},{"type":"link","label":"Privacy Policy","href":"/docs/privacy-policy","docId":"privacy-policy"},{"type":"link","label":"Role Session Name","href":"/docs/role-session-name","docId":"role-session-name"},{"type":"link","label":"Examples","href":"/docs/examples","docId":"examples"},{"type":"link","label":"Network Connections","href":"/docs/network-connections","docId":"network-connections"}]},"docs":{"cli":{"id":"cli","title":"CLI","description":"As stated previously, in most cases you don\'t use the vegas-credentials command directly but if you need to troubleshoot something or you wish to configure the user output (such as colors or emoji), then this section is for you.","sidebar":"tutorialSidebar"},"examples":{"id":"examples","title":"Examples","description":"Multiple Source Profiles","sidebar":"tutorialSidebar"},"install":{"id":"install","title":"Install","description":"Additional installation methods may be added later but the below options should cover pretty much every environment.","sidebar":"tutorialSidebar"},"logs":{"id":"logs","title":"Logs","description":"In troubleshooting situtations, it can be useful to investigate application.log: Running any vegas-credentials command logs into that (rotated) log file.","sidebar":"tutorialSidebar"},"network-connections":{"id":"network-connections","title":"Network Connections","description":"If you\'re using a firewall software (such as Little Snitch on macOS) you need to allow following outbound connections to be made by vegas-credentials process:","sidebar":"tutorialSidebar"},"privacy-policy":{"id":"privacy-policy","title":"Privacy Policy","description":"vegas-credentials does not use or send any tracking, metrics or analytics data.","sidebar":"tutorialSidebar"},"role-session-name":{"id":"role-session-name","title":"Role Session Name","description":"You can configure the Role Session Name via shared configuration file:","sidebar":"tutorialSidebar"},"setup":{"id":"setup","title":"Setup","description":"In this document we\'re only setting up the necessary configuration changes into ~/.aws/config to get started using vegas-credentials as a credential source.","sidebar":"tutorialSidebar"},"yubikeys/aws-mfa":{"id":"yubikeys/aws-mfa","title":"AWS MFA Setup","description":"Using Yubikey(s) with vegas-credentials in AWS requires you to set up a new Virtual MFA Device in AWS.","sidebar":"tutorialSidebar"},"yubikeys/setup":{"id":"yubikeys/setup","title":"Yubikey Setup","description":"This section covers how you can configure and use Yubikeys as one of the MFA authentication methods with vegas-credentials. The nice thing is the \u201cone of the methods\u201d part: You can (and should) set your MFA to multiple devices, for example into one Yubikey and into your mobile Authenticator App; This way if you don\'t have your Yubikey device with you at all times, you\'d still be able to authenticate with your mobile device!","sidebar":"tutorialSidebar"},"yubikeys/vegas":{"id":"yubikeys/vegas","title":"Yubikey(s) with Vegas","description":"1. In your ~/.aws/config you should have the mfaserial configured and matching the Virtual MFA Device_ serial we configured in the previous step:","sidebar":"tutorialSidebar"}}}')}}]);