import React from 'react';
import clsx from 'clsx';
import styles from './styles.module.css';

const FeatureList = [
  {
    title: 'Pluggable',
    Svg: require('@site/static/img/features/plugin.svg').default,
    description: (
      <>Into AWS <code>credential_process</code></>
    ),
    details: (<><a href="https://docs.aws.amazon.com/sdkref/latest/guide/feature-process-credentials.html" title="Compatibility with AWS SDKS" target="_blank" rel="noopener noreferer">what's that?</a></>),
  },
  {
    title: 'Automatic Refresh',
    Svg: require('@site/static/img/features/refresh.svg').default,
    description: (
      <>Credential Refresh on Session Expiry</>
    ),
    details: (<>for example <a href="https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html" title="CLI Credentials" target="_blank" rel="noopener noreferer">CLI <code>v2</code></a> and <a href="https://pkg.go.dev/github.com/aws/aws-sdk-go/aws/credentials" title="Go SDK Credentials" target="_blank" rel="noopener noreferer">Go SDK</a></>),
  },
  {
    title: 'Role Chaining',
    Svg: require('@site/static/img/features/role-chaining.svg').default,
    description: (
      <>Start with Vegas, go anywhere</>
    ),
    details: (<><code>IAM User → Role A → Role B</code></>),
  },
  {
    title: 'Yubikey Touch',
    Svg: require('@site/static/img/features/yubikey-mfa.svg').default,
    description: (
      <>Multiple Devices & Password Protection</>
    ),
    details: (<> <a href="https://www.yubico.com/products/yubikey-5-overview/" title="Yubikey Series 5" target="_blank" rel="noopener noreferer">Series 5</a> with OATH TOTP support</>),
  },
  {
    title: 'Authenticator Apps',
    Svg: require('@site/static/img/features/mobile-mfa.svg').default,
    description: (
      <>For <i>copy-pasting</i> TOTP codes</>
    ),
    details: (<><a href="https://authy.com/" title="Twilio Authy" target="_blank" rel="noopener noreferer">Authy</a>, <a href="https://support.google.com/accounts/answer/1066447" title="Google Authenticator" target="_blank" rel="noopener noreferer">Google Authenticator</a>, etc.</>),
  },

  {
    title: 'Multiple MFA inputs',
    Svg: require('@site/static/img/features/mfa-input.svg').default,
    description: (
      <>Yubikey Touch, GUI Prompt & Standard Input</>
    ),
    details: (<>first input wins</>),
  },





  {
    title: 'Encrypted Cache',
    Svg: require('@site/static/img/features/encrypted-cache.svg').default,
    description: (
      <>Protection against Credential Scrapers</>
    ),
    details: (<>... and only <i>temporary</i> credentials cached</>),
  },

  {
    title: 'Cache Invalidation',
    Svg: require('@site/static/img/features/trash.svg').default,
    description: (
      <>On Configuration Change or Credential Expiry</>
    ),
    details: (<>e.g. change of <code>role_session_name</code></>),
  },

  {
    title: 'Fast',
    Svg: require('@site/static/img/features/blazing-fast.svg').default,
    description: (
      <><code>&lt;100ms</code> for Cached Credentials</>
    ),
    details: (<>... does that count as <code>blazing</code>?</>),
  },



  {
    title: 'Parallelism',
    Svg: require('@site/static/img/features/parallelism.svg').default,
    description: (
      <>Parallel calls handled via mutex locking</>
    ),
    details: (<>e.g. Terraform <code>--parallelism=n</code></>),
  },

  {
    title: 'Cross-Platform',
    Svg: require('@site/static/img/features/systems.svg').default,
    description: (
      <>Built with <a href="https://go.dev/" title="Go Programming Language" target="_blank" rel="noopener noreferer">Go</a></>
    ),
    details: (<><code>macos|linux|win</code> @ <code>x86_64|arm64</code></>),
  },


  {
    title: 'Minimal Configuration',
    Svg: require('@site/static/img/features/ini-file.svg').default,
    description: (
      <>Within AWS config file</>
    ),
    details: (<><code>~/.aws/config</code></>),
  },

];

const ToolsList = [
  {
    title: 'AWS CLI',
    Svg: require('@site/static/img/features/aws-cli.svg').default,
    description: (
      <>Since version <code>v2</code>
      <br/>
      <a href="https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html" title="AWS CLI Credential Sourcing" target="_blank" rel="noopener noreferer">docs</a>


      </>
    ),
  },
  {
    title: 'AWS SDKs',
    Svg: require('@site/static/img/features/aws-sdks.svg').default,
    description: (
      <>C++/Go/Java/JS/.NET/PHP/Python/Ruby
      <br/>
      <a href="https://docs.aws.amazon.com/sdkref/latest/guide/feature-process-credentials.html#feature-process-credentials-sdk-compat" title="Compatibility with AWS SDKS" target="_blank" rel="noopener noreferer">compatibility</a>
      </>
    ),
  },
  {
    title: 'AWS CDK',
    Svg: require('@site/static/img/features/aws-cdk.svg').default,
    description: (
      <>Since version <code>v1.73.0</code>
       <br/>
      <a href="https://github.com/aws/aws-cdk/releases/tag/v1.73.0" title="AWS CDK Credential Process Support" target="_blank" rel="noopener noreferer">release info</a>


      </>
    ),
  },

  {
    title: 'Terraform',
    Svg: require('@site/static/img/features/terraform.svg').default,
    description: (
      <>Since <code>terraform-provider-aws</code> version <code>v3.0.0</code>
      <br/>
      <a href="https://github.com/hashicorp/terraform-provider-aws/pull/14077" title="Terraform Credential Process support" target="_blank" rel="noopener noreferer">release info</a>


      </>
    ),
  },

  {
    title: 'Pulumi',
    Svg: require('@site/static/img/features/pulumi.svg').default,
    description: (
      <>Respects and uses your configuration settings
      <br/>
      <a href="https://www.pulumi.com/docs/get-started/aws/begin/#configure-pulumi-to-access-your-aws-account" title="Pulumi will respect and use your configuration settings" target="_blank" rel="noopener noreferer">details</a>


      </>
    ),
  },

  {
    title: 'Ansible',
    Svg: require('@site/static/img/features/ansible.svg').default,
    description: (
      <>Respects and uses your configuration settings
      <br/>
      <a href="https://docs.ansible.com/ansible/latest/collections/community/aws/aws_config_recorder_module.html#parameter-profile" title="Ansible AWS module with Profile configuration" target="_blank" rel="noopener noreferer">details</a>


      </>
    ),
  },
]

function Feature({Svg, title, description, details}) {
  return (
    <div className={clsx('col col--4', styles.feature)}>
      <div className="text--center">
        <Svg className={styles.featureSvg} role="img" />
      </div>
      <div className="text--center padding-horiz--md">
        <h3 className={styles.featureName}>{title}</h3>
        <p className={styles.featureDescription}>{description}</p>
        <p className={styles.featureDetails}>{details}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <>
    <section className={styles.features}>
      <div className="container">
        <h2>Features</h2>
        <p>List of most noteworthy features shipping with <code>vegas-credentials</code></p>
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
            ))}
        </div>
      </div>
    </section>
    <section className={clsx(styles.features, styles.tools)}>
      <div className="container">
        <h2>Supported Tools</h2>
        <p>Partial list of tools that work with <code>vegas-credentials</code></p>
        <div className="row">
          {ToolsList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
    </>
  );
}
