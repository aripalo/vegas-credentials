---
sidebar_position: 10
---


# Role Session Name

You can configure the Role Session Name via shared configuration file:
```ini
# ~/.aws/config
[profile frank@concerts]
credential_process = vegas-credentials assume --profile=frank@concerts
vegas_role_arn = arn:aws:iam::222222222222:role/SingerRole
vegas_source_profile = work
role_session_name = SinatraAtTheSands
```

## Fallback value

[AssumeRole](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html) requires Role Session Name value to be set. If you do not provide one, Vegas Credentials will use a fallback value which is one of the following (in the following priority):
1. User Full Name: e.g. `Frank Sinatra`
2. User (System) Name: e.g. `frank`
3. System Hostname: e.g. `franks-laptop`
4. Combination of OS & Architecture: e.g. `darwin_amd64`

It also removes disallowed characters and truncates the string to match the [requirements](https://aws.amazon.com/blogs/security/easily-control-naming-individual-iam-role-sessions/).
