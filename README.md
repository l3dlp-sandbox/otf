![OTF logo](/readme_inkscape_logo.png)

OTF is an open source alternative to Terraform Enterprise. Includes SSO, team management, agents, and no per-resource pricing.

Docs: https://docs.otf.ninja/

[![Slack](https://img.shields.io/badge/join-slack%20community-brightgreen)](https://join.slack.com/t/otf-pg29376/shared_invite/zt-1jga4k1cl-bzmJg71f4uUB9fJhxdT~gQ)

## Quickstart Demo

To quickly try out OTF you can sign into the demo server using your github account:

https://demo.otf.ninja

Once you're signed in, create an organization and make a note of its name - you'll be using it below.

Now we'll login to the account in your terminal. You'll need terraform installed.

NOTE: only terraform version 1.2.0 and later is supported.

Setup local credentials:

```bash
terraform login demo.otf.ninja
```

Confirm with `yes` to proceed and you'll be asked to give consent for `terraform` to access your account on OTF. After you give consent, you should be notified that you can close the browser and return to the terminal.

In the terminal `terraform login` should have printed out confirmation of success:

```
Success! Terraform has obtained and saved an API token.
```

Write some terraform configuration to a file, setting the organization to your username:

```terraform
terraform {
  backend "remote" {
    hostname     = "demo.otf.ninja"
    organization = "<your username>"

    workspaces {
      name = "dev"
    }
  }
}

resource "null_resource" "demo" {}
```

Initialize terraform:

```bash
terraform init
```

> NOTE: if you're using Mac or Windows, you may receive an error similar to the following error:
>
> > Error: Failed to install provider
> >
> > Error while installing hashicorp/null v3.2.1: the current package for registry.terraform.io/hashicorp/null 3.2.1
> > doesn't match any of the checksums previously recorded in the dependency lock file; for more information:
> > https://www.terraform.io/language/provider-checksum-verification
>
> If so, you need to update your lockfile (`.terraform.lock.hcl`) to include hashes for the platform that the OTF demo is hosted on (linux):
>
> ```
> terraform providers lock -platform=linux_amd64
> ```
>
> Then re-run `terraform init`

Now create a plan:

```bash
terraform plan
```

After you've invoked `terraform plan`, you'll see the plan output along with a link. Click on the link to the run to view the status and logs:

<p align="center">
    <kbd>
        <img src="./docs/images/run_page_planned_and_finished_state.png" width="600px" />
    </kbd>
</p>

You can optionally run `terraform apply` to apply the changes:

```bash
terraform apply
```

You've reached the end of this quickstart demo. See the [docs](https://docs.otf.ninja) for instructions on deploying OTF.

## Legal

OTF is in no way affiliated with Hashicorp. Terraform and Terraform Enterprise are trademarks of Hashicorp.
