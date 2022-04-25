import React from 'react';
import clsx from 'clsx';
import styles from './styles.module.css';


import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import CodeBlock from '@theme/CodeBlock';

export default function HowItWorks() {
  return (
    <section className={styles.howItWorks}>
      <div className="container">
        <h2>How it Works</h2>

        <figure className={styles.diagram}>
          <figcaption>
            <strong>You just need to set couple lines of configurations and then use any AWS tools as you would normally with named profiles.</strong> It's almost like <i>magic</i>. But if you don't believe in magic, here's all the tricks <code>vegas-credentials</code> does in the background:
          </figcaption>
          <img className={styles.diagramImg} src="/img/how-vegas-credentials-work.svg" alt="How it works" />
        </figure>

        <p class={styles.tools}>As far as various AWS tools are concerned, nothing special needs to be done! No bash script wrappers, no need to export any environment variables and no need for custom code to handle the credentials:</p>

        <Tabs>

          <TabItem value="awscli" label="CLI">
            <CodeBlock language="bash">
            $ aws aws s3api create-bucket --bucket=my-bucket --profile=my-profile
            </CodeBlock>
          </TabItem>

          <TabItem value="awscdk" label="CDK">
            <CodeBlock language="bash">
            $ npx cdk deploy --profile=my-profile
            </CodeBlock>
          </TabItem>

          <TabItem value="awssdkboto3" label="SDK">
            <CodeBlock language="bash">
            $ AWS_PROFILE=my-profile python3 create-bucket.py
            </CodeBlock>
          </TabItem>

          <TabItem value="terraform" label="Terraform">
            <CodeBlock language="bash">
            $ terraform apply -var="aws_profile=my-profile
            </CodeBlock>
            <CodeBlock language="hcl">{`// in your Terraform AWS Provider configuration
variable "aws_profile" {
  type = string
}
provider "aws" {
  profile = var.aws_profile
}`}
            </CodeBlock>
          </TabItem>

          <TabItem value="ansible" label="Ansible">
            <CodeBlock language="bash">
            $ ansible-playbook -i hosts tasks.yml --extra-vars "aws_profile=my-profile"
            </CodeBlock>
            <CodeBlock language="yml">{`---
- name: Create Bucket
  hosts: local
  connection: local
  vars:
    aws_profile:
  tasks:
    - name: Create new bucket
      aws_s3:
        bucket: my-bucket
        mode: create
        region: eu-west-1
        profile: "{{ aws_profile }}"`}
            </CodeBlock>
          </TabItem>

          <TabItem value="pulumi" label="Pulumi">
            <CodeBlock language="bash">
            $ pulumi config set aws:profile=my-profile && pulumi up
            </CodeBlock>
          </TabItem>
        </Tabs>
      </div>
    </section>
  )
}
