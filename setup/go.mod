module github.com/replicatedhq/kots-field-labs/setup

go 1.16

require (
	github.com/gosimple/slug v1.9.0
	github.com/johandry/terranova v0.0.3
	github.com/manifoldco/promptui v0.7.0
	github.com/pkg/errors v0.9.1
	github.com/replicatedhq/replicated v0.33.8
	github.com/stretchr/testify v1.7.0
	github.com/terraform-providers/terraform-provider-aws v1.60.1-0.20191003145700-f8707a46c6ec
	google.golang.org/api v0.43.0 // indirect
)

replace github.com/terraform-providers/terraform-provider-tls => github.com/terraform-providers/terraform-provider-tls v1.2.1-0.20190816230231-0790c4b40281
