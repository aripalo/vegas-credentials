import React from 'react';
import clsx from 'clsx';
import Layout from '@theme/Layout';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import styles from './index.module.css';
import HomepageFeatures from '@site/src/components/HomepageFeatures';
import HowItWorks from '@site/src/components/HowItWorks';

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import CodeBlock from '@theme/CodeBlock';

function HomepageHeader() {
  const {siteConfig} = useDocusaurusContext();
  return (
    <header className={clsx('hero hero--primary', styles.heroBanner)}>
      <div className="container">
        <h1 className={styles.heroTitle}>
          <span className={styles.textVegas}>Vegas</span> <span className={styles.textCredentials}>Credentials</span>
        </h1>
        <aside>
          <blockquote className={styles.heroQuote}>
          <i>Much like spending a week in Las Vegas at AWS re:Invent</i>, using multiple AWS tools (SDKs, CLI, CDK, Terraform, etc) via command-line to assume IAM roles in different accounts with Multi-Factor Authentication can be an exhausting experience: <code>vegas-credentials</code> aims to simplify the credential process! <i>And just like you shouldn't stay too long in Las Vegas at once</i>, this tool only deals with temporary sesssion credentials.
          </blockquote>
        </aside>

        <div className={styles.heroDescription}>
        AWS <a href="https://docs.aws.amazon.com/sdkref/latest/guide/feature-process-credentials.html" title="Sourcing AWS Credentials via External Process" target="_blank" rel="noopener noreferer"><code>credential_process</code></a> utility to request STS Temporary Security Credentials by assuming an IAM role with TOTP MFA via either Yubikey Touch or Authenticator Apps.
        </div>
        <div className={styles.install}>
          <Tabs>
            <TabItem value="brew" label={<><span className={styles.installTabPrefix}>macos/linux/wsl&nbsp;</span><span className={styles.installTabTitle}>brew</span></>} attributes={{className: styles.installTab}}>
              <CodeBlock language="bash">brew install aripalo/tap/vegas-credentials</CodeBlock>
            </TabItem>
            <TabItem value="scoop" label={<><span className={styles.installTabPrefix}>windows&nbsp;</span><span className={styles.installTabTitle}>scoop</span></>} attributes={{className: styles.installTab}}>
              <CodeBlock language="bash">
                scoop bucket add aripalo https://github.com/aripalo/scoops.git &&
                scoop install vegas-credentials
              </CodeBlock>
            </TabItem>
          </Tabs>
        </div>
        <footer className={styles.afterInstall}>
          <ol>
            <li>After installation, see: <a href="/docs/setup">Setup</a></li>
            <li>After configuration, use any AWS tool normally with <code>--profile=your-profile</code> or see <a href="/docs/examples">Examples</a> for more</li>
          </ol>
        </footer>

      </div>
    </header>
  );
}

export default function Home() {
  // const {siteConfig} = useDocusaurusContext();
  return (
    <Layout>
      <HomepageHeader />
      <main>
        <HowItWorks />
        <HomepageFeatures />
      </main>
    </Layout>
  );
}
