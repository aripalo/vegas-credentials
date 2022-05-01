// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion


const lightCodeTheme = require("prism-react-renderer/themes/github");
const darkCodeTheme = require("prism-react-renderer/themes/dracula");

const socialImageUrl = '/img/vegas-credentials-hero.jpg';
const socialTitle = 'vegas-credentials';
const socialDescription = 'AWS credential_process utility to request STS Temporary Security Credentials by assuming an IAM role with TOTP MFA via either Yubikey Touch or Authenticator Apps.';

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'vegas-credentials',
  tagline: 'AWS credential_process utility',
  url: 'https://credentials.vegas',
  baseUrl: '/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.ico',
  organizationName: 'aripalo', // Usually your GitHub org/user name.
  projectName: 'vegas-credentials', // Usually your repo name.
  trailingSlash: false,

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          editUrl: 'https://github.com/aripalo/vegas-credentials/tree/main/packages/create-docusaurus/templates/shared/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      metadata: [
        {name: 'description', content: socialDescription},
        {property: 'og:description', content: socialDescription},
        {name: 'twitter:description', content: socialDescription},
        {itemprop: 'description', content: socialDescription},

        {name: 'keywords', content: 'amazon web services, aws, iam, sts, mfa, multi-factor authentication, yubikey'},

        {property: 'og:title', content: socialTitle},
        {name: 'twitter:title', content: socialTitle},
        {itemprop: 'title', content: socialTitle},

        {property: 'og:image', content: socialImageUrl},
        {name: 'twitter:image', content: socialImageUrl},
        {name: 'twitter:card', content: 'summary_large_image'},
        {itemprop: 'image', content: socialImageUrl},
      ],
      colorMode: {
        defaultMode: 'dark',
        disableSwitch: false,
        respectPrefersColorScheme: true,
      },
      navbar: {
        title: 'Vegas Credentials',
        /*
        logo: {
          alt: 'My Site Logo',
          src: 'img/logo.svg',
        },
        */
        items: [
          {
            type: 'doc',
            docId: 'setup',
            position: 'left',
            label: 'Docs',
          },
          {
            href: '/design-principles',
            label: 'Design Principles',
          },
          {
            href: '/alternatives',
            label: 'Alternatives',
          },
          {
            href: 'https://github.com/aripalo/vegas-credentials',
            label: 'GitHub',
            position: 'right',
          },
        ],
      },
      footer: {
        style: 'dark',
        links: [
          {
            title: 'Docs',
            items: [
              {
                label: 'Setup',
                to: '/docs/setup',
              },
            ],
          },
          {
            title: 'Community',
            items: [
              {
                label: 'Stack Overflow',
                href: 'https://stackoverflow.com/questions/tagged/vegas-credentials',
              },
              {
                label: 'Issues',
                href: 'https://github.com/aripalo/vegas-credentials/issues',
              },
              {
                label: 'Discussions',
                href: 'https://github.com/aripalo/vegas-credentials/discussions',
              },
            ],
          },
          {
            title: 'More',
            items: [
              {
                label: 'GitHub',
                href: 'https://github.com/aripalo/vegas-credentials',
              },
              {
                label: 'Author',
                href: 'https://aripalo.com',
              },
            ],
          },
        ],
        copyright: `Copyright © ${new Date().getFullYear()} Ari Palo. Built with Docusaurus.`,
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
        additionalLanguages: ['ini', 'hcl', 'yaml', 'log'],
      },
    }),
};

module.exports = config;
