module github.com/replicatedhq/kots-field-labs/setup

go 1.19

require (
	github.com/aws/aws-lambda-go v1.31.1
	github.com/gosimple/slug v1.9.0
	github.com/pkg/errors v0.9.1
	github.com/replicatedhq/replicated v0.40.2
)

require (
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	github.com/rainycape/unidecode v0.0.0-20150907023854-cb7f23ec59be // indirect
	github.com/stretchr/testify v1.7.0 // indirect
)

replace github.com/terraform-providers/terraform-provider-tls => github.com/terraform-providers/terraform-provider-tls v1.2.1-0.20190816230231-0790c4b40281
