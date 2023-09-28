# Static Assets Modules

[![Maintained by SimplyCubed.com](https://img.shields.io/badge/maintained%20by-simplycubed.com-%235849a6.svg)](https://simplycubed.com/?ref=repo_google_static_assets)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/simplycubed/terraform-google-static-assets.svg?label=latest)](https://github.com/simplycubed/terraform-google-static-assets/releases/latest)
![Terraform Version](https://img.shields.io/badge/tf-%3E%3D1.0.x-blue.svg)

This repo contains modules for managing static assets (CSS, JS, images) in GCP.

## Updates

This project was forked from the great work that Gruntwork has done and has been extended to include new functionality.

- Custom Header Support Added for Buckets

```text
# Accepts a list of headers
variable "custom_headers" {
  description = "A list of custom Headers that the HTTP/S load balancer should add to proxied responses"
  type        = list(string)
  default     = ["Referrer-Policy: origin"]
}
```

## Quickstart

If you want to quickly launch a static website using [Google Cloud Storage](https://cloud.google.com/storage/),
you can run the example that is in the root of this repo. Check out the [cloud-storage-static-website example documentation](https://github.com/simplycubed/terraform-google-static-assets/blob/master/examples/cloud-storage-static-website) for instructions.

## What's in this repo

This repo has the following folder structure:

- [root](https://github.com/simplycubed/terraform-google-static-assets/tree/master): The root folder contains an example of how to launch a static website using [Google Cloud Storage](https://cloud.google.com/storage/). See [cloud-storage-static-website example documentation](https://github.com/simplycubed/terraform-google-static-assets/blob/master/examples/cloud-storage-static-website) for the documentation.

- [modules](https://github.com/simplycubed/terraform-google-static-assets/blob/master/modules): This folder contains the main implementation code for this Module.

  The primary modules are:

  - [cloud-storage-static-website](https://github.com/simplycubed/terraform-google-static-assets/blob/master/modules/cloud-storage-static-website):
    The Cloud Storage Static Website module is used to create a [Google Cloud Storage](https://cloud.google.com/storage/)
    bucket that can be used to host a [static website](https://cloud.google.com/storage/docs/hosting-static-website).

  - [http-load-balancer-website](https://github.com/simplycubed/terraform-google-static-assets/blob/master/modules/http-load-balancer-website):
    The HTTP Load Balancer Website module is used to create a [HTTP Load Balancer](https://cloud.google.com/load-balancing/docs/https/)
    that routes requests to a [Google Cloud Storage](https://cloud.google.com/storage/) bucket for static content hosting,
    allowing you to also configure SSL with a custom domain name.

- [examples](https://github.com/simplycubed/terraform-google-static-assets/blob/master/examples): This folder contains examples of how to use the submodules.

- [test](https://github.com/simplycubed/terraform-google-static-assets/blob/master/test): Automated tests for the submodules and examples.

## Who maintains this Module?

This Module and its Submodules are maintained by [SimplyCubed](http://www.simplycubed.com/). If you are looking for help or commercial support, send an email to
[support@simplycubed.com](mailto:support@simplycubed.com?Subject=Google%20Static%20Assets%20Module).

SimplyCubed can help with:

- Setup, customization, and support for this Module.
- Modules and submodules for other types of infrastructure, such as VPCs, Docker clusters, databases, and continuous
  integration.
- Modules and Submodules that meet compliance requirements.
- Consulting & Training on GCP, Terraform, and DevOps.

## How do I contribute to this Module?

Contributions are very welcome! Check out the [Contribution Guidelines](https://github.com/simplycubed/terraform-google-static-assets/blob/master/CONTRIBUTING.md) for instructions.

## How is this Module versioned?

This Module follows the principles of [Semantic Versioning](http://semver.org/). You can find each new release, along
with the changelog, in the [Releases Page](https://github.com/simplycubed/terraform-google-static-assets/releases).

During initial development, the major version will be 0 (e.g., `0.x.y`), which indicates the code does not yet have a stable API. Once we hit `1.0.0`, we will make every effort to maintain a backwards compatible API and use the MAJOR, MINOR, and PATCH versions on each release to indicate any incompatibilities.

## License

Please see [LICENSE.txt](https://github.com/simplycubed/terraform-google-static-assets/blob/master/LICENSE.txt) for details on how the code in this repo is licensed.
